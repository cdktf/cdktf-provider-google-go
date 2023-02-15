package logginglogview


type LoggingLogViewTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_log_view#create LoggingLogView#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_log_view#delete LoggingLogView#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_log_view#update LoggingLogView#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

