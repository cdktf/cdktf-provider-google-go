// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcgoogleapisource


type EventarcGoogleApiSourceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/eventarc_google_api_source#create EventarcGoogleApiSource#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/eventarc_google_api_source#delete EventarcGoogleApiSource#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/eventarc_google_api_source#update EventarcGoogleApiSource#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

