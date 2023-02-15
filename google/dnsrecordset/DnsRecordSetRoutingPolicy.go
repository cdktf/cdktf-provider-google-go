package dnsrecordset


type DnsRecordSetRoutingPolicy struct {
	// Specifies whether to enable fencing for geo queries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#enable_geo_fencing DnsRecordSet#enable_geo_fencing}
	EnableGeoFencing interface{} `field:"optional" json:"enableGeoFencing" yaml:"enableGeoFencing"`
	// geo block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#geo DnsRecordSet#geo}
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// primary_backup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#primary_backup DnsRecordSet#primary_backup}
	PrimaryBackup *DnsRecordSetRoutingPolicyPrimaryBackup `field:"optional" json:"primaryBackup" yaml:"primaryBackup"`
	// wrr block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#wrr DnsRecordSet#wrr}
	Wrr interface{} `field:"optional" json:"wrr" yaml:"wrr"`
}

