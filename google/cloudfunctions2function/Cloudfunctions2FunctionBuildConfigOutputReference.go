// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfunctions2function

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/cloudfunctions2function/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudfunctions2FunctionBuildConfigOutputReference interface {
	cdktf.ComplexObject
	AutomaticUpdatePolicy() Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicyOutputReference
	AutomaticUpdatePolicyInput() *Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy
	BuildAttribute() *string
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
	DockerRepository() *string
	SetDockerRepository(val *string)
	DockerRepositoryInput() *string
	EntryPoint() *string
	SetEntryPoint(val *string)
	EntryPointInput() *string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *Cloudfunctions2FunctionBuildConfig
	SetInternalValue(val *Cloudfunctions2FunctionBuildConfig)
	OnDeployUpdatePolicy() Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicyOutputReference
	OnDeployUpdatePolicyInput() *Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	Source() Cloudfunctions2FunctionBuildConfigSourceOutputReference
	SourceInput() *Cloudfunctions2FunctionBuildConfigSource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerPool() *string
	SetWorkerPool(val *string)
	WorkerPoolInput() *string
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
	PutAutomaticUpdatePolicy(value *Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy)
	PutOnDeployUpdatePolicy(value *Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy)
	PutSource(value *Cloudfunctions2FunctionBuildConfigSource)
	ResetAutomaticUpdatePolicy()
	ResetDockerRepository()
	ResetEntryPoint()
	ResetEnvironmentVariables()
	ResetOnDeployUpdatePolicy()
	ResetRuntime()
	ResetServiceAccount()
	ResetSource()
	ResetWorkerPool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Cloudfunctions2FunctionBuildConfigOutputReference
type jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) AutomaticUpdatePolicy() Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicyOutputReference {
	var returns Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"automaticUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) AutomaticUpdatePolicyInput() *Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy {
	var returns *Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy
	_jsii_.Get(
		j,
		"automaticUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) BuildAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) DockerRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) DockerRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) InternalValue() *Cloudfunctions2FunctionBuildConfig {
	var returns *Cloudfunctions2FunctionBuildConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) OnDeployUpdatePolicy() Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicyOutputReference {
	var returns Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"onDeployUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) OnDeployUpdatePolicyInput() *Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy {
	var returns *Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy
	_jsii_.Get(
		j,
		"onDeployUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) Source() Cloudfunctions2FunctionBuildConfigSourceOutputReference {
	var returns Cloudfunctions2FunctionBuildConfigSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) SourceInput() *Cloudfunctions2FunctionBuildConfigSource {
	var returns *Cloudfunctions2FunctionBuildConfigSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) WorkerPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) WorkerPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPoolInput",
		&returns,
	)
	return returns
}


func NewCloudfunctions2FunctionBuildConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Cloudfunctions2FunctionBuildConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfunctions2FunctionBuildConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctions2Function.Cloudfunctions2FunctionBuildConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfunctions2FunctionBuildConfigOutputReference_Override(c Cloudfunctions2FunctionBuildConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctions2Function.Cloudfunctions2FunctionBuildConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetDockerRepository(val *string) {
	if err := j.validateSetDockerRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRepository",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetInternalValue(val *Cloudfunctions2FunctionBuildConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference)SetWorkerPool(val *string) {
	if err := j.validateSetWorkerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerPool",
		val,
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) PutAutomaticUpdatePolicy(value *Cloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy) {
	if err := c.validatePutAutomaticUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutomaticUpdatePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) PutOnDeployUpdatePolicy(value *Cloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy) {
	if err := c.validatePutOnDeployUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnDeployUpdatePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) PutSource(value *Cloudfunctions2FunctionBuildConfigSource) {
	if err := c.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetAutomaticUpdatePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomaticUpdatePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetDockerRepository() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerRepository",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		c,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetOnDeployUpdatePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetOnDeployUpdatePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetRuntime() {
	_jsii_.InvokeVoid(
		c,
		"resetRuntime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		c,
		"resetSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ResetWorkerPool() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkerPool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_Cloudfunctions2FunctionBuildConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

