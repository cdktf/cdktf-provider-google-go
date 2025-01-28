// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigHostConfig struct {
	// DNS search domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gkeonprem_vmware_cluster#dns_search_domains GkeonpremVmwareCluster#dns_search_domains}
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// DNS servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gkeonprem_vmware_cluster#dns_servers GkeonpremVmwareCluster#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// NTP servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gkeonprem_vmware_cluster#ntp_servers GkeonpremVmwareCluster#ntp_servers}
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
}

