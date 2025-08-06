// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentEndpointConfig struct {
	// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com. The limitations will be removed soon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#dedicated_endpoint_enabled VertexAiEndpointWithModelGardenDeployment#dedicated_endpoint_enabled}
	DedicatedEndpointEnabled interface{} `field:"optional" json:"dedicatedEndpointEnabled" yaml:"dedicatedEndpointEnabled"`
	// The user-specified display name of the endpoint. If not set, a default name will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#endpoint_display_name VertexAiEndpointWithModelGardenDeployment#endpoint_display_name}
	EndpointDisplayName *string `field:"optional" json:"endpointDisplayName" yaml:"endpointDisplayName"`
}

