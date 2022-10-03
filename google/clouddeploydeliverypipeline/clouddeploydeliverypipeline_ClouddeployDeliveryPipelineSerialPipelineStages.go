package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStages struct {
	// Skaffold profiles to use when rendering the manifest for this stage's `Target`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/clouddeploy_delivery_pipeline#profiles ClouddeployDeliveryPipeline#profiles}
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
	// The target_id to which this stage points.
	//
	// This field refers exclusively to the last segment of a target name. For example, this field would just be `my-target` (rather than `projects/project/locations/location/targets/my-target`). The location of the `Target` is inferred to be the same as the location of the `DeliveryPipeline` that contains this `Stage`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/clouddeploy_delivery_pipeline#target_id ClouddeployDeliveryPipeline#target_id}
	TargetId *string `field:"optional" json:"targetId" yaml:"targetId"`
}

