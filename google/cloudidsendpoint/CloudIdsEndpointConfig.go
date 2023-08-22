package cloudidsendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudIdsEndpointConfig struct {
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
	// The location for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#location CloudIdsEndpoint#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the endpoint in the format projects/{project_id}/locations/{locationId}/endpoints/{endpointId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#name CloudIdsEndpoint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the VPC network that is connected to the IDS endpoint.
	//
	// This can either contain the VPC network name itself (like "src-net") or the full URL to the network (like "projects/{project_id}/global/networks/src-net").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#network CloudIdsEndpoint#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The minimum alert severity level that is reported by the endpoint. Possible values: ["INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "CRITICAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#severity CloudIdsEndpoint#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// An optional description of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#description CloudIdsEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#id CloudIdsEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#project CloudIdsEndpoint#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#threat_exceptions CloudIdsEndpoint#threat_exceptions}
	ThreatExceptions *[]*string `field:"optional" json:"threatExceptions" yaml:"threatExceptions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_ids_endpoint#timeouts CloudIdsEndpoint#timeouts}
	Timeouts *CloudIdsEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

