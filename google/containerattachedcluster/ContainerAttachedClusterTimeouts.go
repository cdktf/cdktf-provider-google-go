package containerattachedcluster


type ContainerAttachedClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_attached_cluster#create ContainerAttachedCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_attached_cluster#delete ContainerAttachedCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_attached_cluster#update ContainerAttachedCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

