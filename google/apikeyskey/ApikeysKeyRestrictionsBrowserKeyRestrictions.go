// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apikeyskey


type ApikeysKeyRestrictionsBrowserKeyRestrictions struct {
	// A list of regular expressions for the referrer URLs that are allowed to make API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/apikeys_key#allowed_referrers ApikeysKey#allowed_referrers}
	AllowedReferrers *[]*string `field:"required" json:"allowedReferrers" yaml:"allowedReferrers"`
}

