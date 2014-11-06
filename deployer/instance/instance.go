package instance

import (
	"time"

	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	boshsys "github.com/cloudfoundry/bosh-agent/system"

	bmcloud "github.com/cloudfoundry/bosh-micro-cli/cloud"
	bmagentclient "github.com/cloudfoundry/bosh-micro-cli/deployer/agentclient"
	bmas "github.com/cloudfoundry/bosh-micro-cli/deployer/applyspec"
	bmdisk "github.com/cloudfoundry/bosh-micro-cli/deployer/disk"
	bmretrystrategy "github.com/cloudfoundry/bosh-micro-cli/deployer/retrystrategy"
	bmstemcell "github.com/cloudfoundry/bosh-micro-cli/deployer/stemcell"
	bmdepl "github.com/cloudfoundry/bosh-micro-cli/deployment"
)

type instance struct {
	vmCID                  string
	agentClient            bmagentclient.AgentClient
	cloud                  bmcloud.Cloud
	templatesSpecGenerator TemplatesSpecGenerator
	applySpecFactory       bmas.Factory
	mbusURL                string
	fs                     boshsys.FileSystem
	logger                 boshlog.Logger
	logTag                 string
}

type Instance interface {
	WaitToBeReady(maxAttempts int, delay time.Duration) error
	Apply(bmstemcell.ApplySpec, bmdepl.Deployment) error
	Start() error
	WaitToBeRunning(maxAttempts int, delay time.Duration) error
	AttachDisk(bmdisk.Disk) error
}

func NewInstance(
	vmCID string,
	agentClient bmagentclient.AgentClient,
	cloud bmcloud.Cloud,
	templatesSpecGenerator TemplatesSpecGenerator,
	applySpecFactory bmas.Factory,
	mbusURL string,
	fs boshsys.FileSystem,
	logger boshlog.Logger,
) Instance {
	return &instance{
		vmCID:       vmCID,
		agentClient: agentClient,
		cloud:       cloud,
		templatesSpecGenerator: templatesSpecGenerator,
		applySpecFactory:       applySpecFactory,
		mbusURL:                mbusURL,
		fs:                     fs,
		logger:                 logger,
		logTag:                 "instanceUpdater",
	}
}

func (i *instance) WaitToBeReady(maxAttempts int, delay time.Duration) error {
	agentPingRetryable := bmagentclient.NewPingRetryable(i.agentClient)
	agentPingRetryStrategy := bmretrystrategy.NewAttemptRetryStrategy(maxAttempts, delay, agentPingRetryable, i.logger)
	return agentPingRetryStrategy.Try()
}

func (i *instance) Apply(stemcellApplySpec bmstemcell.ApplySpec, deployment bmdepl.Deployment) error {
	i.logger.Debug(i.logTag, "Stopping agent")

	err := i.agentClient.Stop()
	if err != nil {
		return bosherr.WrapError(err, "Stopping agent")
	}

	i.logger.Debug(i.logTag, "Rendering job templates")
	renderedJobDir, err := i.fs.TempDir("instance-updater-render-job")
	if err != nil {
		return bosherr.WrapError(err, "Creating rendered job directory")
	}
	defer i.fs.RemoveAll(renderedJobDir)

	deploymentJob := deployment.Jobs[0]
	jobProperties, err := deploymentJob.Properties()
	if err != nil {
		return bosherr.WrapError(err, "Stringifying job properties")
	}

	networksSpec, err := deployment.NetworksSpec(deploymentJob.Name)
	if err != nil {
		return bosherr.WrapError(err, "Stringifying job properties")
	}

	templatesSpec, err := i.templatesSpecGenerator.Create(
		deploymentJob,
		stemcellApplySpec.Job,
		deployment.Name,
		jobProperties,
		i.mbusURL,
	)
	if err != nil {
		return bosherr.WrapError(err, "Generating templates spec")
	}

	i.logger.Debug(i.logTag, "Creating apply spec")
	agentApplySpec := i.applySpecFactory.Create(
		stemcellApplySpec,
		deployment.Name,
		deploymentJob.Name,
		networksSpec,
		templatesSpec.BlobID,
		templatesSpec.ArchiveSha1,
		templatesSpec.ConfigurationHash,
	)

	i.logger.Debug(i.logTag, "Sending apply message to the agent with '%#v'", agentApplySpec)
	err = i.agentClient.Apply(agentApplySpec)
	if err != nil {
		return bosherr.WrapError(err, "Sending apply spec to agent")
	}

	return nil
}

func (i *instance) Start() error {
	return i.agentClient.Start()
}

func (i *instance) WaitToBeRunning(maxAttempts int, delay time.Duration) error {
	agentGetStateRetryable := bmagentclient.NewGetStateRetryable(i.agentClient)
	agentGetStateRetryStrategy := bmretrystrategy.NewAttemptRetryStrategy(maxAttempts, delay, agentGetStateRetryable, i.logger)
	return agentGetStateRetryStrategy.Try()
}

func (i *instance) AttachDisk(disk bmdisk.Disk) error {
	err := i.cloud.AttachDisk(i.vmCID, disk.CID)
	if err != nil {
		return bosherr.WrapError(err, "Attaching disk in the cloud")
	}

	err = i.agentClient.MountDisk(disk.CID)
	if err != nil {
		return bosherr.WrapError(err, "Mounting disk")
	}

	return nil
}