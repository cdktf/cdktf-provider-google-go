// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexairagengineconfig


type VertexAiRagEngineConfigRagManagedDbConfig struct {
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_rag_engine_config#basic VertexAiRagEngineConfig#basic}
	Basic *VertexAiRagEngineConfigRagManagedDbConfigBasic `field:"optional" json:"basic" yaml:"basic"`
	// scaled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_rag_engine_config#scaled VertexAiRagEngineConfig#scaled}
	Scaled *VertexAiRagEngineConfigRagManagedDbConfigScaled `field:"optional" json:"scaled" yaml:"scaled"`
	// unprovisioned block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_rag_engine_config#unprovisioned VertexAiRagEngineConfig#unprovisioned}
	Unprovisioned *VertexAiRagEngineConfigRagManagedDbConfigUnprovisioned `field:"optional" json:"unprovisioned" yaml:"unprovisioned"`
}

