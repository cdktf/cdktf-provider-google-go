package computeinstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/computeinstancegroupmanager/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceGroupManagerUpdatePolicyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeInstanceGroupManagerUpdatePolicy
	SetInternalValue(val *ComputeInstanceGroupManagerUpdatePolicy)
	MaxSurgeFixed() *float64
	SetMaxSurgeFixed(val *float64)
	MaxSurgeFixedInput() *float64
	MaxSurgePercent() *float64
	SetMaxSurgePercent(val *float64)
	MaxSurgePercentInput() *float64
	MaxUnavailableFixed() *float64
	SetMaxUnavailableFixed(val *float64)
	MaxUnavailableFixedInput() *float64
	MaxUnavailablePercent() *float64
	SetMaxUnavailablePercent(val *float64)
	MaxUnavailablePercentInput() *float64
	MinimalAction() *string
	SetMinimalAction(val *string)
	MinimalActionInput() *string
	MostDisruptiveAllowedAction() *string
	SetMostDisruptiveAllowedAction(val *string)
	MostDisruptiveAllowedActionInput() *string
	ReplacementMethod() *string
	SetReplacementMethod(val *string)
	ReplacementMethodInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetMaxSurgeFixed()
	ResetMaxSurgePercent()
	ResetMaxUnavailableFixed()
	ResetMaxUnavailablePercent()
	ResetMostDisruptiveAllowedAction()
	ResetReplacementMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupManagerUpdatePolicyOutputReference
type jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) InternalValue() *ComputeInstanceGroupManagerUpdatePolicy {
	var returns *ComputeInstanceGroupManagerUpdatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgeFixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgeFixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeFixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailableFixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailableFixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableFixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailablePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailablePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MinimalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MinimalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MostDisruptiveAllowedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) MostDisruptiveAllowedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ReplacementMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ReplacementMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupManagerUpdatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupManagerUpdatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceGroupManagerUpdatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManagerUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupManagerUpdatePolicyOutputReference_Override(c ComputeInstanceGroupManagerUpdatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManagerUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetInternalValue(val *ComputeInstanceGroupManagerUpdatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxSurgeFixed(val *float64) {
	if err := j.validateSetMaxSurgeFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurgeFixed",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxSurgePercent(val *float64) {
	if err := j.validateSetMaxSurgePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurgePercent",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxUnavailableFixed(val *float64) {
	if err := j.validateSetMaxUnavailableFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailableFixed",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxUnavailablePercent(val *float64) {
	if err := j.validateSetMaxUnavailablePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailablePercent",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMinimalAction(val *string) {
	if err := j.validateSetMinimalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalAction",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMostDisruptiveAllowedAction(val *string) {
	if err := j.validateSetMostDisruptiveAllowedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mostDisruptiveAllowedAction",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetReplacementMethod(val *string) {
	if err := j.validateSetReplacementMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementMethod",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxSurgeFixed() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxSurgeFixed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxSurgePercent() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxSurgePercent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxUnavailableFixed() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxUnavailableFixed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxUnavailablePercent() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxUnavailablePercent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMostDisruptiveAllowedAction() {
	_jsii_.InvokeVoid(
		c,
		"resetMostDisruptiveAllowedAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetReplacementMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetReplacementMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManagerUpdatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

