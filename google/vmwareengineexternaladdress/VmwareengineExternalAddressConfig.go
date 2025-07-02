// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineexternaladdress

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareengineExternalAddressConfig struct {
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
	// The internal IP address of a workload VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/vmwareengine_external_address#internal_ip VmwareengineExternalAddress#internal_ip}
	InternalIp *string `field:"required" json:"internalIp" yaml:"internalIp"`
	// The ID of the external IP Address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/vmwareengine_external_address#name VmwareengineExternalAddress#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource name of the private cloud to create a new external address in.
	//
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/vmwareengine_external_address#parent VmwareengineExternalAddress#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// User-provided description for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/vmwareengine_external_address#description VmwareengineExternalAddress#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/vmwareengine_external_address#id VmwareengineExternalAddress#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/vmwareengine_external_address#timeouts VmwareengineExternalAddress#timeouts}
	Timeouts *VmwareengineExternalAddressTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

