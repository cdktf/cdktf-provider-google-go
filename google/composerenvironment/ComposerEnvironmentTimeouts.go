package composerenvironment


type ComposerEnvironmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#create ComposerEnvironment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#delete ComposerEnvironment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#update ComposerEnvironment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

