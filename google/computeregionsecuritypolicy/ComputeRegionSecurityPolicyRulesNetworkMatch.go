// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy


type ComputeRegionSecurityPolicyRulesNetworkMatch struct {
	// Destination IPv4/IPv6 addresses or CIDR prefixes, in standard text format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#dest_ip_ranges ComputeRegionSecurityPolicy#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// Destination port numbers for TCP/UDP/SCTP.
	//
	// Each element can be a 16-bit unsigned decimal number (e.g. "80") or range (e.g. "0-1023").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#dest_ports ComputeRegionSecurityPolicy#dest_ports}
	DestPorts *[]*string `field:"optional" json:"destPorts" yaml:"destPorts"`
	// IPv4 protocol / IPv6 next header (after extension headers).
	//
	// Each element can be an 8-bit unsigned decimal number (e.g. "6"), range (e.g. "253-254"), or one of the following protocol names: "tcp", "udp", "icmp", "esp", "ah", "ipip", or "sctp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#ip_protocols ComputeRegionSecurityPolicy#ip_protocols}
	IpProtocols *[]*string `field:"optional" json:"ipProtocols" yaml:"ipProtocols"`
	// BGP Autonomous System Number associated with the source IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#src_asns ComputeRegionSecurityPolicy#src_asns}
	SrcAsns *[]*float64 `field:"optional" json:"srcAsns" yaml:"srcAsns"`
	// Source IPv4/IPv6 addresses or CIDR prefixes, in standard text format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#src_ip_ranges ComputeRegionSecurityPolicy#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
	// Source port numbers for TCP/UDP/SCTP.
	//
	// Each element can be a 16-bit unsigned decimal number (e.g. "80") or range (e.g. "0-1023").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#src_ports ComputeRegionSecurityPolicy#src_ports}
	SrcPorts *[]*string `field:"optional" json:"srcPorts" yaml:"srcPorts"`
	// Two-letter ISO 3166-1 alpha-2 country code associated with the source IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#src_region_codes ComputeRegionSecurityPolicy#src_region_codes}
	SrcRegionCodes *[]*string `field:"optional" json:"srcRegionCodes" yaml:"srcRegionCodes"`
	// user_defined_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#user_defined_fields ComputeRegionSecurityPolicy#user_defined_fields}
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
}

