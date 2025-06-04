// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/notebooksruntime/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebooksRuntimeSoftwareConfigOutputReference interface {
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
	CustomGpuDriverPath() *string
	SetCustomGpuDriverPath(val *string)
	CustomGpuDriverPathInput() *string
	EnableHealthMonitoring() interface{}
	SetEnableHealthMonitoring(val interface{})
	EnableHealthMonitoringInput() interface{}
	// Experimental.
	Fqn() *string
	IdleShutdown() interface{}
	SetIdleShutdown(val interface{})
	IdleShutdownInput() interface{}
	IdleShutdownTimeout() *float64
	SetIdleShutdownTimeout(val *float64)
	IdleShutdownTimeoutInput() *float64
	InstallGpuDriver() interface{}
	SetInstallGpuDriver(val interface{})
	InstallGpuDriverInput() interface{}
	InternalValue() *NotebooksRuntimeSoftwareConfig
	SetInternalValue(val *NotebooksRuntimeSoftwareConfig)
	Kernels() NotebooksRuntimeSoftwareConfigKernelsList
	KernelsInput() interface{}
	NotebookUpgradeSchedule() *string
	SetNotebookUpgradeSchedule(val *string)
	NotebookUpgradeScheduleInput() *string
	PostStartupScript() *string
	SetPostStartupScript(val *string)
	PostStartupScriptBehavior() *string
	SetPostStartupScriptBehavior(val *string)
	PostStartupScriptBehaviorInput() *string
	PostStartupScriptInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upgradeable() cdktf.IResolvable
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
	PutKernels(value interface{})
	ResetCustomGpuDriverPath()
	ResetEnableHealthMonitoring()
	ResetIdleShutdown()
	ResetIdleShutdownTimeout()
	ResetInstallGpuDriver()
	ResetKernels()
	ResetNotebookUpgradeSchedule()
	ResetPostStartupScript()
	ResetPostStartupScriptBehavior()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotebooksRuntimeSoftwareConfigOutputReference
type jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) CustomGpuDriverPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) CustomGpuDriverPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) EnableHealthMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHealthMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) EnableHealthMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHealthMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) IdleShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idleShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) IdleShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idleShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) IdleShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) IdleShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) InstallGpuDriver() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) InstallGpuDriverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) InternalValue() *NotebooksRuntimeSoftwareConfig {
	var returns *NotebooksRuntimeSoftwareConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) Kernels() NotebooksRuntimeSoftwareConfigKernelsList {
	var returns NotebooksRuntimeSoftwareConfigKernelsList
	_jsii_.Get(
		j,
		"kernels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) KernelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kernelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) NotebookUpgradeSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookUpgradeSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) NotebookUpgradeScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookUpgradeScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) PostStartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) PostStartupScriptBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) PostStartupScriptBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) PostStartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) Upgradeable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"upgradeable",
		&returns,
	)
	return returns
}


func NewNotebooksRuntimeSoftwareConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotebooksRuntimeSoftwareConfigOutputReference {
	_init_.Initialize()

	if err := validateNewNotebooksRuntimeSoftwareConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotebooksRuntimeSoftwareConfigOutputReference_Override(n NotebooksRuntimeSoftwareConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetCustomGpuDriverPath(val *string) {
	if err := j.validateSetCustomGpuDriverPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customGpuDriverPath",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetEnableHealthMonitoring(val interface{}) {
	if err := j.validateSetEnableHealthMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHealthMonitoring",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetIdleShutdown(val interface{}) {
	if err := j.validateSetIdleShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleShutdown",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetIdleShutdownTimeout(val *float64) {
	if err := j.validateSetIdleShutdownTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetInstallGpuDriver(val interface{}) {
	if err := j.validateSetInstallGpuDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installGpuDriver",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetInternalValue(val *NotebooksRuntimeSoftwareConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetNotebookUpgradeSchedule(val *string) {
	if err := j.validateSetNotebookUpgradeScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookUpgradeSchedule",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetPostStartupScript(val *string) {
	if err := j.validateSetPostStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postStartupScript",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetPostStartupScriptBehavior(val *string) {
	if err := j.validateSetPostStartupScriptBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postStartupScriptBehavior",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) PutKernels(value interface{}) {
	if err := n.validatePutKernelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putKernels",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetCustomGpuDriverPath() {
	_jsii_.InvokeVoid(
		n,
		"resetCustomGpuDriverPath",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetEnableHealthMonitoring() {
	_jsii_.InvokeVoid(
		n,
		"resetEnableHealthMonitoring",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetIdleShutdown() {
	_jsii_.InvokeVoid(
		n,
		"resetIdleShutdown",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetIdleShutdownTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetIdleShutdownTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetInstallGpuDriver() {
	_jsii_.InvokeVoid(
		n,
		"resetInstallGpuDriver",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetKernels() {
	_jsii_.InvokeVoid(
		n,
		"resetKernels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetNotebookUpgradeSchedule() {
	_jsii_.InvokeVoid(
		n,
		"resetNotebookUpgradeSchedule",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetPostStartupScript() {
	_jsii_.InvokeVoid(
		n,
		"resetPostStartupScript",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ResetPostStartupScriptBehavior() {
	_jsii_.InvokeVoid(
		n,
		"resetPostStartupScriptBehavior",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeSoftwareConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

