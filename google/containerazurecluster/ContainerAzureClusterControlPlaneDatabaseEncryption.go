package containerazurecluster


type ContainerAzureClusterControlPlaneDatabaseEncryption struct {
	// The ARM ID of the Azure Key Vault key to encrypt / decrypt data.
	//
	// For example: `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.KeyVault/vaults/<key-vault-id>/keys/<key-name>` Encryption will always take the latest version of the key and hence specific version is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#key_id ContainerAzureCluster#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

