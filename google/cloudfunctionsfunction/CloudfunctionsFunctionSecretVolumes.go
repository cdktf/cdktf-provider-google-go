package cloudfunctionsfunction


type CloudfunctionsFunctionSecretVolumes struct {
	// The path within the container to mount the secret volume.
	//
	// For example, setting the mount_path as "/etc/secrets" would mount the secret value files under the "/etc/secrets" directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount paths: "/etc/secrets" Restricted mount paths: "/cloudsql", "/dev/log", "/pod", "/proc", "/var/log".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#mount_path CloudfunctionsFunction#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// ID of the secret in secret manager (not the full resource name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#secret CloudfunctionsFunction#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret.
	//
	// If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#project_id CloudfunctionsFunction#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#versions CloudfunctionsFunction#versions}
	Versions interface{} `field:"optional" json:"versions" yaml:"versions"`
}

