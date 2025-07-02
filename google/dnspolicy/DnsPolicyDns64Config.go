// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnspolicy


type DnsPolicyDns64Config struct {
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dns_policy#scope DnsPolicy#scope}
	Scope *DnsPolicyDns64ConfigScope `field:"required" json:"scope" yaml:"scope"`
}

