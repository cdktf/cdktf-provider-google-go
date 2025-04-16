// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpoint


type VertexAiEndpointPrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/vertex_ai_endpoint#enable_private_service_connect VertexAiEndpoint#enable_private_service_connect}
	EnablePrivateServiceConnect interface{} `field:"required" json:"enablePrivateServiceConnect" yaml:"enablePrivateServiceConnect"`
	// If set to true, enable secure private service connect with IAM authorization.
	//
	// Otherwise, private service connect will be done without authorization. Note latency will be slightly increased if authorization is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/vertex_ai_endpoint#enable_secure_private_service_connect VertexAiEndpoint#enable_secure_private_service_connect}
	EnableSecurePrivateServiceConnect interface{} `field:"optional" json:"enableSecurePrivateServiceConnect" yaml:"enableSecurePrivateServiceConnect"`
	// A list of Projects from which the forwarding rule will target the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/vertex_ai_endpoint#project_allowlist VertexAiEndpoint#project_allowlist}
	ProjectAllowlist *[]*string `field:"optional" json:"projectAllowlist" yaml:"projectAllowlist"`
}

