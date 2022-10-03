package filestoreinstance


type FilestoreInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#create FilestoreInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#delete FilestoreInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#update FilestoreInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

