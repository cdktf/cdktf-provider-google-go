package cloudbuildtrigger


type CloudbuildTriggerBuildAvailableSecrets struct {
	// secret_manager block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#secret_manager CloudbuildTrigger#secret_manager}
	SecretManager interface{} `field:"required" json:"secretManager" yaml:"secretManager"`
}

