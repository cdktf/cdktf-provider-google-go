package containercluster


type ContainerClusterGatewayApiConfig struct {
	// The Gateway API release channel to use for Gateway API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#channel ContainerCluster#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

