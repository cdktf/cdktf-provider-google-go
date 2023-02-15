package iapclient

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IapClientConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Identifier of the brand to which this client is attached to. The format is 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_client#brand IapClient#brand}
	Brand *string `field:"required" json:"brand" yaml:"brand"`
	// Human-friendly name given to the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_client#display_name IapClient#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_client#id IapClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_client#timeouts IapClient#timeouts}
	Timeouts *IapClientTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

