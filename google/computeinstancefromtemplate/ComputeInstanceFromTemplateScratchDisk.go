// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateScratchDisk struct {
	// The disk interface used for attaching this disk. One of SCSI or NVME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_instance_from_template#interface ComputeInstanceFromTemplate#interface}
	Interface *string `field:"required" json:"interface" yaml:"interface"`
	// Name with which the attached disk is accessible under /dev/disk/by-id/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_instance_from_template#device_name ComputeInstanceFromTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The size of the disk in gigabytes. One of 375 or 3000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_instance_from_template#size ComputeInstanceFromTemplate#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

