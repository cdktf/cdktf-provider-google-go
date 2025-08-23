// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexairagengineconfig


type VertexAiRagEngineConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_rag_engine_config#create VertexAiRagEngineConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_rag_engine_config#delete VertexAiRagEngineConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vertex_ai_rag_engine_config#update VertexAiRagEngineConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

