package secretmanagersecretversion


type SecretManagerSecretVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret_version#create SecretManagerSecretVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret_version#delete SecretManagerSecretVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

