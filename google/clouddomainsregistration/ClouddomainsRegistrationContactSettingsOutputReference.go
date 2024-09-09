// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/clouddomainsregistration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddomainsRegistrationContactSettingsOutputReference interface {
	cdktf.ComplexObject
	AdminContact() ClouddomainsRegistrationContactSettingsAdminContactOutputReference
	AdminContactInput() *ClouddomainsRegistrationContactSettingsAdminContact
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
	InternalValue() *ClouddomainsRegistrationContactSettings
	SetInternalValue(val *ClouddomainsRegistrationContactSettings)
	Privacy() *string
	SetPrivacy(val *string)
	PrivacyInput() *string
	RegistrantContact() ClouddomainsRegistrationContactSettingsRegistrantContactOutputReference
	RegistrantContactInput() *ClouddomainsRegistrationContactSettingsRegistrantContact
	TechnicalContact() ClouddomainsRegistrationContactSettingsTechnicalContactOutputReference
	TechnicalContactInput() *ClouddomainsRegistrationContactSettingsTechnicalContact
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
	PutAdminContact(value *ClouddomainsRegistrationContactSettingsAdminContact)
	PutRegistrantContact(value *ClouddomainsRegistrationContactSettingsRegistrantContact)
	PutTechnicalContact(value *ClouddomainsRegistrationContactSettingsTechnicalContact)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddomainsRegistrationContactSettingsOutputReference
type jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) AdminContact() ClouddomainsRegistrationContactSettingsAdminContactOutputReference {
	var returns ClouddomainsRegistrationContactSettingsAdminContactOutputReference
	_jsii_.Get(
		j,
		"adminContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) AdminContactInput() *ClouddomainsRegistrationContactSettingsAdminContact {
	var returns *ClouddomainsRegistrationContactSettingsAdminContact
	_jsii_.Get(
		j,
		"adminContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) InternalValue() *ClouddomainsRegistrationContactSettings {
	var returns *ClouddomainsRegistrationContactSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) Privacy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) PrivacyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) RegistrantContact() ClouddomainsRegistrationContactSettingsRegistrantContactOutputReference {
	var returns ClouddomainsRegistrationContactSettingsRegistrantContactOutputReference
	_jsii_.Get(
		j,
		"registrantContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) RegistrantContactInput() *ClouddomainsRegistrationContactSettingsRegistrantContact {
	var returns *ClouddomainsRegistrationContactSettingsRegistrantContact
	_jsii_.Get(
		j,
		"registrantContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) TechnicalContact() ClouddomainsRegistrationContactSettingsTechnicalContactOutputReference {
	var returns ClouddomainsRegistrationContactSettingsTechnicalContactOutputReference
	_jsii_.Get(
		j,
		"technicalContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) TechnicalContactInput() *ClouddomainsRegistrationContactSettingsTechnicalContact {
	var returns *ClouddomainsRegistrationContactSettingsTechnicalContact
	_jsii_.Get(
		j,
		"technicalContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddomainsRegistrationContactSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddomainsRegistrationContactSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewClouddomainsRegistrationContactSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddomainsRegistrationContactSettingsOutputReference_Override(c ClouddomainsRegistrationContactSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference)SetInternalValue(val *ClouddomainsRegistrationContactSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference)SetPrivacy(val *string) {
	if err := j.validateSetPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privacy",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) PutAdminContact(value *ClouddomainsRegistrationContactSettingsAdminContact) {
	if err := c.validatePutAdminContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdminContact",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) PutRegistrantContact(value *ClouddomainsRegistrationContactSettingsRegistrantContact) {
	if err := c.validatePutRegistrantContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegistrantContact",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) PutTechnicalContact(value *ClouddomainsRegistrationContactSettingsTechnicalContact) {
	if err := c.validatePutTechnicalContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTechnicalContact",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

