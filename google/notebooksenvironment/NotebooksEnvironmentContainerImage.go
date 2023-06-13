package notebooksenvironment


type NotebooksEnvironmentContainerImage struct {
	// The path to the container image repository. For example: gcr.io/{project_id}/{imageName}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/notebooks_environment#repository NotebooksEnvironment#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The tag of the container image. If not specified, this defaults to the latest tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/notebooks_environment#tag NotebooksEnvironment#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

