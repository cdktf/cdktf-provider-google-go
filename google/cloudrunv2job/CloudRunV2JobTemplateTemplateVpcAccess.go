package cloudrunv2job


type CloudRunV2JobTemplateTemplateVpcAccess struct {
	// VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}, where {project} can be project id or number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_job#connector CloudRunV2Job#connector}
	Connector *string `field:"optional" json:"connector" yaml:"connector"`
	// Traffic VPC egress settings. Possible values: ["ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_job#egress CloudRunV2Job#egress}
	Egress *string `field:"optional" json:"egress" yaml:"egress"`
}
