package computeinstancegroupnamedport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupNamedPortAConfig struct {
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
	// The name of the instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_group_named_port#group ComputeInstanceGroupNamedPortA#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// The name for this named port. The name must be 1-63 characters long, and comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_group_named_port#name ComputeInstanceGroupNamedPortA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number, which can be a value between 1 and 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_group_named_port#port ComputeInstanceGroupNamedPortA#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_group_named_port#id ComputeInstanceGroupNamedPortA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_group_named_port#project ComputeInstanceGroupNamedPortA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_group_named_port#timeouts ComputeInstanceGroupNamedPortA#timeouts}
	Timeouts *ComputeInstanceGroupNamedPortTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone of the instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_group_named_port#zone ComputeInstanceGroupNamedPortA#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

