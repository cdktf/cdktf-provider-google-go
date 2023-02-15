package containercluster


type ContainerClusterGatewayApiConfig struct {
	// The Gateway API release channel to use for Gateway API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#channel ContainerCluster#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

