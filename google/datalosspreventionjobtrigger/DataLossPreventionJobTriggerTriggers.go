package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerTriggers struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.0/docs/resources/data_loss_prevention_job_trigger#schedule DataLossPreventionJobTrigger#schedule}
	Schedule *DataLossPreventionJobTriggerTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

