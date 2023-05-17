package cloudiotdevice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudiotDeviceConfig struct {
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
	// A unique name for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#name CloudiotDevice#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the device registry where this device should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#registry CloudiotDevice#registry}
	Registry *string `field:"required" json:"registry" yaml:"registry"`
	// If a device is blocked, connections or requests from this device will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#blocked CloudiotDevice#blocked}
	Blocked interface{} `field:"optional" json:"blocked" yaml:"blocked"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#credentials CloudiotDevice#credentials}
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// gateway_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#gateway_config CloudiotDevice#gateway_config}
	GatewayConfig *CloudiotDeviceGatewayConfig `field:"optional" json:"gatewayConfig" yaml:"gatewayConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#id CloudiotDevice#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The logging verbosity for device activity. Possible values: ["NONE", "ERROR", "INFO", "DEBUG"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#log_level CloudiotDevice#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The metadata key-value pairs assigned to the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#metadata CloudiotDevice#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudiot_device#timeouts CloudiotDevice#timeouts}
	Timeouts *CloudiotDeviceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

