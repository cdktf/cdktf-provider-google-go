// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress struct {
	// Required.
	//
	// CLDR region code of the country/region of the address. This is never inferred and it is up to the user to
	// ensure the value is correct. See https://cldr.unicode.org/ and
	// https://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/clouddomains_registration#region_code ClouddomainsRegistration#region_code}
	RegionCode *string `field:"required" json:"regionCode" yaml:"regionCode"`
	// Unstructured address lines describing the lower levels of an address.
	//
	// Because values in addressLines do not have type information and may sometimes contain multiple values in a single
	// field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be
	// "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language
	// is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way,
	// the most specific line of an address can be selected based on the language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/clouddomains_registration#address_lines ClouddomainsRegistration#address_lines}
	AddressLines *[]*string `field:"optional" json:"addressLines" yaml:"addressLines"`
	// Highest administrative subdivision which is used for postal addresses of a country or region.
	//
	// For example, this can be a state,
	// a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community
	// (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland
	// this should be left unpopulated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/clouddomains_registration#administrative_area ClouddomainsRegistration#administrative_area}
	AdministrativeArea *string `field:"optional" json:"administrativeArea" yaml:"administrativeArea"`
	// Generally refers to the city/town portion of the address.
	//
	// Examples: US city, IT comune, UK post town. In regions of the world
	// where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/clouddomains_registration#locality ClouddomainsRegistration#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The name of the organization at the address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/clouddomains_registration#organization ClouddomainsRegistration#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Postal code of the address.
	//
	// Not all countries use or require postal codes to be present, but where they are used,
	// they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/clouddomains_registration#postal_code ClouddomainsRegistration#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The recipient at the address.
	//
	// This field may, under certain circumstances, contain multiline information. For example,
	// it might contain "care of" information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/clouddomains_registration#recipients ClouddomainsRegistration#recipients}
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
}

