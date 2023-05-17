package cloudbuildtrigger


type CloudbuildTriggerWebhookConfig struct {
	// Resource name for the secret required as a URL parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudbuild_trigger#secret CloudbuildTrigger#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

