package cloudrunservice


type CloudRunServiceTraffic struct {
	// Percent specifies percent of the traffic to this Revision or Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#percent CloudRunService#percent}
	Percent *float64 `field:"required" json:"percent" yaml:"percent"`
	// LatestRevision may be optionally provided to indicate that the latest ready Revision of the Configuration should be used for this traffic target.
	//
	// When
	// provided LatestRevision must be true if RevisionName is empty; it must be
	// false when RevisionName is non-empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#latest_revision CloudRunService#latest_revision}
	LatestRevision interface{} `field:"optional" json:"latestRevision" yaml:"latestRevision"`
	// RevisionName of a specific revision to which to send this portion of traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#revision_name CloudRunService#revision_name}
	RevisionName *string `field:"optional" json:"revisionName" yaml:"revisionName"`
	// Tag is optionally used to expose a dedicated url for referencing this target exclusively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#tag CloudRunService#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

