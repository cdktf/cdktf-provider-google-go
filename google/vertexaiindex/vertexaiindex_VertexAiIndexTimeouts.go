package vertexaiindex


type VertexAiIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#create VertexAiIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#delete VertexAiIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_index#update VertexAiIndex#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

