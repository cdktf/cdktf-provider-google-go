package computenodegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeNodeGroupConfig struct {
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
	// The URL of the node template to which this node group belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#node_template ComputeNodeGroup#node_template}
	NodeTemplate *string `field:"required" json:"nodeTemplate" yaml:"nodeTemplate"`
	// autoscaling_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#autoscaling_policy ComputeNodeGroup#autoscaling_policy}
	AutoscalingPolicy *ComputeNodeGroupAutoscalingPolicy `field:"optional" json:"autoscalingPolicy" yaml:"autoscalingPolicy"`
	// An optional textual description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#description ComputeNodeGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#id ComputeNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The initial number of nodes in the node group. One of 'initial_size' or 'size' must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#initial_size ComputeNodeGroup#initial_size}
	InitialSize *float64 `field:"optional" json:"initialSize" yaml:"initialSize"`
	// Specifies how to handle instances when a node in the group undergoes maintenance.
	//
	// Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#maintenance_policy ComputeNodeGroup#maintenance_policy}
	MaintenancePolicy *string `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#maintenance_window ComputeNodeGroup#maintenance_window}
	MaintenanceWindow *ComputeNodeGroupMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#name ComputeNodeGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#project ComputeNodeGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// share_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#share_settings ComputeNodeGroup#share_settings}
	ShareSettings *ComputeNodeGroupShareSettings `field:"optional" json:"shareSettings" yaml:"shareSettings"`
	// The total number of nodes in the node group. One of 'initial_size' or 'size' must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#size ComputeNodeGroup#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#timeouts ComputeNodeGroup#timeouts}
	Timeouts *ComputeNodeGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zone where this node group is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_node_group#zone ComputeNodeGroup#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

