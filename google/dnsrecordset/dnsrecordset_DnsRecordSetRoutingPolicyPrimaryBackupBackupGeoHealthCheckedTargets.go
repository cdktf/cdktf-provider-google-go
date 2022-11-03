package dnsrecordset


type DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets struct {
	// internal_load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_record_set#internal_load_balancers DnsRecordSet#internal_load_balancers}
	InternalLoadBalancers interface{} `field:"required" json:"internalLoadBalancers" yaml:"internalLoadBalancers"`
}

