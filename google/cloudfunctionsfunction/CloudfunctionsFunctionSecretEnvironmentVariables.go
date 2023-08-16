package cloudfunctionsfunction


type CloudfunctionsFunctionSecretEnvironmentVariables struct {
	// Name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions_function#key CloudfunctionsFunction#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// ID of the secret in secret manager (not the full resource name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions_function#secret CloudfunctionsFunction#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// Version of the secret (version number or the string "latest").
	//
	// It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new clones start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions_function#version CloudfunctionsFunction#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret.
	//
	// If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions_function#project_id CloudfunctionsFunction#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

