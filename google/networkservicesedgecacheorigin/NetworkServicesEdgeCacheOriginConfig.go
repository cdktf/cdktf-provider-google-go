package networkservicesedgecacheorigin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesEdgeCacheOriginConfig struct {
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
	// Name of the resource;
	//
	// provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#name NetworkServicesEdgeCacheOrigin#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	//
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com, IPv4: 35.218.1.1, IPv6: 2607:f8b0:4012:809::200e, Cloud Storage: gs://bucketname
	//
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.  It must not contain a protocol (e.g., https://) and it must not contain any slashes.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#origin_address NetworkServicesEdgeCacheOrigin#origin_address}
	OriginAddress *string `field:"required" json:"originAddress" yaml:"originAddress"`
	// aws_v4_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#aws_v4_authentication NetworkServicesEdgeCacheOrigin#aws_v4_authentication}
	AwsV4Authentication *NetworkServicesEdgeCacheOriginAwsV4Authentication `field:"optional" json:"awsV4Authentication" yaml:"awsV4Authentication"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#description NetworkServicesEdgeCacheOrigin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Origin resource to try when the current origin cannot be reached.
	//
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	//
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#failover_origin NetworkServicesEdgeCacheOrigin#failover_origin}
	FailoverOrigin *string `field:"optional" json:"failoverOrigin" yaml:"failoverOrigin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#id NetworkServicesEdgeCacheOrigin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the EdgeCache resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#labels NetworkServicesEdgeCacheOrigin#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The maximum number of attempts to cache fill from this origin.
	//
	// Another attempt is made when a cache fill fails with one of the retryConditions.
	//
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	//
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	//
	// The last valid, non-retried response from all origins will be returned to the client.
	// If no origin returns a valid response, an HTTP 502 will be returned to the client.
	//
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#max_attempts NetworkServicesEdgeCacheOrigin#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// origin_override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#origin_override_action NetworkServicesEdgeCacheOrigin#origin_override_action}
	OriginOverrideAction *NetworkServicesEdgeCacheOriginOriginOverrideAction `field:"optional" json:"originOverrideAction" yaml:"originOverrideAction"`
	// origin_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#origin_redirect NetworkServicesEdgeCacheOrigin#origin_redirect}
	OriginRedirect *NetworkServicesEdgeCacheOriginOriginRedirect `field:"optional" json:"originRedirect" yaml:"originRedirect"`
	// The port to connect to the origin on.
	//
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#port NetworkServicesEdgeCacheOrigin#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#project NetworkServicesEdgeCacheOrigin#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The protocol to use to connect to the configured origin.
	//
	// Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	//
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server. Possible values: ["HTTP2", "HTTPS", "HTTP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#protocol NetworkServicesEdgeCacheOrigin#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Specifies one or more retry conditions for the configured origin.
	//
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	//
	// The default retryCondition is "CONNECT_FAILURE".
	//
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	//
	// Valid values are:
	//
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
	// - FORBIDDEN: Retry if the origin returns a HTTP 403 (Forbidden). Possible values: ["CONNECT_FAILURE", "HTTP_5XX", "GATEWAY_ERROR", "RETRIABLE_4XX", "NOT_FOUND", "FORBIDDEN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#retry_conditions NetworkServicesEdgeCacheOrigin#retry_conditions}
	RetryConditions *[]*string `field:"optional" json:"retryConditions" yaml:"retryConditions"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#timeout NetworkServicesEdgeCacheOrigin#timeout}
	Timeout *NetworkServicesEdgeCacheOriginTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_origin#timeouts NetworkServicesEdgeCacheOrigin#timeouts}
	Timeouts *NetworkServicesEdgeCacheOriginTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

