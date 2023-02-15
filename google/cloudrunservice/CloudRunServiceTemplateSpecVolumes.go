package cloudrunservice


type CloudRunServiceTemplateSpecVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#secret CloudRunService#secret}
	Secret *CloudRunServiceTemplateSpecVolumesSecret `field:"required" json:"secret" yaml:"secret"`
}

