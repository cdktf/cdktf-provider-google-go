package networkconnectivityspoke


type NetworkConnectivitySpokeLinkedRouterApplianceInstances struct {
	// instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#instances NetworkConnectivitySpoke#instances}
	Instances interface{} `field:"required" json:"instances" yaml:"instances"`
	// A value that controls whether site-to-site data transfer is enabled for these resources.
	//
	// Note that data transfer is available only in supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#site_to_site_data_transfer NetworkConnectivitySpoke#site_to_site_data_transfer}
	SiteToSiteDataTransfer interface{} `field:"required" json:"siteToSiteDataTransfer" yaml:"siteToSiteDataTransfer"`
}

