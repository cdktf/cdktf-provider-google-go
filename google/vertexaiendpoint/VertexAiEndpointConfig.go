// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiEndpointConfig struct {
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
	// Required.
	//
	// The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#display_name VertexAiEndpoint#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#location VertexAiEndpoint#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the Endpoint.
	//
	// The name must be numeric with no leading zeros and can be at most 10 digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#name VertexAiEndpoint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If true, the endpoint will be exposed through a dedicated DNS [Endpoint.dedicated_endpoint_dns]. Your request to the dedicated DNS will be isolated from other users' traffic and will have better performance and reliability. Note: Once you enabled dedicated endpoint, you won't be able to send request to the shared DNS {region}-aiplatform.googleapis.com. The limitation will be removed soon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#dedicated_endpoint_enabled VertexAiEndpoint#dedicated_endpoint_enabled}
	DedicatedEndpointEnabled interface{} `field:"optional" json:"dedicatedEndpointEnabled" yaml:"dedicatedEndpointEnabled"`
	// The description of the Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#description VertexAiEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#encryption_spec VertexAiEndpoint#encryption_spec}
	EncryptionSpec *VertexAiEndpointEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#id VertexAiEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels with user-defined metadata to organize your Endpoints.
	//
	// Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#labels VertexAiEndpoint#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is network name. Only one of the fields, 'network' or 'privateServiceConnectConfig', can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#network VertexAiEndpoint#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// predict_request_response_logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#predict_request_response_logging_config VertexAiEndpoint#predict_request_response_logging_config}
	PredictRequestResponseLoggingConfig *VertexAiEndpointPredictRequestResponseLoggingConfig `field:"optional" json:"predictRequestResponseLoggingConfig" yaml:"predictRequestResponseLoggingConfig"`
	// private_service_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#private_service_connect_config VertexAiEndpoint#private_service_connect_config}
	PrivateServiceConnectConfig *VertexAiEndpointPrivateServiceConnectConfig `field:"optional" json:"privateServiceConnectConfig" yaml:"privateServiceConnectConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#project VertexAiEndpoint#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#region VertexAiEndpoint#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#timeouts VertexAiEndpoint#timeouts}
	Timeouts *VertexAiEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A map from a DeployedModel's id to the percentage of this Endpoint's traffic that should be forwarded to that DeployedModel.
	//
	// If a DeployedModel's id is not listed in this map, then it receives no traffic.
	// The traffic percentage values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment. See
	// the 'deployModel' [example](https://cloud.google.com/vertex-ai/docs/general/deployment#deploy_a_model_to_an_endpoint) and
	// [documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints/deployModel) for more information.
	//
	// ~> **Note:** To set the map to empty, set '"{}"', apply, and then remove the field from your config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/vertex_ai_endpoint#traffic_split VertexAiEndpoint#traffic_split}
	TrafficSplit *string `field:"optional" json:"trafficSplit" yaml:"trafficSplit"`
}

