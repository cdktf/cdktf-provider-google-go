// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsresponsepolicyrule


type DnsResponsePolicyRuleLocalData struct {
	// local_datas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dns_response_policy_rule#local_datas DnsResponsePolicyRule#local_datas}
	LocalDatas interface{} `field:"required" json:"localDatas" yaml:"localDatas"`
}

