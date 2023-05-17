package sccsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccSourceConfig struct {
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
	// The source’s display name.
	//
	// A source’s display name must be unique
	// amongst its siblings, for example, two sources with the same parent
	// can't share the same display name. The display name must start and end
	// with a letter or digit, may contain letters, digits, spaces, hyphens,
	// and underscores, and can be no longer than 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/scc_source#display_name SccSource#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The organization whose Cloud Security Command Center the Source lives in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/scc_source#organization SccSource#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The description of the source (max of 1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/scc_source#description SccSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/scc_source#id SccSource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/scc_source#timeouts SccSource#timeouts}
	Timeouts *SccSourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

