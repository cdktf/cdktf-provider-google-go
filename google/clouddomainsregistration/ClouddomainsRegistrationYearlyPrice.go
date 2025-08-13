// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationYearlyPrice struct {
	// The three-letter currency code defined in ISO 4217.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/clouddomains_registration#currency_code ClouddomainsRegistration#currency_code}
	CurrencyCode *string `field:"optional" json:"currencyCode" yaml:"currencyCode"`
	// The whole units of the amount. For example if currencyCode is "USD", then 1 unit is one US dollar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/clouddomains_registration#units ClouddomainsRegistration#units}
	Units *string `field:"optional" json:"units" yaml:"units"`
}

