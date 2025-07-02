// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containeranalysisnote

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/containeranalysisnote/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/container_analysis_note google_container_analysis_note}.
type ContainerAnalysisNote interface {
	cdktf.TerraformResource
	AttestationAuthority() ContainerAnalysisNoteAttestationAuthorityOutputReference
	AttestationAuthorityInput() *ContainerAnalysisNoteAttestationAuthority
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExpirationTime() *string
	SetExpirationTime(val *string)
	ExpirationTimeInput() *string
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
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LongDescription() *string
	SetLongDescription(val *string)
	LongDescriptionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RelatedNoteNames() *[]*string
	SetRelatedNoteNames(val *[]*string)
	RelatedNoteNamesInput() *[]*string
	RelatedUrl() ContainerAnalysisNoteRelatedUrlList
	RelatedUrlInput() interface{}
	ShortDescription() *string
	SetShortDescription(val *string)
	ShortDescriptionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerAnalysisNoteTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutAttestationAuthority(value *ContainerAnalysisNoteAttestationAuthority)
	PutRelatedUrl(value interface{})
	PutTimeouts(value *ContainerAnalysisNoteTimeouts)
	ResetExpirationTime()
	ResetId()
	ResetLongDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRelatedNoteNames()
	ResetRelatedUrl()
	ResetShortDescription()
	ResetTimeouts()
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

// The jsii proxy struct for ContainerAnalysisNote
type jsiiProxy_ContainerAnalysisNote struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerAnalysisNote) AttestationAuthority() ContainerAnalysisNoteAttestationAuthorityOutputReference {
	var returns ContainerAnalysisNoteAttestationAuthorityOutputReference
	_jsii_.Get(
		j,
		"attestationAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) AttestationAuthorityInput() *ContainerAnalysisNoteAttestationAuthority {
	var returns *ContainerAnalysisNoteAttestationAuthority
	_jsii_.Get(
		j,
		"attestationAuthorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) ExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) ExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) LongDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) LongDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) RelatedNoteNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relatedNoteNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) RelatedNoteNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relatedNoteNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) RelatedUrl() ContainerAnalysisNoteRelatedUrlList {
	var returns ContainerAnalysisNoteRelatedUrlList
	_jsii_.Get(
		j,
		"relatedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) RelatedUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) ShortDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) ShortDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) Timeouts() ContainerAnalysisNoteTimeoutsOutputReference {
	var returns ContainerAnalysisNoteTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAnalysisNote) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/container_analysis_note google_container_analysis_note} Resource.
func NewContainerAnalysisNote(scope constructs.Construct, id *string, config *ContainerAnalysisNoteConfig) ContainerAnalysisNote {
	_init_.Initialize()

	if err := validateNewContainerAnalysisNoteParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAnalysisNote{}

	_jsii_.Create(
		"@cdktf/provider-google.containerAnalysisNote.ContainerAnalysisNote",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/container_analysis_note google_container_analysis_note} Resource.
func NewContainerAnalysisNote_Override(c ContainerAnalysisNote, scope constructs.Construct, id *string, config *ContainerAnalysisNoteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerAnalysisNote.ContainerAnalysisNote",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetExpirationTime(val *string) {
	if err := j.validateSetExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTime",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetLongDescription(val *string) {
	if err := j.validateSetLongDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longDescription",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetRelatedNoteNames(val *[]*string) {
	if err := j.validateSetRelatedNoteNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relatedNoteNames",
		val,
	)
}

func (j *jsiiProxy_ContainerAnalysisNote)SetShortDescription(val *string) {
	if err := j.validateSetShortDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shortDescription",
		val,
	)
}

// Generates CDKTF code for importing a ContainerAnalysisNote resource upon running "cdktf plan <stack-name>".
func ContainerAnalysisNote_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateContainerAnalysisNote_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAnalysisNote.ContainerAnalysisNote",
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
func ContainerAnalysisNote_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAnalysisNote_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAnalysisNote.ContainerAnalysisNote",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAnalysisNote_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAnalysisNote_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAnalysisNote.ContainerAnalysisNote",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAnalysisNote_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAnalysisNote_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAnalysisNote.ContainerAnalysisNote",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerAnalysisNote_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.containerAnalysisNote.ContainerAnalysisNote",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerAnalysisNote) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAnalysisNote) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAnalysisNote) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAnalysisNote) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAnalysisNote) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) PutAttestationAuthority(value *ContainerAnalysisNoteAttestationAuthority) {
	if err := c.validatePutAttestationAuthorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAttestationAuthority",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) PutRelatedUrl(value interface{}) {
	if err := c.validatePutRelatedUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRelatedUrl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) PutTimeouts(value *ContainerAnalysisNoteTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetExpirationTime() {
	_jsii_.InvokeVoid(
		c,
		"resetExpirationTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetLongDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetLongDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetRelatedNoteNames() {
	_jsii_.InvokeVoid(
		c,
		"resetRelatedNoteNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetRelatedUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetRelatedUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetShortDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetShortDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAnalysisNote) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAnalysisNote) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAnalysisNote) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAnalysisNote) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAnalysisNote) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAnalysisNote) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

