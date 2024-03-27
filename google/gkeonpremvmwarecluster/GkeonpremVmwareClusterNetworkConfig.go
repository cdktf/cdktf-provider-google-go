// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfig struct {
	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// Only a single range is supported. This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/gkeonprem_vmware_cluster#pod_address_cidr_blocks GkeonpremVmwareCluster#pod_address_cidr_blocks}
	PodAddressCidrBlocks *[]*string `field:"required" json:"podAddressCidrBlocks" yaml:"podAddressCidrBlocks"`
	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// Only a single range is supported.. This field
	// cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/gkeonprem_vmware_cluster#service_address_cidr_blocks GkeonpremVmwareCluster#service_address_cidr_blocks}
	ServiceAddressCidrBlocks *[]*string `field:"required" json:"serviceAddressCidrBlocks" yaml:"serviceAddressCidrBlocks"`
	// control_plane_v2_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/gkeonprem_vmware_cluster#control_plane_v2_config GkeonpremVmwareCluster#control_plane_v2_config}
	ControlPlaneV2Config *GkeonpremVmwareClusterNetworkConfigControlPlaneV2Config `field:"optional" json:"controlPlaneV2Config" yaml:"controlPlaneV2Config"`
	// dhcp_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/gkeonprem_vmware_cluster#dhcp_ip_config GkeonpremVmwareCluster#dhcp_ip_config}
	DhcpIpConfig *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig `field:"optional" json:"dhcpIpConfig" yaml:"dhcpIpConfig"`
	// host_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/gkeonprem_vmware_cluster#host_config GkeonpremVmwareCluster#host_config}
	HostConfig *GkeonpremVmwareClusterNetworkConfigHostConfig `field:"optional" json:"hostConfig" yaml:"hostConfig"`
	// static_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/gkeonprem_vmware_cluster#static_ip_config GkeonpremVmwareCluster#static_ip_config}
	StaticIpConfig *GkeonpremVmwareClusterNetworkConfigStaticIpConfig `field:"optional" json:"staticIpConfig" yaml:"staticIpConfig"`
	// vcenter_network specifies vCenter network name. Inherited from the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/gkeonprem_vmware_cluster#vcenter_network GkeonpremVmwareCluster#vcenter_network}
	VcenterNetwork *string `field:"optional" json:"vcenterNetwork" yaml:"vcenterNetwork"`
}

