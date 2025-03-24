// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterResourceUsageExportConfigBigqueryDestination struct {
	// The ID of a BigQuery Dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/container_cluster#dataset_id ContainerCluster#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

