package datafusioninstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFusionInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#name DataFusionInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Represents the type of Data Fusion instance.
	//
	// Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	// using point and click UI. However, there are certain limitations, such as fewer number
	// of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	// available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	// with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	// pipelines at low cost. Possible values: ["BASIC", "ENTERPRISE", "DEVELOPER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#type DataFusionInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#accelerators DataFusionInstance#accelerators}
	Accelerators interface{} `field:"optional" json:"accelerators" yaml:"accelerators"`
	// crypto_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#crypto_key_config DataFusionInstance#crypto_key_config}
	CryptoKeyConfig *DataFusionInstanceCryptoKeyConfig `field:"optional" json:"cryptoKeyConfig" yaml:"cryptoKeyConfig"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#dataproc_service_account DataFusionInstance#dataproc_service_account}
	DataprocServiceAccount *string `field:"optional" json:"dataprocServiceAccount" yaml:"dataprocServiceAccount"`
	// An optional description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#description DataFusionInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#display_name DataFusionInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Option to enable granular role-based access control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#enable_rbac DataFusionInstance#enable_rbac}
	EnableRbac interface{} `field:"optional" json:"enableRbac" yaml:"enableRbac"`
	// Option to enable Stackdriver Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#enable_stackdriver_logging DataFusionInstance#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#enable_stackdriver_monitoring DataFusionInstance#enable_stackdriver_monitoring}
	EnableStackdriverMonitoring interface{} `field:"optional" json:"enableStackdriverMonitoring" yaml:"enableStackdriverMonitoring"`
	// event_publish_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#event_publish_config DataFusionInstance#event_publish_config}
	EventPublishConfig *DataFusionInstanceEventPublishConfig `field:"optional" json:"eventPublishConfig" yaml:"eventPublishConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#id DataFusionInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#labels DataFusionInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#network_config DataFusionInstance#network_config}
	NetworkConfig *DataFusionInstanceNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#options DataFusionInstance#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Specifies whether the Data Fusion instance should be private.
	//
	// If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#private_instance DataFusionInstance#private_instance}
	PrivateInstance interface{} `field:"optional" json:"privateInstance" yaml:"privateInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#project DataFusionInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the Data Fusion instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#region DataFusionInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#timeouts DataFusionInstance#timeouts}
	Timeouts *DataFusionInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Current version of the Data Fusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#version DataFusionInstance#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#zone DataFusionInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

