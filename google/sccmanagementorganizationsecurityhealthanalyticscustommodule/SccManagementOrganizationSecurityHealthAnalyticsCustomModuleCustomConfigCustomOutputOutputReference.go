// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementorganizationsecurityhealthanalyticscustommodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/sccmanagementorganizationsecurityhealthanalyticscustommodule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference interface {
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
	InternalValue() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	SetInternalValue(val *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput)
	Properties() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesList
	PropertiesInput() interface{}
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
	PutProperties(value interface{})
	ResetProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference
type jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) InternalValue() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput {
	var returns *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) Properties() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesList {
	var returns SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesList
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) PropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference {
	_init_.Initialize()

	if err := validateNewSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementOrganizationSecurityHealthAnalyticsCustomModule.SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference_Override(s SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementOrganizationSecurityHealthAnalyticsCustomModule.SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference)SetInternalValue(val *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) PutProperties(value interface{}) {
	if err := s.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

