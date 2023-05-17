package sourcereporepository


type SourcerepoRepositoryPubsubConfigs struct {
	// The format of the Cloud Pub/Sub messages.
	//
	// - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
	// - JSON: The message payload is a JSON string of SourceRepoEvent. Possible values: ["PROTOBUF", "JSON"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/sourcerepo_repository#message_format SourcerepoRepository#message_format}
	MessageFormat *string `field:"required" json:"messageFormat" yaml:"messageFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/sourcerepo_repository#topic SourcerepoRepository#topic}.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Email address of the service account used for publishing Cloud Pub/Sub messages.
	//
	// This service account needs to be in the same project as the PubsubConfig. When added,
	// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
	// If unspecified, it defaults to the compute engine default service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/sourcerepo_repository#service_account_email SourcerepoRepository#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

