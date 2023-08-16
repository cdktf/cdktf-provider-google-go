package secretmanagersecret


type SecretManagerSecretReplication struct {
	// The Secret will automatically be replicated without any restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/secret_manager_secret#automatic SecretManagerSecret#automatic}
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// user_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/secret_manager_secret#user_managed SecretManagerSecret#user_managed}
	UserManaged *SecretManagerSecretReplicationUserManaged `field:"optional" json:"userManaged" yaml:"userManaged"`
}

