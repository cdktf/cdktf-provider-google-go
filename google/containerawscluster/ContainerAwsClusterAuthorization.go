package containerawscluster


type ContainerAwsClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_aws_cluster#admin_users ContainerAwsCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

