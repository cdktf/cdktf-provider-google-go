package containerazurecluster


type ContainerAzureClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_azure_cluster#username ContainerAzureCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

