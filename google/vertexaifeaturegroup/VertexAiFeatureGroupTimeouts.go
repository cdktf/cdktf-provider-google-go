// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeaturegroup


type VertexAiFeatureGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vertex_ai_feature_group#create VertexAiFeatureGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vertex_ai_feature_group#delete VertexAiFeatureGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/vertex_ai_feature_group#update VertexAiFeatureGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

