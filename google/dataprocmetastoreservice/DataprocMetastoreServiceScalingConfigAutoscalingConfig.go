// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceScalingConfigAutoscalingConfig struct {
	// Defines whether autoscaling is enabled. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_metastore_service#autoscaling_enabled DataprocMetastoreService#autoscaling_enabled}
	AutoscalingEnabled interface{} `field:"optional" json:"autoscalingEnabled" yaml:"autoscalingEnabled"`
	// limit_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_metastore_service#limit_config DataprocMetastoreService#limit_config}
	LimitConfig *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig `field:"optional" json:"limitConfig" yaml:"limitConfig"`
}

