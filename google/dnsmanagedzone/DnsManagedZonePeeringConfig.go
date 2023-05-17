package dnsmanagedzone


type DnsManagedZonePeeringConfig struct {
	// target_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dns_managed_zone#target_network DnsManagedZone#target_network}
	TargetNetwork *DnsManagedZonePeeringConfigTargetNetwork `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
}

