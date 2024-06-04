// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceslbrouteextension


type NetworkServicesLbRouteExtensionExtensionChainsExtensions struct {
	// The name for this extension.
	//
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,
	// and can have a maximum length of 63 characters. Additionally, the first character must be a letter
	// and the last a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_services_lb_route_extension#name NetworkServicesLbRouteExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The reference to the service that runs the extension. Must be a reference to a backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_services_lb_route_extension#service NetworkServicesLbRouteExtension#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// The :authority header in the gRPC request sent from Envoy to the extension service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_services_lb_route_extension#authority NetworkServicesLbRouteExtension#authority}
	Authority *string `field:"optional" json:"authority" yaml:"authority"`
	// Determines how the proxy behaves if the call to the extension fails or times out.
	//
	// When set to TRUE, request or response processing continues without error.
	// Any subsequent extensions in the extension chain are also executed.
	// When set to FALSE: * If response headers have not been delivered to the downstream client,
	// a generic 500 error is returned to the client. The error response can be tailored by
	// configuring a custom error response in the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_services_lb_route_extension#fail_open NetworkServicesLbRouteExtension#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// List of the HTTP headers to forward to the extension (from the client or backend).
	//
	// If omitted, all headers are sent. Each element is a string indicating the header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_services_lb_route_extension#forward_headers NetworkServicesLbRouteExtension#forward_headers}
	ForwardHeaders *[]*string `field:"optional" json:"forwardHeaders" yaml:"forwardHeaders"`
	// Specifies the timeout for each individual message on the stream.
	//
	// The timeout must be between 10-1000 milliseconds.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_services_lb_route_extension#timeout NetworkServicesLbRouteExtension#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

