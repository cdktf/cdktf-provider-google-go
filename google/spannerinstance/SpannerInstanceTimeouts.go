package spannerinstance


type SpannerInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_instance#create SpannerInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_instance#delete SpannerInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/spanner_instance#update SpannerInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

