// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeendpointattachment


type ApigeeEndpointAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/apigee_endpoint_attachment#create ApigeeEndpointAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/apigee_endpoint_attachment#delete ApigeeEndpointAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

