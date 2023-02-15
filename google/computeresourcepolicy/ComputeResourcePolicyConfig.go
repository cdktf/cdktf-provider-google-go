package computeresourcepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResourcePolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the resource, provided by the client when initially creating the resource.
	//
	// The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#name ComputeResourcePolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#description ComputeResourcePolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// group_placement_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#group_placement_policy ComputeResourcePolicy#group_placement_policy}
	GroupPlacementPolicy *ComputeResourcePolicyGroupPlacementPolicy `field:"optional" json:"groupPlacementPolicy" yaml:"groupPlacementPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#id ComputeResourcePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_schedule_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#instance_schedule_policy ComputeResourcePolicy#instance_schedule_policy}
	InstanceSchedulePolicy *ComputeResourcePolicyInstanceSchedulePolicy `field:"optional" json:"instanceSchedulePolicy" yaml:"instanceSchedulePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#project ComputeResourcePolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where resource policy resides.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#region ComputeResourcePolicy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// snapshot_schedule_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#snapshot_schedule_policy ComputeResourcePolicy#snapshot_schedule_policy}
	SnapshotSchedulePolicy *ComputeResourcePolicySnapshotSchedulePolicy `field:"optional" json:"snapshotSchedulePolicy" yaml:"snapshotSchedulePolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#timeouts ComputeResourcePolicy#timeouts}
	Timeouts *ComputeResourcePolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

