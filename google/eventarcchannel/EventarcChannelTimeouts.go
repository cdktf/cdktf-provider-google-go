// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcchannel


type EventarcChannelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/eventarc_channel#create EventarcChannel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/eventarc_channel#delete EventarcChannel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/eventarc_channel#update EventarcChannel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

