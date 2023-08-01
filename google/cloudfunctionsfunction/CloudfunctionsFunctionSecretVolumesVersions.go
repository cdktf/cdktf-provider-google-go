package cloudfunctionsfunction


type CloudfunctionsFunctionSecretVolumesVersions struct {
	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available.
	//
	// For example, setting the mount_path as "/etc/secrets" and path as "/secret_foo" would mount the secret value file at "/etc/secrets/secret_foo".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudfunctions_function#path CloudfunctionsFunction#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Version of the secret (version number or the string "latest").
	//
	// It is preferable to use "latest" version with secret volumes as secret value changes are reflected immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudfunctions_function#version CloudfunctionsFunction#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

