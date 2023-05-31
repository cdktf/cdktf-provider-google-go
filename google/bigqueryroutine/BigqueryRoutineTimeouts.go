package bigqueryroutine


type BigqueryRoutineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/bigquery_routine#create BigqueryRoutine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/bigquery_routine#delete BigqueryRoutine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/bigquery_routine#update BigqueryRoutine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

