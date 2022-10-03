package bigqueryroutine


type BigqueryRoutineTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_routine#create BigqueryRoutine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_routine#delete BigqueryRoutine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_routine#update BigqueryRoutine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

