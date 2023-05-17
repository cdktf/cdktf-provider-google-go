package appengineflexibleappversion


type AppEngineFlexibleAppVersionResourcesVolumes struct {
	// Unique name for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#name AppEngineFlexibleAppVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Volume size in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#size_gb AppEngineFlexibleAppVersion#size_gb}
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Underlying volume type, e.g. 'tmpfs'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#volume_type AppEngineFlexibleAppVersion#volume_type}
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
}

