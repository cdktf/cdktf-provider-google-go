package bigquerydatasetaccess


type BigqueryDatasetAccessRoutineA struct {
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_dataset_access#dataset_id BigqueryDatasetAccessA#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_dataset_access#project_id BigqueryDatasetAccessA#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ID of the routine.
	//
	// The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_dataset_access#routine_id BigqueryDatasetAccessA#routine_id}
	RoutineId *string `field:"required" json:"routineId" yaml:"routineId"`
}

