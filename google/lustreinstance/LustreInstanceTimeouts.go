// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lustreinstance


type LustreInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/lustre_instance#create LustreInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/lustre_instance#delete LustreInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/lustre_instance#update LustreInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

