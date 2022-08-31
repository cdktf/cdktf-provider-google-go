// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type IapTunnelInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_tunnel_instance_iam_binding#expression IapTunnelInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_tunnel_instance_iam_binding#title IapTunnelInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_tunnel_instance_iam_binding#description IapTunnelInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

