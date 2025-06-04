// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationConnectorsConnectionConfig struct {
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
	// connectorVersion of the Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#connector_version IntegrationConnectorsConnection#connector_version}
	ConnectorVersion *string `field:"required" json:"connectorVersion" yaml:"connectorVersion"`
	// Location in which Connection needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#location IntegrationConnectorsConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of Connection needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#name IntegrationConnectorsConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#auth_config IntegrationConnectorsConnection#auth_config}
	AuthConfig *IntegrationConnectorsConnectionAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// config_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#config_variable IntegrationConnectorsConnection#config_variable}
	ConfigVariable interface{} `field:"optional" json:"configVariable" yaml:"configVariable"`
	// An arbitrary description for the Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#description IntegrationConnectorsConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#destination_config IntegrationConnectorsConnection#destination_config}
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// eventing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#eventing_config IntegrationConnectorsConnection#eventing_config}
	EventingConfig *IntegrationConnectorsConnectionEventingConfig `field:"optional" json:"eventingConfig" yaml:"eventingConfig"`
	// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION", "ONLY_EVENTING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#eventing_enablement_type IntegrationConnectorsConnection#eventing_enablement_type}
	EventingEnablementType *string `field:"optional" json:"eventingEnablementType" yaml:"eventingEnablementType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#id IntegrationConnectorsConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#labels IntegrationConnectorsConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// lock_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#lock_config IntegrationConnectorsConnection#lock_config}
	LockConfig *IntegrationConnectorsConnectionLockConfig `field:"optional" json:"lockConfig" yaml:"lockConfig"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#log_config IntegrationConnectorsConnection#log_config}
	LogConfig *IntegrationConnectorsConnectionLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#node_config IntegrationConnectorsConnection#node_config}
	NodeConfig *IntegrationConnectorsConnectionNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#project IntegrationConnectorsConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Service account needed for runtime plane to access Google Cloud resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#service_account IntegrationConnectorsConnection#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// ssl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#ssl_config IntegrationConnectorsConnection#ssl_config}
	SslConfig *IntegrationConnectorsConnectionSslConfig `field:"optional" json:"sslConfig" yaml:"sslConfig"`
	// Suspended indicates if a user has suspended a connection or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#suspended IntegrationConnectorsConnection#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/integration_connectors_connection#timeouts IntegrationConnectorsConnection#timeouts}
	Timeouts *IntegrationConnectorsConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

