// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubservice


type ApphubServiceAttributesBusinessOwners struct {
	// Required. Email address of the contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/apphub_service#email ApphubService#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Contact's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/apphub_service#display_name ApphubService#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

