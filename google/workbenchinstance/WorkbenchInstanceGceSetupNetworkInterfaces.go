// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance


type WorkbenchInstanceGceSetupNetworkInterfaces struct {
	// Optional. The name of the VPC that this VM instance is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/workbench_instance#network WorkbenchInstance#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Optional.
	//
	// The type of vNIC to be used on this interface. This
	// may be gVNIC or VirtioNet. Possible values: ["VIRTIO_NET", "GVNIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/workbench_instance#nic_type WorkbenchInstance#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// Optional. The name of the subnet that this VM instance is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/workbench_instance#subnet WorkbenchInstance#subnet}
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
}

