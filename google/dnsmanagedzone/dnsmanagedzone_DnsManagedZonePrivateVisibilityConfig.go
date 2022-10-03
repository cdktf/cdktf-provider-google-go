package dnsmanagedzone


type DnsManagedZonePrivateVisibilityConfig struct {
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_managed_zone#networks DnsManagedZone#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
}

