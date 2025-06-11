// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket


type StorageBucketIpFilter struct {
	// The mode of the IP filter. Valid values are 'Enabled' and 'Disabled'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_bucket#mode StorageBucket#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// public_network_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_bucket#public_network_source StorageBucket#public_network_source}
	PublicNetworkSource *StorageBucketIpFilterPublicNetworkSource `field:"optional" json:"publicNetworkSource" yaml:"publicNetworkSource"`
	// vpc_network_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/storage_bucket#vpc_network_sources StorageBucket#vpc_network_sources}
	VpcNetworkSources interface{} `field:"optional" json:"vpcNetworkSources" yaml:"vpcNetworkSources"`
}

