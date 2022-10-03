package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerUpdatePolicy struct {
	// Minimal action to be taken on an instance.
	//
	// You can specify either REFRESH to update without stopping instances, RESTART to restart existing instances or REPLACE to delete and create new instances from the target template. If you specify a REFRESH, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#minimal_action ComputeRegionInstanceGroupManager#minimal_action}
	MinimalAction *string `field:"required" json:"minimalAction" yaml:"minimalAction"`
	// The type of update process.
	//
	// You can specify either PROACTIVE so that the instance group manager proactively executes actions in order to bring instances to their target versions or OPPORTUNISTIC so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#type ComputeRegionInstanceGroupManager#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The instance redistribution policy for regional managed instance groups.
	//
	// Valid values are: "PROACTIVE", "NONE". If PROACTIVE (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If NONE, proactive redistribution is disabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#instance_redistribution_type ComputeRegionInstanceGroupManager#instance_redistribution_type}
	InstanceRedistributionType *string `field:"optional" json:"instanceRedistributionType" yaml:"instanceRedistributionType"`
	// The maximum number of instances that can be created above the specified targetSize during the update process.
	//
	// Conflicts with max_surge_percent. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of max_unavailable_fixed or max_surge_fixed must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#max_surge_fixed ComputeRegionInstanceGroupManager#max_surge_fixed}
	MaxSurgeFixed *float64 `field:"optional" json:"maxSurgeFixed" yaml:"maxSurgeFixed"`
	// The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process.
	//
	// Conflicts with max_surge_fixed. Percent value is only allowed for regional managed instance groups with size at least 10.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#max_surge_percent ComputeRegionInstanceGroupManager#max_surge_percent}
	MaxSurgePercent *float64 `field:"optional" json:"maxSurgePercent" yaml:"maxSurgePercent"`
	// The maximum number of instances that can be unavailable during the update process.
	//
	// Conflicts with max_unavailable_percent. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of max_unavailable_fixed or max_surge_fixed must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#max_unavailable_fixed ComputeRegionInstanceGroupManager#max_unavailable_fixed}
	MaxUnavailableFixed *float64 `field:"optional" json:"maxUnavailableFixed" yaml:"maxUnavailableFixed"`
	// The maximum number of instances(calculated as percentage) that can be unavailable during the update process.
	//
	// Conflicts with max_unavailable_fixed. Percent value is only allowed for regional managed instance groups with size at least 10.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#max_unavailable_percent ComputeRegionInstanceGroupManager#max_unavailable_percent}
	MaxUnavailablePercent *float64 `field:"optional" json:"maxUnavailablePercent" yaml:"maxUnavailablePercent"`
	// Most disruptive action that is allowed to be taken on an instance.
	//
	// You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#most_disruptive_allowed_action ComputeRegionInstanceGroupManager#most_disruptive_allowed_action}
	MostDisruptiveAllowedAction *string `field:"optional" json:"mostDisruptiveAllowedAction" yaml:"mostDisruptiveAllowedAction"`
	// The instance replacement method for regional managed instance groups.
	//
	// Valid values are: "RECREATE", "SUBSTITUTE". If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#replacement_method ComputeRegionInstanceGroupManager#replacement_method}
	ReplacementMethod *string `field:"optional" json:"replacementMethod" yaml:"replacementMethod"`
}

