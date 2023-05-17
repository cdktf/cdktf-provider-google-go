package accesscontextmanagerauthorizedorgsdesc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerAuthorizedOrgsDescConfig struct {
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
	// Resource name for the 'AuthorizedOrgsDesc'.
	//
	// Format:
	// 'accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}'.
	// The 'authorized_orgs_desc' component must begin with a letter, followed by
	// alphanumeric characters or '_'.
	// After you create an 'AuthorizedOrgsDesc', you cannot change its 'name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#name AccessContextManagerAuthorizedOrgsDesc#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. Resource name for the access policy which owns this 'AuthorizedOrgsDesc'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#parent AccessContextManagerAuthorizedOrgsDesc#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// The type of entities that need to use the authorization relationship during evaluation, such as a device.
	//
	// Valid values are "ASSET_TYPE_DEVICE" and
	// "ASSET_TYPE_CREDENTIAL_STRENGTH". Possible values: ["ASSET_TYPE_DEVICE", "ASSET_TYPE_CREDENTIAL_STRENGTH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#asset_type AccessContextManagerAuthorizedOrgsDesc#asset_type}
	AssetType *string `field:"optional" json:"assetType" yaml:"assetType"`
	// The direction of the authorization relationship between this organization and the organizations listed in the "orgs" field.
	//
	// The valid values for this
	// field include the following:
	//
	// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
	// in the organizations listed in the 'orgs' field.
	//
	// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the 'orgs'
	// field to evaluate the traffic in this organization.
	//
	// For the authorization relationship to take effect, all of the organizations
	// must authorize and specify the appropriate relationship direction. For
	// example, if organization A authorized organization B and C to evaluate its
	// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
	// direction, organizations B and C must specify
	// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
	// "AuthorizedOrgsDesc" resource. Possible values: ["AUTHORIZATION_DIRECTION_TO", "AUTHORIZATION_DIRECTION_FROM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#authorization_direction AccessContextManagerAuthorizedOrgsDesc#authorization_direction}
	AuthorizationDirection *string `field:"optional" json:"authorizationDirection" yaml:"authorizationDirection"`
	// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST". Possible values: ["AUTHORIZATION_TYPE_TRUST"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#authorization_type AccessContextManagerAuthorizedOrgsDesc#authorization_type}
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#id AccessContextManagerAuthorizedOrgsDesc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The list of organization ids in this AuthorizedOrgsDesc. Format: 'organizations/<org_number>' Example: 'organizations/123456'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#orgs AccessContextManagerAuthorizedOrgsDesc#orgs}
	Orgs *[]*string `field:"optional" json:"orgs" yaml:"orgs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_authorized_orgs_desc#timeouts AccessContextManagerAuthorizedOrgsDesc#timeouts}
	Timeouts *AccessContextManagerAuthorizedOrgsDescTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

