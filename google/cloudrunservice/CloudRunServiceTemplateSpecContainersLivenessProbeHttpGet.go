package cloudrunservice


type CloudRunServiceTemplateSpecContainersLivenessProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#http_headers CloudRunService#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server. If set, it should not be empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#path CloudRunService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port number to access on the container.
	//
	// Number must be in the range 1 to 65535.
	// If not specified, defaults to the same value as container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_run_service#port CloudRunService#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

