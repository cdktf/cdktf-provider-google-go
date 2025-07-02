// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket


type StorageBucketIpFilterVpcNetworkSources struct {
	// The list of public or private IPv4 and IPv6 CIDR ranges that can access the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/storage_bucket#allowed_ip_cidr_ranges StorageBucket#allowed_ip_cidr_ranges}
	AllowedIpCidrRanges *[]*string `field:"required" json:"allowedIpCidrRanges" yaml:"allowedIpCidrRanges"`
	// Name of the network. Format: projects/{PROJECT_ID}/global/networks/{NETWORK_NAME}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/storage_bucket#network StorageBucket#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

