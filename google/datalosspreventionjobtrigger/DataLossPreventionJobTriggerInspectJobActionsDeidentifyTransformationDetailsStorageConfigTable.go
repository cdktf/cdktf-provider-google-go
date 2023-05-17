package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationDetailsStorageConfigTable struct {
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#dataset_id DataLossPreventionJobTrigger#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#project_id DataLossPreventionJobTrigger#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ID of the table.
	//
	// The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#table_id DataLossPreventionJobTrigger#table_id}
	TableId *string `field:"optional" json:"tableId" yaml:"tableId"`
}

