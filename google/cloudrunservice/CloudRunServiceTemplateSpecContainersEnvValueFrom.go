package cloudrunservice


type CloudRunServiceTemplateSpecContainersEnvValueFrom struct {
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#secret_key_ref CloudRunService#secret_key_ref}
	SecretKeyRef *CloudRunServiceTemplateSpecContainersEnvValueFromSecretKeyRef `field:"required" json:"secretKeyRef" yaml:"secretKeyRef"`
}

