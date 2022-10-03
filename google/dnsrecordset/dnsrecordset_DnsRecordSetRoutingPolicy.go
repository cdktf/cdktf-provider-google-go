package dnsrecordset


type DnsRecordSetRoutingPolicy struct {
	// geo block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#geo DnsRecordSet#geo}
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// wrr block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#wrr DnsRecordSet#wrr}
	Wrr interface{} `field:"optional" json:"wrr" yaml:"wrr"`
}

