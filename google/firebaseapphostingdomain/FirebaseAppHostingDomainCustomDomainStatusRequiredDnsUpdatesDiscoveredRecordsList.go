// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/firebaseapphostingdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList
type jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList {
	_init_.Initialize()

	if err := validateNewFirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList{}

	_jsii_.Create(
		"@cdktf/provider-google.firebaseAppHostingDomain.FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList_Override(f FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.firebaseAppHostingDomain.FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := f.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		f,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) Get(index *float64) FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsOutputReference {
	if err := f.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FirebaseAppHostingDomainCustomDomainStatusRequiredDnsUpdatesDiscoveredRecordsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

