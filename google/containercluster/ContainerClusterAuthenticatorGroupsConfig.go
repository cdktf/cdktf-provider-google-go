package containercluster


type ContainerClusterAuthenticatorGroupsConfig struct {
	// The name of the RBAC security group for use with Google security groups in Kubernetes RBAC.
	//
	// Group name must be in format gke-security-groups@yourdomain.com.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#security_group ContainerCluster#security_group}
	SecurityGroup *string `field:"required" json:"securityGroup" yaml:"securityGroup"`
}

