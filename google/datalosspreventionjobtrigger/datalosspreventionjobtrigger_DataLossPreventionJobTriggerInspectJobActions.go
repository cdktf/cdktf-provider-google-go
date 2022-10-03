package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActions struct {
	// save_findings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#save_findings DataLossPreventionJobTrigger#save_findings}
	SaveFindings *DataLossPreventionJobTriggerInspectJobActionsSaveFindings `field:"required" json:"saveFindings" yaml:"saveFindings"`
}

