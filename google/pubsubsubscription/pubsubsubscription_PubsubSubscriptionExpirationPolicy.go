package pubsubsubscription


type PubsubSubscriptionExpirationPolicy struct {
	// Specifies the "time-to-live" duration for an associated resource.
	//
	// The
	// resource expires if it is not active for a period of ttl.
	// If ttl is not set, the associated resource never expires.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// Example - "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_subscription#ttl PubsubSubscription#ttl}
	Ttl *string `field:"required" json:"ttl" yaml:"ttl"`
}

