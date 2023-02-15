package notebooksinstance


type NotebooksInstanceContainerImage struct {
	// The path to the container image repository. For example: gcr.io/{project_id}/{imageName}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#repository NotebooksInstance#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The tag of the container image. If not specified, this defaults to the latest tag.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#tag NotebooksInstance#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

