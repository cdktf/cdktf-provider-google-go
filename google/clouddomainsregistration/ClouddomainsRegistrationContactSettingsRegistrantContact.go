// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationContactSettingsRegistrantContact struct {
	// Required. Email address of the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/clouddomains_registration#email ClouddomainsRegistration#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Required. Phone number of the contact in international format. For example, "+1-800-555-0123".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/clouddomains_registration#phone_number ClouddomainsRegistration#phone_number}
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// postal_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/clouddomains_registration#postal_address ClouddomainsRegistration#postal_address}
	PostalAddress *ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress `field:"required" json:"postalAddress" yaml:"postalAddress"`
	// Fax number of the contact in international format. For example, "+1-800-555-0123".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/clouddomains_registration#fax_number ClouddomainsRegistration#fax_number}
	FaxNumber *string `field:"optional" json:"faxNumber" yaml:"faxNumber"`
}

