package cloudrunservice


type CloudRunServiceTemplateSpecVolumesSecretItems struct {
	// The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#key CloudRunService#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The relative path of the file to map the key to.
	//
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#path CloudRunService#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Mode bits to use on this file, must be a value between 0000 and 0777.
	//
	// If
	// not specified, the volume defaultMode will be used. This might be in
	// conflict with other options that affect the file mode, like fsGroup, and
	// the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#mode CloudRunService#mode}
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
}

