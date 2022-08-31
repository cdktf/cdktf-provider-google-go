// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigqueryConnectionAwsAccessRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection#iam_role_id BigqueryConnection#iam_role_id}
	IamRoleId *string `field:"required" json:"iamRoleId" yaml:"iamRoleId"`
}

