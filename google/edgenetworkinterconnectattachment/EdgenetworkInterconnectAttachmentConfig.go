// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgenetworkinterconnectattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgenetworkInterconnectAttachmentConfig struct {
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
	// The ID of the underlying interconnect that this attachment's traffic will traverse through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#interconnect EdgenetworkInterconnectAttachment#interconnect}
	Interconnect *string `field:"required" json:"interconnect" yaml:"interconnect"`
	// A unique ID that identifies this interconnect attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#interconnect_attachment_id EdgenetworkInterconnectAttachment#interconnect_attachment_id}
	InterconnectAttachmentId *string `field:"required" json:"interconnectAttachmentId" yaml:"interconnectAttachmentId"`
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#location EdgenetworkInterconnectAttachment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the network to which this interconnect attachment belongs. Must be of the form: 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#network EdgenetworkInterconnectAttachment#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// VLAN ID provided by user. Must be site-wise unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#vlan_id EdgenetworkInterconnectAttachment#vlan_id}
	VlanId *float64 `field:"required" json:"vlanId" yaml:"vlanId"`
	// The name of the target Distributed Cloud Edge zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#zone EdgenetworkInterconnectAttachment#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#description EdgenetworkInterconnectAttachment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#id EdgenetworkInterconnectAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels associated with this resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#labels EdgenetworkInterconnectAttachment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// IP (L3) MTU value of the virtual edge cloud. Default value is '1500'. Possible values are: '1500', '9000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#mtu EdgenetworkInterconnectAttachment#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#project EdgenetworkInterconnectAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/edgenetwork_interconnect_attachment#timeouts EdgenetworkInterconnectAttachment#timeouts}
	Timeouts *EdgenetworkInterconnectAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

