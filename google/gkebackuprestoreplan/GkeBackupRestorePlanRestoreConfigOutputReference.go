// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkebackuprestoreplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupRestorePlanRestoreConfigOutputReference interface {
	cdktf.ComplexObject
	AllNamespaces() interface{}
	SetAllNamespaces(val interface{})
	AllNamespacesInput() interface{}
	ClusterResourceConflictPolicy() *string
	SetClusterResourceConflictPolicy(val *string)
	ClusterResourceConflictPolicyInput() *string
	ClusterResourceRestoreScope() GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference
	ClusterResourceRestoreScopeInput() *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope
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
	ExcludedNamespaces() GkeBackupRestorePlanRestoreConfigExcludedNamespacesOutputReference
	ExcludedNamespacesInput() *GkeBackupRestorePlanRestoreConfigExcludedNamespaces
	// Experimental.
	Fqn() *string
	InternalValue() *GkeBackupRestorePlanRestoreConfig
	SetInternalValue(val *GkeBackupRestorePlanRestoreConfig)
	NamespacedResourceRestoreMode() *string
	SetNamespacedResourceRestoreMode(val *string)
	NamespacedResourceRestoreModeInput() *string
	NoNamespaces() interface{}
	SetNoNamespaces(val interface{})
	NoNamespacesInput() interface{}
	RestoreOrder() GkeBackupRestorePlanRestoreConfigRestoreOrderOutputReference
	RestoreOrderInput() *GkeBackupRestorePlanRestoreConfigRestoreOrder
	SelectedApplications() GkeBackupRestorePlanRestoreConfigSelectedApplicationsOutputReference
	SelectedApplicationsInput() *GkeBackupRestorePlanRestoreConfigSelectedApplications
	SelectedNamespaces() GkeBackupRestorePlanRestoreConfigSelectedNamespacesOutputReference
	SelectedNamespacesInput() *GkeBackupRestorePlanRestoreConfigSelectedNamespaces
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformationRules() GkeBackupRestorePlanRestoreConfigTransformationRulesList
	TransformationRulesInput() interface{}
	VolumeDataRestorePolicy() *string
	SetVolumeDataRestorePolicy(val *string)
	VolumeDataRestorePolicyBindings() GkeBackupRestorePlanRestoreConfigVolumeDataRestorePolicyBindingsList
	VolumeDataRestorePolicyBindingsInput() interface{}
	VolumeDataRestorePolicyInput() *string
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
	PutClusterResourceRestoreScope(value *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope)
	PutExcludedNamespaces(value *GkeBackupRestorePlanRestoreConfigExcludedNamespaces)
	PutRestoreOrder(value *GkeBackupRestorePlanRestoreConfigRestoreOrder)
	PutSelectedApplications(value *GkeBackupRestorePlanRestoreConfigSelectedApplications)
	PutSelectedNamespaces(value *GkeBackupRestorePlanRestoreConfigSelectedNamespaces)
	PutTransformationRules(value interface{})
	PutVolumeDataRestorePolicyBindings(value interface{})
	ResetAllNamespaces()
	ResetClusterResourceConflictPolicy()
	ResetClusterResourceRestoreScope()
	ResetExcludedNamespaces()
	ResetNamespacedResourceRestoreMode()
	ResetNoNamespaces()
	ResetRestoreOrder()
	ResetSelectedApplications()
	ResetSelectedNamespaces()
	ResetTransformationRules()
	ResetVolumeDataRestorePolicy()
	ResetVolumeDataRestorePolicyBindings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeBackupRestorePlanRestoreConfigOutputReference
type jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) AllNamespaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) AllNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceConflictPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceConflictPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceConflictPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceConflictPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceRestoreScope() GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference {
	var returns GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference
	_jsii_.Get(
		j,
		"clusterResourceRestoreScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceRestoreScopeInput() *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope {
	var returns *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope
	_jsii_.Get(
		j,
		"clusterResourceRestoreScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ExcludedNamespaces() GkeBackupRestorePlanRestoreConfigExcludedNamespacesOutputReference {
	var returns GkeBackupRestorePlanRestoreConfigExcludedNamespacesOutputReference
	_jsii_.Get(
		j,
		"excludedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ExcludedNamespacesInput() *GkeBackupRestorePlanRestoreConfigExcludedNamespaces {
	var returns *GkeBackupRestorePlanRestoreConfigExcludedNamespaces
	_jsii_.Get(
		j,
		"excludedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) InternalValue() *GkeBackupRestorePlanRestoreConfig {
	var returns *GkeBackupRestorePlanRestoreConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) NamespacedResourceRestoreMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespacedResourceRestoreMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) NamespacedResourceRestoreModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespacedResourceRestoreModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) NoNamespaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) NoNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) RestoreOrder() GkeBackupRestorePlanRestoreConfigRestoreOrderOutputReference {
	var returns GkeBackupRestorePlanRestoreConfigRestoreOrderOutputReference
	_jsii_.Get(
		j,
		"restoreOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) RestoreOrderInput() *GkeBackupRestorePlanRestoreConfigRestoreOrder {
	var returns *GkeBackupRestorePlanRestoreConfigRestoreOrder
	_jsii_.Get(
		j,
		"restoreOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) SelectedApplications() GkeBackupRestorePlanRestoreConfigSelectedApplicationsOutputReference {
	var returns GkeBackupRestorePlanRestoreConfigSelectedApplicationsOutputReference
	_jsii_.Get(
		j,
		"selectedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) SelectedApplicationsInput() *GkeBackupRestorePlanRestoreConfigSelectedApplications {
	var returns *GkeBackupRestorePlanRestoreConfigSelectedApplications
	_jsii_.Get(
		j,
		"selectedApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) SelectedNamespaces() GkeBackupRestorePlanRestoreConfigSelectedNamespacesOutputReference {
	var returns GkeBackupRestorePlanRestoreConfigSelectedNamespacesOutputReference
	_jsii_.Get(
		j,
		"selectedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) SelectedNamespacesInput() *GkeBackupRestorePlanRestoreConfigSelectedNamespaces {
	var returns *GkeBackupRestorePlanRestoreConfigSelectedNamespaces
	_jsii_.Get(
		j,
		"selectedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) TransformationRules() GkeBackupRestorePlanRestoreConfigTransformationRulesList {
	var returns GkeBackupRestorePlanRestoreConfigTransformationRulesList
	_jsii_.Get(
		j,
		"transformationRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) TransformationRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformationRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) VolumeDataRestorePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeDataRestorePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) VolumeDataRestorePolicyBindings() GkeBackupRestorePlanRestoreConfigVolumeDataRestorePolicyBindingsList {
	var returns GkeBackupRestorePlanRestoreConfigVolumeDataRestorePolicyBindingsList
	_jsii_.Get(
		j,
		"volumeDataRestorePolicyBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) VolumeDataRestorePolicyBindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeDataRestorePolicyBindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) VolumeDataRestorePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeDataRestorePolicyInput",
		&returns,
	)
	return returns
}


func NewGkeBackupRestorePlanRestoreConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeBackupRestorePlanRestoreConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeBackupRestorePlanRestoreConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupRestorePlan.GkeBackupRestorePlanRestoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeBackupRestorePlanRestoreConfigOutputReference_Override(g GkeBackupRestorePlanRestoreConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupRestorePlan.GkeBackupRestorePlanRestoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetAllNamespaces(val interface{}) {
	if err := j.validateSetAllNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allNamespaces",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetClusterResourceConflictPolicy(val *string) {
	if err := j.validateSetClusterResourceConflictPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterResourceConflictPolicy",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetInternalValue(val *GkeBackupRestorePlanRestoreConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetNamespacedResourceRestoreMode(val *string) {
	if err := j.validateSetNamespacedResourceRestoreModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespacedResourceRestoreMode",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetNoNamespaces(val interface{}) {
	if err := j.validateSetNoNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNamespaces",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference)SetVolumeDataRestorePolicy(val *string) {
	if err := j.validateSetVolumeDataRestorePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeDataRestorePolicy",
		val,
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) PutClusterResourceRestoreScope(value *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope) {
	if err := g.validatePutClusterResourceRestoreScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterResourceRestoreScope",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) PutExcludedNamespaces(value *GkeBackupRestorePlanRestoreConfigExcludedNamespaces) {
	if err := g.validatePutExcludedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludedNamespaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) PutRestoreOrder(value *GkeBackupRestorePlanRestoreConfigRestoreOrder) {
	if err := g.validatePutRestoreOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestoreOrder",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) PutSelectedApplications(value *GkeBackupRestorePlanRestoreConfigSelectedApplications) {
	if err := g.validatePutSelectedApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedApplications",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) PutSelectedNamespaces(value *GkeBackupRestorePlanRestoreConfigSelectedNamespaces) {
	if err := g.validatePutSelectedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedNamespaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) PutTransformationRules(value interface{}) {
	if err := g.validatePutTransformationRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTransformationRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) PutVolumeDataRestorePolicyBindings(value interface{}) {
	if err := g.validatePutVolumeDataRestorePolicyBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVolumeDataRestorePolicyBindings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetAllNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetAllNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetClusterResourceConflictPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterResourceConflictPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetClusterResourceRestoreScope() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterResourceRestoreScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetExcludedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetNamespacedResourceRestoreMode() {
	_jsii_.InvokeVoid(
		g,
		"resetNamespacedResourceRestoreMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetNoNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetNoNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetRestoreOrder() {
	_jsii_.InvokeVoid(
		g,
		"resetRestoreOrder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetSelectedApplications() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedApplications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetSelectedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetTransformationRules() {
	_jsii_.InvokeVoid(
		g,
		"resetTransformationRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetVolumeDataRestorePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeDataRestorePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ResetVolumeDataRestorePolicyBindings() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeDataRestorePolicyBindings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeBackupRestorePlanRestoreConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

