// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcmessagebus

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventarcMessageBusConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#location EventarcMessageBus#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required. The user-provided ID to be assigned to the MessageBus. It should match the format '^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#message_bus_id EventarcMessageBus#message_bus_id}
	MessageBusId *string `field:"required" json:"messageBusId" yaml:"messageBusId"`
	// Optional. Resource annotations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#annotations EventarcMessageBus#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data.
	//
	// It must match the pattern
	// 'projects/* /locations/* /keyRings/* /cryptoKeys/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#crypto_key_name EventarcMessageBus#crypto_key_name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	CryptoKeyName *string `field:"optional" json:"cryptoKeyName" yaml:"cryptoKeyName"`
	// Optional. Resource display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#display_name EventarcMessageBus#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#id EventarcMessageBus#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Resource labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#labels EventarcMessageBus#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#logging_config EventarcMessageBus#logging_config}
	LoggingConfig *EventarcMessageBusLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#project EventarcMessageBus#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/eventarc_message_bus#timeouts EventarcMessageBus#timeouts}
	Timeouts *EventarcMessageBusTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

