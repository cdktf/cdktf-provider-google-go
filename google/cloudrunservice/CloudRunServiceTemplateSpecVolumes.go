package cloudrunservice


type CloudRunServiceTemplateSpecVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#secret CloudRunService#secret}
	Secret *CloudRunServiceTemplateSpecVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

