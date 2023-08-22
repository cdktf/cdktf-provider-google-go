package vpcaccessconnector


type VpcAccessConnectorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/vpc_access_connector#create VpcAccessConnector#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/vpc_access_connector#delete VpcAccessConnector#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

