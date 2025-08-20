// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbeTcpSocket struct {
	// Optional: Host name to connect to, defaults to the model serving container's IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#host VertexAiEndpointWithModelGardenDeployment#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Number of the port to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#port VertexAiEndpointWithModelGardenDeployment#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

