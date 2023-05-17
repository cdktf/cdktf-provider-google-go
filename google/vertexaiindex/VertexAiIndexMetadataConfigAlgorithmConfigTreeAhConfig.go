package vertexaiindex


type VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig struct {
	// Number of embeddings on each leaf node. The default value is 1000 if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#leaf_node_embedding_count VertexAiIndex#leaf_node_embedding_count}
	LeafNodeEmbeddingCount *float64 `field:"optional" json:"leafNodeEmbeddingCount" yaml:"leafNodeEmbeddingCount"`
	// The default percentage of leaf nodes that any query may be searched.
	//
	// Must be in
	// range 1-100, inclusive. The default value is 10 (means 10%) if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#leaf_nodes_to_search_percent VertexAiIndex#leaf_nodes_to_search_percent}
	LeafNodesToSearchPercent *float64 `field:"optional" json:"leafNodesToSearchPercent" yaml:"leafNodesToSearchPercent"`
}

