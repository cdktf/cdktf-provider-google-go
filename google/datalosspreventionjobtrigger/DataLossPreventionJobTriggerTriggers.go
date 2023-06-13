package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerTriggers struct {
	// manual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#manual DataLossPreventionJobTrigger#manual}
	Manual *DataLossPreventionJobTriggerTriggersManual `field:"optional" json:"manual" yaml:"manual"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#schedule DataLossPreventionJobTrigger#schedule}
	Schedule *DataLossPreventionJobTriggerTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

