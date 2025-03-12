// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginecluster


type VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds struct {
	// The utilization triggering the scale-in operation in percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/vmwareengine_cluster#scale_in VmwareengineCluster#scale_in}
	ScaleIn *float64 `field:"required" json:"scaleIn" yaml:"scaleIn"`
	// The utilization triggering the scale-out operation in percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/vmwareengine_cluster#scale_out VmwareengineCluster#scale_out}
	ScaleOut *float64 `field:"required" json:"scaleOut" yaml:"scaleOut"`
}

