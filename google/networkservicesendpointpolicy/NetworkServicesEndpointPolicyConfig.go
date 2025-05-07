// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesendpointpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesEndpointPolicyConfig struct {
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
	// endpoint_matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#endpoint_matcher NetworkServicesEndpointPolicy#endpoint_matcher}
	EndpointMatcher *NetworkServicesEndpointPolicyEndpointMatcher `field:"required" json:"endpointMatcher" yaml:"endpointMatcher"`
	// Name of the EndpointPolicy resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#name NetworkServicesEndpointPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of endpoint policy. This is primarily used to validate the configuration. Possible values: ["SIDECAR_PROXY", "GRPC_SERVER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#type NetworkServicesEndpointPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#authorization_policy NetworkServicesEndpointPolicy#authorization_policy}
	AuthorizationPolicy *string `field:"optional" json:"authorizationPolicy" yaml:"authorizationPolicy"`
	// A URL referring to a ClientTlsPolicy resource.
	//
	// ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#client_tls_policy NetworkServicesEndpointPolicy#client_tls_policy}
	ClientTlsPolicy *string `field:"optional" json:"clientTlsPolicy" yaml:"clientTlsPolicy"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#description NetworkServicesEndpointPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#id NetworkServicesEndpointPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the TcpRoute resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#labels NetworkServicesEndpointPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#project NetworkServicesEndpointPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A URL referring to ServerTlsPolicy resource.
	//
	// ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#server_tls_policy NetworkServicesEndpointPolicy#server_tls_policy}
	ServerTlsPolicy *string `field:"optional" json:"serverTlsPolicy" yaml:"serverTlsPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#timeouts NetworkServicesEndpointPolicy#timeouts}
	Timeouts *NetworkServicesEndpointPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_port_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_services_endpoint_policy#traffic_port_selector NetworkServicesEndpointPolicy#traffic_port_selector}
	TrafficPortSelector *NetworkServicesEndpointPolicyTrafficPortSelector `field:"optional" json:"trafficPortSelector" yaml:"trafficPortSelector"`
}

