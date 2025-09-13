// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket


type StorageBucketIpFilter struct {
	// The mode of the IP filter. Valid values are 'Enabled' and 'Disabled'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_bucket#mode StorageBucket#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Whether to allow all service agents to access the bucket regardless of the IP filter configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_bucket#allow_all_service_agent_access StorageBucket#allow_all_service_agent_access}
	AllowAllServiceAgentAccess interface{} `field:"optional" json:"allowAllServiceAgentAccess" yaml:"allowAllServiceAgentAccess"`
	// Whether to allow cross-org VPCs in the bucket's IP filter configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_bucket#allow_cross_org_vpcs StorageBucket#allow_cross_org_vpcs}
	AllowCrossOrgVpcs interface{} `field:"optional" json:"allowCrossOrgVpcs" yaml:"allowCrossOrgVpcs"`
	// public_network_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_bucket#public_network_source StorageBucket#public_network_source}
	PublicNetworkSource *StorageBucketIpFilterPublicNetworkSource `field:"optional" json:"publicNetworkSource" yaml:"publicNetworkSource"`
	// vpc_network_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/storage_bucket#vpc_network_sources StorageBucket#vpc_network_sources}
	VpcNetworkSources interface{} `field:"optional" json:"vpcNetworkSources" yaml:"vpcNetworkSources"`
}

