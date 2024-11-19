// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterVcenter struct {
	// Contains the vCenter CA certificate public key for SSL verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_vmware_cluster#ca_cert_data GkeonpremVmwareCluster#ca_cert_data}
	CaCertData *string `field:"optional" json:"caCertData" yaml:"caCertData"`
	// The name of the vCenter cluster for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_vmware_cluster#cluster GkeonpremVmwareCluster#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// The name of the vCenter datacenter for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_vmware_cluster#datacenter GkeonpremVmwareCluster#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// The name of the vCenter datastore for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_vmware_cluster#datastore GkeonpremVmwareCluster#datastore}
	Datastore *string `field:"optional" json:"datastore" yaml:"datastore"`
	// The name of the vCenter folder for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_vmware_cluster#folder GkeonpremVmwareCluster#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// The name of the vCenter resource pool for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_vmware_cluster#resource_pool GkeonpremVmwareCluster#resource_pool}
	ResourcePool *string `field:"optional" json:"resourcePool" yaml:"resourcePool"`
	// The name of the vCenter storage policy for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/gkeonprem_vmware_cluster#storage_policy_name GkeonpremVmwareCluster#storage_policy_name}
	StoragePolicyName *string `field:"optional" json:"storagePolicyName" yaml:"storagePolicyName"`
}

