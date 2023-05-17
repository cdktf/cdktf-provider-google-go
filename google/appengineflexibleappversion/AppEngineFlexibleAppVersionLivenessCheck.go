package appengineflexibleappversion


type AppEngineFlexibleAppVersionLivenessCheck struct {
	// The request path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#path AppEngineFlexibleAppVersion#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Interval between health checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#check_interval AppEngineFlexibleAppVersion#check_interval}
	CheckInterval *string `field:"optional" json:"checkInterval" yaml:"checkInterval"`
	// Number of consecutive failed checks required before considering the VM unhealthy. Default: 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#failure_threshold AppEngineFlexibleAppVersion#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Host header to send when performing a HTTP Readiness check. Example: "myapp.appspot.com".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#host AppEngineFlexibleAppVersion#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The initial delay before starting to execute the checks. Default: "300s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#initial_delay AppEngineFlexibleAppVersion#initial_delay}
	InitialDelay *string `field:"optional" json:"initialDelay" yaml:"initialDelay"`
	// Number of consecutive successful checks required before considering the VM healthy. Default: 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#success_threshold AppEngineFlexibleAppVersion#success_threshold}
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Time before the check is considered failed. Default: "4s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#timeout AppEngineFlexibleAppVersion#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

