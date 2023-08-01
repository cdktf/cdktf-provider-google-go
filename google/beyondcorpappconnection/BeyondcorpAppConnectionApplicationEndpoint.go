package beyondcorpappconnection


type BeyondcorpAppConnectionApplicationEndpoint struct {
	// Hostname or IP address of the remote application endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_connection#host BeyondcorpAppConnection#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port of the remote application endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/beyondcorp_app_connection#port BeyondcorpAppConnection#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

