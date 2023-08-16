package networkconnectivityserviceconnectionpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivityServiceConnectionPolicyConfig struct {
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
	// The location of the ServiceConnectionPolicy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#location NetworkConnectivityServiceConnectionPolicy#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#name NetworkConnectivityServiceConnectionPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#network NetworkConnectivityServiceConnectionPolicy#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The service class identifier for which this ServiceConnectionPolicy is for.
	//
	// The service class identifier is a unique, symbolic representation of a ServiceClass.
	// It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#service_class NetworkConnectivityServiceConnectionPolicy#service_class}
	ServiceClass *string `field:"required" json:"serviceClass" yaml:"serviceClass"`
	// Free-text description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#description NetworkConnectivityServiceConnectionPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#id NetworkConnectivityServiceConnectionPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#labels NetworkConnectivityServiceConnectionPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#project NetworkConnectivityServiceConnectionPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// psc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#psc_config NetworkConnectivityServiceConnectionPolicy#psc_config}
	PscConfig *NetworkConnectivityServiceConnectionPolicyPscConfig `field:"optional" json:"pscConfig" yaml:"pscConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_service_connection_policy#timeouts NetworkConnectivityServiceConnectionPolicy#timeouts}
	Timeouts *NetworkConnectivityServiceConnectionPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

