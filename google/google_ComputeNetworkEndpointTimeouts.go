// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeNetworkEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_network_endpoint#create ComputeNetworkEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_network_endpoint#delete ComputeNetworkEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

