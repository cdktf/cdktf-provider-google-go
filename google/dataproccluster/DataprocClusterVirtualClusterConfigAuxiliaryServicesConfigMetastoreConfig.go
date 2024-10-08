// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig struct {
	// The Hive Metastore configuration for this workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/dataproc_cluster#dataproc_metastore_service DataprocCluster#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"optional" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

