package composerenvironment


type ComposerEnvironmentConfigMasterAuthorizedNetworksConfig struct {
	// Whether or not master authorized networks is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#enabled ComposerEnvironment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// cidr_blocks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#cidr_blocks ComposerEnvironment#cidr_blocks}
	CidrBlocks interface{} `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
}

