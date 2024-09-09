// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/gkebackuprestoreplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference interface {
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
	GroupKinds() GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterGroupKindsList
	GroupKindsInput() interface{}
	InternalValue() *GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter
	SetInternalValue(val *GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter)
	JsonPath() *string
	SetJsonPath(val *string)
	JsonPathInput() *string
	Namespaces() *[]*string
	SetNamespaces(val *[]*string)
	NamespacesInput() *[]*string
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
	PutGroupKinds(value interface{})
	ResetGroupKinds()
	ResetJsonPath()
	ResetNamespaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference
type jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GroupKinds() GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterGroupKindsList {
	var returns GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterGroupKindsList
	_jsii_.Get(
		j,
		"groupKinds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GroupKindsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupKindsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) InternalValue() *GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter {
	var returns *GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) JsonPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) JsonPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) Namespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"namespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) NamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"namespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupRestorePlan.GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference_Override(g GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupRestorePlan.GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference)SetInternalValue(val *GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference)SetJsonPath(val *string) {
	if err := j.validateSetJsonPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonPath",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference)SetNamespaces(val *[]*string) {
	if err := j.validateSetNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaces",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) PutGroupKinds(value interface{}) {
	if err := g.validatePutGroupKindsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGroupKinds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) ResetGroupKinds() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupKinds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) ResetJsonPath() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) ResetNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

