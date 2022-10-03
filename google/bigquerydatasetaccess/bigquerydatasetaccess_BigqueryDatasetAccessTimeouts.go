package bigquerydatasetaccess


type BigqueryDatasetAccessTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset_access#create BigqueryDatasetAccessA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset_access#delete BigqueryDatasetAccessA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

