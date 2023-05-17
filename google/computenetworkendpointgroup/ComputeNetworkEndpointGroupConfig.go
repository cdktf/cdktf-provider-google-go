package computenetworkendpointgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeNetworkEndpointGroupConfig struct {
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
	// Name of the resource;
	//
	// provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#name ComputeNetworkEndpointGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#network ComputeNetworkEndpointGroup#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The default port used if the port number is not specified in the network endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#default_port ComputeNetworkEndpointGroup#default_port}
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#description ComputeNetworkEndpointGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#id ComputeNetworkEndpointGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Type of network endpoints in this network endpoint group.
	//
	// NON_GCP_PRIVATE_IP_PORT is used for hybrid connectivity network
	// endpoint groups (see https://cloud.google.com/load-balancing/docs/hybrid).
	// Note that NON_GCP_PRIVATE_IP_PORT can only be used with Backend Services
	// that 1) have the following load balancing schemes: EXTERNAL, EXTERNAL_MANAGED,
	// INTERNAL_MANAGED, and INTERNAL_SELF_MANAGED and 2) support the RATE or
	// CONNECTION balancing modes.
	//
	// Possible values include: GCE_VM_IP, GCE_VM_IP_PORT, and NON_GCP_PRIVATE_IP_PORT. Default value: "GCE_VM_IP_PORT" Possible values: ["GCE_VM_IP", "GCE_VM_IP_PORT", "NON_GCP_PRIVATE_IP_PORT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#network_endpoint_type ComputeNetworkEndpointGroup#network_endpoint_type}
	NetworkEndpointType *string `field:"optional" json:"networkEndpointType" yaml:"networkEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#project ComputeNetworkEndpointGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional subnetwork to which all network endpoints in the NEG belong.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#subnetwork ComputeNetworkEndpointGroup#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#timeouts ComputeNetworkEndpointGroup#timeouts}
	Timeouts *ComputeNetworkEndpointGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zone where the network endpoint group is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_endpoint_group#zone ComputeNetworkEndpointGroup#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

