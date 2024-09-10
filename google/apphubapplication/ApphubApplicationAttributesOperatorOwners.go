// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubapplication


type ApphubApplicationAttributesOperatorOwners struct {
	// Required. Email address of the contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/apphub_application#email ApphubApplication#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Optional. Contact's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/apphub_application#display_name ApphubApplication#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

