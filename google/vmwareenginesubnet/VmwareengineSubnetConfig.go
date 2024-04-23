// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginesubnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareengineSubnetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The IP address range of the subnet in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/vmwareengine_subnet#ip_cidr_range VmwareengineSubnet#ip_cidr_range}
	IpCidrRange *string `field:"required" json:"ipCidrRange" yaml:"ipCidrRange"`
	// The ID of the subnet.
	//
	// For userDefined subnets, this name should be in the format of "service-n",
	// where n ranges from 1 to 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/vmwareengine_subnet#name VmwareengineSubnet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource name of the private cloud to create a new subnet in.
	//
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/vmwareengine_subnet#parent VmwareengineSubnet#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/vmwareengine_subnet#id VmwareengineSubnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/vmwareengine_subnet#timeouts VmwareengineSubnet#timeouts}
	Timeouts *VmwareengineSubnetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

