package spannerdatabase


type SpannerDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database#create SpannerDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database#delete SpannerDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_database#update SpannerDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

