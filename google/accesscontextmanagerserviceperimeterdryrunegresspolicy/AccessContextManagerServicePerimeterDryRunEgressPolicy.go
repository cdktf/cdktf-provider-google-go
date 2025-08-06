// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeterdryrunegresspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/accesscontextmanagerserviceperimeterdryrunegresspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy google_access_context_manager_service_perimeter_dry_run_egress_policy}.
type AccessContextManagerServicePerimeterDryRunEgressPolicy interface {
	cdktf.TerraformResource
	AccessPolicyId() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EgressFrom() AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFromOutputReference
	EgressFromInput() *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFrom
	EgressTo() AccessContextManagerServicePerimeterDryRunEgressPolicyEgressToOutputReference
	EgressToInput() *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressTo
	Etag() *string
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
	// The tree node.
	Node() constructs.Node
	Perimeter() *string
	SetPerimeter(val *string)
	PerimeterInput() *string
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
	Timeouts() AccessContextManagerServicePerimeterDryRunEgressPolicyTimeoutsOutputReference
	TimeoutsInput() interface{}
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	PutEgressFrom(value *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFrom)
	PutEgressTo(value *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressTo)
	PutTimeouts(value *AccessContextManagerServicePerimeterDryRunEgressPolicyTimeouts)
	ResetEgressFrom()
	ResetEgressTo()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetTitle()
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

// The jsii proxy struct for AccessContextManagerServicePerimeterDryRunEgressPolicy
type jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) AccessPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) EgressFrom() AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFromOutputReference {
	var returns AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFromOutputReference
	_jsii_.Get(
		j,
		"egressFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) EgressFromInput() *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFrom {
	var returns *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFrom
	_jsii_.Get(
		j,
		"egressFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) EgressTo() AccessContextManagerServicePerimeterDryRunEgressPolicyEgressToOutputReference {
	var returns AccessContextManagerServicePerimeterDryRunEgressPolicyEgressToOutputReference
	_jsii_.Get(
		j,
		"egressTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) EgressToInput() *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressTo {
	var returns *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressTo
	_jsii_.Get(
		j,
		"egressToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Perimeter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perimeter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) PerimeterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perimeterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Timeouts() AccessContextManagerServicePerimeterDryRunEgressPolicyTimeoutsOutputReference {
	var returns AccessContextManagerServicePerimeterDryRunEgressPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy google_access_context_manager_service_perimeter_dry_run_egress_policy} Resource.
func NewAccessContextManagerServicePerimeterDryRunEgressPolicy(scope constructs.Construct, id *string, config *AccessContextManagerServicePerimeterDryRunEgressPolicyConfig) AccessContextManagerServicePerimeterDryRunEgressPolicy {
	_init_.Initialize()

	if err := validateNewAccessContextManagerServicePerimeterDryRunEgressPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeterDryRunEgressPolicy.AccessContextManagerServicePerimeterDryRunEgressPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy google_access_context_manager_service_perimeter_dry_run_egress_policy} Resource.
func NewAccessContextManagerServicePerimeterDryRunEgressPolicy_Override(a AccessContextManagerServicePerimeterDryRunEgressPolicy, scope constructs.Construct, id *string, config *AccessContextManagerServicePerimeterDryRunEgressPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeterDryRunEgressPolicy.AccessContextManagerServicePerimeterDryRunEgressPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetPerimeter(val *string) {
	if err := j.validateSetPerimeterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perimeter",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

// Generates CDKTF code for importing a AccessContextManagerServicePerimeterDryRunEgressPolicy resource upon running "cdktf plan <stack-name>".
func AccessContextManagerServicePerimeterDryRunEgressPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAccessContextManagerServicePerimeterDryRunEgressPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.accessContextManagerServicePerimeterDryRunEgressPolicy.AccessContextManagerServicePerimeterDryRunEgressPolicy",
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
func AccessContextManagerServicePerimeterDryRunEgressPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessContextManagerServicePerimeterDryRunEgressPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.accessContextManagerServicePerimeterDryRunEgressPolicy.AccessContextManagerServicePerimeterDryRunEgressPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessContextManagerServicePerimeterDryRunEgressPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessContextManagerServicePerimeterDryRunEgressPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.accessContextManagerServicePerimeterDryRunEgressPolicy.AccessContextManagerServicePerimeterDryRunEgressPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessContextManagerServicePerimeterDryRunEgressPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessContextManagerServicePerimeterDryRunEgressPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.accessContextManagerServicePerimeterDryRunEgressPolicy.AccessContextManagerServicePerimeterDryRunEgressPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AccessContextManagerServicePerimeterDryRunEgressPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.accessContextManagerServicePerimeterDryRunEgressPolicy.AccessContextManagerServicePerimeterDryRunEgressPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) PutEgressFrom(value *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFrom) {
	if err := a.validatePutEgressFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEgressFrom",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) PutEgressTo(value *AccessContextManagerServicePerimeterDryRunEgressPolicyEgressTo) {
	if err := a.validatePutEgressToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEgressTo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) PutTimeouts(value *AccessContextManagerServicePerimeterDryRunEgressPolicyTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ResetEgressFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetEgressFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ResetEgressTo() {
	_jsii_.InvokeVoid(
		a,
		"resetEgressTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterDryRunEgressPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

