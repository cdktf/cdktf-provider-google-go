// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apikeyskey


type ApikeysKeyRestrictionsAndroidKeyRestrictions struct {
	// allowed_applications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/apikeys_key#allowed_applications ApikeysKey#allowed_applications}
	AllowedApplications interface{} `field:"required" json:"allowedApplications" yaml:"allowedApplications"`
}

