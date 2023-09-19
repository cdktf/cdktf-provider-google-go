// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterRestoreContinuousBackupSource struct {
	// The name of the source cluster that this cluster is restored from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/alloydb_cluster#cluster AlloydbCluster#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The point in time that this cluster is restored to, in RFC 3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/alloydb_cluster#point_in_time AlloydbCluster#point_in_time}
	PointInTime *string `field:"required" json:"pointInTime" yaml:"pointInTime"`
}

