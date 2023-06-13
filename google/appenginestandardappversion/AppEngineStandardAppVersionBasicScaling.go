package appenginestandardappversion


type AppEngineStandardAppVersionBasicScaling struct {
	// Maximum number of instances to create for this version. Must be in the range [1.0, 200.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_standard_app_version#max_instances AppEngineStandardAppVersion#max_instances}
	MaxInstances *float64 `field:"required" json:"maxInstances" yaml:"maxInstances"`
	// Duration of time after the last request that an instance must wait before the instance is shut down.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_standard_app_version#idle_timeout AppEngineStandardAppVersion#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

