// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigAuxiliaryNodeGroupsNodeGroupNodeGroupConfig struct {
	// accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/dataproc_cluster#accelerators DataprocCluster#accelerators}
	Accelerators interface{} `field:"optional" json:"accelerators" yaml:"accelerators"`
	// disk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/dataproc_cluster#disk_config DataprocCluster#disk_config}
	DiskConfig *DataprocClusterClusterConfigAuxiliaryNodeGroupsNodeGroupNodeGroupConfigDiskConfig `field:"optional" json:"diskConfig" yaml:"diskConfig"`
	// The name of a Google Compute Engine machine type to create for the master.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/dataproc_cluster#machine_type DataprocCluster#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The name of a minimum generation of CPU family for the auxiliary node group.
	//
	// If not specified, GCP will default to a predetermined computed value for each zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/dataproc_cluster#min_cpu_platform DataprocCluster#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Specifies the number of auxiliary nodes to create. If not specified, GCP will default to a predetermined computed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/dataproc_cluster#num_instances DataprocCluster#num_instances}
	NumInstances *float64 `field:"optional" json:"numInstances" yaml:"numInstances"`
}

