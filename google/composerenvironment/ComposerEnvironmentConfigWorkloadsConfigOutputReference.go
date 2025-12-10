// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigWorkloadsConfigOutputReference interface {
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
	DagProcessor() ComposerEnvironmentConfigWorkloadsConfigDagProcessorOutputReference
	DagProcessorInput() *ComposerEnvironmentConfigWorkloadsConfigDagProcessor
	// Experimental.
	Fqn() *string
	InternalValue() *ComposerEnvironmentConfigWorkloadsConfig
	SetInternalValue(val *ComposerEnvironmentConfigWorkloadsConfig)
	Scheduler() ComposerEnvironmentConfigWorkloadsConfigSchedulerOutputReference
	SchedulerInput() *ComposerEnvironmentConfigWorkloadsConfigScheduler
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Triggerer() ComposerEnvironmentConfigWorkloadsConfigTriggererOutputReference
	TriggererInput() *ComposerEnvironmentConfigWorkloadsConfigTriggerer
	WebServer() ComposerEnvironmentConfigWorkloadsConfigWebServerOutputReference
	WebServerInput() *ComposerEnvironmentConfigWorkloadsConfigWebServer
	Worker() ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference
	WorkerInput() *ComposerEnvironmentConfigWorkloadsConfigWorker
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
	PutDagProcessor(value *ComposerEnvironmentConfigWorkloadsConfigDagProcessor)
	PutScheduler(value *ComposerEnvironmentConfigWorkloadsConfigScheduler)
	PutTriggerer(value *ComposerEnvironmentConfigWorkloadsConfigTriggerer)
	PutWebServer(value *ComposerEnvironmentConfigWorkloadsConfigWebServer)
	PutWorker(value *ComposerEnvironmentConfigWorkloadsConfigWorker)
	ResetDagProcessor()
	ResetScheduler()
	ResetTriggerer()
	ResetWebServer()
	ResetWorker()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigWorkloadsConfigOutputReference
type jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) DagProcessor() ComposerEnvironmentConfigWorkloadsConfigDagProcessorOutputReference {
	var returns ComposerEnvironmentConfigWorkloadsConfigDagProcessorOutputReference
	_jsii_.Get(
		j,
		"dagProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) DagProcessorInput() *ComposerEnvironmentConfigWorkloadsConfigDagProcessor {
	var returns *ComposerEnvironmentConfigWorkloadsConfigDagProcessor
	_jsii_.Get(
		j,
		"dagProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) InternalValue() *ComposerEnvironmentConfigWorkloadsConfig {
	var returns *ComposerEnvironmentConfigWorkloadsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) Scheduler() ComposerEnvironmentConfigWorkloadsConfigSchedulerOutputReference {
	var returns ComposerEnvironmentConfigWorkloadsConfigSchedulerOutputReference
	_jsii_.Get(
		j,
		"scheduler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) SchedulerInput() *ComposerEnvironmentConfigWorkloadsConfigScheduler {
	var returns *ComposerEnvironmentConfigWorkloadsConfigScheduler
	_jsii_.Get(
		j,
		"schedulerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) Triggerer() ComposerEnvironmentConfigWorkloadsConfigTriggererOutputReference {
	var returns ComposerEnvironmentConfigWorkloadsConfigTriggererOutputReference
	_jsii_.Get(
		j,
		"triggerer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) TriggererInput() *ComposerEnvironmentConfigWorkloadsConfigTriggerer {
	var returns *ComposerEnvironmentConfigWorkloadsConfigTriggerer
	_jsii_.Get(
		j,
		"triggererInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) WebServer() ComposerEnvironmentConfigWorkloadsConfigWebServerOutputReference {
	var returns ComposerEnvironmentConfigWorkloadsConfigWebServerOutputReference
	_jsii_.Get(
		j,
		"webServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) WebServerInput() *ComposerEnvironmentConfigWorkloadsConfigWebServer {
	var returns *ComposerEnvironmentConfigWorkloadsConfigWebServer
	_jsii_.Get(
		j,
		"webServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) Worker() ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference {
	var returns ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference
	_jsii_.Get(
		j,
		"worker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) WorkerInput() *ComposerEnvironmentConfigWorkloadsConfigWorker {
	var returns *ComposerEnvironmentConfigWorkloadsConfigWorker
	_jsii_.Get(
		j,
		"workerInput",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigWorkloadsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComposerEnvironmentConfigWorkloadsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigWorkloadsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigWorkloadsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigWorkloadsConfigOutputReference_Override(c ComposerEnvironmentConfigWorkloadsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigWorkloadsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference)SetInternalValue(val *ComposerEnvironmentConfigWorkloadsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) PutDagProcessor(value *ComposerEnvironmentConfigWorkloadsConfigDagProcessor) {
	if err := c.validatePutDagProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDagProcessor",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) PutScheduler(value *ComposerEnvironmentConfigWorkloadsConfigScheduler) {
	if err := c.validatePutSchedulerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScheduler",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) PutTriggerer(value *ComposerEnvironmentConfigWorkloadsConfigTriggerer) {
	if err := c.validatePutTriggererParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTriggerer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) PutWebServer(value *ComposerEnvironmentConfigWorkloadsConfigWebServer) {
	if err := c.validatePutWebServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWebServer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) PutWorker(value *ComposerEnvironmentConfigWorkloadsConfigWorker) {
	if err := c.validatePutWorkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorker",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetDagProcessor() {
	_jsii_.InvokeVoid(
		c,
		"resetDagProcessor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetScheduler() {
	_jsii_.InvokeVoid(
		c,
		"resetScheduler",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetTriggerer() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetWebServer() {
	_jsii_.InvokeVoid(
		c,
		"resetWebServer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetWorker() {
	_jsii_.InvokeVoid(
		c,
		"resetWorker",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

