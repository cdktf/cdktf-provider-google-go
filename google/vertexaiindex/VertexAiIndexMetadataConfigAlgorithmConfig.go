package vertexaiindex


type VertexAiIndexMetadataConfigAlgorithmConfig struct {
	// brute_force_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/vertex_ai_index#brute_force_config VertexAiIndex#brute_force_config}
	BruteForceConfig *VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig `field:"optional" json:"bruteForceConfig" yaml:"bruteForceConfig"`
	// tree_ah_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/vertex_ai_index#tree_ah_config VertexAiIndex#tree_ah_config}
	TreeAhConfig *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig `field:"optional" json:"treeAhConfig" yaml:"treeAhConfig"`
}

