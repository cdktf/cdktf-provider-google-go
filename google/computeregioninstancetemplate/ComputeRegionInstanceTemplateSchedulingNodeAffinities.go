// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate


type ComputeRegionInstanceTemplateSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_region_instance_template#key ComputeRegionInstanceTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_region_instance_template#operator ComputeRegionInstanceTemplate#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_region_instance_template#values ComputeRegionInstanceTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

