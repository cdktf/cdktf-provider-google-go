package appengineapplication


type AppEngineApplicationFeatureSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#split_health_checks AppEngineApplication#split_health_checks}.
	SplitHealthChecks interface{} `field:"required" json:"splitHealthChecks" yaml:"splitHealthChecks"`
}

