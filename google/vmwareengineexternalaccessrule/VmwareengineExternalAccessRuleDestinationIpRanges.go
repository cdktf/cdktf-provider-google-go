// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineexternalaccessrule


type VmwareengineExternalAccessRuleDestinationIpRanges struct {
	// The name of an 'ExternalAddress' resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/vmwareengine_external_access_rule#external_address VmwareengineExternalAccessRule#external_address}
	ExternalAddress *string `field:"optional" json:"externalAddress" yaml:"externalAddress"`
	// An IP address range in the CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/vmwareengine_external_access_rule#ip_address_range VmwareengineExternalAccessRule#ip_address_range}
	IpAddressRange *string `field:"optional" json:"ipAddressRange" yaml:"ipAddressRange"`
}

