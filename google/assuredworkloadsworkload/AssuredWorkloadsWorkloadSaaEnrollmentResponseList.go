// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package assuredworkloadsworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/assuredworkloadsworkload/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AssuredWorkloadsWorkloadSaaEnrollmentResponseList interface {
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
	Get(index *float64) AssuredWorkloadsWorkloadSaaEnrollmentResponseOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AssuredWorkloadsWorkloadSaaEnrollmentResponseList
type jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAssuredWorkloadsWorkloadSaaEnrollmentResponseList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AssuredWorkloadsWorkloadSaaEnrollmentResponseList {
	_init_.Initialize()

	if err := validateNewAssuredWorkloadsWorkloadSaaEnrollmentResponseListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList{}

	_jsii_.Create(
		"@cdktf/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkloadSaaEnrollmentResponseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAssuredWorkloadsWorkloadSaaEnrollmentResponseList_Override(a AssuredWorkloadsWorkloadSaaEnrollmentResponseList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkloadSaaEnrollmentResponseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) Get(index *float64) AssuredWorkloadsWorkloadSaaEnrollmentResponseOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AssuredWorkloadsWorkloadSaaEnrollmentResponseOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadSaaEnrollmentResponseList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

