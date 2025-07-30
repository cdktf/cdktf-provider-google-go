// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnspolicy


type DnsPolicyDns64ConfigScope struct {
	// Controls whether DNS64 is enabled globally at the network level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dns_policy#all_queries DnsPolicy#all_queries}
	AllQueries interface{} `field:"optional" json:"allQueries" yaml:"allQueries"`
}

