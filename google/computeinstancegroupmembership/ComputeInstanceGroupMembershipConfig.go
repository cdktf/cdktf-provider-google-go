// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancegroupmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupMembershipConfig struct {
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
	// An instance being added to the InstanceGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/compute_instance_group_membership#instance ComputeInstanceGroupMembership#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Represents an Instance Group resource name that the instance belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/compute_instance_group_membership#instance_group ComputeInstanceGroupMembership#instance_group}
	InstanceGroup *string `field:"required" json:"instanceGroup" yaml:"instanceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/compute_instance_group_membership#id ComputeInstanceGroupMembership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/compute_instance_group_membership#project ComputeInstanceGroupMembership#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/compute_instance_group_membership#timeouts ComputeInstanceGroupMembership#timeouts}
	Timeouts *ComputeInstanceGroupMembershipTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A reference to the zone where the instance group resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/compute_instance_group_membership#zone ComputeInstanceGroupMembership#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

