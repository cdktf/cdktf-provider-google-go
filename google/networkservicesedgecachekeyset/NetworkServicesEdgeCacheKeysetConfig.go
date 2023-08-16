package networkservicesedgecachekeyset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesEdgeCacheKeysetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#name NetworkServicesEdgeCacheKeyset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#description NetworkServicesEdgeCacheKeyset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#id NetworkServicesEdgeCacheKeyset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the EdgeCache resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#labels NetworkServicesEdgeCacheKeyset#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#project NetworkServicesEdgeCacheKeyset#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#public_key NetworkServicesEdgeCacheKeyset#public_key}
	PublicKey interface{} `field:"optional" json:"publicKey" yaml:"publicKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#timeouts NetworkServicesEdgeCacheKeyset#timeouts}
	Timeouts *NetworkServicesEdgeCacheKeysetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// validation_shared_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_keyset#validation_shared_keys NetworkServicesEdgeCacheKeyset#validation_shared_keys}
	ValidationSharedKeys interface{} `field:"optional" json:"validationSharedKeys" yaml:"validationSharedKeys"`
}

