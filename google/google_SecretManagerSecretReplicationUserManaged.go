// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type SecretManagerSecretReplicationUserManaged struct {
	// replicas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/secret_manager_secret#replicas SecretManagerSecret#replicas}
	Replicas interface{} `field:"required" json:"replicas" yaml:"replicas"`
}

