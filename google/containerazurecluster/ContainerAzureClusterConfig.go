package containerazurecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAzureClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#authorization ContainerAzureCluster#authorization}
	Authorization *ContainerAzureClusterAuthorization `field:"required" json:"authorization" yaml:"authorization"`
	// The Azure region where the cluster runs.
	//
	// Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#azure_region ContainerAzureCluster#azure_region}
	AzureRegion *string `field:"required" json:"azureRegion" yaml:"azureRegion"`
	// control_plane block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#control_plane ContainerAzureCluster#control_plane}
	ControlPlane *ContainerAzureClusterControlPlane `field:"required" json:"controlPlane" yaml:"controlPlane"`
	// fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#fleet ContainerAzureCluster#fleet}
	Fleet *ContainerAzureClusterFleet `field:"required" json:"fleet" yaml:"fleet"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#location ContainerAzureCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#name ContainerAzureCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// networking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#networking ContainerAzureCluster#networking}
	Networking *ContainerAzureClusterNetworking `field:"required" json:"networking" yaml:"networking"`
	// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*\/resourceGroups/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#resource_group_id ContainerAzureCluster#resource_group_id}
	ResourceGroupId *string `field:"required" json:"resourceGroupId" yaml:"resourceGroupId"`
	// Optional.
	//
	// Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#annotations ContainerAzureCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// azure_services_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#azure_services_authentication ContainerAzureCluster#azure_services_authentication}
	AzureServicesAuthentication *ContainerAzureClusterAzureServicesAuthentication `field:"optional" json:"azureServicesAuthentication" yaml:"azureServicesAuthentication"`
	// Name of the AzureClient.
	//
	// The `AzureClient` resource must reside on the same GCP project and region as the `AzureCluster`. `AzureClient` names are formatted as `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#client ContainerAzureCluster#client}
	Client *string `field:"optional" json:"client" yaml:"client"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#description ContainerAzureCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#id ContainerAzureCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#project ContainerAzureCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_cluster#timeouts ContainerAzureCluster#timeouts}
	Timeouts *ContainerAzureClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

