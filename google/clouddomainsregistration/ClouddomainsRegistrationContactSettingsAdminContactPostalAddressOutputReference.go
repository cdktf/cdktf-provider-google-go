// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/clouddomainsregistration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference interface {
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
	InternalValue() *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress
	SetInternalValue(val *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAddressLines()
	ResetAdministrativeArea()
	ResetLocality()
	ResetOrganization()
	ResetPostalCode()
	ResetRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference
type jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) AddressLines() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) AddressLinesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) AdministrativeArea() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrativeArea",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) AdministrativeAreaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrativeAreaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) InternalValue() *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress {
	var returns *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) Recipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) RecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) RegionCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) RegionCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference {
	_init_.Initialize()

	if err := validateNewClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference_Override(c ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetAddressLines(val *[]*string) {
	if err := j.validateSetAddressLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLines",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetAdministrativeArea(val *string) {
	if err := j.validateSetAdministrativeAreaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administrativeArea",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetInternalValue(val *ClouddomainsRegistrationContactSettingsAdminContactPostalAddress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetRecipients(val *[]*string) {
	if err := j.validateSetRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipients",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetRegionCode(val *string) {
	if err := j.validateSetRegionCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionCode",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ResetAddressLines() {
	_jsii_.InvokeVoid(
		c,
		"resetAddressLines",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ResetAdministrativeArea() {
	_jsii_.InvokeVoid(
		c,
		"resetAdministrativeArea",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ResetLocality() {
	_jsii_.InvokeVoid(
		c,
		"resetLocality",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ResetOrganization() {
	_jsii_.InvokeVoid(
		c,
		"resetOrganization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		c,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ResetRecipients() {
	_jsii_.InvokeVoid(
		c,
		"resetRecipients",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationContactSettingsAdminContactPostalAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

