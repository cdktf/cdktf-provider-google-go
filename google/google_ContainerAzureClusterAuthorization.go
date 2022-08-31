// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerAzureClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_cluster#admin_users ContainerAzureCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

