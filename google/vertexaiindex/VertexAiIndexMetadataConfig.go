package vertexaiindex


type VertexAiIndexMetadataConfig struct {
	// The number of dimensions of the input vectors.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#dimensions VertexAiIndex#dimensions}
	Dimensions *float64 `field:"required" json:"dimensions" yaml:"dimensions"`
	// algorithm_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#algorithm_config VertexAiIndex#algorithm_config}
	AlgorithmConfig *VertexAiIndexMetadataConfigAlgorithmConfig `field:"optional" json:"algorithmConfig" yaml:"algorithmConfig"`
	// The default number of neighbors to find via approximate search before exact reordering is performed.
	//
	// Exact reordering is a procedure where results returned by an
	// approximate search algorithm are reordered via a more expensive distance computation.
	// Required if tree-AH algorithm is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#approximate_neighbors_count VertexAiIndex#approximate_neighbors_count}
	ApproximateNeighborsCount *float64 `field:"optional" json:"approximateNeighborsCount" yaml:"approximateNeighborsCount"`
	// The distance measure used in nearest neighbor search.
	//
	// The value must be one of the followings:
	// SQUARED_L2_DISTANCE: Euclidean (L_2) Distance
	// L1_DISTANCE: Manhattan (L_1) Distance
	// COSINE_DISTANCE: Cosine Distance. Defined as 1 - cosine similarity.
	// DOT_PRODUCT_DISTANCE: Dot Product Distance. Defined as a negative of the dot product
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#distance_measure_type VertexAiIndex#distance_measure_type}
	DistanceMeasureType *string `field:"optional" json:"distanceMeasureType" yaml:"distanceMeasureType"`
	// Type of normalization to be carried out on each vector.
	//
	// The value must be one of the followings:
	// UNIT_L2_NORM: Unit L2 normalization type
	// NONE: No normalization type is specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#feature_norm_type VertexAiIndex#feature_norm_type}
	FeatureNormType *string `field:"optional" json:"featureNormType" yaml:"featureNormType"`
}
