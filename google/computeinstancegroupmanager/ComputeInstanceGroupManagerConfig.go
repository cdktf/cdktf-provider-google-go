package computeinstancegroupmanager

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupManagerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#base_instance_name ComputeInstanceGroupManager#base_instance_name}
	BaseInstanceName *string `field:"required" json:"baseInstanceName" yaml:"baseInstanceName"`
	// The name of the instance group manager.
	//
	// Must be 1-63 characters long and comply with RFC1035. Supported characters include lowercase letters, numbers, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#name ComputeInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#version ComputeInstanceGroupManager#version}
	Version interface{} `field:"required" json:"version" yaml:"version"`
	// auto_healing_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#auto_healing_policies ComputeInstanceGroupManager#auto_healing_policies}
	AutoHealingPolicies *ComputeInstanceGroupManagerAutoHealingPolicies `field:"optional" json:"autoHealingPolicies" yaml:"autoHealingPolicies"`
	// An optional textual description of the instance group manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#description ComputeInstanceGroupManager#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#id ComputeInstanceGroupManager#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Pagination behavior of the listManagedInstances API method for this managed instance group.
	//
	// Valid values are: "PAGELESS", "PAGINATED". If PAGELESS (default), Pagination is disabled for the group's listManagedInstances API method. maxResults and pageToken query parameters are ignored and all instances are returned in a single response. If PAGINATED, pagination is enabled, maxResults and pageToken query parameters are respected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#list_managed_instances_results ComputeInstanceGroupManager#list_managed_instances_results}
	ListManagedInstancesResults *string `field:"optional" json:"listManagedInstancesResults" yaml:"listManagedInstancesResults"`
	// named_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#named_port ComputeInstanceGroupManager#named_port}
	NamedPort interface{} `field:"optional" json:"namedPort" yaml:"namedPort"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#project ComputeInstanceGroupManager#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// stateful_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#stateful_disk ComputeInstanceGroupManager#stateful_disk}
	StatefulDisk interface{} `field:"optional" json:"statefulDisk" yaml:"statefulDisk"`
	// The full URL of all target pools to which new instances in the group are added.
	//
	// Updating the target pools attribute does not affect existing instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#target_pools ComputeInstanceGroupManager#target_pools}
	TargetPools *[]*string `field:"optional" json:"targetPools" yaml:"targetPools"`
	// The target number of running instances for this managed instance group.
	//
	// This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#target_size ComputeInstanceGroupManager#target_size}
	TargetSize *float64 `field:"optional" json:"targetSize" yaml:"targetSize"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#timeouts ComputeInstanceGroupManager#timeouts}
	Timeouts *ComputeInstanceGroupManagerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#update_policy ComputeInstanceGroupManager#update_policy}
	UpdatePolicy *ComputeInstanceGroupManagerUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Whether to wait for all instances to be created/updated before returning.
	//
	// Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#wait_for_instances ComputeInstanceGroupManager#wait_for_instances}
	WaitForInstances interface{} `field:"optional" json:"waitForInstances" yaml:"waitForInstances"`
	// When used with wait_for_instances specifies the status to wait for.
	//
	// When STABLE is specified this resource will wait until the instances are stable before returning. When UPDATED is set, it will wait for the version target to be reached and any per instance configs to be effective as well as all instances to be stable before returning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#wait_for_instances_status ComputeInstanceGroupManager#wait_for_instances_status}
	WaitForInstancesStatus *string `field:"optional" json:"waitForInstancesStatus" yaml:"waitForInstancesStatus"`
	// The zone that instances in this group should be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_group_manager#zone ComputeInstanceGroupManager#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

