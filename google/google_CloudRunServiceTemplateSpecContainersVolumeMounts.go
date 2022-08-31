// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type CloudRunServiceTemplateSpecContainersVolumeMounts struct {
	// Path within the container at which the volume should be mounted.  Must not contain ':'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#mount_path CloudRunService#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// This must match the Name of a Volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

