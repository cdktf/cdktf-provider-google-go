// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration


type ClouddomainsRegistrationDnsSettingsCustomDnsDsRecords struct {
	// The algorithm used to generate the referenced DNSKEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/clouddomains_registration#algorithm ClouddomainsRegistration#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The digest generated from the referenced DNSKEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/clouddomains_registration#digest ClouddomainsRegistration#digest}
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// The hash function used to generate the digest of the referenced DNSKEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/clouddomains_registration#digest_type ClouddomainsRegistration#digest_type}
	DigestType *string `field:"optional" json:"digestType" yaml:"digestType"`
	// The key tag of the record. Must be set in range 0 -- 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/clouddomains_registration#key_tag ClouddomainsRegistration#key_tag}
	KeyTag *float64 `field:"optional" json:"keyTag" yaml:"keyTag"`
}

