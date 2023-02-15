package secretmanagersecret


type SecretManagerSecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret#create SecretManagerSecret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret#delete SecretManagerSecret#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret#update SecretManagerSecret#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

