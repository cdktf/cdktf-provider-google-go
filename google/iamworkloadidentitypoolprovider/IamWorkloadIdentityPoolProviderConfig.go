package iamworkloadidentitypoolprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkloadIdentityPoolProviderConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID used for the pool, which is the final component of the pool resource name.
	//
	// This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// 'gcp-' is reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#workload_identity_pool_id IamWorkloadIdentityPoolProvider#workload_identity_pool_id}
	WorkloadIdentityPoolId *string `field:"required" json:"workloadIdentityPoolId" yaml:"workloadIdentityPoolId"`
	// The ID for the provider, which becomes the final component of the resource name.
	//
	// This
	// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// 'gcp-' is reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#workload_identity_pool_provider_id IamWorkloadIdentityPoolProvider#workload_identity_pool_provider_id}
	WorkloadIdentityPoolProviderId *string `field:"required" json:"workloadIdentityPoolProviderId" yaml:"workloadIdentityPoolProviderId"`
	// [A Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted.
	//
	// The expression must output a boolean representing whether to allow the federation.
	//
	// The following keywords may be referenced in the expressions:
	// 'assertion': JSON representing the authentication credential issued by the provider.
	// 'google': The Google attributes mapped from the assertion in the 'attribute_mappings'.
	// 'attribute': The custom attributes mapped from the assertion in the 'attribute_mappings'.
	//
	// The maximum length of the attribute condition expression is 4096 characters. If
	// unspecified, all valid authentication credential are accepted.
	//
	// The following example shows how to only allow credentials with a mapped 'google.groups'
	// value of 'admins':
	// ```
	// "'admins' in google.groups"
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#attribute_condition IamWorkloadIdentityPoolProvider#attribute_condition}
	AttributeCondition *string `field:"optional" json:"attributeCondition" yaml:"attributeCondition"`
	// Maps attributes from authentication credentials issued by an external identity provider to Google Cloud attributes, such as 'subject' and 'segment'.
	//
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	//
	// The following keys are supported:
	// 'google.subject': The principal IAM is authenticating. You can reference this value
	// in IAM bindings. This is also the subject that appears in Cloud Logging logs.
	// Cannot exceed 127 characters.
	// 'google.groups': Groups the external identity belongs to. You can grant groups
	// access to resources using an IAM 'principalSet' binding; access applies to all
	// members of the group.
	//
	// You can also provide custom attributes by specifying 'attribute.{custom_attribute}',
	// where '{custom_attribute}' is the name of the custom attribute to be mapped. You can
	// define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
	// is 100 characters, and the key may only contain the characters [a-z0-9_].
	//
	// You can reference these attributes in IAM policies to define fine-grained access for a
	// workload to Google Cloud resources. For example:
	// 'google.subject':
	// 'principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}'
	// 'google.groups':
	// 'principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}'
	// 'attribute.{custom_attribute}':
	// 'principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}'
	//
	// Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	// function that maps an identity provider credential to the normalized attribute specified
	// by the corresponding map key.
	//
	// You can use the 'assertion' keyword in the expression to access a JSON representation of
	// the authentication credential issued by the provider.
	//
	// The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	// the total size of all mapped attributes must not exceed 8KB.
	//
	// For AWS providers, the following rules apply:
	// - If no attribute mapping is defined, the following default mapping applies:
	// ```
	// {
	//   "google.subject":"assertion.arn",
	//   "attribute.aws_role":
	//     "assertion.arn.contains('assumed-role')"
	//     " ? assertion.arn.extract('{account_arn}assumed-role/')"
	//     "   + 'assumed-role/'"
	//     "   + assertion.arn.extract('assumed-role/{role_name}/')"
	//     " : assertion.arn",
	// }
	// ```
	// - If any custom attribute mappings are defined, they must include a mapping to the
	// 'google.subject' attribute.
	//
	// For OIDC providers, the following rules apply:
	// - Custom attribute mappings must be defined, and must include a mapping to the
	// 'google.subject' attribute. For example, the following maps the 'sub' claim of the
	// incoming credential to the 'subject' attribute on a Google token.
	// ```
	// {"google.subject": "assertion.sub"}
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#attribute_mapping IamWorkloadIdentityPoolProvider#attribute_mapping}
	AttributeMapping *map[string]*string `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#aws IamWorkloadIdentityPoolProvider#aws}
	Aws *IamWorkloadIdentityPoolProviderAws `field:"optional" json:"aws" yaml:"aws"`
	// A description for the provider. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#description IamWorkloadIdentityPoolProvider#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#disabled IamWorkloadIdentityPoolProvider#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A display name for the provider. Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#display_name IamWorkloadIdentityPoolProvider#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#id IamWorkloadIdentityPoolProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#oidc IamWorkloadIdentityPoolProvider#oidc}
	Oidc *IamWorkloadIdentityPoolProviderOidc `field:"optional" json:"oidc" yaml:"oidc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#project IamWorkloadIdentityPoolProvider#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_workload_identity_pool_provider#timeouts IamWorkloadIdentityPoolProvider#timeouts}
	Timeouts *IamWorkloadIdentityPoolProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

