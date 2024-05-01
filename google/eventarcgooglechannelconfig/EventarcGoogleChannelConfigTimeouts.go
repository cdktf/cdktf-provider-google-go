// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcgooglechannelconfig


type EventarcGoogleChannelConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/eventarc_google_channel_config#create EventarcGoogleChannelConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/eventarc_google_channel_config#delete EventarcGoogleChannelConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/eventarc_google_channel_config#update EventarcGoogleChannelConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

