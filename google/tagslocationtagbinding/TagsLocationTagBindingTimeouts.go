// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagslocationtagbinding


type TagsLocationTagBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/tags_location_tag_binding#create TagsLocationTagBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/tags_location_tag_binding#delete TagsLocationTagBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

