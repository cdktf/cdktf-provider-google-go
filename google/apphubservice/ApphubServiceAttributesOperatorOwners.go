// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubservice


type ApphubServiceAttributesOperatorOwners struct {
	// Required. Email address of the contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/apphub_service#email ApphubService#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Contact's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/apphub_service#display_name ApphubService#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

