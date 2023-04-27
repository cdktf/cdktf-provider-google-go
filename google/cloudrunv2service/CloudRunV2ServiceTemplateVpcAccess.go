package cloudrunv2service


type CloudRunV2ServiceTemplateVpcAccess struct {
	// VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}, where {project} can be project id or number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/cloud_run_v2_service#connector CloudRunV2Service#connector}
	Connector *string `field:"optional" json:"connector" yaml:"connector"`
	// Traffic VPC egress settings. Possible values: ["ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/cloud_run_v2_service#egress CloudRunV2Service#egress}
	Egress *string `field:"optional" json:"egress" yaml:"egress"`
}

