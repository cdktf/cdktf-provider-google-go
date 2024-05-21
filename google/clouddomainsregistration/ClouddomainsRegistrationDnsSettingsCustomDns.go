// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationDnsSettingsCustomDns struct {
	// Required.
	//
	// A list of name servers that store the DNS zone for this domain. Each name server is a domain
	// name, with Unicode domain names expressed in Punycode format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/clouddomains_registration#name_servers ClouddomainsRegistration#name_servers}
	NameServers *[]*string `field:"required" json:"nameServers" yaml:"nameServers"`
	// ds_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/clouddomains_registration#ds_records ClouddomainsRegistration#ds_records}
	DsRecords interface{} `field:"optional" json:"dsRecords" yaml:"dsRecords"`
}

