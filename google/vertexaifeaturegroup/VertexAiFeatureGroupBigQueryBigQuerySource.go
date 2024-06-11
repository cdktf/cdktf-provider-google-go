// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeaturegroup


type VertexAiFeatureGroupBigQueryBigQuerySource struct {
	// BigQuery URI to a table, up to 2000 characters long. For example: 'bq://projectId.bqDatasetId.bqTableId.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.33.0/docs/resources/vertex_ai_feature_group#input_uri VertexAiFeatureGroup#input_uri}
	InputUri *string `field:"required" json:"inputUri" yaml:"inputUri"`
}

