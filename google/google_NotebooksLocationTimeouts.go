// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type NotebooksLocationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_location#create NotebooksLocation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_location#delete NotebooksLocation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_location#update NotebooksLocation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

