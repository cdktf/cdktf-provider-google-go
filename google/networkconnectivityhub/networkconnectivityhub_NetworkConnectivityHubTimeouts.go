package networkconnectivityhub


type NetworkConnectivityHubTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_hub#create NetworkConnectivityHub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_hub#delete NetworkConnectivityHub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_hub#update NetworkConnectivityHub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

