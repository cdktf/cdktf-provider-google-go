// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeatureonlinestore


type VertexAiFeatureOnlineStoreDedicatedServingEndpoint struct {
	// private_service_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_feature_online_store#private_service_connect_config VertexAiFeatureOnlineStore#private_service_connect_config}
	PrivateServiceConnectConfig *VertexAiFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig `field:"optional" json:"privateServiceConnectConfig" yaml:"privateServiceConnectConfig"`
}

