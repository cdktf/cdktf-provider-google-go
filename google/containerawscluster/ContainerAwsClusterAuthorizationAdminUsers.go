package containerawscluster


type ContainerAwsClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_cluster#username ContainerAwsCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

