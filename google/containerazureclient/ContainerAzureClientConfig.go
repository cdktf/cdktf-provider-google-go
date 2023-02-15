package containerazureclient

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAzureClientConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Azure Active Directory Application ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#application_id ContainerAzureClient#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#location ContainerAzureClient#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#name ContainerAzureClient#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Azure Active Directory Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#tenant_id ContainerAzureClient#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#id ContainerAzureClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#project ContainerAzureClient#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#timeouts ContainerAzureClient#timeouts}
	Timeouts *ContainerAzureClientTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

