package clouddeploytarget


type ClouddeployTargetAnthosCluster struct {
	// Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/clouddeploy_target#membership ClouddeployTarget#membership}
	Membership *string `field:"optional" json:"membership" yaml:"membership"`
}

