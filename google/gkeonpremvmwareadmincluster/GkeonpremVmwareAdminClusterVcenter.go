// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterVcenter struct {
	// The vCenter IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#address GkeonpremVmwareAdminCluster#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Contains the vCenter CA certificate public key for SSL verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#ca_cert_data GkeonpremVmwareAdminCluster#ca_cert_data}
	CaCertData *string `field:"optional" json:"caCertData" yaml:"caCertData"`
	// The name of the vCenter cluster for the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#cluster GkeonpremVmwareAdminCluster#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// The name of the vCenter datacenter for the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#datacenter GkeonpremVmwareAdminCluster#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// The name of the virtual machine disk (VMDK) for the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#data_disk GkeonpremVmwareAdminCluster#data_disk}
	DataDisk *string `field:"optional" json:"dataDisk" yaml:"dataDisk"`
	// The name of the vCenter datastore for the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#datastore GkeonpremVmwareAdminCluster#datastore}
	Datastore *string `field:"optional" json:"datastore" yaml:"datastore"`
	// The name of the vCenter folder for the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#folder GkeonpremVmwareAdminCluster#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// The name of the vCenter resource pool for the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#resource_pool GkeonpremVmwareAdminCluster#resource_pool}
	ResourcePool *string `field:"optional" json:"resourcePool" yaml:"resourcePool"`
	// The name of the vCenter storage policy for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gkeonprem_vmware_admin_cluster#storage_policy_name GkeonpremVmwareAdminCluster#storage_policy_name}
	StoragePolicyName *string `field:"optional" json:"storagePolicyName" yaml:"storagePolicyName"`
}

