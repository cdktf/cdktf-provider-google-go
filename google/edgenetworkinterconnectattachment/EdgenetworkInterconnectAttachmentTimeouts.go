// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgenetworkinterconnectattachment


type EdgenetworkInterconnectAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/edgenetwork_interconnect_attachment#create EdgenetworkInterconnectAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/edgenetwork_interconnect_attachment#delete EdgenetworkInterconnectAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/edgenetwork_interconnect_attachment#update EdgenetworkInterconnectAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

