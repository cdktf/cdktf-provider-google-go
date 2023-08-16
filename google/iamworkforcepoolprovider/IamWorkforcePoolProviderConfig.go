package iamworkforcepoolprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolProviderConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#location IamWorkforcePoolProvider#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID for the provider, which becomes the final component of the resource name.
	//
	// This value must be 4-32 characters, and may contain the characters [a-z0-9-].
	// The prefix 'gcp-' is reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#provider_id IamWorkforcePoolProvider#provider_id}
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// The ID to use for the pool, which becomes the final component of the resource name.
	//
	// The IDs must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens.
	// It must start with a letter, and cannot have a trailing hyphen.
	// The prefix 'gcp-' is reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#workforce_pool_id IamWorkforcePoolProvider#workforce_pool_id}
	WorkforcePoolId *string `field:"required" json:"workforcePoolId" yaml:"workforcePoolId"`
	// A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted.
	//
	// The expression must output a boolean representing whether to allow the federation.
	//
	// The following keywords may be referenced in the expressions:
	// 'assertion': JSON representing the authentication credential issued by the provider.
	// 'google': The Google attributes mapped from the assertion in the 'attribute_mappings'.
	// 'google.profile_photo' and 'google.display_name' are not supported.
	// 'attribute': The custom attributes mapped from the assertion in the 'attribute_mappings'.
	//
	// The maximum length of the attribute condition expression is 4096 characters.
	// If unspecified, all valid authentication credentials will be accepted.
	//
	// The following example shows how to only allow credentials with a mapped 'google.groups' value of 'admins':
	// ```
	// "'admins' in google.groups"
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#attribute_condition IamWorkforcePoolProvider#attribute_condition}
	AttributeCondition *string `field:"optional" json:"attributeCondition" yaml:"attributeCondition"`
	// Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as 'subject' and 'segment'.
	//
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	//
	// The following keys are supported:
	// 'google.subject': The principal IAM is authenticating. You can reference this value in IAM bindings.
	// This is also the subject that appears in Cloud Logging logs. This is a required field and
	// the mapped subject cannot exceed 127 bytes.
	// 'google.groups': Groups the authenticating user belongs to. You can grant groups access to
	// resources using an IAM 'principalSet' binding; access applies to all members of the group.
	// 'google.display_name': The name of the authenticated user. This is an optional field and
	// the mapped display name cannot exceed 100 bytes. If not set, 'google.subject' will be displayed instead.
	// This attribute cannot be referenced in IAM bindings.
	// 'google.profile_photo': The URL that specifies the authenticated user's thumbnail photo.
	// This is an optional field. When set, the image will be visible as the user's profile picture.
	// If not set, a generic user icon will be displayed instead.
	// This attribute cannot be referenced in IAM bindings.
	//
	// You can also provide custom attributes by specifying 'attribute.{custom_attribute}', where {custom_attribute}
	// is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes.
	// The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_].
	//
	// You can reference these attributes in IAM policies to define fine-grained access for a workforce pool
	// to Google Cloud resources. For example:
	// 'google.subject':
	// 'principal://iam.googleapis.com/locations/{location}/workforcePools/{pool}/subject/{value}'
	// 'google.groups':
	// 'principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/group/{value}'
	// 'attribute.{custom_attribute}':
	// 'principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/attribute.{custom_attribute}/{value}'
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
	// For OIDC providers, you must supply a custom mapping that includes the 'google.subject' attribute.
	// For example, the following maps the sub claim of the incoming credential to the 'subject' attribute
	// on a Google token:
	// ```
	// {"google.subject": "assertion.sub"}
	// ```
	//
	// An object containing a list of '"key": value' pairs.
	// Example: '{ "name": "wrench", "mass": "1.3kg", "count": "3" }'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#attribute_mapping IamWorkforcePoolProvider#attribute_mapping}
	AttributeMapping *map[string]*string `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// A user-specified description of the provider. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#description IamWorkforcePoolProvider#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#disabled IamWorkforcePoolProvider#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A user-specified display name for the provider. Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#display_name IamWorkforcePoolProvider#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#id IamWorkforcePoolProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#oidc IamWorkforcePoolProvider#oidc}
	Oidc *IamWorkforcePoolProviderOidc `field:"optional" json:"oidc" yaml:"oidc"`
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#saml IamWorkforcePoolProvider#saml}
	Saml *IamWorkforcePoolProviderSaml `field:"optional" json:"saml" yaml:"saml"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_workforce_pool_provider#timeouts IamWorkforcePoolProvider#timeouts}
	Timeouts *IamWorkforcePoolProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

