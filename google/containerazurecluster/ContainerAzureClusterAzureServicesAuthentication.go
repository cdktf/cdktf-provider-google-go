// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerazurecluster


type ContainerAzureClusterAzureServicesAuthentication struct {
	// The Azure Active Directory Application ID for Authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_azure_cluster#application_id ContainerAzureCluster#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Azure Active Directory Tenant ID for Authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_azure_cluster#tenant_id ContainerAzureCluster#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

