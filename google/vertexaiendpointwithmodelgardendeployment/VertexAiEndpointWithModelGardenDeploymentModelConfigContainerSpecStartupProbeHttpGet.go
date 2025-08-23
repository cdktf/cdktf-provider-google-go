// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeHttpGet struct {
	// Host name to connect to, defaults to the model serving container's IP.
	//
	// You probably want to set "Host" in httpHeaders instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#host VertexAiEndpointWithModelGardenDeployment#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#http_headers VertexAiEndpointWithModelGardenDeployment#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#path VertexAiEndpointWithModelGardenDeployment#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Number of the port to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#port VertexAiEndpointWithModelGardenDeployment#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Scheme to use for connecting to the host. Defaults to HTTP. Acceptable values are "HTTP" or "HTTPS".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#scheme VertexAiEndpointWithModelGardenDeployment#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

