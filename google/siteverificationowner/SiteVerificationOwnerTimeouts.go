// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siteverificationowner


type SiteVerificationOwnerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/site_verification_owner#create SiteVerificationOwner#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/site_verification_owner#delete SiteVerificationOwner#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

