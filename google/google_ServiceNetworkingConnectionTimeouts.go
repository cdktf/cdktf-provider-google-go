// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ServiceNetworkingConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_networking_connection#create ServiceNetworkingConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_networking_connection#delete ServiceNetworkingConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_networking_connection#update ServiceNetworkingConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

