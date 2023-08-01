package cloudbuildtrigger


type CloudbuildTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#create CloudbuildTrigger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#delete CloudbuildTrigger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_trigger#update CloudbuildTrigger#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

