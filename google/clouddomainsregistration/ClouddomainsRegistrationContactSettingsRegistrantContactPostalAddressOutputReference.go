// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/clouddomainsregistration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference interface {
	cdktf.ComplexObject
	AddressLines() *[]*string
	SetAddressLines(val *[]*string)
	AddressLinesInput() *[]*string
	AdministrativeArea() *string
	SetAdministrativeArea(val *string)
	AdministrativeAreaInput() *string
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
	InternalValue() *ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress
	SetInternalValue(val *ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress)
	Locality() *string
	SetLocality(val *string)
	LocalityInput() *string
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	Recipients() *[]*string
	SetRecipients(val *[]*string)
	RecipientsInput() *[]*string
	RegionCode() *string
	SetRegionCode(val *string)
	RegionCodeInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAddressLines()
	ResetAdministrativeArea()
	ResetLocality()
	ResetOrganization()
	ResetPostalCode()
	ResetRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference
type jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) AddressLines() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) AddressLinesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) AdministrativeArea() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrativeArea",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) AdministrativeAreaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrativeAreaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) InternalValue() *ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress {
	var returns *ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) Recipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) RecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) RegionCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) RegionCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference {
	_init_.Initialize()

	if err := validateNewClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference_Override(c ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetAddressLines(val *[]*string) {
	if err := j.validateSetAddressLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLines",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetAdministrativeArea(val *string) {
	if err := j.validateSetAdministrativeAreaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administrativeArea",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetInternalValue(val *ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetRecipients(val *[]*string) {
	if err := j.validateSetRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipients",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetRegionCode(val *string) {
	if err := j.validateSetRegionCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionCode",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ResetAddressLines() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressLines",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ResetAdministrativeArea() {
	_jsii_.InvokeVoid(
		c,
		"resetAdministrativeArea",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ResetLocality() {
	_jsii_.InvokeVoid(
		c,
		"resetLocality",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ResetOrganization() {
	_jsii_.InvokeVoid(
		c,
		"resetOrganization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		c,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ResetRecipients() {
	_jsii_.InvokeVoid(
		c,
		"resetRecipients",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

