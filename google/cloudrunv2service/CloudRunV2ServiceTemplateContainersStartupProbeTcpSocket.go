package cloudrunv2service


type CloudRunV2ServiceTemplateContainersStartupProbeTcpSocket struct {
	// Port number to access on the container.
	//
	// Must be in the range 1 to 65535. If not specified, defaults to 8080.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_service#port CloudRunV2Service#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

