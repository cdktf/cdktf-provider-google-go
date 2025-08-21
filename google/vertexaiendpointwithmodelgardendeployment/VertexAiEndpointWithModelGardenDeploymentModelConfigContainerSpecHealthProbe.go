// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#exec VertexAiEndpointWithModelGardenDeployment#exec}
	Exec *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeExec `field:"optional" json:"exec" yaml:"exec"`
	// Number of consecutive failures before the probe is considered failed. Defaults to 3. Minimum value is 1.
	//
	// Maps to Kubernetes probe argument 'failureThreshold'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#failure_threshold VertexAiEndpointWithModelGardenDeployment#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#grpc VertexAiEndpointWithModelGardenDeployment#grpc}
	Grpc *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#http_get VertexAiEndpointWithModelGardenDeployment#http_get}
	HttpGet *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Number of seconds to wait before starting the probe. Defaults to 0. Minimum value is 0.
	//
	// Maps to Kubernetes probe argument 'initialDelaySeconds'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#initial_delay_seconds VertexAiEndpointWithModelGardenDeployment#initial_delay_seconds}
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds.
	// Minimum value is 1. Must be less than timeout_seconds.
	//
	// Maps to Kubernetes probe argument 'periodSeconds'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#period_seconds VertexAiEndpointWithModelGardenDeployment#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Number of consecutive successes before the probe is considered successful. Defaults to 1. Minimum value is 1.
	//
	// Maps to Kubernetes probe argument 'successThreshold'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#success_threshold VertexAiEndpointWithModelGardenDeployment#success_threshold}
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#tcp_socket VertexAiEndpointWithModelGardenDeployment#tcp_socket}
	TcpSocket *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second.
	// Minimum value is 1. Must be greater or equal to period_seconds.
	//
	// Maps to Kubernetes probe argument 'timeoutSeconds'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#timeout_seconds VertexAiEndpointWithModelGardenDeployment#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

