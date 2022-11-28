package dnsmanagedzone


type DnsManagedZonePrivateVisibilityConfig struct {
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_managed_zone#networks DnsManagedZone#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// gke_clusters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_managed_zone#gke_clusters DnsManagedZone#gke_clusters}
	GkeClusters interface{} `field:"optional" json:"gkeClusters" yaml:"gkeClusters"`
}

