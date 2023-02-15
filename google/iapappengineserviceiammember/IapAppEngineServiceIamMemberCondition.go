package iapappengineserviceiammember


type IapAppEngineServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_app_engine_service_iam_member#expression IapAppEngineServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_app_engine_service_iam_member#title IapAppEngineServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_app_engine_service_iam_member#description IapAppEngineServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

