package bigqueryjob


type BigqueryJobExtractSourceTable struct {
	// The table. Can be specified '{{table_id}}' if 'project_id' and 'dataset_id' are also set, or of the form 'projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}' if not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#table_id BigqueryJob#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#dataset_id BigqueryJob#dataset_id}
	DatasetId *string `field:"optional" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#project_id BigqueryJob#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

