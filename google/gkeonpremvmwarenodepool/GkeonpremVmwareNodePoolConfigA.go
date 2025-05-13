// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarenodepool


type GkeonpremVmwareNodePoolConfigA struct {
	// The OS image to be used for each node in a node pool.
	//
	// Currently 'cos', 'cos_cgv2', 'ubuntu', 'ubuntu_cgv2', 'ubuntu_containerd' and 'windows' are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#image_type GkeonpremVmwareNodePool#image_type}
	ImageType *string `field:"required" json:"imageType" yaml:"imageType"`
	// VMware disk size to be used during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#boot_disk_size_gb GkeonpremVmwareNodePool#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// The number of CPUs for each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#cpus GkeonpremVmwareNodePool#cpus}
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// Allow node pool traffic to be load balanced. Only works for clusters with MetalLB load balancers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#enable_load_balancer GkeonpremVmwareNodePool#enable_load_balancer}
	EnableLoadBalancer interface{} `field:"optional" json:"enableLoadBalancer" yaml:"enableLoadBalancer"`
	// The OS image name in vCenter, only valid when using Windows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#image GkeonpremVmwareNodePool#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s) that
	// Kubernetes may apply to the node.
	// In case of conflict in label keys, the applied set may differ depending on
	// the Kubernetes version -- it's best to assume the behavior is undefined
	// and conflicts should be avoided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#labels GkeonpremVmwareNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The megabytes of memory for each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#memory_mb GkeonpremVmwareNodePool#memory_mb}
	MemoryMb *float64 `field:"optional" json:"memoryMb" yaml:"memoryMb"`
	// The number of nodes in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#replicas GkeonpremVmwareNodePool#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#taints GkeonpremVmwareNodePool#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// vsphere_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/gkeonprem_vmware_node_pool#vsphere_config GkeonpremVmwareNodePool#vsphere_config}
	VsphereConfig *GkeonpremVmwareNodePoolConfigVsphereConfig `field:"optional" json:"vsphereConfig" yaml:"vsphereConfig"`
}

