// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpappconnection


type BeyondcorpAppConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/beyondcorp_app_connection#create BeyondcorpAppConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/beyondcorp_app_connection#delete BeyondcorpAppConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/beyondcorp_app_connection#update BeyondcorpAppConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

