package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelBasicConditions struct {
	// device_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_level#device_policy AccessContextManagerAccessLevel#device_policy}
	DevicePolicy *AccessContextManagerAccessLevelBasicConditionsDevicePolicy `field:"optional" json:"devicePolicy" yaml:"devicePolicy"`
	// A list of CIDR block IP subnetwork specification.
	//
	// May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_level#ip_subnetworks AccessContextManagerAccessLevel#ip_subnetworks}
	IpSubnetworks *[]*string `field:"optional" json:"ipSubnetworks" yaml:"ipSubnetworks"`
	// An allowed list of members (users, service accounts). Using groups is not supported yet.
	//
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: 'user:{emailid}', 'serviceAccount:{emailid}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_level#members AccessContextManagerAccessLevel#members}
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
	// Whether to negate the Condition.
	//
	// If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_level#negate AccessContextManagerAccessLevel#negate}
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// The request must originate from one of the provided countries/regions. Format: A valid ISO 3166-1 alpha-2 code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_level#regions AccessContextManagerAccessLevel#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// A list of other access levels defined in the same Policy, referenced by resource name.
	//
	// Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_level#required_access_levels AccessContextManagerAccessLevel#required_access_levels}
	RequiredAccessLevels *[]*string `field:"optional" json:"requiredAccessLevels" yaml:"requiredAccessLevels"`
}

