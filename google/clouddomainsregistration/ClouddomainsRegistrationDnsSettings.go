// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationDnsSettings struct {
	// custom_dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/clouddomains_registration#custom_dns ClouddomainsRegistration#custom_dns}
	CustomDns *ClouddomainsRegistrationDnsSettingsCustomDns `field:"optional" json:"customDns" yaml:"customDns"`
	// glue_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/clouddomains_registration#glue_records ClouddomainsRegistration#glue_records}
	GlueRecords interface{} `field:"optional" json:"glueRecords" yaml:"glueRecords"`
}

