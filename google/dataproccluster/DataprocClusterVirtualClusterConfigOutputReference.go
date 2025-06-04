// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/dataproccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterVirtualClusterConfigOutputReference interface {
	cdktf.ComplexObject
	AuxiliaryServicesConfig() DataprocClusterVirtualClusterConfigAuxiliaryServicesConfigOutputReference
	AuxiliaryServicesConfigInput() *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfig
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
	InternalValue() *DataprocClusterVirtualClusterConfig
	SetInternalValue(val *DataprocClusterVirtualClusterConfig)
	KubernetesClusterConfig() DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference
	KubernetesClusterConfigInput() *DataprocClusterVirtualClusterConfigKubernetesClusterConfig
	StagingBucket() *string
	SetStagingBucket(val *string)
	StagingBucketInput() *string
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
	PutAuxiliaryServicesConfig(value *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfig)
	PutKubernetesClusterConfig(value *DataprocClusterVirtualClusterConfigKubernetesClusterConfig)
	ResetAuxiliaryServicesConfig()
	ResetKubernetesClusterConfig()
	ResetStagingBucket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterVirtualClusterConfigOutputReference
type jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) AuxiliaryServicesConfig() DataprocClusterVirtualClusterConfigAuxiliaryServicesConfigOutputReference {
	var returns DataprocClusterVirtualClusterConfigAuxiliaryServicesConfigOutputReference
	_jsii_.Get(
		j,
		"auxiliaryServicesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) AuxiliaryServicesConfigInput() *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfig {
	var returns *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfig
	_jsii_.Get(
		j,
		"auxiliaryServicesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) InternalValue() *DataprocClusterVirtualClusterConfig {
	var returns *DataprocClusterVirtualClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) KubernetesClusterConfig() DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference {
	var returns DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference
	_jsii_.Get(
		j,
		"kubernetesClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) KubernetesClusterConfigInput() *DataprocClusterVirtualClusterConfigKubernetesClusterConfig {
	var returns *DataprocClusterVirtualClusterConfigKubernetesClusterConfig
	_jsii_.Get(
		j,
		"kubernetesClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) StagingBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) StagingBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocClusterVirtualClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterVirtualClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterVirtualClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterVirtualClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterVirtualClusterConfigOutputReference_Override(d DataprocClusterVirtualClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterVirtualClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference)SetInternalValue(val *DataprocClusterVirtualClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference)SetStagingBucket(val *string) {
	if err := j.validateSetStagingBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingBucket",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) PutAuxiliaryServicesConfig(value *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfig) {
	if err := d.validatePutAuxiliaryServicesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuxiliaryServicesConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) PutKubernetesClusterConfig(value *DataprocClusterVirtualClusterConfigKubernetesClusterConfig) {
	if err := d.validatePutKubernetesClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKubernetesClusterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) ResetAuxiliaryServicesConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAuxiliaryServicesConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) ResetKubernetesClusterConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetKubernetesClusterConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) ResetStagingBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetStagingBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

