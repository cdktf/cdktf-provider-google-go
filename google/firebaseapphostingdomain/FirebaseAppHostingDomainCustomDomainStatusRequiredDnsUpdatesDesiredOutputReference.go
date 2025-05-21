// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/firebaseapphostingdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference interface {
	cdktf.ComplexObject
	CheckError() FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredCheckErrorList
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
	DomainName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesired
	SetInternalValue(val *FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesired)
	Records() FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredRecordsList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference
type jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) CheckError() FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredCheckErrorList {
	var returns FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredCheckErrorList
	_jsii_.Get(
		j,
		"checkError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) InternalValue() *FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesired {
	var returns *FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesired
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) Records() FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredRecordsList {
	var returns FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredRecordsList
	_jsii_.Get(
		j,
		"records",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference {
	_init_.Initialize()

	if err := validateNewFirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.firebaseAppHostingDomain.FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference_Override(f FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.firebaseAppHostingDomain.FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference)SetInternalValue(val *FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesired) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDesiredOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

