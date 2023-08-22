package appengineflexibleappversion


type AppEngineFlexibleAppVersionEndpointsApiService struct {
	// Endpoints service name which is the name of the "service" resource in the Service Management API. For example "myapi.endpoints.myproject.cloud.goog".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_flexible_app_version#name AppEngineFlexibleAppVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Endpoints service configuration ID as specified by the Service Management API. For example "2016-09-19r1".
	//
	// By default, the rollout strategy for Endpoints is "FIXED". This means that Endpoints starts up with a particular configuration ID.
	// When a new configuration is rolled out, Endpoints must be given the new configuration ID. The configId field is used to give the configuration ID
	// and is required in this case.
	//
	// Endpoints also has a rollout strategy called "MANAGED". When using this, Endpoints fetches the latest configuration and does not need
	// the configuration ID. In this case, configId must be omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_flexible_app_version#config_id AppEngineFlexibleAppVersion#config_id}
	ConfigId *string `field:"optional" json:"configId" yaml:"configId"`
	// Enable or disable trace sampling. By default, this is set to false for enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_flexible_app_version#disable_trace_sampling AppEngineFlexibleAppVersion#disable_trace_sampling}
	DisableTraceSampling interface{} `field:"optional" json:"disableTraceSampling" yaml:"disableTraceSampling"`
	// Endpoints rollout strategy.
	//
	// If FIXED, configId must be specified. If MANAGED, configId must be omitted. Default value: "FIXED" Possible values: ["FIXED", "MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_flexible_app_version#rollout_strategy AppEngineFlexibleAppVersion#rollout_strategy}
	RolloutStrategy *string `field:"optional" json:"rolloutStrategy" yaml:"rolloutStrategy"`
}

