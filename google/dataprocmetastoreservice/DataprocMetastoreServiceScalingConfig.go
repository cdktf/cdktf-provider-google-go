// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceScalingConfig struct {
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataproc_metastore_service#autoscaling_config DataprocMetastoreService#autoscaling_config}
	AutoscalingConfig *DataprocMetastoreServiceScalingConfigAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// Metastore instance sizes. Possible values: ["EXTRA_SMALL", "SMALL", "MEDIUM", "LARGE", "EXTRA_LARGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataproc_metastore_service#instance_size DataprocMetastoreService#instance_size}
	InstanceSize *string `field:"optional" json:"instanceSize" yaml:"instanceSize"`
	// Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataproc_metastore_service#scaling_factor DataprocMetastoreService#scaling_factor}
	ScalingFactor *float64 `field:"optional" json:"scalingFactor" yaml:"scalingFactor"`
}

