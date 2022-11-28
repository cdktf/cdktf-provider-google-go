package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActions struct {
	// pub_sub block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#pub_sub DataLossPreventionJobTrigger#pub_sub}
	PubSub *DataLossPreventionJobTriggerInspectJobActionsPubSub `field:"optional" json:"pubSub" yaml:"pubSub"`
	// save_findings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#save_findings DataLossPreventionJobTrigger#save_findings}
	SaveFindings *DataLossPreventionJobTriggerInspectJobActionsSaveFindings `field:"optional" json:"saveFindings" yaml:"saveFindings"`
}

