package dnsrecordset


type DnsRecordSetRoutingPolicyWrr struct {
	// The ratio of traffic routed to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_record_set#weight DnsRecordSet#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// health_checked_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_record_set#health_checked_targets DnsRecordSet#health_checked_targets}
	HealthCheckedTargets *DnsRecordSetRoutingPolicyWrrHealthCheckedTargets `field:"optional" json:"healthCheckedTargets" yaml:"healthCheckedTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_record_set#rrdatas DnsRecordSet#rrdatas}.
	Rrdatas *[]*string `field:"optional" json:"rrdatas" yaml:"rrdatas"`
}

