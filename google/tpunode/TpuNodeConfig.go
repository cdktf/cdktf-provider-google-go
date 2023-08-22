package tpunode

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TpuNodeConfig struct {
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
	// The type of hardware accelerators associated with this node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#accelerator_type TpuNode#accelerator_type}
	AcceleratorType *string `field:"required" json:"acceleratorType" yaml:"acceleratorType"`
	// The immutable name of the TPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#name TpuNode#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of Tensorflow running in the Node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#tensorflow_version TpuNode#tensorflow_version}
	TensorflowVersion *string `field:"required" json:"tensorflowVersion" yaml:"tensorflowVersion"`
	// The CIDR block that the TPU node will use when selecting an IP address.
	//
	// This CIDR block must be a /29 block; the Compute Engine
	// networks API forbids a smaller block, and using a larger block would
	// be wasteful (a node can only consume one IP address).
	//
	// Errors will occur if the CIDR block has already been used for a
	// currently existing TPU node, the CIDR block conflicts with any
	// subnetworks in the user's provided network, or the provided network
	// is peered with another network that is using that CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#cidr_block TpuNode#cidr_block}
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#description TpuNode#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#id TpuNode#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#labels TpuNode#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of a network to peer the TPU node to.
	//
	// It must be a
	// preexisting Compute Engine network inside of the project on which
	// this API has been activated. If none is provided, "default" will be
	// used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#network TpuNode#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#project TpuNode#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// scheduling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#scheduling_config TpuNode#scheduling_config}
	SchedulingConfig *TpuNodeSchedulingConfig `field:"optional" json:"schedulingConfig" yaml:"schedulingConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#timeouts TpuNode#timeouts}
	Timeouts *TpuNodeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Whether the VPC peering for the node is set up through Service Networking API.
	//
	// The VPC Peering should be set up before provisioning the node. If this field is set,
	// cidr_block field should not be specified. If the network that you want to peer the
	// TPU Node to is a Shared VPC network, the node must be created with this this field enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#use_service_networking TpuNode#use_service_networking}
	UseServiceNetworking interface{} `field:"optional" json:"useServiceNetworking" yaml:"useServiceNetworking"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/tpu_node#zone TpuNode#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

