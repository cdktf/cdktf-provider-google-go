package composerenvironment


type ComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocks struct {
	// cidr_block must be specified in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#cidr_block ComposerEnvironment#cidr_block}
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// display_name is a field for users to identify CIDR blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#display_name ComposerEnvironment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

