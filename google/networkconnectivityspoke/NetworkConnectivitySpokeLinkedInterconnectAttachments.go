package networkconnectivityspoke


type NetworkConnectivitySpokeLinkedInterconnectAttachments struct {
	// A value that controls whether site-to-site data transfer is enabled for these resources.
	//
	// Note that data transfer is available only in supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_connectivity_spoke#site_to_site_data_transfer NetworkConnectivitySpoke#site_to_site_data_transfer}
	SiteToSiteDataTransfer interface{} `field:"required" json:"siteToSiteDataTransfer" yaml:"siteToSiteDataTransfer"`
	// The URIs of linked interconnect attachment resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_connectivity_spoke#uris NetworkConnectivitySpoke#uris}
	Uris *[]*string `field:"required" json:"uris" yaml:"uris"`
}

