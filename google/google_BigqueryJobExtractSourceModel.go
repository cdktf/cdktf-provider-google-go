// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigqueryJobExtractSourceModel struct {
	// The ID of the dataset containing this model.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#dataset_id BigqueryJob#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the model.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#model_id BigqueryJob#model_id}
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// The ID of the project containing this model.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#project_id BigqueryJob#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

