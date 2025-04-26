// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomaintrust


type ActiveDirectoryDomainTrustTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/active_directory_domain_trust#create ActiveDirectoryDomainTrust#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/active_directory_domain_trust#delete ActiveDirectoryDomainTrust#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/active_directory_domain_trust#update ActiveDirectoryDomainTrust#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

