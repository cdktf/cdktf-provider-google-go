// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkebackuprestoreplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference interface {
	cdktf.ComplexObject
	AllGroupKinds() interface{}
	SetAllGroupKinds(val interface{})
	AllGroupKindsInput() interface{}
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
	ExcludedGroupKinds() GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeExcludedGroupKindsList
	ExcludedGroupKindsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope
	SetInternalValue(val *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope)
	NoGroupKinds() interface{}
	SetNoGroupKinds(val interface{})
	NoGroupKindsInput() interface{}
	SelectedGroupKinds() GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindsList
	SelectedGroupKindsInput() interface{}
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
	PutExcludedGroupKinds(value interface{})
	PutSelectedGroupKinds(value interface{})
	ResetAllGroupKinds()
	ResetExcludedGroupKinds()
	ResetNoGroupKinds()
	ResetSelectedGroupKinds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference
type jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) AllGroupKinds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allGroupKinds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) AllGroupKindsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allGroupKindsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ExcludedGroupKinds() GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeExcludedGroupKindsList {
	var returns GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeExcludedGroupKindsList
	_jsii_.Get(
		j,
		"excludedGroupKinds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ExcludedGroupKindsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedGroupKindsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) InternalValue() *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope {
	var returns *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) NoGroupKinds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noGroupKinds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) NoGroupKindsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noGroupKindsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) SelectedGroupKinds() GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindsList {
	var returns GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindsList
	_jsii_.Get(
		j,
		"selectedGroupKinds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) SelectedGroupKindsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectedGroupKindsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference {
	_init_.Initialize()

	if err := validateNewGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupRestorePlan.GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference_Override(g GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupRestorePlan.GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference)SetAllGroupKinds(val interface{}) {
	if err := j.validateSetAllGroupKindsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allGroupKinds",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference)SetInternalValue(val *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference)SetNoGroupKinds(val interface{}) {
	if err := j.validateSetNoGroupKindsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noGroupKinds",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) PutExcludedGroupKinds(value interface{}) {
	if err := g.validatePutExcludedGroupKindsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludedGroupKinds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) PutSelectedGroupKinds(value interface{}) {
	if err := g.validatePutSelectedGroupKindsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedGroupKinds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ResetAllGroupKinds() {
	_jsii_.InvokeVoid(
		g,
		"resetAllGroupKinds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ResetExcludedGroupKinds() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedGroupKinds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ResetNoGroupKinds() {
	_jsii_.InvokeVoid(
		g,
		"resetNoGroupKinds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ResetSelectedGroupKinds() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedGroupKinds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

