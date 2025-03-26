// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeimage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeImageShieldedInstanceInitialStateOutputReference interface {
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
	Dbs() ComputeImageShieldedInstanceInitialStateDbsList
	DbsInput() interface{}
	Dbxs() ComputeImageShieldedInstanceInitialStateDbxsList
	DbxsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeImageShieldedInstanceInitialState
	SetInternalValue(val *ComputeImageShieldedInstanceInitialState)
	Keks() ComputeImageShieldedInstanceInitialStateKeksList
	KeksInput() interface{}
	Pk() ComputeImageShieldedInstanceInitialStatePkOutputReference
	PkInput() *ComputeImageShieldedInstanceInitialStatePk
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
	PutDbs(value interface{})
	PutDbxs(value interface{})
	PutKeks(value interface{})
	PutPk(value *ComputeImageShieldedInstanceInitialStatePk)
	ResetDbs()
	ResetDbxs()
	ResetKeks()
	ResetPk()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeImageShieldedInstanceInitialStateOutputReference
type jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) Dbs() ComputeImageShieldedInstanceInitialStateDbsList {
	var returns ComputeImageShieldedInstanceInitialStateDbsList
	_jsii_.Get(
		j,
		"dbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) DbsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) Dbxs() ComputeImageShieldedInstanceInitialStateDbxsList {
	var returns ComputeImageShieldedInstanceInitialStateDbxsList
	_jsii_.Get(
		j,
		"dbxs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) DbxsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbxsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) InternalValue() *ComputeImageShieldedInstanceInitialState {
	var returns *ComputeImageShieldedInstanceInitialState
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) Keks() ComputeImageShieldedInstanceInitialStateKeksList {
	var returns ComputeImageShieldedInstanceInitialStateKeksList
	_jsii_.Get(
		j,
		"keks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) KeksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) Pk() ComputeImageShieldedInstanceInitialStatePkOutputReference {
	var returns ComputeImageShieldedInstanceInitialStatePkOutputReference
	_jsii_.Get(
		j,
		"pk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) PkInput() *ComputeImageShieldedInstanceInitialStatePk {
	var returns *ComputeImageShieldedInstanceInitialStatePk
	_jsii_.Get(
		j,
		"pkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeImageShieldedInstanceInitialStateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeImageShieldedInstanceInitialStateOutputReference {
	_init_.Initialize()

	if err := validateNewComputeImageShieldedInstanceInitialStateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeImage.ComputeImageShieldedInstanceInitialStateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeImageShieldedInstanceInitialStateOutputReference_Override(c ComputeImageShieldedInstanceInitialStateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeImage.ComputeImageShieldedInstanceInitialStateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference)SetInternalValue(val *ComputeImageShieldedInstanceInitialState) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) PutDbs(value interface{}) {
	if err := c.validatePutDbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDbs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) PutDbxs(value interface{}) {
	if err := c.validatePutDbxsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDbxs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) PutKeks(value interface{}) {
	if err := c.validatePutKeksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKeks",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) PutPk(value *ComputeImageShieldedInstanceInitialStatePk) {
	if err := c.validatePutPkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ResetDbs() {
	_jsii_.InvokeVoid(
		c,
		"resetDbs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ResetDbxs() {
	_jsii_.InvokeVoid(
		c,
		"resetDbxs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ResetKeks() {
	_jsii_.InvokeVoid(
		c,
		"resetKeks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ResetPk() {
	_jsii_.InvokeVoid(
		c,
		"resetPk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeImageShieldedInstanceInitialStateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

