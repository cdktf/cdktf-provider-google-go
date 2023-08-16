package eventarcgooglechannelconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventarcGoogleChannelConfigConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_google_channel_config#location EventarcGoogleChannelConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required. The resource name of the config. Must be in the format of, `projects/{project}/locations/{location}/googleChannelConfig`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_google_channel_config#name EventarcGoogleChannelConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional.
	//
	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*\/locations/*\/keyRings/*\/cryptoKeys/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_google_channel_config#crypto_key_name EventarcGoogleChannelConfig#crypto_key_name}
	CryptoKeyName *string `field:"optional" json:"cryptoKeyName" yaml:"cryptoKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_google_channel_config#id EventarcGoogleChannelConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_google_channel_config#project EventarcGoogleChannelConfig#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_google_channel_config#timeouts EventarcGoogleChannelConfig#timeouts}
	Timeouts *EventarcGoogleChannelConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

