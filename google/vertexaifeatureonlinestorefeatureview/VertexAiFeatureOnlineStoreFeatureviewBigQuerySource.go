// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeatureonlinestorefeatureview


type VertexAiFeatureOnlineStoreFeatureviewBigQuerySource struct {
	// Columns to construct entityId / row keys. Start by supporting 1 only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/vertex_ai_feature_online_store_featureview#entity_id_columns VertexAiFeatureOnlineStoreFeatureview#entity_id_columns}
	EntityIdColumns *[]*string `field:"required" json:"entityIdColumns" yaml:"entityIdColumns"`
	// The BigQuery view URI that will be materialized on each sync trigger based on FeatureView.SyncConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/vertex_ai_feature_online_store_featureview#uri VertexAiFeatureOnlineStoreFeatureview#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

