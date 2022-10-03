package endpointsservice


type EndpointsServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service#create EndpointsService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service#delete EndpointsService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service#update EndpointsService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

