package containercluster


type ContainerClusterMasterAuthClientCertificateConfig struct {
	// Whether client certificate authorization is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#issue_client_certificate ContainerCluster#issue_client_certificate}
	IssueClientCertificate interface{} `field:"required" json:"issueClientCertificate" yaml:"issueClientCertificate"`
}

