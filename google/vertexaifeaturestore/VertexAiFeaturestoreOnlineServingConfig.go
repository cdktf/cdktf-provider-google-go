package vertexaifeaturestore


type VertexAiFeaturestoreOnlineServingConfig struct {
	// The number of nodes for each cluster.
	//
	// The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_featurestore#fixed_node_count VertexAiFeaturestore#fixed_node_count}
	FixedNodeCount *float64 `field:"optional" json:"fixedNodeCount" yaml:"fixedNodeCount"`
	// scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_featurestore#scaling VertexAiFeaturestore#scaling}
	Scaling *VertexAiFeaturestoreOnlineServingConfigScaling `field:"optional" json:"scaling" yaml:"scaling"`
}

