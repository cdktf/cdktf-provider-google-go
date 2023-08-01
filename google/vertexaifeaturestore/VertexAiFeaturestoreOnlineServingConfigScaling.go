package vertexaifeaturestore


type VertexAiFeaturestoreOnlineServingConfigScaling struct {
	// The maximum number of nodes to scale up to.
	//
	// Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_featurestore#max_node_count VertexAiFeaturestore#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// The minimum number of nodes to scale down to. Must be greater than or equal to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_featurestore#min_node_count VertexAiFeaturestore#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

