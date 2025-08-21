// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apikeyskey


type ApikeysKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/apikeys_key#create ApikeysKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/apikeys_key#delete ApikeysKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/apikeys_key#update ApikeysKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

