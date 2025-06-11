// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremVmwareClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The admin cluster this VMware User Cluster belongs to.
	//
	// This is the full resource name of the admin cluster's hub membership.
	// In the future, references to other resource types might be allowed if
	// admin clusters are modeled as their own resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#admin_cluster_membership GkeonpremVmwareCluster#admin_cluster_membership}
	AdminClusterMembership *string `field:"required" json:"adminClusterMembership" yaml:"adminClusterMembership"`
	// control_plane_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#control_plane_node GkeonpremVmwareCluster#control_plane_node}
	ControlPlaneNode *GkeonpremVmwareClusterControlPlaneNode `field:"required" json:"controlPlaneNode" yaml:"controlPlaneNode"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#location GkeonpremVmwareCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The VMware cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#name GkeonpremVmwareCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Anthos clusters on the VMware version for your user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#on_prem_version GkeonpremVmwareCluster#on_prem_version}
	OnPremVersion *string `field:"required" json:"onPremVersion" yaml:"onPremVersion"`
	// Annotations on the VMware User Cluster.
	//
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#annotations GkeonpremVmwareCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// anti_affinity_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#anti_affinity_groups GkeonpremVmwareCluster#anti_affinity_groups}
	AntiAffinityGroups *GkeonpremVmwareClusterAntiAffinityGroups `field:"optional" json:"antiAffinityGroups" yaml:"antiAffinityGroups"`
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#authorization GkeonpremVmwareCluster#authorization}
	Authorization *GkeonpremVmwareClusterAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// auto_repair_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#auto_repair_config GkeonpremVmwareCluster#auto_repair_config}
	AutoRepairConfig *GkeonpremVmwareClusterAutoRepairConfig `field:"optional" json:"autoRepairConfig" yaml:"autoRepairConfig"`
	// dataplane_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#dataplane_v2 GkeonpremVmwareCluster#dataplane_v2}
	DataplaneV2 *GkeonpremVmwareClusterDataplaneV2 `field:"optional" json:"dataplaneV2" yaml:"dataplaneV2"`
	// A human readable description of this VMware User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#description GkeonpremVmwareCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disable bundled ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#disable_bundled_ingress GkeonpremVmwareCluster#disable_bundled_ingress}
	DisableBundledIngress interface{} `field:"optional" json:"disableBundledIngress" yaml:"disableBundledIngress"`
	// Enable advanced cluster. Default to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#enable_advanced_cluster GkeonpremVmwareCluster#enable_advanced_cluster}
	EnableAdvancedCluster interface{} `field:"optional" json:"enableAdvancedCluster" yaml:"enableAdvancedCluster"`
	// Enable control plane V2. Default to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#enable_control_plane_v2 GkeonpremVmwareCluster#enable_control_plane_v2}
	EnableControlPlaneV2 interface{} `field:"optional" json:"enableControlPlaneV2" yaml:"enableControlPlaneV2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#id GkeonpremVmwareCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#load_balancer GkeonpremVmwareCluster#load_balancer}
	LoadBalancer *GkeonpremVmwareClusterLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#network_config GkeonpremVmwareCluster#network_config}
	NetworkConfig *GkeonpremVmwareClusterNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#project GkeonpremVmwareCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#storage GkeonpremVmwareCluster#storage}
	Storage *GkeonpremVmwareClusterStorage `field:"optional" json:"storage" yaml:"storage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#timeouts GkeonpremVmwareCluster#timeouts}
	Timeouts *GkeonpremVmwareClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#upgrade_policy GkeonpremVmwareCluster#upgrade_policy}
	UpgradePolicy *GkeonpremVmwareClusterUpgradePolicy `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// vcenter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#vcenter GkeonpremVmwareCluster#vcenter}
	Vcenter *GkeonpremVmwareClusterVcenter `field:"optional" json:"vcenter" yaml:"vcenter"`
	// Enable VM tracking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/gkeonprem_vmware_cluster#vm_tracking_enabled GkeonpremVmwareCluster#vm_tracking_enabled}
	VmTrackingEnabled interface{} `field:"optional" json:"vmTrackingEnabled" yaml:"vmTrackingEnabled"`
}

