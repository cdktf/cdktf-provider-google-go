// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeatureonlinestorefeatureview


type VertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource struct {
	// feature_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/vertex_ai_feature_online_store_featureview#feature_groups VertexAiFeatureOnlineStoreFeatureview#feature_groups}
	FeatureGroups interface{} `field:"required" json:"featureGroups" yaml:"featureGroups"`
	// The project number of the parent project of the feature Groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/vertex_ai_feature_online_store_featureview#project_number VertexAiFeatureOnlineStoreFeatureview#project_number}
	ProjectNumber *string `field:"optional" json:"projectNumber" yaml:"projectNumber"`
}

