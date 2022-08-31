// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigqueryDatasetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#create BigqueryDataset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#delete BigqueryDataset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#update BigqueryDataset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

