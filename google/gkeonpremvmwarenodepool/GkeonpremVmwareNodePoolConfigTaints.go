// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarenodepool


type GkeonpremVmwareNodePoolConfigTaints struct {
	// Key associated with the effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/gkeonprem_vmware_node_pool#key GkeonpremVmwareNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value associated with the effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/gkeonprem_vmware_node_pool#value GkeonpremVmwareNodePool#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Available taint effects. Possible values: ["EFFECT_UNSPECIFIED", "NO_SCHEDULE", "PREFER_NO_SCHEDULE", "NO_EXECUTE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/gkeonprem_vmware_node_pool#effect GkeonpremVmwareNodePool#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
}

