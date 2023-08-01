package computeregionsslpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionSslPolicyConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#name ComputeRegionSslPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The region where the regional SSL policy resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#region ComputeRegionSslPolicy#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// A list of features enabled when the selected profile is CUSTOM.
	//
	// The
	// method returns the set of features that can be specified in this
	// list. This field must be empty if the profile is not CUSTOM.
	//
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// must* be present when using the 'CUSTOM' profile. This argument
	// must not* be present when using any other profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#custom_features ComputeRegionSslPolicy#custom_features}
	CustomFeatures *[]*string `field:"optional" json:"customFeatures" yaml:"customFeatures"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#description ComputeRegionSslPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#id ComputeRegionSslPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer.
	//
	// Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#min_tls_version ComputeRegionSslPolicy#min_tls_version}
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients.
	//
	// If using 'CUSTOM',
	// the set of SSL features to enable must be specified in the
	// 'customFeatures' field.
	//
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// 'CUSTOM' is used, the 'custom_features' attribute **must be set**. Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#profile ComputeRegionSslPolicy#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#project ComputeRegionSslPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_ssl_policy#timeouts ComputeRegionSslPolicy#timeouts}
	Timeouts *ComputeRegionSslPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

