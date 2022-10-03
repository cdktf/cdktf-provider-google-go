package bigqueryjob


type BigqueryJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#create BigqueryJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_job#delete BigqueryJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

