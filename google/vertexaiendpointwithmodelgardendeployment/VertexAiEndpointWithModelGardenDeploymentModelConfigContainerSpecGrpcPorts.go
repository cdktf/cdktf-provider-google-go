// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPorts struct {
	// The number of the port to expose on the pod's IP address.
	//
	// Must be a valid port number, between 1 and 65535 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#container_port VertexAiEndpointWithModelGardenDeployment#container_port}
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
}

