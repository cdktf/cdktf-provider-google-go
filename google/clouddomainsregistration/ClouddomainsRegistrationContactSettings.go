// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationContactSettings struct {
	// admin_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/clouddomains_registration#admin_contact ClouddomainsRegistration#admin_contact}
	AdminContact *ClouddomainsRegistrationContactSettingsAdminContact `field:"required" json:"adminContact" yaml:"adminContact"`
	// Required. Privacy setting for the contacts associated with the Registration. Values are PUBLIC_CONTACT_DATA, PRIVATE_CONTACT_DATA, and REDACTED_CONTACT_DATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/clouddomains_registration#privacy ClouddomainsRegistration#privacy}
	Privacy *string `field:"required" json:"privacy" yaml:"privacy"`
	// registrant_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/clouddomains_registration#registrant_contact ClouddomainsRegistration#registrant_contact}
	RegistrantContact *ClouddomainsRegistrationContactSettingsRegistrantContact `field:"required" json:"registrantContact" yaml:"registrantContact"`
	// technical_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/clouddomains_registration#technical_contact ClouddomainsRegistration#technical_contact}
	TechnicalContact *ClouddomainsRegistrationContactSettingsTechnicalContact `field:"required" json:"technicalContact" yaml:"technicalContact"`
}

