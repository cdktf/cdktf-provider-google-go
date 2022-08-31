// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataLossPreventionJobTriggerTriggers struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#schedule DataLossPreventionJobTrigger#schedule}
	Schedule *DataLossPreventionJobTriggerTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

