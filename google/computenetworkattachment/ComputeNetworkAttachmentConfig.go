// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeNetworkAttachmentConfig struct {
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
	// The connection preference of service attachment.
	//
	// The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules. Possible values: ["ACCEPT_AUTOMATIC", "ACCEPT_MANUAL", "INVALID"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#connection_preference ComputeNetworkAttachment#connection_preference}
	ConnectionPreference *string `field:"required" json:"connectionPreference" yaml:"connectionPreference"`
	// Name of the resource.
	//
	// Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#name ComputeNetworkAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#subnetworks ComputeNetworkAttachment#subnetworks}
	Subnetworks *[]*string `field:"required" json:"subnetworks" yaml:"subnetworks"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#description ComputeNetworkAttachment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Projects that are allowed to connect to this network attachment.
	//
	// The project can be specified using its id or number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#producer_accept_lists ComputeNetworkAttachment#producer_accept_lists}
	ProducerAcceptLists *[]*string `field:"optional" json:"producerAcceptLists" yaml:"producerAcceptLists"`
	// Projects that are not allowed to connect to this network attachment.
	//
	// The project can be specified using its id or number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#producer_reject_lists ComputeNetworkAttachment#producer_reject_lists}
	ProducerRejectLists *[]*string `field:"optional" json:"producerRejectLists" yaml:"producerRejectLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#project ComputeNetworkAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// URL of the region where the network attachment resides.
	//
	// This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#region ComputeNetworkAttachment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_network_attachment#timeouts ComputeNetworkAttachment#timeouts}
	Timeouts *ComputeNetworkAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

