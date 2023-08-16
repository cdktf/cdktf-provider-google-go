package datagooglecomputeforwardingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeForwardingRuleConfig struct {
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
	// The name must be 1-63 characters long, and comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	//
	// Specifically, the name must be 1-63 characters long and match the regular
	// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit, except the last character, which
	// cannot be a dash.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, the forwarding rule name must be a 1-20 characters string with
	// lowercase letters and numbers and must start with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_forwarding_rule#name DataGoogleComputeForwardingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_forwarding_rule#id DataGoogleComputeForwardingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_forwarding_rule#project DataGoogleComputeForwardingRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region where the regional forwarding rule resides.
	//
	// This field is not applicable to global forwarding rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/compute_forwarding_rule#region DataGoogleComputeForwardingRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

