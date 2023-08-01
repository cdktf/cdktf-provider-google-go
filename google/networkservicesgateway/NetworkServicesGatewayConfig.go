package networkservicesgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesGatewayConfig struct {
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
	// Short name of the Gateway resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#name NetworkServicesGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or more port numbers (1-65535), on which the Gateway will receive traffic.
	//
	// The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are
	// limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#ports NetworkServicesGateway#ports}
	Ports *[]*float64 `field:"required" json:"ports" yaml:"ports"`
	// Immutable. The type of the customer-managed gateway. Possible values are: * OPEN_MESH * SECURE_WEB_GATEWAY. Possible values: ["TYPE_UNSPECIFIED", "OPEN_MESH", "SECURE_WEB_GATEWAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#type NetworkServicesGateway#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Zero or one IPv4-address on which the Gateway will receive the traffic.
	//
	// When no address is provided,
	// an IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'.
	// Gateways of type 'OPEN_MESH' listen on 0.0.0.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#addresses NetworkServicesGateway#addresses}
	Addresses *[]*string `field:"optional" json:"addresses" yaml:"addresses"`
	// A fully-qualified Certificates URL reference.
	//
	// The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection.
	// This feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#certificate_urls NetworkServicesGateway#certificate_urls}
	CertificateUrls *[]*string `field:"optional" json:"certificateUrls" yaml:"certificateUrls"`
	// When deleting a gateway of type 'SECURE_WEB_GATEWAY', this boolean option will also delete auto generated router by the gateway creation.
	//
	// If there is no other gateway of type 'SECURE_WEB_GATEWAY' remaining for that region and network it will be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#delete_swg_autogen_router_on_destroy NetworkServicesGateway#delete_swg_autogen_router_on_destroy}
	DeleteSwgAutogenRouterOnDestroy interface{} `field:"optional" json:"deleteSwgAutogenRouterOnDestroy" yaml:"deleteSwgAutogenRouterOnDestroy"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#description NetworkServicesGateway#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A fully-qualified GatewaySecurityPolicy URL reference.
	//
	// Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections.
	// For example: 'projects/*\/locations/*\/gatewaySecurityPolicies/swg-policy'.
	// This policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#gateway_security_policy NetworkServicesGateway#gateway_security_policy}
	GatewaySecurityPolicy *string `field:"optional" json:"gatewaySecurityPolicy" yaml:"gatewaySecurityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#id NetworkServicesGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the Gateway resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#labels NetworkServicesGateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the gateway. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#location NetworkServicesGateway#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The relative resource name identifying the VPC network that is using this configuration.
	//
	// For example: 'projects/*\/global/networks/network-1'.
	// Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#network NetworkServicesGateway#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#project NetworkServicesGateway#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Immutable.
	//
	// Scope determines how configuration across multiple Gateway instances are merged.
	// The configuration for multiple Gateway instances with the same scope will be merged as presented as
	// a single coniguration to the proxy/load balancer.
	// Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#scope NetworkServicesGateway#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#server_tls_policy NetworkServicesGateway#server_tls_policy}
	ServerTlsPolicy *string `field:"optional" json:"serverTlsPolicy" yaml:"serverTlsPolicy"`
	// The relative resource name identifying the subnetwork in which this SWG is allocated.
	//
	// For example: 'projects/*\/regions/us-central1/subnetworks/network-1'.
	// Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#subnetwork NetworkServicesGateway#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_gateway#timeouts NetworkServicesGateway#timeouts}
	Timeouts *NetworkServicesGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

