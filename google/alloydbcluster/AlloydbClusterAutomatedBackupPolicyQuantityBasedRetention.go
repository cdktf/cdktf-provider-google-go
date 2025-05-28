// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterAutomatedBackupPolicyQuantityBasedRetention struct {
	// The number of backups to retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/alloydb_cluster#count AlloydbCluster#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

