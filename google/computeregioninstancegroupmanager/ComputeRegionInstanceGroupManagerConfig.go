package computeregioninstancegroupmanager

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionInstanceGroupManagerConfig struct {
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
	// The base instance name to use for instances in this group.
	//
	// The value must be a valid RFC1035 name. Supported characters are lowercase letters, numbers, and hyphens (-). Instances are named by appending a hyphen and a random four-character string to the base instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#base_instance_name ComputeRegionInstanceGroupManager#base_instance_name}
	BaseInstanceName *string `field:"required" json:"baseInstanceName" yaml:"baseInstanceName"`
	// The name of the instance group manager.
	//
	// Must be 1-63 characters long and comply with RFC1035. Supported characters include lowercase letters, numbers, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#name ComputeRegionInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#version ComputeRegionInstanceGroupManager#version}
	Version interface{} `field:"required" json:"version" yaml:"version"`
	// auto_healing_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#auto_healing_policies ComputeRegionInstanceGroupManager#auto_healing_policies}
	AutoHealingPolicies *ComputeRegionInstanceGroupManagerAutoHealingPolicies `field:"optional" json:"autoHealingPolicies" yaml:"autoHealingPolicies"`
	// An optional textual description of the instance group manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#description ComputeRegionInstanceGroupManager#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The shape to which the group converges either proactively or on resize events (depending on the value set in updatePolicy.instanceRedistributionType).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#distribution_policy_target_shape ComputeRegionInstanceGroupManager#distribution_policy_target_shape}
	DistributionPolicyTargetShape *string `field:"optional" json:"distributionPolicyTargetShape" yaml:"distributionPolicyTargetShape"`
	// The distribution policy for this managed instance group. You can specify one or more values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#distribution_policy_zones ComputeRegionInstanceGroupManager#distribution_policy_zones}
	DistributionPolicyZones *[]*string `field:"optional" json:"distributionPolicyZones" yaml:"distributionPolicyZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#id ComputeRegionInstanceGroupManager#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Pagination behavior of the listManagedInstances API method for this managed instance group.
	//
	// Valid values are: "PAGELESS", "PAGINATED". If PAGELESS (default), Pagination is disabled for the group's listManagedInstances API method. maxResults and pageToken query parameters are ignored and all instances are returned in a single response. If PAGINATED, pagination is enabled, maxResults and pageToken query parameters are respected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#list_managed_instances_results ComputeRegionInstanceGroupManager#list_managed_instances_results}
	ListManagedInstancesResults *string `field:"optional" json:"listManagedInstancesResults" yaml:"listManagedInstancesResults"`
	// named_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#named_port ComputeRegionInstanceGroupManager#named_port}
	NamedPort interface{} `field:"optional" json:"namedPort" yaml:"namedPort"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#project ComputeRegionInstanceGroupManager#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region where the managed instance group resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#region ComputeRegionInstanceGroupManager#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// stateful_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#stateful_disk ComputeRegionInstanceGroupManager#stateful_disk}
	StatefulDisk interface{} `field:"optional" json:"statefulDisk" yaml:"statefulDisk"`
	// The full URL of all target pools to which new instances in the group are added.
	//
	// Updating the target pools attribute does not affect existing instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#target_pools ComputeRegionInstanceGroupManager#target_pools}
	TargetPools *[]*string `field:"optional" json:"targetPools" yaml:"targetPools"`
	// The target number of running instances for this managed instance group.
	//
	// This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#target_size ComputeRegionInstanceGroupManager#target_size}
	TargetSize *float64 `field:"optional" json:"targetSize" yaml:"targetSize"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#timeouts ComputeRegionInstanceGroupManager#timeouts}
	Timeouts *ComputeRegionInstanceGroupManagerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#update_policy ComputeRegionInstanceGroupManager#update_policy}
	UpdatePolicy *ComputeRegionInstanceGroupManagerUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Whether to wait for all instances to be created/updated before returning.
	//
	// Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#wait_for_instances ComputeRegionInstanceGroupManager#wait_for_instances}
	WaitForInstances interface{} `field:"optional" json:"waitForInstances" yaml:"waitForInstances"`
	// When used with wait_for_instances specifies the status to wait for.
	//
	// When STABLE is specified this resource will wait until the instances are stable before returning. When UPDATED is set, it will wait for the version target to be reached and any per instance configs to be effective as well as all instances to be stable before returning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_instance_group_manager#wait_for_instances_status ComputeRegionInstanceGroupManager#wait_for_instances_status}
	WaitForInstancesStatus *string `field:"optional" json:"waitForInstancesStatus" yaml:"waitForInstancesStatus"`
}

