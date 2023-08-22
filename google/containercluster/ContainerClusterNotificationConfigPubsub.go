package containercluster


type ContainerClusterNotificationConfigPubsub struct {
	// Whether or not the notification config is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#filter ContainerCluster#filter}
	Filter *ContainerClusterNotificationConfigPubsubFilter `field:"optional" json:"filter" yaml:"filter"`
	// The pubsub topic to push upgrade notifications to.
	//
	// Must be in the same project as the cluster. Must be in the format: projects/{project}/topics/{topic}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#topic ContainerCluster#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

