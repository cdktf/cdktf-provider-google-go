// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type CloudRunServiceTemplateSpecContainersEnvFrom struct {
	// config_map_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#config_map_ref CloudRunService#config_map_ref}
	ConfigMapRef *CloudRunServiceTemplateSpecContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifier to prepend to each key in the ConfigMap.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#prefix CloudRunService#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#secret_ref CloudRunService#secret_ref}
	SecretRef *CloudRunServiceTemplateSpecContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

