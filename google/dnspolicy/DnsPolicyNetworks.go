// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnspolicy


type DnsPolicyNetworks struct {
	// The id or fully qualified URL of the VPC network to forward queries to.
	//
	// This should be formatted like 'projects/{project}/global/networks/{network}' or
	// 'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dns_policy#network_url DnsPolicy#network_url}
	NetworkUrl *string `field:"required" json:"networkUrl" yaml:"networkUrl"`
}

