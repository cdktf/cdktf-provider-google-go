package cloudfunctions2function


type Cloudfunctions2FunctionServiceConfigSecretVolumesVersions struct {
	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available.
	//
	// For example, setting the mountPath as '/etc/secrets' and path as secret_foo would mount the secret value file at /etc/secrets/secret_foo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#path Cloudfunctions2Function#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Version of the secret (version number or the string 'latest').
	//
	// It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#version Cloudfunctions2Function#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

