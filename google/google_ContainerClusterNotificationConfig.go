// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerClusterNotificationConfig struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#pubsub ContainerCluster#pubsub}
	Pubsub *ContainerClusterNotificationConfigPubsub `field:"required" json:"pubsub" yaml:"pubsub"`
}

