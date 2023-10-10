// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apikeyskey


type ApikeysKeyRestrictionsIosKeyRestrictions struct {
	// A list of bundle IDs that are allowed when making API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/apikeys_key#allowed_bundle_ids ApikeysKey#allowed_bundle_ids}
	AllowedBundleIds *[]*string `field:"required" json:"allowedBundleIds" yaml:"allowedBundleIds"`
}

