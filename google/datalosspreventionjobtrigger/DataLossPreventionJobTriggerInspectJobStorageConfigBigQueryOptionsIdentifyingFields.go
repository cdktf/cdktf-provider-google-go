package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields struct {
	// Name of a BigQuery field to be returned with the findings.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#name DataLossPreventionJobTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

