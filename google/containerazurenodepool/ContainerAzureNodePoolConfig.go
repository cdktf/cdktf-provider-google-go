package containerazurenodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAzureNodePoolConfig struct {
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
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#autoscaling ContainerAzureNodePool#autoscaling}
	Autoscaling *ContainerAzureNodePoolAutoscaling `field:"required" json:"autoscaling" yaml:"autoscaling"`
	// The azureCluster for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#cluster ContainerAzureNodePool#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#config ContainerAzureNodePool#config}
	Config *ContainerAzureNodePoolConfigA `field:"required" json:"config" yaml:"config"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#location ContainerAzureNodePool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// max_pods_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#max_pods_constraint ContainerAzureNodePool#max_pods_constraint}
	MaxPodsConstraint *ContainerAzureNodePoolMaxPodsConstraint `field:"required" json:"maxPodsConstraint" yaml:"maxPodsConstraint"`
	// The name of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#name ContainerAzureNodePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARM ID of the subnet where the node pool VMs run.
	//
	// Make sure it's a subnet under the virtual network in the cluster configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#subnet_id ContainerAzureNodePool#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The Kubernetes version (e.g. `1.19.10-gke.1000`) running on this node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#version ContainerAzureNodePool#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Optional.
	//
	// Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#annotations ContainerAzureNodePool#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Optional. The Azure availability zone of the nodes in this nodepool. When unspecified, it defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#azure_availability_zone ContainerAzureNodePool#azure_availability_zone}
	AzureAvailabilityZone *string `field:"optional" json:"azureAvailabilityZone" yaml:"azureAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#id ContainerAzureNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#project ContainerAzureNodePool#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_azure_node_pool#timeouts ContainerAzureNodePool#timeouts}
	Timeouts *ContainerAzureNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

