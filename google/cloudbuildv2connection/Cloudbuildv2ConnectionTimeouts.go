package cloudbuildv2connection


type Cloudbuildv2ConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#create Cloudbuildv2Connection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#delete Cloudbuildv2Connection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#update Cloudbuildv2Connection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

