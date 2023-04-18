package cloudrunv2service


type CloudRunV2ServiceTemplateContainersStartupProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.0/docs/resources/cloud_run_v2_service#http_headers CloudRunV2Service#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server. Defaults to '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.0/docs/resources/cloud_run_v2_service#path CloudRunV2Service#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

