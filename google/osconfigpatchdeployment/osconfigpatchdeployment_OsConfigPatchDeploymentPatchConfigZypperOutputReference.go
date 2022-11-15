package osconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/osconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigPatchDeploymentPatchConfigZypperOutputReference interface {
	cdktf.ComplexObject
	Categories() *[]*string
	SetCategories(val *[]*string)
	CategoriesInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Excludes() *[]*string
	SetExcludes(val *[]*string)
	ExcludesInput() *[]*string
	ExclusivePatches() *[]*string
	SetExclusivePatches(val *[]*string)
	ExclusivePatchesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *OsConfigPatchDeploymentPatchConfigZypper
	SetInternalValue(val *OsConfigPatchDeploymentPatchConfigZypper)
	Severities() *[]*string
	SetSeverities(val *[]*string)
	SeveritiesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WithOptional() interface{}
	SetWithOptional(val interface{})
	WithOptionalInput() interface{}
	WithUpdate() interface{}
	SetWithUpdate(val interface{})
	WithUpdateInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCategories()
	ResetExcludes()
	ResetExclusivePatches()
	ResetSeverities()
	ResetWithOptional()
	ResetWithUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigPatchDeploymentPatchConfigZypperOutputReference
type jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) Categories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) CategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) Excludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ExclusivePatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusivePatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ExclusivePatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusivePatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) InternalValue() *OsConfigPatchDeploymentPatchConfigZypper {
	var returns *OsConfigPatchDeploymentPatchConfigZypper
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) Severities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"severities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) SeveritiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"severitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) WithOptional() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withOptional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) WithOptionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withOptionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) WithUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) WithUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withUpdateInput",
		&returns,
	)
	return returns
}


func NewOsConfigPatchDeploymentPatchConfigZypperOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigPatchDeploymentPatchConfigZypperOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigPatchDeploymentPatchConfigZypperOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentPatchConfigZypperOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigPatchDeploymentPatchConfigZypperOutputReference_Override(o OsConfigPatchDeploymentPatchConfigZypperOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentPatchConfigZypperOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetCategories(val *[]*string) {
	if err := j.validateSetCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"categories",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetExcludes(val *[]*string) {
	if err := j.validateSetExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetExclusivePatches(val *[]*string) {
	if err := j.validateSetExclusivePatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusivePatches",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetInternalValue(val *OsConfigPatchDeploymentPatchConfigZypper) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetSeverities(val *[]*string) {
	if err := j.validateSetSeveritiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severities",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetWithOptional(val interface{}) {
	if err := j.validateSetWithOptionalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withOptional",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference)SetWithUpdate(val interface{}) {
	if err := j.validateSetWithUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withUpdate",
		val,
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetCategories() {
	_jsii_.InvokeVoid(
		o,
		"resetCategories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetExcludes() {
	_jsii_.InvokeVoid(
		o,
		"resetExcludes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetExclusivePatches() {
	_jsii_.InvokeVoid(
		o,
		"resetExclusivePatches",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetSeverities() {
	_jsii_.InvokeVoid(
		o,
		"resetSeverities",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetWithOptional() {
	_jsii_.InvokeVoid(
		o,
		"resetWithOptional",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ResetWithUpdate() {
	_jsii_.InvokeVoid(
		o,
		"resetWithUpdate",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigZypperOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

