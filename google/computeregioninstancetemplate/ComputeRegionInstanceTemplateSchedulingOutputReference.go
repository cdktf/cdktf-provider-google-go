// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeregioninstancetemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionInstanceTemplateSchedulingOutputReference interface {
	cdktf.ComplexObject
	AutomaticRestart() interface{}
	SetAutomaticRestart(val interface{})
	AutomaticRestartInput() interface{}
	AvailabilityDomain() *float64
	SetAvailabilityDomain(val *float64)
	AvailabilityDomainInput() *float64
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
	InstanceTerminationAction() *string
	SetInstanceTerminationAction(val *string)
	InstanceTerminationActionInput() *string
	InternalValue() *ComputeRegionInstanceTemplateScheduling
	SetInternalValue(val *ComputeRegionInstanceTemplateScheduling)
	LocalSsdRecoveryTimeout() ComputeRegionInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList
	LocalSsdRecoveryTimeoutInput() interface{}
	MaxRunDuration() ComputeRegionInstanceTemplateSchedulingMaxRunDurationOutputReference
	MaxRunDurationInput() *ComputeRegionInstanceTemplateSchedulingMaxRunDuration
	MinNodeCpus() *float64
	SetMinNodeCpus(val *float64)
	MinNodeCpusInput() *float64
	NodeAffinities() ComputeRegionInstanceTemplateSchedulingNodeAffinitiesList
	NodeAffinitiesInput() interface{}
	OnHostMaintenance() *string
	SetOnHostMaintenance(val *string)
	OnHostMaintenanceInput() *string
	OnInstanceStopAction() ComputeRegionInstanceTemplateSchedulingOnInstanceStopActionOutputReference
	OnInstanceStopActionInput() *ComputeRegionInstanceTemplateSchedulingOnInstanceStopAction
	Preemptible() interface{}
	SetPreemptible(val interface{})
	PreemptibleInput() interface{}
	ProvisioningModel() *string
	SetProvisioningModel(val *string)
	ProvisioningModelInput() *string
	TerminationTime() *string
	SetTerminationTime(val *string)
	TerminationTimeInput() *string
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
	PutLocalSsdRecoveryTimeout(value interface{})
	PutMaxRunDuration(value *ComputeRegionInstanceTemplateSchedulingMaxRunDuration)
	PutNodeAffinities(value interface{})
	PutOnInstanceStopAction(value *ComputeRegionInstanceTemplateSchedulingOnInstanceStopAction)
	ResetAutomaticRestart()
	ResetAvailabilityDomain()
	ResetInstanceTerminationAction()
	ResetLocalSsdRecoveryTimeout()
	ResetMaxRunDuration()
	ResetMinNodeCpus()
	ResetNodeAffinities()
	ResetOnHostMaintenance()
	ResetOnInstanceStopAction()
	ResetPreemptible()
	ResetProvisioningModel()
	ResetTerminationTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionInstanceTemplateSchedulingOutputReference
type jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) AutomaticRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) AutomaticRestartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) AvailabilityDomain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) AvailabilityDomainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) InstanceTerminationAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) InstanceTerminationActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) InternalValue() *ComputeRegionInstanceTemplateScheduling {
	var returns *ComputeRegionInstanceTemplateScheduling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) LocalSsdRecoveryTimeout() ComputeRegionInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList {
	var returns ComputeRegionInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) LocalSsdRecoveryTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) MaxRunDuration() ComputeRegionInstanceTemplateSchedulingMaxRunDurationOutputReference {
	var returns ComputeRegionInstanceTemplateSchedulingMaxRunDurationOutputReference
	_jsii_.Get(
		j,
		"maxRunDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) MaxRunDurationInput() *ComputeRegionInstanceTemplateSchedulingMaxRunDuration {
	var returns *ComputeRegionInstanceTemplateSchedulingMaxRunDuration
	_jsii_.Get(
		j,
		"maxRunDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) MinNodeCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) MinNodeCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) NodeAffinities() ComputeRegionInstanceTemplateSchedulingNodeAffinitiesList {
	var returns ComputeRegionInstanceTemplateSchedulingNodeAffinitiesList
	_jsii_.Get(
		j,
		"nodeAffinities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) NodeAffinitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeAffinitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) OnHostMaintenance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) OnHostMaintenanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) OnInstanceStopAction() ComputeRegionInstanceTemplateSchedulingOnInstanceStopActionOutputReference {
	var returns ComputeRegionInstanceTemplateSchedulingOnInstanceStopActionOutputReference
	_jsii_.Get(
		j,
		"onInstanceStopAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) OnInstanceStopActionInput() *ComputeRegionInstanceTemplateSchedulingOnInstanceStopAction {
	var returns *ComputeRegionInstanceTemplateSchedulingOnInstanceStopAction
	_jsii_.Get(
		j,
		"onInstanceStopActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ProvisioningModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ProvisioningModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) TerminationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) TerminationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionInstanceTemplateSchedulingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionInstanceTemplateSchedulingOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionInstanceTemplateSchedulingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionInstanceTemplate.ComputeRegionInstanceTemplateSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionInstanceTemplateSchedulingOutputReference_Override(c ComputeRegionInstanceTemplateSchedulingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionInstanceTemplate.ComputeRegionInstanceTemplateSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetAutomaticRestart(val interface{}) {
	if err := j.validateSetAutomaticRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticRestart",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetAvailabilityDomain(val *float64) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetInstanceTerminationAction(val *string) {
	if err := j.validateSetInstanceTerminationActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTerminationAction",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetInternalValue(val *ComputeRegionInstanceTemplateScheduling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetMinNodeCpus(val *float64) {
	if err := j.validateSetMinNodeCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCpus",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetOnHostMaintenance(val *string) {
	if err := j.validateSetOnHostMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onHostMaintenance",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetProvisioningModel(val *string) {
	if err := j.validateSetProvisioningModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningModel",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetTerminationTime(val *string) {
	if err := j.validateSetTerminationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationTime",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) PutLocalSsdRecoveryTimeout(value interface{}) {
	if err := c.validatePutLocalSsdRecoveryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocalSsdRecoveryTimeout",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) PutMaxRunDuration(value *ComputeRegionInstanceTemplateSchedulingMaxRunDuration) {
	if err := c.validatePutMaxRunDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxRunDuration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) PutNodeAffinities(value interface{}) {
	if err := c.validatePutNodeAffinitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeAffinities",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) PutOnInstanceStopAction(value *ComputeRegionInstanceTemplateSchedulingOnInstanceStopAction) {
	if err := c.validatePutOnInstanceStopActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnInstanceStopAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetAutomaticRestart() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomaticRestart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetInstanceTerminationAction() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceTerminationAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetLocalSsdRecoveryTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalSsdRecoveryTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetMaxRunDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRunDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetMinNodeCpus() {
	_jsii_.InvokeVoid(
		c,
		"resetMinNodeCpus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetNodeAffinities() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeAffinities",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetOnHostMaintenance() {
	_jsii_.InvokeVoid(
		c,
		"resetOnHostMaintenance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetOnInstanceStopAction() {
	_jsii_.InvokeVoid(
		c,
		"resetOnInstanceStopAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		c,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetProvisioningModel() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisioningModel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ResetTerminationTime() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateSchedulingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

