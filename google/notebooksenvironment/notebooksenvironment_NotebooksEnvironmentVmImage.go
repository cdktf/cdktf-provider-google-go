package notebooksenvironment


type NotebooksEnvironmentVmImage struct {
	// The name of the Google Cloud project that this VM image belongs to. Format: projects/{project_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_environment#project NotebooksEnvironment#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Use this VM image family to find the image; the newest image in this family will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_environment#image_family NotebooksEnvironment#image_family}
	ImageFamily *string `field:"optional" json:"imageFamily" yaml:"imageFamily"`
	// Use VM image name to find the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_environment#image_name NotebooksEnvironment#image_name}
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
}

