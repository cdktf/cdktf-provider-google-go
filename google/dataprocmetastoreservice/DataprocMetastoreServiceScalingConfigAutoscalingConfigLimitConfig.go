// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig struct {
	// The maximum scaling factor that the service will autoscale to. The default value is 6.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/dataproc_metastore_service#max_scaling_factor DataprocMetastoreService#max_scaling_factor}
	MaxScalingFactor *float64 `field:"optional" json:"maxScalingFactor" yaml:"maxScalingFactor"`
	// The minimum scaling factor that the service will autoscale to. The default value is 0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/dataproc_metastore_service#min_scaling_factor DataprocMetastoreService#min_scaling_factor}
	MinScalingFactor *float64 `field:"optional" json:"minScalingFactor" yaml:"minScalingFactor"`
}

