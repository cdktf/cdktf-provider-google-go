package notebooksenvironment


type NotebooksEnvironmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_environment#create NotebooksEnvironment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_environment#delete NotebooksEnvironment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_environment#update NotebooksEnvironment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

