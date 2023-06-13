package bigqueryjob


type BigqueryJobQueryDefaultDataset struct {
	// The dataset. Can be specified '{{dataset_id}}' if 'project_id' is also set, or of the form 'projects/{{project}}/datasets/{{dataset_id}}' if not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#dataset_id BigqueryJob#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#project_id BigqueryJob#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

