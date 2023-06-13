package notebooksinstance


type NotebooksInstanceVmImage struct {
	// The name of the Google Cloud project that this VM image belongs to. Format: projects/{project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/notebooks_instance#project NotebooksInstance#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Use this VM image family to find the image; the newest image in this family will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/notebooks_instance#image_family NotebooksInstance#image_family}
	ImageFamily *string `field:"optional" json:"imageFamily" yaml:"imageFamily"`
	// Use VM image name to find the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/notebooks_instance#image_name NotebooksInstance#image_name}
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
}

