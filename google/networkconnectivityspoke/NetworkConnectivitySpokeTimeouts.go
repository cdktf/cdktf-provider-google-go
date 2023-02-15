package networkconnectivityspoke


type NetworkConnectivitySpokeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_spoke#create NetworkConnectivitySpoke#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_spoke#delete NetworkConnectivitySpoke#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_spoke#update NetworkConnectivitySpoke#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

