package containercluster


type ContainerClusterNotificationConfig struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#pubsub ContainerCluster#pubsub}
	Pubsub *ContainerClusterNotificationConfigPubsub `field:"required" json:"pubsub" yaml:"pubsub"`
}

