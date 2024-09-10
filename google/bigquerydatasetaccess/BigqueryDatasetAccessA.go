// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatasetaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/bigquerydatasetaccess/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/bigquery_dataset_access google_bigquery_dataset_access}.
type BigqueryDatasetAccessA interface {
	cdktf.TerraformResource
	ApiUpdatedMember() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Dataset() BigqueryDatasetAccessDatasetAOutputReference
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DatasetInput() *BigqueryDatasetAccessDatasetA
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupByEmail() *string
	SetGroupByEmail(val *string)
	GroupByEmailInput() *string
	IamMember() *string
	SetIamMember(val *string)
	IamMemberInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Routine() BigqueryDatasetAccessRoutineAOutputReference
	RoutineInput() *BigqueryDatasetAccessRoutineA
	SpecialGroup() *string
	SetSpecialGroup(val *string)
	SpecialGroupInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BigqueryDatasetAccessTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserByEmail() *string
	SetUserByEmail(val *string)
	UserByEmailInput() *string
	View() BigqueryDatasetAccessViewAOutputReference
	ViewInput() *BigqueryDatasetAccessViewA
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDataset(value *BigqueryDatasetAccessDatasetA)
	PutRoutine(value *BigqueryDatasetAccessRoutineA)
	PutTimeouts(value *BigqueryDatasetAccessTimeouts)
	PutView(value *BigqueryDatasetAccessViewA)
	ResetDataset()
	ResetDomain()
	ResetGroupByEmail()
	ResetIamMember()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRole()
	ResetRoutine()
	ResetSpecialGroup()
	ResetTimeouts()
	ResetUserByEmail()
	ResetView()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for BigqueryDatasetAccessA
type jsiiProxy_BigqueryDatasetAccessA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BigqueryDatasetAccessA) ApiUpdatedMember() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"apiUpdatedMember",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Dataset() BigqueryDatasetAccessDatasetAOutputReference {
	var returns BigqueryDatasetAccessDatasetAOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) DatasetInput() *BigqueryDatasetAccessDatasetA {
	var returns *BigqueryDatasetAccessDatasetA
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) GroupByEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) GroupByEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) IamMember() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamMember",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) IamMemberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamMemberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Routine() BigqueryDatasetAccessRoutineAOutputReference {
	var returns BigqueryDatasetAccessRoutineAOutputReference
	_jsii_.Get(
		j,
		"routine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) RoutineInput() *BigqueryDatasetAccessRoutineA {
	var returns *BigqueryDatasetAccessRoutineA
	_jsii_.Get(
		j,
		"routineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) SpecialGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specialGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) SpecialGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specialGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) Timeouts() BigqueryDatasetAccessTimeoutsOutputReference {
	var returns BigqueryDatasetAccessTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) UserByEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userByEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) UserByEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userByEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) View() BigqueryDatasetAccessViewAOutputReference {
	var returns BigqueryDatasetAccessViewAOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatasetAccessA) ViewInput() *BigqueryDatasetAccessViewA {
	var returns *BigqueryDatasetAccessViewA
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/bigquery_dataset_access google_bigquery_dataset_access} Resource.
func NewBigqueryDatasetAccessA(scope constructs.Construct, id *string, config *BigqueryDatasetAccessAConfig) BigqueryDatasetAccessA {
	_init_.Initialize()

	if err := validateNewBigqueryDatasetAccessAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryDatasetAccessA{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDatasetAccess.BigqueryDatasetAccessA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/bigquery_dataset_access google_bigquery_dataset_access} Resource.
func NewBigqueryDatasetAccessA_Override(b BigqueryDatasetAccessA, scope constructs.Construct, id *string, config *BigqueryDatasetAccessAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDatasetAccess.BigqueryDatasetAccessA",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetGroupByEmail(val *string) {
	if err := j.validateSetGroupByEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupByEmail",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetIamMember(val *string) {
	if err := j.validateSetIamMemberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamMember",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetSpecialGroup(val *string) {
	if err := j.validateSetSpecialGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specialGroup",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatasetAccessA)SetUserByEmail(val *string) {
	if err := j.validateSetUserByEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userByEmail",
		val,
	)
}

// Generates CDKTF code for importing a BigqueryDatasetAccessA resource upon running "cdktf plan <stack-name>".
func BigqueryDatasetAccessA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBigqueryDatasetAccessA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatasetAccess.BigqueryDatasetAccessA",
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
func BigqueryDatasetAccessA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDatasetAccessA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatasetAccess.BigqueryDatasetAccessA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDatasetAccessA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDatasetAccessA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatasetAccess.BigqueryDatasetAccessA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDatasetAccessA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDatasetAccessA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatasetAccess.BigqueryDatasetAccessA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BigqueryDatasetAccessA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.bigqueryDatasetAccess.BigqueryDatasetAccessA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) PutDataset(value *BigqueryDatasetAccessDatasetA) {
	if err := b.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDataset",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) PutRoutine(value *BigqueryDatasetAccessRoutineA) {
	if err := b.validatePutRoutineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRoutine",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) PutTimeouts(value *BigqueryDatasetAccessTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) PutView(value *BigqueryDatasetAccessViewA) {
	if err := b.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putView",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetDataset() {
	_jsii_.InvokeVoid(
		b,
		"resetDataset",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetDomain() {
	_jsii_.InvokeVoid(
		b,
		"resetDomain",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetGroupByEmail() {
	_jsii_.InvokeVoid(
		b,
		"resetGroupByEmail",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetIamMember() {
	_jsii_.InvokeVoid(
		b,
		"resetIamMember",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetProject() {
	_jsii_.InvokeVoid(
		b,
		"resetProject",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetRole() {
	_jsii_.InvokeVoid(
		b,
		"resetRole",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetRoutine() {
	_jsii_.InvokeVoid(
		b,
		"resetRoutine",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetSpecialGroup() {
	_jsii_.InvokeVoid(
		b,
		"resetSpecialGroup",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetUserByEmail() {
	_jsii_.InvokeVoid(
		b,
		"resetUserByEmail",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ResetView() {
	_jsii_.InvokeVoid(
		b,
		"resetView",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatasetAccessA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatasetAccessA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

