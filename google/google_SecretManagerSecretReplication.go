// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type SecretManagerSecretReplication struct {
	// The Secret will automatically be replicated without any restrictions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret#automatic SecretManagerSecret#automatic}
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// user_managed block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret#user_managed SecretManagerSecret#user_managed}
	UserManaged *SecretManagerSecretReplicationUserManaged `field:"optional" json:"userManaged" yaml:"userManaged"`
}

