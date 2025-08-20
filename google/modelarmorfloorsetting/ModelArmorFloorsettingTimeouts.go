// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting


type ModelArmorFloorsettingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/model_armor_floorsetting#create ModelArmorFloorsetting#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/model_armor_floorsetting#delete ModelArmorFloorsetting#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/model_armor_floorsetting#update ModelArmorFloorsetting#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

