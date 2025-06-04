// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/osconfigospolicyassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference interface {
	cdktf.ComplexObject
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	GpgKeys() *[]*string
	SetGpgKeys(val *[]*string)
	GpgKeysInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum)
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
	ResetDisplayName()
	ResetGpgKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference
type jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GpgKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gpgKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GpgKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gpgKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference_Override(o OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetBaseUrl(val *string) {
	if err := j.validateSetBaseUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetGpgKeys(val *[]*string) {
	if err := j.validateSetGpgKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpgKeys",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		o,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) ResetGpgKeys() {
	_jsii_.InvokeVoid(
		o,
		"resetGpgKeys",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

