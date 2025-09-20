// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecloudquotasquotainfo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagooglecloudquotasquotainfo/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/data-sources/cloud_quotas_quota_info google_cloud_quotas_quota_info}.
type DataGoogleCloudQuotasQuotaInfo interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerType() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Dimensions() *[]*string
	DimensionsInfos() DataGoogleCloudQuotasQuotaInfoDimensionsInfosList
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
	IsConcurrent() cdktf.IResolvable
	IsFixed() cdktf.IResolvable
	IsPrecise() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metric() *string
	MetricDisplayName() *string
	MetricUnit() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	QuotaDisplayName() *string
	QuotaId() *string
	SetQuotaId(val *string)
	QuotaIdInput() *string
	QuotaIncreaseEligibility() DataGoogleCloudQuotasQuotaInfoQuotaIncreaseEligibilityList
	// Experimental.
	RawOverrides() interface{}
	RefreshInterval() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ServiceRequestQuotaUri() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGoogleCloudQuotasQuotaInfo
type jsiiProxy_DataGoogleCloudQuotasQuotaInfo struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Dimensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) DimensionsInfos() DataGoogleCloudQuotasQuotaInfoDimensionsInfosList {
	var returns DataGoogleCloudQuotasQuotaInfoDimensionsInfosList
	_jsii_.Get(
		j,
		"dimensionsInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) IsConcurrent() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isConcurrent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) IsFixed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) IsPrecise() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isPrecise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) MetricDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) MetricUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) QuotaDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) QuotaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) QuotaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) QuotaIncreaseEligibility() DataGoogleCloudQuotasQuotaInfoQuotaIncreaseEligibilityList {
	var returns DataGoogleCloudQuotasQuotaInfoQuotaIncreaseEligibilityList
	_jsii_.Get(
		j,
		"quotaIncreaseEligibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) RefreshInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ServiceRequestQuotaUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRequestQuotaUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/data-sources/cloud_quotas_quota_info google_cloud_quotas_quota_info} Data Source.
func NewDataGoogleCloudQuotasQuotaInfo(scope constructs.Construct, id *string, config *DataGoogleCloudQuotasQuotaInfoConfig) DataGoogleCloudQuotasQuotaInfo {
	_init_.Initialize()

	if err := validateNewDataGoogleCloudQuotasQuotaInfoParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleCloudQuotasQuotaInfo{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfo.DataGoogleCloudQuotasQuotaInfo",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/data-sources/cloud_quotas_quota_info google_cloud_quotas_quota_info} Data Source.
func NewDataGoogleCloudQuotasQuotaInfo_Override(d DataGoogleCloudQuotasQuotaInfo, scope constructs.Construct, id *string, config *DataGoogleCloudQuotasQuotaInfoConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfo.DataGoogleCloudQuotasQuotaInfo",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetQuotaId(val *string) {
	if err := j.validateSetQuotaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaId",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfo)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

// Generates CDKTF code for importing a DataGoogleCloudQuotasQuotaInfo resource upon running "cdktf plan <stack-name>".
func DataGoogleCloudQuotasQuotaInfo_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleCloudQuotasQuotaInfo_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfo.DataGoogleCloudQuotasQuotaInfo",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DataGoogleCloudQuotasQuotaInfo_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleCloudQuotasQuotaInfo_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfo.DataGoogleCloudQuotasQuotaInfo",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleCloudQuotasQuotaInfo_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleCloudQuotasQuotaInfo_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfo.DataGoogleCloudQuotasQuotaInfo",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleCloudQuotasQuotaInfo_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleCloudQuotasQuotaInfo_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfo.DataGoogleCloudQuotasQuotaInfo",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleCloudQuotasQuotaInfo_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfo.DataGoogleCloudQuotasQuotaInfo",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfo) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

