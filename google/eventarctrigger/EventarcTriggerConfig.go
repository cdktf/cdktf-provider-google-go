package eventarctrigger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventarcTriggerConfig struct {
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
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#destination EventarcTrigger#destination}
	Destination *EventarcTriggerDestination `field:"required" json:"destination" yaml:"destination"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#location EventarcTrigger#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// matching_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#matching_criteria EventarcTrigger#matching_criteria}
	MatchingCriteria interface{} `field:"required" json:"matchingCriteria" yaml:"matchingCriteria"`
	// Required. The resource name of the trigger. Must be unique within the location on the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#name EventarcTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional.
	//
	// The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#channel EventarcTrigger#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// Optional.
	//
	// EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#event_data_content_type EventarcTrigger#event_data_content_type}
	EventDataContentType *string `field:"optional" json:"eventDataContentType" yaml:"eventDataContentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#id EventarcTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#labels EventarcTrigger#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#project EventarcTrigger#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional.
	//
	// The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#service_account EventarcTrigger#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#timeouts EventarcTrigger#timeouts}
	Timeouts *EventarcTriggerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// transport block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#transport EventarcTrigger#transport}
	Transport *EventarcTriggerTransport `field:"optional" json:"transport" yaml:"transport"`
}

