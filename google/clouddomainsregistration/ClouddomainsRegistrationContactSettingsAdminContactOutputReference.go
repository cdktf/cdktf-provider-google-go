// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/clouddomainsregistration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddomainsRegistrationContactSettingsAdminContactOutputReference interface {
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
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	FaxNumber() *string
	SetFaxNumber(val *string)
	FaxNumberInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ClouddomainsRegistrationContactSettingsAdminContact
	SetInternalValue(val *ClouddomainsRegistrationContactSettingsAdminContact)
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
	PostalAddress() ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference
	PostalAddressInput() *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress
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
	PutPostalAddress(value *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress)
	ResetFaxNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddomainsRegistrationContactSettingsAdminContactOutputReference
type jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) FaxNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) FaxNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) InternalValue() *ClouddomainsRegistrationContactSettingsAdminContact {
	var returns *ClouddomainsRegistrationContactSettingsAdminContact
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) PostalAddress() ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference {
	var returns ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference
	_jsii_.Get(
		j,
		"postalAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) PostalAddressInput() *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress {
	var returns *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress
	_jsii_.Get(
		j,
		"postalAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddomainsRegistrationContactSettingsAdminContactOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddomainsRegistrationContactSettingsAdminContactOutputReference {
	_init_.Initialize()

	if err := validateNewClouddomainsRegistrationContactSettingsAdminContactOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsAdminContactOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddomainsRegistrationContactSettingsAdminContactOutputReference_Override(c ClouddomainsRegistrationContactSettingsAdminContactOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsAdminContactOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetFaxNumber(val *string) {
	if err := j.validateSetFaxNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faxNumber",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetInternalValue(val *ClouddomainsRegistrationContactSettingsAdminContact) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) PutPostalAddress(value *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress) {
	if err := c.validatePutPostalAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostalAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) ResetFaxNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetFaxNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

