package cloudbuildtrigger


type CloudbuildTriggerWebhookConfig struct {
	// Resource name for the secret required as a URL parameter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#secret CloudbuildTrigger#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

