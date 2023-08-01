package containerawscluster


type ContainerAwsClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_aws_cluster#username ContainerAwsCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

