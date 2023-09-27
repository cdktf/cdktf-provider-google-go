// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudassetfolderfeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v9/cloudassetfolderfeed/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/cloud_asset_folder_feed google_cloud_asset_folder_feed}.
type CloudAssetFolderFeed interface {
	cdktf.TerraformResource
	AssetNames() *[]*string
	SetAssetNames(val *[]*string)
	AssetNamesInput() *[]*string
	AssetTypes() *[]*string
	SetAssetTypes(val *[]*string)
	AssetTypesInput() *[]*string
	BillingProject() *string
	SetBillingProject(val *string)
	BillingProjectInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Condition() CloudAssetFolderFeedConditionOutputReference
	ConditionInput() *CloudAssetFolderFeedCondition
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FeedId() *string
	SetFeedId(val *string)
	FeedIdInput() *string
	FeedOutputConfig() CloudAssetFolderFeedFeedOutputConfigOutputReference
	FeedOutputConfigInput() *CloudAssetFolderFeedFeedOutputConfig
	Folder() *string
	SetFolder(val *string)
	FolderId() *string
	FolderInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CloudAssetFolderFeedTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCondition(value *CloudAssetFolderFeedCondition)
	PutFeedOutputConfig(value *CloudAssetFolderFeedFeedOutputConfig)
	PutTimeouts(value *CloudAssetFolderFeedTimeouts)
	ResetAssetNames()
	ResetAssetTypes()
	ResetCondition()
	ResetContentType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudAssetFolderFeed
type jsiiProxy_CloudAssetFolderFeed struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudAssetFolderFeed) AssetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) AssetNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) AssetTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) AssetTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) BillingProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) BillingProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Condition() CloudAssetFolderFeedConditionOutputReference {
	var returns CloudAssetFolderFeedConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) ConditionInput() *CloudAssetFolderFeedCondition {
	var returns *CloudAssetFolderFeedCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) FeedId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) FeedIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) FeedOutputConfig() CloudAssetFolderFeedFeedOutputConfigOutputReference {
	var returns CloudAssetFolderFeedFeedOutputConfigOutputReference
	_jsii_.Get(
		j,
		"feedOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) FeedOutputConfigInput() *CloudAssetFolderFeedFeedOutputConfig {
	var returns *CloudAssetFolderFeedFeedOutputConfig
	_jsii_.Get(
		j,
		"feedOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) FolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) Timeouts() CloudAssetFolderFeedTimeoutsOutputReference {
	var returns CloudAssetFolderFeedTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetFolderFeed) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/cloud_asset_folder_feed google_cloud_asset_folder_feed} Resource.
func NewCloudAssetFolderFeed(scope constructs.Construct, id *string, config *CloudAssetFolderFeedConfig) CloudAssetFolderFeed {
	_init_.Initialize()

	if err := validateNewCloudAssetFolderFeedParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAssetFolderFeed{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudAssetFolderFeed.CloudAssetFolderFeed",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/cloud_asset_folder_feed google_cloud_asset_folder_feed} Resource.
func NewCloudAssetFolderFeed_Override(c CloudAssetFolderFeed, scope constructs.Construct, id *string, config *CloudAssetFolderFeedConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudAssetFolderFeed.CloudAssetFolderFeed",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetAssetNames(val *[]*string) {
	if err := j.validateSetAssetNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetNames",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetAssetTypes(val *[]*string) {
	if err := j.validateSetAssetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetTypes",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetBillingProject(val *string) {
	if err := j.validateSetBillingProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingProject",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetFeedId(val *string) {
	if err := j.validateSetFeedIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedId",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudAssetFolderFeed)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CloudAssetFolderFeed_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAssetFolderFeed_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudAssetFolderFeed.CloudAssetFolderFeed",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAssetFolderFeed_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAssetFolderFeed_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudAssetFolderFeed.CloudAssetFolderFeed",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAssetFolderFeed_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAssetFolderFeed_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudAssetFolderFeed.CloudAssetFolderFeed",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudAssetFolderFeed_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.cloudAssetFolderFeed.CloudAssetFolderFeed",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) PutCondition(value *CloudAssetFolderFeedCondition) {
	if err := c.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) PutFeedOutputConfig(value *CloudAssetFolderFeedFeedOutputConfig) {
	if err := c.validatePutFeedOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFeedOutputConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) PutTimeouts(value *CloudAssetFolderFeedTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) ResetAssetNames() {
	_jsii_.InvokeVoid(
		c,
		"resetAssetNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) ResetAssetTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetAssetTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) ResetCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) ResetContentType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetFolderFeed) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetFolderFeed) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

