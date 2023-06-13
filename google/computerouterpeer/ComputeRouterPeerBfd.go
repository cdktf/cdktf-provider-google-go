package computerouterpeer


type ComputeRouterPeerBfd struct {
	// The BFD session initialization mode for this BGP peer.
	//
	// If set to 'ACTIVE', the Cloud Router will initiate the BFD session
	// for this BGP peer. If set to 'PASSIVE', the Cloud Router will wait
	// for the peer router to initiate the BFD session for this BGP peer.
	// If set to 'DISABLED', BFD is disabled for this BGP peer. Possible values: ["ACTIVE", "DISABLED", "PASSIVE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_peer#session_initialization_mode ComputeRouterPeer#session_initialization_mode}
	SessionInitializationMode *string `field:"required" json:"sessionInitializationMode" yaml:"sessionInitializationMode"`
	// The minimum interval, in milliseconds, between BFD control packets received from the peer router.
	//
	// The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the transmit interval of the other router. If set, this value
	// must be between 1000 and 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_peer#min_receive_interval ComputeRouterPeer#min_receive_interval}
	MinReceiveInterval *float64 `field:"optional" json:"minReceiveInterval" yaml:"minReceiveInterval"`
	// The minimum interval, in milliseconds, between BFD control packets transmitted to the peer router.
	//
	// The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the corresponding receive interval of the other router. If set,
	// this value must be between 1000 and 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_peer#min_transmit_interval ComputeRouterPeer#min_transmit_interval}
	MinTransmitInterval *float64 `field:"optional" json:"minTransmitInterval" yaml:"minTransmitInterval"`
	// The number of consecutive BFD packets that must be missed before BFD declares that a peer is unavailable.
	//
	// If set, the value must
	// be a value between 5 and 16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_peer#multiplier ComputeRouterPeer#multiplier}
	Multiplier *float64 `field:"optional" json:"multiplier" yaml:"multiplier"`
}

