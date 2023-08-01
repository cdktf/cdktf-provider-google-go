package bigqueryjob


type BigqueryJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_job#create BigqueryJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_job#delete BigqueryJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

