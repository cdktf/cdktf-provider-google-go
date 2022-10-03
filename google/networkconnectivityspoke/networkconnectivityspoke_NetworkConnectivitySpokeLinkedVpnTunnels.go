package networkconnectivityspoke


type NetworkConnectivitySpokeLinkedVpnTunnels struct {
	// A value that controls whether site-to-site data transfer is enabled for these resources.
	//
	// Note that data transfer is available only in supported locations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_spoke#site_to_site_data_transfer NetworkConnectivitySpoke#site_to_site_data_transfer}
	SiteToSiteDataTransfer interface{} `field:"required" json:"siteToSiteDataTransfer" yaml:"siteToSiteDataTransfer"`
	// The URIs of linked VPN tunnel resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_connectivity_spoke#uris NetworkConnectivitySpoke#uris}
	Uris *[]*string `field:"required" json:"uris" yaml:"uris"`
}

