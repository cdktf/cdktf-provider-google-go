package lookerinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LookerInstanceConfig struct {
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
	// The ID of the instance or a fully qualified identifier for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#name LookerInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// admin_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#admin_settings LookerInstance#admin_settings}
	AdminSettings *LookerInstanceAdminSettings `field:"optional" json:"adminSettings" yaml:"adminSettings"`
	// Network name in the consumer project in the format of: projects/{project}/global/networks/{network} Note that the consumer network may be in a different GCP project than the consumer project that is hosting the Looker Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#consumer_network LookerInstance#consumer_network}
	ConsumerNetwork *string `field:"optional" json:"consumerNetwork" yaml:"consumerNetwork"`
	// deny_maintenance_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#deny_maintenance_period LookerInstance#deny_maintenance_period}
	DenyMaintenancePeriod *LookerInstanceDenyMaintenancePeriod `field:"optional" json:"denyMaintenancePeriod" yaml:"denyMaintenancePeriod"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#encryption_config LookerInstance#encryption_config}
	EncryptionConfig *LookerInstanceEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#id LookerInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#maintenance_window LookerInstance#maintenance_window}
	MaintenanceWindow *LookerInstanceMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// oauth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#oauth_config LookerInstance#oauth_config}
	OauthConfig *LookerInstanceOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// Platform editions for a Looker instance.
	//
	// Each edition maps to a set of instance features, like its size. Must be one of these values:
	// - LOOKER_CORE_TRIAL: trial instance
	// - LOOKER_CORE_STANDARD: pay as you go standard instance
	// - LOOKER_CORE_STANDARD_ANNUAL: subscription standard instance
	// - LOOKER_CORE_ENTERPRISE_ANNUAL: subscription enterprise instance
	// - LOOKER_CORE_EMBED_ANNUAL: subscription embed instance
	// - LOOKER_MODELER: standalone modeling service Default value: "LOOKER_CORE_TRIAL" Possible values: ["LOOKER_CORE_TRIAL", "LOOKER_CORE_STANDARD", "LOOKER_CORE_STANDARD_ANNUAL", "LOOKER_CORE_ENTERPRISE_ANNUAL", "LOOKER_CORE_EMBED_ANNUAL", "LOOKER_MODELER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#platform_edition LookerInstance#platform_edition}
	PlatformEdition *string `field:"optional" json:"platformEdition" yaml:"platformEdition"`
	// Whether private IP is enabled on the Looker instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#private_ip_enabled LookerInstance#private_ip_enabled}
	PrivateIpEnabled interface{} `field:"optional" json:"privateIpEnabled" yaml:"privateIpEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#project LookerInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Whether public IP is enabled on the Looker instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#public_ip_enabled LookerInstance#public_ip_enabled}
	PublicIpEnabled interface{} `field:"optional" json:"publicIpEnabled" yaml:"publicIpEnabled"`
	// The name of the Looker region of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#region LookerInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Name of a reserved IP address range within the consumer network, to be used for private service access connection.
	//
	// User may or may not specify this in a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#reserved_range LookerInstance#reserved_range}
	ReservedRange *string `field:"optional" json:"reservedRange" yaml:"reservedRange"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#timeouts LookerInstance#timeouts}
	Timeouts *LookerInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#user_metadata LookerInstance#user_metadata}
	UserMetadata *LookerInstanceUserMetadata `field:"optional" json:"userMetadata" yaml:"userMetadata"`
}

