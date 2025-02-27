// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudquotasquotapreference


type CloudQuotasQuotaPreferenceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_quotas_quota_preference#create CloudQuotasQuotaPreference#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_quotas_quota_preference#delete CloudQuotasQuotaPreference#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_quotas_quota_preference#update CloudQuotasQuotaPreference#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

