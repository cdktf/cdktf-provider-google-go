// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudquotasquotapreference


type CloudQuotasQuotaPreferenceQuotaConfig struct {
	// The preferred value.
	//
	// Must be greater than or equal to -1. If set to -1, it means the value is "unlimited".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/cloud_quotas_quota_preference#preferred_value CloudQuotasQuotaPreference#preferred_value}
	PreferredValue *string `field:"required" json:"preferredValue" yaml:"preferredValue"`
	// The annotations map for clients to store small amounts of arbitrary data.
	//
	// Do not put PII or other sensitive information here. See https://google.aip.dev/128#annotations.
	//
	// An object containing a list of "key: value" pairs. Example: '{ "name": "wrench", "mass": "1.3kg", "count": "3" }'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/cloud_quotas_quota_preference#annotations CloudQuotasQuotaPreference#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
}

