// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafusioninstance


type DataFusionInstanceNetworkConfigPrivateServiceConnectConfig struct {
	// Optional.
	//
	// The reference to the network attachment used to establish private connectivity.
	// It will be of the form projects/{project-id}/regions/{region}/networkAttachments/{network-attachment-id}.
	// This is required only when using connection type PRIVATE_SERVICE_CONNECT_INTERFACES.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/data_fusion_instance#network_attachment DataFusionInstance#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
	// Optional.
	//
	// Input only. The CIDR block to which the CDF instance can't route traffic to in the consumer project VPC.
	// The size of this block should be at least /25. This range should not overlap with the primary address range of any subnetwork used by the network attachment.
	// This range can be used for other purposes in the consumer VPC as long as there is no requirement for CDF to reach destinations using these addresses.
	// If this value is not provided, the server chooses a non RFC 1918 address range. The format of this field is governed by RFC 4632.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/data_fusion_instance#unreachable_cidr_block DataFusionInstance#unreachable_cidr_block}
	UnreachableCidrBlock *string `field:"optional" json:"unreachableCidrBlock" yaml:"unreachableCidrBlock"`
}

