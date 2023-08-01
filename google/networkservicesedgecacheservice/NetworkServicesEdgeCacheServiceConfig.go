package networkservicesedgecacheservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesEdgeCacheServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#name NetworkServicesEdgeCacheService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#routing NetworkServicesEdgeCacheService#routing}
	Routing *NetworkServicesEdgeCacheServiceRouting `field:"required" json:"routing" yaml:"routing"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#description NetworkServicesEdgeCacheService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disables HTTP/2.
	//
	// HTTP/2 (h2) is enabled by default and recommended for performance. HTTP/2 improves connection re-use and reduces connection setup overhead by sending multiple streams over the same connection.
	//
	// Some legacy HTTP clients may have issues with HTTP/2 connections due to broken HTTP/2 implementations. Setting this to true will prevent HTTP/2 from being advertised and negotiated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#disable_http2 NetworkServicesEdgeCacheService#disable_http2}
	DisableHttp2 interface{} `field:"optional" json:"disableHttp2" yaml:"disableHttp2"`
	// HTTP/3 (IETF QUIC) and Google QUIC are enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#disable_quic NetworkServicesEdgeCacheService#disable_quic}
	DisableQuic interface{} `field:"optional" json:"disableQuic" yaml:"disableQuic"`
	// Resource URL that points at the Cloud Armor edge security policy that is applied on each request against the EdgeCacheService.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#edge_security_policy NetworkServicesEdgeCacheService#edge_security_policy}
	EdgeSecurityPolicy *string `field:"optional" json:"edgeSecurityPolicy" yaml:"edgeSecurityPolicy"`
	// URLs to sslCertificate resources that are used to authenticate connections between users and the EdgeCacheService.
	//
	// Note that only "global" certificates with a "scope" of "EDGE_CACHE" can be attached to an EdgeCacheService.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#edge_ssl_certificates NetworkServicesEdgeCacheService#edge_ssl_certificates}
	EdgeSslCertificates *[]*string `field:"optional" json:"edgeSslCertificates" yaml:"edgeSslCertificates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#id NetworkServicesEdgeCacheService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the EdgeCache resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#labels NetworkServicesEdgeCacheService#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#log_config NetworkServicesEdgeCacheService#log_config}
	LogConfig *NetworkServicesEdgeCacheServiceLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#project NetworkServicesEdgeCacheService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Require TLS (HTTPS) for all clients connecting to this service.
	//
	// Clients who connect over HTTP (port 80) will receive a HTTP 301 to the same URL over HTTPS (port 443).
	// You must have at least one (1) edgeSslCertificate specified to enable this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#require_tls NetworkServicesEdgeCacheService#require_tls}
	RequireTls interface{} `field:"optional" json:"requireTls" yaml:"requireTls"`
	// URL of the SslPolicy resource that will be associated with the EdgeCacheService.
	//
	// If not set, the EdgeCacheService has no SSL policy configured, and will default to the "COMPATIBLE" policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#ssl_policy NetworkServicesEdgeCacheService#ssl_policy}
	SslPolicy *string `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#timeouts NetworkServicesEdgeCacheService#timeouts}
	Timeouts *NetworkServicesEdgeCacheServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

