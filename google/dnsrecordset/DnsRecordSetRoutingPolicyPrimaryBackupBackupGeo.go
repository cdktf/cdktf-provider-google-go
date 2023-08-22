package dnsrecordset


type DnsRecordSetRoutingPolicyPrimaryBackupBackupGeo struct {
	// The location name defined in Google Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_record_set#location DnsRecordSet#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// health_checked_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_record_set#health_checked_targets DnsRecordSet#health_checked_targets}
	HealthCheckedTargets *DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets `field:"optional" json:"healthCheckedTargets" yaml:"healthCheckedTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_record_set#rrdatas DnsRecordSet#rrdatas}.
	Rrdatas *[]*string `field:"optional" json:"rrdatas" yaml:"rrdatas"`
}

