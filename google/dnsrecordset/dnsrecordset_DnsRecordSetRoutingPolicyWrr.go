package dnsrecordset


type DnsRecordSetRoutingPolicyWrr struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#rrdatas DnsRecordSet#rrdatas}.
	Rrdatas *[]*string `field:"required" json:"rrdatas" yaml:"rrdatas"`
	// The ratio of traffic routed to the target.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#weight DnsRecordSet#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

