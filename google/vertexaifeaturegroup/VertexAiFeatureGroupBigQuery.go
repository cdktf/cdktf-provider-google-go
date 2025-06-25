// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeaturegroup


type VertexAiFeatureGroupBigQuery struct {
	// big_query_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/vertex_ai_feature_group#big_query_source VertexAiFeatureGroup#big_query_source}
	BigQuerySource *VertexAiFeatureGroupBigQueryBigQuerySource `field:"required" json:"bigQuerySource" yaml:"bigQuerySource"`
	// Columns to construct entityId / row keys. If not provided defaults to entityId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/vertex_ai_feature_group#entity_id_columns VertexAiFeatureGroup#entity_id_columns}
	EntityIdColumns *[]*string `field:"optional" json:"entityIdColumns" yaml:"entityIdColumns"`
}

