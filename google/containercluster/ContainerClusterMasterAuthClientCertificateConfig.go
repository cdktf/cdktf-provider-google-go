package containercluster


type ContainerClusterMasterAuthClientCertificateConfig struct {
	// Whether client certificate authorization is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#issue_client_certificate ContainerCluster#issue_client_certificate}
	IssueClientCertificate interface{} `field:"required" json:"issueClientCertificate" yaml:"issueClientCertificate"`
}

