package apigeeflowhook


type ApigeeFlowhookTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_flowhook#create ApigeeFlowhook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_flowhook#delete ApigeeFlowhook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

