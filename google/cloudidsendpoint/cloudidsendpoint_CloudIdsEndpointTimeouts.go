package cloudidsendpoint


type CloudIdsEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_ids_endpoint#create CloudIdsEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_ids_endpoint#delete CloudIdsEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_ids_endpoint#update CloudIdsEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

