package computeregionperinstanceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionPerInstanceConfigConfig struct {
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
	// The name for this per-instance config and its corresponding instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#name ComputeRegionPerInstanceConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The region instance group manager this instance config is part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#region_instance_group_manager ComputeRegionPerInstanceConfig#region_instance_group_manager}
	RegionInstanceGroupManager *string `field:"required" json:"regionInstanceGroupManager" yaml:"regionInstanceGroupManager"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#id ComputeRegionPerInstanceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The minimal action to perform on the instance during an update. Default is 'NONE'. Possible values are: REPLACE RESTART REFRESH NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#minimal_action ComputeRegionPerInstanceConfig#minimal_action}
	MinimalAction *string `field:"optional" json:"minimalAction" yaml:"minimalAction"`
	// The most disruptive action to perform on the instance during an update. Default is 'REPLACE'. Possible values are: REPLACE RESTART REFRESH NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#most_disruptive_allowed_action ComputeRegionPerInstanceConfig#most_disruptive_allowed_action}
	MostDisruptiveAllowedAction *string `field:"optional" json:"mostDisruptiveAllowedAction" yaml:"mostDisruptiveAllowedAction"`
	// preserved_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#preserved_state ComputeRegionPerInstanceConfig#preserved_state}
	PreservedState *ComputeRegionPerInstanceConfigPreservedState `field:"optional" json:"preservedState" yaml:"preservedState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#project ComputeRegionPerInstanceConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the containing instance group manager is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#region ComputeRegionPerInstanceConfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	//
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#remove_instance_state_on_destroy ComputeRegionPerInstanceConfig#remove_instance_state_on_destroy}
	RemoveInstanceStateOnDestroy interface{} `field:"optional" json:"removeInstanceStateOnDestroy" yaml:"removeInstanceStateOnDestroy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_per_instance_config#timeouts ComputeRegionPerInstanceConfig#timeouts}
	Timeouts *ComputeRegionPerInstanceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

