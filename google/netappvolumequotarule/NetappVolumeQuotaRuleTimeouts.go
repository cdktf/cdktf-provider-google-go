// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumequotarule


type NetappVolumeQuotaRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/netapp_volume_quota_rule#create NetappVolumeQuotaRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/netapp_volume_quota_rule#delete NetappVolumeQuotaRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/netapp_volume_quota_rule#update NetappVolumeQuotaRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

