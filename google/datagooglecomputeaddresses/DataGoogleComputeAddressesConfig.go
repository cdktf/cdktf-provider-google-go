package datagooglecomputeaddresses

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeAddressesConfig struct {
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
	// Filter sets the optional parameter "filter": A filter expression that filters resources listed in the response.
	//
	// The expression must specify
	// the field name, an operator, and the value that you want to use for
	// filtering. The value must be a string, a number, or a boolean. The
	// operator must be either "=", "!=", ">", "<", "<=", ">=" or ":". For
	// example, if you are filtering Compute Engine instances, you can
	// exclude instances named "example-instance" by specifying "name !=
	// example-instance". The ":" operator can be used with string fields to
	// match substrings. For non-string fields it is equivalent to the "="
	// operator. The ":*" comparison can be used to test whether a key has
	// been defined. For example, to find all objects with "owner" label
	// use: """ labels.owner:* """ You can also filter nested fields. For
	// example, you could specify "scheduling.automaticRestart = false" to
	// include instances only if they are not scheduled for automatic
	// restarts. You can use filtering on nested fields to filter based on
	// resource labels. To filter on multiple expressions, provide each
	// separate expression within parentheses. For example: """
	// (scheduling.automaticRestart = true) (cpuPlatform = "Intel Skylake")
	// """ By default, each expression is an "AND" expression. However, you
	// can include "AND" and "OR" expressions explicitly. For example: """
	// (cpuPlatform = "Intel Skylake") OR (cpuPlatform = "Intel Broadwell")
	// AND (scheduling.automaticRestart = true) """
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_addresses#filter DataGoogleComputeAddresses#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_addresses#id DataGoogleComputeAddresses#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The google project in which addresses are listed. Defaults to provider's configuration if missing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_addresses#project DataGoogleComputeAddresses#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region that should be considered to search addresses. All regions are considered if missing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_addresses#region DataGoogleComputeAddresses#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

