// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationManagementSettings struct {
	// The desired renewal method for this Registration.
	//
	// The actual renewalMethod is automatically updated to reflect this choice.
	// If unset or equal to RENEWAL_METHOD_UNSPECIFIED, the actual renewalMethod is treated as if it were set to AUTOMATIC_RENEWAL.
	// You cannot use RENEWAL_DISABLED during resource creation, and you can update the renewal status only when the Registration
	// resource has state ACTIVE or SUSPENDED.
	//
	// When preferredRenewalMethod is set to AUTOMATIC_RENEWAL, the actual renewalMethod can be set to RENEWAL_DISABLED in case of
	// problems with the billing account or reported domain abuse. In such cases, check the issues field on the Registration. After
	// the problem is resolved, the renewalMethod is automatically updated to preferredRenewalMethod in a few hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/clouddomains_registration#preferred_renewal_method ClouddomainsRegistration#preferred_renewal_method}
	PreferredRenewalMethod *string `field:"optional" json:"preferredRenewalMethod" yaml:"preferredRenewalMethod"`
	// Controls whether the domain can be transferred to another registrar. Values are UNLOCKED or LOCKED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/clouddomains_registration#transfer_lock_state ClouddomainsRegistration#transfer_lock_state}
	TransferLockState *string `field:"optional" json:"transferLockState" yaml:"transferLockState"`
}

