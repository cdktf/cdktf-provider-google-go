// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket


type StorageBucketIpFilterPublicNetworkSource struct {
	// The list of public IPv4, IPv6 cidr ranges that are allowed to access the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_bucket#allowed_ip_cidr_ranges StorageBucket#allowed_ip_cidr_ranges}
	AllowedIpCidrRanges *[]*string `field:"required" json:"allowedIpCidrRanges" yaml:"allowedIpCidrRanges"`
}

