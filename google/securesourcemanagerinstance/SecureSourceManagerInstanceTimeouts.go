// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerinstance


type SecureSourceManagerInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/secure_source_manager_instance#create SecureSourceManagerInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/secure_source_manager_instance#delete SecureSourceManagerInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/secure_source_manager_instance#update SecureSourceManagerInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

