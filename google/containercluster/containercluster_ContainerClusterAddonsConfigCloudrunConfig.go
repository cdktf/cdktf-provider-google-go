package containercluster


type ContainerClusterAddonsConfigCloudrunConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#disabled ContainerCluster#disabled}.
	Disabled interface{} `field:"required" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#load_balancer_type ContainerCluster#load_balancer_type}.
	LoadBalancerType *string `field:"optional" json:"loadBalancerType" yaml:"loadBalancerType"`
}

