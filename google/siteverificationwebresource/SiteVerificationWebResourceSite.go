// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siteverificationwebresource


type SiteVerificationWebResourceSite struct {
	// The site identifier.
	//
	// If the type is set to SITE, the identifier is a URL. If the type is
	// set to INET_DOMAIN, the identifier is a domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/site_verification_web_resource#identifier SiteVerificationWebResource#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The type of resource to be verified. Possible values: ["INET_DOMAIN", "SITE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/site_verification_web_resource#type SiteVerificationWebResource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

