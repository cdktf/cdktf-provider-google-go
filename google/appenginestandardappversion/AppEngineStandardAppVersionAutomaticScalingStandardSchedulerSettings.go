// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appenginestandardappversion


type AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings struct {
	// Maximum number of instances to run for this version. Set to zero to disable maxInstances configuration.
	//
	// **Note:** Starting from March 2025, App Engine sets the maxInstances default for standard environment deployments to 20. This change doesn't impact existing apps. To override the default, specify a new value between 0 and 2147483647, and deploy a new version or redeploy over an existing version. To disable the maxInstances default configuration setting, specify the maximum permitted value 2147483647.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/app_engine_standard_app_version#max_instances AppEngineStandardAppVersion#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// Minimum number of instances to run for this version. Set to zero to disable minInstances configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/app_engine_standard_app_version#min_instances AppEngineStandardAppVersion#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// Target CPU utilization ratio to maintain when scaling.
	//
	// Should be a value in the range [0.50, 0.95], zero, or a negative value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/app_engine_standard_app_version#target_cpu_utilization AppEngineStandardAppVersion#target_cpu_utilization}
	TargetCpuUtilization *float64 `field:"optional" json:"targetCpuUtilization" yaml:"targetCpuUtilization"`
	// Target throughput utilization ratio to maintain when scaling.
	//
	// Should be a value in the range [0.50, 0.95], zero, or a negative value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/app_engine_standard_app_version#target_throughput_utilization AppEngineStandardAppVersion#target_throughput_utilization}
	TargetThroughputUtilization *float64 `field:"optional" json:"targetThroughputUtilization" yaml:"targetThroughputUtilization"`
}

