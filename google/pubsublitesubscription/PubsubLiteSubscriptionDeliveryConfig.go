package pubsublitesubscription


type PubsubLiteSubscriptionDeliveryConfig struct {
	// When this subscription should send messages to subscribers relative to messages persistence in storage. Possible values: ["DELIVER_IMMEDIATELY", "DELIVER_AFTER_STORED", "DELIVERY_REQUIREMENT_UNSPECIFIED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/pubsub_lite_subscription#delivery_requirement PubsubLiteSubscription#delivery_requirement}
	DeliveryRequirement *string `field:"required" json:"deliveryRequirement" yaml:"deliveryRequirement"`
}

