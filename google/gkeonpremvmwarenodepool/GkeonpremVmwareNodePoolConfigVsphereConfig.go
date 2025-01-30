// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarenodepool


type GkeonpremVmwareNodePoolConfigVsphereConfig struct {
	// The name of the vCenter datastore. Inherited from the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/gkeonprem_vmware_node_pool#datastore GkeonpremVmwareNodePool#datastore}
	Datastore *string `field:"optional" json:"datastore" yaml:"datastore"`
	// Vsphere host groups to apply to all VMs in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/gkeonprem_vmware_node_pool#host_groups GkeonpremVmwareNodePool#host_groups}
	HostGroups *[]*string `field:"optional" json:"hostGroups" yaml:"hostGroups"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/gkeonprem_vmware_node_pool#tags GkeonpremVmwareNodePool#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

