package vertexaiindex


type VertexAiIndexMetadata struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_index#config VertexAiIndex#config}
	Config *VertexAiIndexMetadataConfig `field:"optional" json:"config" yaml:"config"`
	// Allows inserting, updating  or deleting the contents of the Matching Engine Index.
	//
	// The string must be a valid Cloud Storage directory path. If this
	// field is set when calling IndexService.UpdateIndex, then no other
	// Index field can be also updated as part of the same call.
	// The expected structure and format of the files this URI points to is
	// described at https://cloud.google.com/vertex-ai/docs/matching-engine/using-matching-engine#input-data-format
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_index#contents_delta_uri VertexAiIndex#contents_delta_uri}
	ContentsDeltaUri *string `field:"optional" json:"contentsDeltaUri" yaml:"contentsDeltaUri"`
	// If this field is set together with contentsDeltaUri when calling IndexService.UpdateIndex, then existing content of the Index will be replaced by the data from the contentsDeltaUri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_index#is_complete_overwrite VertexAiIndex#is_complete_overwrite}
	IsCompleteOverwrite interface{} `field:"optional" json:"isCompleteOverwrite" yaml:"isCompleteOverwrite"`
}

