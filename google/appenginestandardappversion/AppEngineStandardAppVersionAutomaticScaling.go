package appenginestandardappversion


type AppEngineStandardAppVersionAutomaticScaling struct {
	// Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
	//
	// Defaults to a runtime-specific value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#max_concurrent_requests AppEngineStandardAppVersion#max_concurrent_requests}
	MaxConcurrentRequests *float64 `field:"optional" json:"maxConcurrentRequests" yaml:"maxConcurrentRequests"`
	// Maximum number of idle instances that should be maintained for this version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#max_idle_instances AppEngineStandardAppVersion#max_idle_instances}
	MaxIdleInstances *float64 `field:"optional" json:"maxIdleInstances" yaml:"maxIdleInstances"`
	// Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#max_pending_latency AppEngineStandardAppVersion#max_pending_latency}
	MaxPendingLatency *string `field:"optional" json:"maxPendingLatency" yaml:"maxPendingLatency"`
	// Minimum number of idle instances that should be maintained for this version.
	//
	// Only applicable for the default version of a service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#min_idle_instances AppEngineStandardAppVersion#min_idle_instances}
	MinIdleInstances *float64 `field:"optional" json:"minIdleInstances" yaml:"minIdleInstances"`
	// Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#min_pending_latency AppEngineStandardAppVersion#min_pending_latency}
	MinPendingLatency *string `field:"optional" json:"minPendingLatency" yaml:"minPendingLatency"`
	// standard_scheduler_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#standard_scheduler_settings AppEngineStandardAppVersion#standard_scheduler_settings}
	StandardSchedulerSettings *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings `field:"optional" json:"standardSchedulerSettings" yaml:"standardSchedulerSettings"`
}

