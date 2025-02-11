// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigPreemptibleWorkerConfig struct {
	// disk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dataproc_cluster#disk_config DataprocCluster#disk_config}
	DiskConfig *DataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig `field:"optional" json:"diskConfig" yaml:"diskConfig"`
	// instance_flexibility_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dataproc_cluster#instance_flexibility_policy DataprocCluster#instance_flexibility_policy}
	InstanceFlexibilityPolicy *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy `field:"optional" json:"instanceFlexibilityPolicy" yaml:"instanceFlexibilityPolicy"`
	// Specifies the number of preemptible nodes to create. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dataproc_cluster#num_instances DataprocCluster#num_instances}
	NumInstances *float64 `field:"optional" json:"numInstances" yaml:"numInstances"`
	// Specifies the preemptibility of the secondary nodes. Defaults to PREEMPTIBLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dataproc_cluster#preemptibility DataprocCluster#preemptibility}
	Preemptibility *string `field:"optional" json:"preemptibility" yaml:"preemptibility"`
}

