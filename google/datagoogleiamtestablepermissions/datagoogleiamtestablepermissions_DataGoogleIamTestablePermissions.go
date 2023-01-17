package datagoogleiamtestablepermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v5/datagoogleiamtestablepermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/google/d/iam_testable_permissions google_iam_testable_permissions}.
type DataGoogleIamTestablePermissions interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CustomSupportLevel() *string
	SetCustomSupportLevel(val *string)
	CustomSupportLevelInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullResourceName() *string
	SetFullResourceName(val *string)
	FullResourceNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Permissions() DataGoogleIamTestablePermissionsPermissionsList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Stages() *[]*string
	SetStages(val *[]*string)
	StagesInput() *[]*string
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
	ResetCustomSupportLevel()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStages()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGoogleIamTestablePermissions
type jsiiProxy_DataGoogleIamTestablePermissions struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) CustomSupportLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSupportLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) CustomSupportLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSupportLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) FullResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) FullResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Permissions() DataGoogleIamTestablePermissionsPermissionsList {
	var returns DataGoogleIamTestablePermissionsPermissionsList
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) StagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google/d/iam_testable_permissions google_iam_testable_permissions} Data Source.
func NewDataGoogleIamTestablePermissions(scope constructs.Construct, id *string, config *DataGoogleIamTestablePermissionsConfig) DataGoogleIamTestablePermissions {
	_init_.Initialize()

	if err := validateNewDataGoogleIamTestablePermissionsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleIamTestablePermissions{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleIamTestablePermissions.DataGoogleIamTestablePermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google/d/iam_testable_permissions google_iam_testable_permissions} Data Source.
func NewDataGoogleIamTestablePermissions_Override(d DataGoogleIamTestablePermissions, scope constructs.Construct, id *string, config *DataGoogleIamTestablePermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleIamTestablePermissions.DataGoogleIamTestablePermissions",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetCustomSupportLevel(val *string) {
	if err := j.validateSetCustomSupportLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSupportLevel",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetFullResourceName(val *string) {
	if err := j.validateSetFullResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullResourceName",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamTestablePermissions)SetStages(val *[]*string) {
	if err := j.validateSetStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stages",
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
func DataGoogleIamTestablePermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleIamTestablePermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleIamTestablePermissions.DataGoogleIamTestablePermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleIamTestablePermissions_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleIamTestablePermissions_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleIamTestablePermissions.DataGoogleIamTestablePermissions",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleIamTestablePermissions_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleIamTestablePermissions_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleIamTestablePermissions.DataGoogleIamTestablePermissions",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleIamTestablePermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleIamTestablePermissions.DataGoogleIamTestablePermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleIamTestablePermissions) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) ResetCustomSupportLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSupportLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) ResetStages() {
	_jsii_.InvokeVoid(
		d,
		"resetStages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleIamTestablePermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

