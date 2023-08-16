package containerazurecluster


type ContainerAzureClusterControlPlaneProxyConfig struct {
	// The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as `/subscriptions/<subscription-id>/resourceGroups/<resource-group-name>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#resource_group_id ContainerAzureCluster#resource_group_id}
	ResourceGroupId *string `field:"required" json:"resourceGroupId" yaml:"resourceGroupId"`
	// The URL the of the proxy setting secret with its version. Secret ids are formatted as `https:<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#secret_id ContainerAzureCluster#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}

