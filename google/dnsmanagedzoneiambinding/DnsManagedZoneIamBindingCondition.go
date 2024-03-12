// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsmanagedzoneiambinding


type DnsManagedZoneIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/dns_managed_zone_iam_binding#expression DnsManagedZoneIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/dns_managed_zone_iam_binding#title DnsManagedZoneIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/dns_managed_zone_iam_binding#description DnsManagedZoneIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

