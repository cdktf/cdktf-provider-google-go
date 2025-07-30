// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnectattachmentgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInterconnectAttachmentGroupConfig struct {
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
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#intent ComputeInterconnectAttachmentGroup#intent}
	Intent *ComputeInterconnectAttachmentGroupIntent `field:"required" json:"intent" yaml:"intent"`
	// Name of the resource.
	//
	// Provided by the client when the resource is created. The name must be
	// 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#name ComputeInterconnectAttachmentGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// attachments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#attachments ComputeInterconnectAttachmentGroup#attachments}
	Attachments interface{} `field:"optional" json:"attachments" yaml:"attachments"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#description ComputeInterconnectAttachmentGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#id ComputeInterconnectAttachmentGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The URL of an InterconnectGroup that groups these Attachments' Interconnects.
	//
	// Customers do not need to set this unless directed by
	// Google Support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#interconnect_group ComputeInterconnectAttachmentGroup#interconnect_group}
	InterconnectGroup *string `field:"optional" json:"interconnectGroup" yaml:"interconnectGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#project ComputeInterconnectAttachmentGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#timeouts ComputeInterconnectAttachmentGroup#timeouts}
	Timeouts *ComputeInterconnectAttachmentGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

