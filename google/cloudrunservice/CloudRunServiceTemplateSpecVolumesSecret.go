package cloudrunservice


type CloudRunServiceTemplateSpecVolumesSecret struct {
	// The name of the secret in Cloud Secret Manager.
	//
	// By default, the secret
	// is assumed to be in the same project.
	// If the secret is in another project, you must define an alias.
	// An alias definition has the form:
	// {alias}:projects/{project-id|project-number}/secrets/{secret-name}.
	// If multiple alias definitions are needed, they must be separated by
	// commas.
	// The alias definitions must be set on the run.googleapis.com/secrets
	// annotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#secret_name CloudRunService#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0000
	// and 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#default_mode CloudRunService#default_mode}
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#items CloudRunService#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

