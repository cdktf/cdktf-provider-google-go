package loggingmetric


type LoggingMetricTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#create LoggingMetric#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#delete LoggingMetric#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#update LoggingMetric#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

