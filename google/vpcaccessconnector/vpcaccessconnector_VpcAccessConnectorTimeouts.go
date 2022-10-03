package vpcaccessconnector


type VpcAccessConnectorTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vpc_access_connector#create VpcAccessConnector#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vpc_access_connector#delete VpcAccessConnector#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

