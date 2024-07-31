// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocjob


type DataprocJobPlacement struct {
	// The name of the cluster where the job will be submitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/dataproc_job#cluster_name DataprocJob#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

