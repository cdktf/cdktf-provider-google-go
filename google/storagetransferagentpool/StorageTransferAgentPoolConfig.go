package storagetransferagentpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferAgentPoolConfig struct {
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
	// The ID of the agent pool to create.
	//
	// The agentPoolId must meet the following requirements:
	// Length of 128 characters or less.
	// Not start with the string goog.
	// Start with a lowercase ASCII character, followed by:
	// Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
	// One or more numerals or lowercase ASCII characters.
	//
	// As expressed by the regular expression: ^(?!goog)[a-z]([a-z0-9-._~]*[a-z0-9])?$.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/storage_transfer_agent_pool#name StorageTransferAgentPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// bandwidth_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/storage_transfer_agent_pool#bandwidth_limit StorageTransferAgentPool#bandwidth_limit}
	BandwidthLimit *StorageTransferAgentPoolBandwidthLimit `field:"optional" json:"bandwidthLimit" yaml:"bandwidthLimit"`
	// Specifies the client-specified AgentPool description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/storage_transfer_agent_pool#display_name StorageTransferAgentPool#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/storage_transfer_agent_pool#id StorageTransferAgentPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/storage_transfer_agent_pool#project StorageTransferAgentPool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/storage_transfer_agent_pool#timeouts StorageTransferAgentPool#timeouts}
	Timeouts *StorageTransferAgentPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

