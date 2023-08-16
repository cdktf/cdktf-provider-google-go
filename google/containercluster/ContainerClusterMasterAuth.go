package containercluster


type ContainerClusterMasterAuth struct {
	// client_certificate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#client_certificate_config ContainerCluster#client_certificate_config}
	ClientCertificateConfig *ContainerClusterMasterAuthClientCertificateConfig `field:"required" json:"clientCertificateConfig" yaml:"clientCertificateConfig"`
}

