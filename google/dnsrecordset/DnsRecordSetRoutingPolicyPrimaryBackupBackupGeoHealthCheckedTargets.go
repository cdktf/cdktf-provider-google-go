// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsrecordset


type DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets struct {
	// internal_load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/dns_record_set#internal_load_balancers DnsRecordSet#internal_load_balancers}
	InternalLoadBalancers interface{} `field:"required" json:"internalLoadBalancers" yaml:"internalLoadBalancers"`
}

