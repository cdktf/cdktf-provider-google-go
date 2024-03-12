// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateScratchDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/compute_instance_from_template#device_name ComputeInstanceFromTemplate#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/compute_instance_from_template#interface ComputeInstanceFromTemplate#interface}.
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/compute_instance_from_template#size ComputeInstanceFromTemplate#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

