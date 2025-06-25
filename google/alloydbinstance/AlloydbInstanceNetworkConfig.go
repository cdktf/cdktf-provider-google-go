// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstanceNetworkConfig struct {
	// Name of the allocated IP range for the private IP AlloyDB instance, for example: "google-managed-services-default".
	//
	// If set, the instance IPs will be created from this allocated range and will override the IP range used by the parent cluster.
	// The range name must comply with RFC 1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/alloydb_instance#allocated_ip_range_override AlloydbInstance#allocated_ip_range_override}
	AllocatedIpRangeOverride *string `field:"optional" json:"allocatedIpRangeOverride" yaml:"allocatedIpRangeOverride"`
	// authorized_external_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/alloydb_instance#authorized_external_networks AlloydbInstance#authorized_external_networks}
	AuthorizedExternalNetworks interface{} `field:"optional" json:"authorizedExternalNetworks" yaml:"authorizedExternalNetworks"`
	// Enabling outbound public ip for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/alloydb_instance#enable_outbound_public_ip AlloydbInstance#enable_outbound_public_ip}
	EnableOutboundPublicIp interface{} `field:"optional" json:"enableOutboundPublicIp" yaml:"enableOutboundPublicIp"`
	// Enabling public ip for the instance.
	//
	// If a user wishes to disable this,
	// please also clear the list of the authorized external networks set on
	// the same instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/alloydb_instance#enable_public_ip AlloydbInstance#enable_public_ip}
	EnablePublicIp interface{} `field:"optional" json:"enablePublicIp" yaml:"enablePublicIp"`
}

