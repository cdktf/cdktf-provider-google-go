// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiindexendpoint


type VertexAiIndexEndpointPrivateServiceConnectConfig struct {
	// If set to true, the IndexEndpoint is created without private service access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/vertex_ai_index_endpoint#enable_private_service_connect VertexAiIndexEndpoint#enable_private_service_connect}
	EnablePrivateServiceConnect interface{} `field:"required" json:"enablePrivateServiceConnect" yaml:"enablePrivateServiceConnect"`
	// A list of Projects from which the forwarding rule will target the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/vertex_ai_index_endpoint#project_allowlist VertexAiIndexEndpoint#project_allowlist}
	ProjectAllowlist *[]*string `field:"optional" json:"projectAllowlist" yaml:"projectAllowlist"`
}

