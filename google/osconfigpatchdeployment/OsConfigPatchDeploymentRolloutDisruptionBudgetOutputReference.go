package osconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/osconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference interface {
	cdktf.ComplexObject
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
	Fixed() *float64
	SetFixed(val *float64)
	FixedInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *OsConfigPatchDeploymentRolloutDisruptionBudget
	SetInternalValue(val *OsConfigPatchDeploymentRolloutDisruptionBudget)
	Percentage() *float64
	SetPercentage(val *float64)
	PercentageInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetFixed()
	ResetPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference
type jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) Fixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) FixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) InternalValue() *OsConfigPatchDeploymentRolloutDisruptionBudget {
	var returns *OsConfigPatchDeploymentRolloutDisruptionBudget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) PercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigPatchDeploymentRolloutDisruptionBudgetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference_Override(o OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference)SetFixed(val *float64) {
	if err := j.validateSetFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixed",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference)SetInternalValue(val *OsConfigPatchDeploymentRolloutDisruptionBudget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference)SetPercentage(val *float64) {
	if err := j.validateSetPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentage",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) ResetFixed() {
	_jsii_.InvokeVoid(
		o,
		"resetFixed",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) ResetPercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetPercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRolloutDisruptionBudgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

