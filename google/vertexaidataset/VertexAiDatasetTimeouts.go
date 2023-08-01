package vertexaidataset


type VertexAiDatasetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_dataset#create VertexAiDataset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_dataset#delete VertexAiDataset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/vertex_ai_dataset#update VertexAiDataset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

