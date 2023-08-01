package dnsmanagedzone


type DnsManagedZoneDnssecConfigDefaultKeySpecs struct {
	// String mnemonic specifying the DNSSEC algorithm of this key Possible values: ["ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_managed_zone#algorithm DnsManagedZone#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Length of the keys in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_managed_zone#key_length DnsManagedZone#key_length}
	KeyLength *float64 `field:"optional" json:"keyLength" yaml:"keyLength"`
	// Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK).
	//
	// Key signing keys have the Secure Entry
	// Point flag set and, when active, will only be used to sign
	// resource record sets of type DNSKEY. Zone signing keys do
	// not have the Secure Entry Point flag set and will be used
	// to sign all other types of resource record sets. Possible values: ["keySigning", "zoneSigning"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_managed_zone#key_type DnsManagedZone#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Identifies what kind of resource this is.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_managed_zone#kind DnsManagedZone#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

