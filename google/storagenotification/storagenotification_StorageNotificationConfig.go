package storagenotification

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageNotificationConfig struct {
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
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_notification#bucket StorageNotification#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The desired content of the Payload. One of "JSON_API_V1" or "NONE".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_notification#payload_format StorageNotification#payload_format}
	PayloadFormat *string `field:"required" json:"payloadFormat" yaml:"payloadFormat"`
	// The Cloud Pub/Sub topic to which this subscription publishes.
	//
	// Expects either the  topic name, assumed to belong to the default GCP provider project, or the project-level name,  i.e. projects/my-gcp-project/topics/my-topic or my-topic. If the project is not set in the provider, you will need to use the project-level name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_notification#topic StorageNotification#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// A set of key/value attribute pairs to attach to each Cloud Pub/Sub message published for this notification subscription.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_notification#custom_attributes StorageNotification#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// List of event type filters for this notification config.
	//
	// If not specified, Cloud Storage will send notifications for all event types. The valid types are: "OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE", "OBJECT_DELETE", "OBJECT_ARCHIVE"
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_notification#event_types StorageNotification#event_types}
	EventTypes *[]*string `field:"optional" json:"eventTypes" yaml:"eventTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_notification#id StorageNotification#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies a prefix path filter for this notification config.
	//
	// Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_notification#object_name_prefix StorageNotification#object_name_prefix}
	ObjectNamePrefix *string `field:"optional" json:"objectNamePrefix" yaml:"objectNamePrefix"`
}

