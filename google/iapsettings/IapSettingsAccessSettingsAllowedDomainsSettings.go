// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings


type IapSettingsAccessSettingsAllowedDomainsSettings struct {
	// List of trusted domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/iap_settings#domains IapSettings#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Configuration for customers to opt in for the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/iap_settings#enable IapSettings#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

