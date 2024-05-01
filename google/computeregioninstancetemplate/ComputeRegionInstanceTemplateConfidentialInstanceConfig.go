// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate


type ComputeRegionInstanceTemplateConfidentialInstanceConfig struct {
	// Defines whether the instance should have confidential compute enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/compute_region_instance_template#enable_confidential_compute ComputeRegionInstanceTemplate#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"required" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

