package containernodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerNodePoolConfig struct {
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
	// The cluster to create the node pool for. Cluster must be present in location provided for zonal clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#cluster ContainerNodePool#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#autoscaling ContainerNodePool#autoscaling}
	Autoscaling *ContainerNodePoolAutoscaling `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#id ContainerNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The initial number of nodes for the pool.
	//
	// In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#initial_node_count ContainerNodePool#initial_node_count}
	InitialNodeCount *float64 `field:"optional" json:"initialNodeCount" yaml:"initialNodeCount"`
	// The location (region or zone) of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#location ContainerNodePool#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#management ContainerNodePool#management}
	Management *ContainerNodePoolManagement `field:"optional" json:"management" yaml:"management"`
	// The maximum number of pods per node in this node pool.
	//
	// Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#max_pods_per_node ContainerNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
	// The name of the node pool. If left blank, Terraform will auto-generate a unique name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#name ContainerNodePool#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#name_prefix ContainerNodePool#name_prefix}
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#network_config ContainerNodePool#network_config}
	NetworkConfig *ContainerNodePoolNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#node_config ContainerNodePool#node_config}
	NodeConfig *ContainerNodePoolNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// The number of nodes per instance group.
	//
	// This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#node_count ContainerNodePool#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// The list of zones in which the node pool's nodes should be located.
	//
	// Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#node_locations ContainerNodePool#node_locations}
	NodeLocations *[]*string `field:"optional" json:"nodeLocations" yaml:"nodeLocations"`
	// placement_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#placement_policy ContainerNodePool#placement_policy}
	PlacementPolicy *ContainerNodePoolPlacementPolicy `field:"optional" json:"placementPolicy" yaml:"placementPolicy"`
	// The ID of the project in which to create the node pool.
	//
	// If blank, the provider-configured project will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#project ContainerNodePool#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#timeouts ContainerNodePool#timeouts}
	Timeouts *ContainerNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#upgrade_settings ContainerNodePool#upgrade_settings}
	UpgradeSettings *ContainerNodePoolUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
	// The Kubernetes version for the nodes in this pool.
	//
	// Note that if this field and auto_upgrade are both specified, they will fight each other for what the node version should be, so setting both is highly discouraged. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_node_pool#version ContainerNodePool#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

