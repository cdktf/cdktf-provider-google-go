package notebooksinstance


type NotebooksInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/notebooks_instance#create NotebooksInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/notebooks_instance#delete NotebooksInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.70.0/docs/resources/notebooks_instance#update NotebooksInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

