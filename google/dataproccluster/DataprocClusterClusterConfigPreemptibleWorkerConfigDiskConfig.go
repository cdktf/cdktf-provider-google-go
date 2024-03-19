// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig struct {
	// Size of the primary disk attached to each preemptible worker node, specified in GB.
	//
	// The smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_cluster#boot_disk_size_gb DataprocCluster#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// The disk type of the primary disk attached to each preemptible worker node.
	//
	// Such as "pd-ssd" or "pd-standard". Defaults to "pd-standard".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_cluster#boot_disk_type DataprocCluster#boot_disk_type}
	BootDiskType *string `field:"optional" json:"bootDiskType" yaml:"bootDiskType"`
	// The amount of local SSD disks that will be attached to each preemptible worker node. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_cluster#num_local_ssds DataprocCluster#num_local_ssds}
	NumLocalSsds *float64 `field:"optional" json:"numLocalSsds" yaml:"numLocalSsds"`
}

