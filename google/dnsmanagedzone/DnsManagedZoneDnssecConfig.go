package dnsmanagedzone


type DnsManagedZoneDnssecConfig struct {
	// default_key_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone#default_key_specs DnsManagedZone#default_key_specs}
	DefaultKeySpecs interface{} `field:"optional" json:"defaultKeySpecs" yaml:"defaultKeySpecs"`
	// Identifies what kind of resource this is.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone#kind DnsManagedZone#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Specifies the mechanism used to provide authenticated denial-of-existence responses.
	//
	// non_existence can only be updated when the state is 'off'. Possible values: ["nsec", "nsec3"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone#non_existence DnsManagedZone#non_existence}
	NonExistence *string `field:"optional" json:"nonExistence" yaml:"nonExistence"`
	// Specifies whether DNSSEC is enabled, and what mode it is in Possible values: ["off", "on", "transfer"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone#state DnsManagedZone#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

