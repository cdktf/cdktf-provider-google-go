package iamworkloadidentitypoolprovider


type IamWorkloadIdentityPoolProviderAws struct {
	// The AWS account ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_workload_identity_pool_provider#account_id IamWorkloadIdentityPoolProvider#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

