// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnspolicy


type DnsPolicyAlternativeNameServerConfig struct {
	// target_name_servers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/dns_policy#target_name_servers DnsPolicy#target_name_servers}
	TargetNameServers interface{} `field:"required" json:"targetNameServers" yaml:"targetNameServers"`
}

