package notebooksruntime


type NotebooksRuntimeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_runtime#create NotebooksRuntime#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_runtime#delete NotebooksRuntime#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_runtime#update NotebooksRuntime#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

