// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecontainerClusterConfig struct {
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
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#authorization EdgecontainerCluster#authorization}
	Authorization *EdgecontainerClusterAuthorization `field:"required" json:"authorization" yaml:"authorization"`
	// fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#fleet EdgecontainerCluster#fleet}
	Fleet *EdgecontainerClusterFleet `field:"required" json:"fleet" yaml:"fleet"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#location EdgecontainerCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The GDCE cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#name EdgecontainerCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// networking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#networking EdgecontainerCluster#networking}
	Networking *EdgecontainerClusterNetworking `field:"required" json:"networking" yaml:"networking"`
	// control_plane block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#control_plane EdgecontainerCluster#control_plane}
	ControlPlane *EdgecontainerClusterControlPlane `field:"optional" json:"controlPlane" yaml:"controlPlane"`
	// control_plane_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#control_plane_encryption EdgecontainerCluster#control_plane_encryption}
	ControlPlaneEncryption *EdgecontainerClusterControlPlaneEncryption `field:"optional" json:"controlPlaneEncryption" yaml:"controlPlaneEncryption"`
	// The default maximum number of pods per node used if a maximum value is not specified explicitly for a node pool in this cluster.
	//
	// If unspecified, the
	// Kubernetes default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#default_max_pods_per_node EdgecontainerCluster#default_max_pods_per_node}
	DefaultMaxPodsPerNode *float64 `field:"optional" json:"defaultMaxPodsPerNode" yaml:"defaultMaxPodsPerNode"`
	// Address pools for cluster data plane external load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#external_load_balancer_ipv4_address_pools EdgecontainerCluster#external_load_balancer_ipv4_address_pools}
	ExternalLoadBalancerIpv4AddressPools *[]*string `field:"optional" json:"externalLoadBalancerIpv4AddressPools" yaml:"externalLoadBalancerIpv4AddressPools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#id EdgecontainerCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the edgecloud cluster.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#labels EdgecontainerCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#maintenance_policy EdgecontainerCluster#maintenance_policy}
	MaintenancePolicy *EdgecontainerClusterMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#project EdgecontainerCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The release channel a cluster is subscribed to. Possible values: ["RELEASE_CHANNEL_UNSPECIFIED", "NONE", "REGULAR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#release_channel EdgecontainerCluster#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// system_addons_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#system_addons_config EdgecontainerCluster#system_addons_config}
	SystemAddonsConfig *EdgecontainerClusterSystemAddonsConfig `field:"optional" json:"systemAddonsConfig" yaml:"systemAddonsConfig"`
	// The target cluster version. For example: "1.5.0".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#target_version EdgecontainerCluster#target_version}
	TargetVersion *string `field:"optional" json:"targetVersion" yaml:"targetVersion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/edgecontainer_cluster#timeouts EdgecontainerCluster#timeouts}
	Timeouts *EdgecontainerClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

