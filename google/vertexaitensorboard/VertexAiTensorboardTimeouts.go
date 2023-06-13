package vertexaitensorboard


type VertexAiTensorboardTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_tensorboard#create VertexAiTensorboard#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_tensorboard#delete VertexAiTensorboard#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_tensorboard#update VertexAiTensorboard#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

