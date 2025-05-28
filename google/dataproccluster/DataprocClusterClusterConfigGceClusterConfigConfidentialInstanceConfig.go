// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigGceClusterConfigConfidentialInstanceConfig struct {
	// Defines whether the instance should have confidential compute enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dataproc_cluster#enable_confidential_compute DataprocCluster#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

