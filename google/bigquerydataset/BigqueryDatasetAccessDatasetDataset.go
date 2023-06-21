package bigquerydataset


type BigqueryDatasetAccessDatasetDataset struct {
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/bigquery_dataset#dataset_id BigqueryDataset#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/bigquery_dataset#project_id BigqueryDataset#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

