package cloudrunv2job


type CloudRunV2JobTemplateTemplateContainersStartupProbeTcpSocket struct {
	// Port number to access on the container.
	//
	// Must be in the range 1 to 65535. If not specified, defaults to 8080.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_job#port CloudRunV2Job#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}
