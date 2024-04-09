// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigWorkerConfig struct {
	// accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/dataproc_cluster#accelerators DataprocCluster#accelerators}
	Accelerators interface{} `field:"optional" json:"accelerators" yaml:"accelerators"`
	// disk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/dataproc_cluster#disk_config DataprocCluster#disk_config}
	DiskConfig *DataprocClusterClusterConfigWorkerConfigDiskConfig `field:"optional" json:"diskConfig" yaml:"diskConfig"`
	// The URI for the image to use for this master/worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/dataproc_cluster#image_uri DataprocCluster#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// The name of a Google Compute Engine machine type to create for the master/worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/dataproc_cluster#machine_type DataprocCluster#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The name of a minimum generation of CPU family for the master/worker.
	//
	// If not specified, GCP will default to a predetermined computed value for each zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/dataproc_cluster#min_cpu_platform DataprocCluster#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// The minimum number of primary worker instances to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/dataproc_cluster#min_num_instances DataprocCluster#min_num_instances}
	MinNumInstances *float64 `field:"optional" json:"minNumInstances" yaml:"minNumInstances"`
	// Specifies the number of worker nodes to create. If not specified, GCP will default to a predetermined computed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/dataproc_cluster#num_instances DataprocCluster#num_instances}
	NumInstances *float64 `field:"optional" json:"numInstances" yaml:"numInstances"`
}

