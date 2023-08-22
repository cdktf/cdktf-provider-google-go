package containerawscluster


type ContainerAwsClusterNetworking struct {
	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// Only a single range is supported. This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_aws_cluster#pod_address_cidr_blocks ContainerAwsCluster#pod_address_cidr_blocks}
	PodAddressCidrBlocks *[]*string `field:"required" json:"podAddressCidrBlocks" yaml:"podAddressCidrBlocks"`
	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// Only a single range is supported. This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_aws_cluster#service_address_cidr_blocks ContainerAwsCluster#service_address_cidr_blocks}
	ServiceAddressCidrBlocks *[]*string `field:"required" json:"serviceAddressCidrBlocks" yaml:"serviceAddressCidrBlocks"`
	// The VPC associated with the cluster.
	//
	// All component clusters (i.e. control plane and node pools) run on a single VPC. This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_aws_cluster#vpc_id ContainerAwsCluster#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Disable the per node pool subnet security group rules on the control plane security group.
	//
	// When set to true, you must also provide one or more security groups that ensure node pools are able to send requests to the control plane on TCP/443 and TCP/8132. Failure to do so may result in unavailable node pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_aws_cluster#per_node_pool_sg_rules_disabled ContainerAwsCluster#per_node_pool_sg_rules_disabled}
	PerNodePoolSgRulesDisabled interface{} `field:"optional" json:"perNodePoolSgRulesDisabled" yaml:"perNodePoolSgRulesDisabled"`
}

