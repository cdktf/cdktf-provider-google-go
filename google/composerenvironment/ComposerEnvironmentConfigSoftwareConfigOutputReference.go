// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigSoftwareConfigOutputReference interface {
	cdktf.ComplexObject
	AirflowConfigOverrides() *map[string]*string
	SetAirflowConfigOverrides(val *map[string]*string)
	AirflowConfigOverridesInput() *map[string]*string
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
	EnvVariables() *map[string]*string
	SetEnvVariables(val *map[string]*string)
	EnvVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	ImageVersion() *string
	SetImageVersion(val *string)
	ImageVersionInput() *string
	InternalValue() *ComposerEnvironmentConfigSoftwareConfig
	SetInternalValue(val *ComposerEnvironmentConfigSoftwareConfig)
	PypiPackages() *map[string]*string
	SetPypiPackages(val *map[string]*string)
	PypiPackagesInput() *map[string]*string
	PythonVersion() *string
	SetPythonVersion(val *string)
	PythonVersionInput() *string
	SchedulerCount() *float64
	SetSchedulerCount(val *float64)
	SchedulerCountInput() *float64
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
	ResetAirflowConfigOverrides()
	ResetEnvVariables()
	ResetImageVersion()
	ResetPypiPackages()
	ResetPythonVersion()
	ResetSchedulerCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigSoftwareConfigOutputReference
type jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) AirflowConfigOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) AirflowConfigOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) EnvVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) EnvVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ImageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ImageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) InternalValue() *ComposerEnvironmentConfigSoftwareConfig {
	var returns *ComposerEnvironmentConfigSoftwareConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) PypiPackages() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pypiPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) PypiPackagesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pypiPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) SchedulerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) SchedulerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigSoftwareConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComposerEnvironmentConfigSoftwareConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigSoftwareConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigSoftwareConfigOutputReference_Override(c ComposerEnvironmentConfigSoftwareConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetAirflowConfigOverrides(val *map[string]*string) {
	if err := j.validateSetAirflowConfigOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"airflowConfigOverrides",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetEnvVariables(val *map[string]*string) {
	if err := j.validateSetEnvVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envVariables",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetImageVersion(val *string) {
	if err := j.validateSetImageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageVersion",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetInternalValue(val *ComposerEnvironmentConfigSoftwareConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetPypiPackages(val *map[string]*string) {
	if err := j.validateSetPypiPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pypiPackages",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetPythonVersion(val *string) {
	if err := j.validateSetPythonVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetSchedulerCount(val *float64) {
	if err := j.validateSetSchedulerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulerCount",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ResetAirflowConfigOverrides() {
	_jsii_.InvokeVoid(
		c,
		"resetAirflowConfigOverrides",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ResetEnvVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ResetImageVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetImageVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ResetPypiPackages() {
	_jsii_.InvokeVoid(
		c,
		"resetPypiPackages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetPythonVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ResetSchedulerCount() {
	_jsii_.InvokeVoid(
		c,
		"resetSchedulerCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigSoftwareConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

