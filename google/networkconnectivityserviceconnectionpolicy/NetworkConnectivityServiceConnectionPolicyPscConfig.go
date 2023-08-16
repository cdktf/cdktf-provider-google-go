package networkconnectivityserviceconnectionpolicy


type NetworkConnectivityServiceConnectionPolicyPscConfig struct {
	// IDs of the subnetworks or fully qualified identifiers for the subnetworks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#subnetworks NetworkConnectivityServiceConnectionPolicy#subnetworks}
	Subnetworks *[]*string `field:"required" json:"subnetworks" yaml:"subnetworks"`
	// Max number of PSC connections for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#limit NetworkConnectivityServiceConnectionPolicy#limit}
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
}

