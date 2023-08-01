package computeinstancegroupmanager


type ComputeInstanceGroupManagerUpdatePolicy struct {
	// Minimal action to be taken on an instance.
	//
	// You can specify either REFRESH to update without stopping instances, RESTART to restart existing instances or REPLACE to delete and create new instances from the target template. If you specify a REFRESH, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#minimal_action ComputeInstanceGroupManager#minimal_action}
	MinimalAction *string `field:"required" json:"minimalAction" yaml:"minimalAction"`
	// The type of update process.
	//
	// You can specify either PROACTIVE so that the instance group manager proactively executes actions in order to bring instances to their target versions or OPPORTUNISTIC so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#type ComputeInstanceGroupManager#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The maximum number of instances that can be created above the specified targetSize during the update process.
	//
	// Conflicts with max_surge_percent. If neither is set, defaults to 1
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#max_surge_fixed ComputeInstanceGroupManager#max_surge_fixed}
	MaxSurgeFixed *float64 `field:"optional" json:"maxSurgeFixed" yaml:"maxSurgeFixed"`
	// The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process.
	//
	// Conflicts with max_surge_fixed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#max_surge_percent ComputeInstanceGroupManager#max_surge_percent}
	MaxSurgePercent *float64 `field:"optional" json:"maxSurgePercent" yaml:"maxSurgePercent"`
	// The maximum number of instances that can be unavailable during the update process.
	//
	// Conflicts with max_unavailable_percent. If neither is set, defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#max_unavailable_fixed ComputeInstanceGroupManager#max_unavailable_fixed}
	MaxUnavailableFixed *float64 `field:"optional" json:"maxUnavailableFixed" yaml:"maxUnavailableFixed"`
	// The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with max_unavailable_fixed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#max_unavailable_percent ComputeInstanceGroupManager#max_unavailable_percent}
	MaxUnavailablePercent *float64 `field:"optional" json:"maxUnavailablePercent" yaml:"maxUnavailablePercent"`
	// Most disruptive action that is allowed to be taken on an instance.
	//
	// You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#most_disruptive_allowed_action ComputeInstanceGroupManager#most_disruptive_allowed_action}
	MostDisruptiveAllowedAction *string `field:"optional" json:"mostDisruptiveAllowedAction" yaml:"mostDisruptiveAllowedAction"`
	// The instance replacement method for managed instance groups.
	//
	// Valid values are: "RECREATE", "SUBSTITUTE". If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#replacement_method ComputeInstanceGroupManager#replacement_method}
	ReplacementMethod *string `field:"optional" json:"replacementMethod" yaml:"replacementMethod"`
}

