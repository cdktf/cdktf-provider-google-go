// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstanceNetworkConfig struct {
	// authorized_external_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/alloydb_instance#authorized_external_networks AlloydbInstance#authorized_external_networks}
	AuthorizedExternalNetworks interface{} `field:"optional" json:"authorizedExternalNetworks" yaml:"authorizedExternalNetworks"`
	// Enabling outbound public ip for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/alloydb_instance#enable_outbound_public_ip AlloydbInstance#enable_outbound_public_ip}
	EnableOutboundPublicIp interface{} `field:"optional" json:"enableOutboundPublicIp" yaml:"enableOutboundPublicIp"`
	// Enabling public ip for the instance.
	//
	// If a user wishes to disable this,
	// please also clear the list of the authorized external networks set on
	// the same instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/alloydb_instance#enable_public_ip AlloydbInstance#enable_public_ip}
	EnablePublicIp interface{} `field:"optional" json:"enablePublicIp" yaml:"enablePublicIp"`
}

