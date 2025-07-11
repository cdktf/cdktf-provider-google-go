// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/billingbudget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BillingBudgetBudgetFilterCustomPeriodOutputReference interface {
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
	EndDate() BillingBudgetBudgetFilterCustomPeriodEndDateOutputReference
	EndDateInput() *BillingBudgetBudgetFilterCustomPeriodEndDate
	// Experimental.
	Fqn() *string
	InternalValue() *BillingBudgetBudgetFilterCustomPeriod
	SetInternalValue(val *BillingBudgetBudgetFilterCustomPeriod)
	StartDate() BillingBudgetBudgetFilterCustomPeriodStartDateOutputReference
	StartDateInput() *BillingBudgetBudgetFilterCustomPeriodStartDate
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
	PutEndDate(value *BillingBudgetBudgetFilterCustomPeriodEndDate)
	PutStartDate(value *BillingBudgetBudgetFilterCustomPeriodStartDate)
	ResetEndDate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BillingBudgetBudgetFilterCustomPeriodOutputReference
type jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) EndDate() BillingBudgetBudgetFilterCustomPeriodEndDateOutputReference {
	var returns BillingBudgetBudgetFilterCustomPeriodEndDateOutputReference
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) EndDateInput() *BillingBudgetBudgetFilterCustomPeriodEndDate {
	var returns *BillingBudgetBudgetFilterCustomPeriodEndDate
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) InternalValue() *BillingBudgetBudgetFilterCustomPeriod {
	var returns *BillingBudgetBudgetFilterCustomPeriod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) StartDate() BillingBudgetBudgetFilterCustomPeriodStartDateOutputReference {
	var returns BillingBudgetBudgetFilterCustomPeriodStartDateOutputReference
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) StartDateInput() *BillingBudgetBudgetFilterCustomPeriodStartDate {
	var returns *BillingBudgetBudgetFilterCustomPeriodStartDate
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBillingBudgetBudgetFilterCustomPeriodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BillingBudgetBudgetFilterCustomPeriodOutputReference {
	_init_.Initialize()

	if err := validateNewBillingBudgetBudgetFilterCustomPeriodOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.billingBudget.BillingBudgetBudgetFilterCustomPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBillingBudgetBudgetFilterCustomPeriodOutputReference_Override(b BillingBudgetBudgetFilterCustomPeriodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.billingBudget.BillingBudgetBudgetFilterCustomPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference)SetInternalValue(val *BillingBudgetBudgetFilterCustomPeriod) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) PutEndDate(value *BillingBudgetBudgetFilterCustomPeriodEndDate) {
	if err := b.validatePutEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEndDate",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) PutStartDate(value *BillingBudgetBudgetFilterCustomPeriodStartDate) {
	if err := b.validatePutStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStartDate",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) ResetEndDate() {
	_jsii_.InvokeVoid(
		b,
		"resetEndDate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetBudgetFilterCustomPeriodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

