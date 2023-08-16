package appenginestandardappversion


type AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings struct {
	// Maximum number of instances to run for this version. Set to zero to disable maxInstances configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#max_instances AppEngineStandardAppVersion#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// Minimum number of instances to run for this version. Set to zero to disable minInstances configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#min_instances AppEngineStandardAppVersion#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// Target CPU utilization ratio to maintain when scaling.
	//
	// Should be a value in the range [0.50, 0.95], zero, or a negative value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#target_cpu_utilization AppEngineStandardAppVersion#target_cpu_utilization}
	TargetCpuUtilization *float64 `field:"optional" json:"targetCpuUtilization" yaml:"targetCpuUtilization"`
	// Target throughput utilization ratio to maintain when scaling.
	//
	// Should be a value in the range [0.50, 0.95], zero, or a negative value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_standard_app_version#target_throughput_utilization AppEngineStandardAppVersion#target_throughput_utilization}
	TargetThroughputUtilization *float64 `field:"optional" json:"targetThroughputUtilization" yaml:"targetThroughputUtilization"`
}

