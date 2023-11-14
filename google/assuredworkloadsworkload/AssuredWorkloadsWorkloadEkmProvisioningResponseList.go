// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package assuredworkloadsworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/assuredworkloadsworkload/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AssuredWorkloadsWorkloadEkmProvisioningResponseList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AssuredWorkloadsWorkloadEkmProvisioningResponseOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AssuredWorkloadsWorkloadEkmProvisioningResponseList
type jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAssuredWorkloadsWorkloadEkmProvisioningResponseList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AssuredWorkloadsWorkloadEkmProvisioningResponseList {
	_init_.Initialize()

	if err := validateNewAssuredWorkloadsWorkloadEkmProvisioningResponseListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList{}

	_jsii_.Create(
		"@cdktf/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkloadEkmProvisioningResponseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAssuredWorkloadsWorkloadEkmProvisioningResponseList_Override(a AssuredWorkloadsWorkloadEkmProvisioningResponseList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkloadEkmProvisioningResponseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) Get(index *float64) AssuredWorkloadsWorkloadEkmProvisioningResponseOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AssuredWorkloadsWorkloadEkmProvisioningResponseOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AssuredWorkloadsWorkloadEkmProvisioningResponseList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

