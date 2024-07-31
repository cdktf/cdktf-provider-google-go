// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterMaintenanceUpdatePolicyMaintenanceWindowsStartTime struct {
	// Hours of day in 24 hour format. Should be from 0 to 23.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/alloydb_cluster#hours AlloydbCluster#hours}
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Minutes of hour of day. Currently, only the value 0 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/alloydb_cluster#minutes AlloydbCluster#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds. Currently, only the value 0 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/alloydb_cluster#nanos AlloydbCluster#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time. Currently, only the value 0 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/alloydb_cluster#seconds AlloydbCluster#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

