// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/vmwareengineprivatecloud/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareenginePrivateCloudManagementClusterOutputReference interface {
	cdktf.ComplexObject
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	InternalValue() *VmwareenginePrivateCloudManagementCluster
	SetInternalValue(val *VmwareenginePrivateCloudManagementCluster)
	NodeTypeConfigs() VmwareenginePrivateCloudManagementClusterNodeTypeConfigsList
	NodeTypeConfigsInput() interface{}
	StretchedClusterConfig() VmwareenginePrivateCloudManagementClusterStretchedClusterConfigOutputReference
	StretchedClusterConfigInput() *VmwareenginePrivateCloudManagementClusterStretchedClusterConfig
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
	PutNodeTypeConfigs(value interface{})
	PutStretchedClusterConfig(value *VmwareenginePrivateCloudManagementClusterStretchedClusterConfig)
	ResetNodeTypeConfigs()
	ResetStretchedClusterConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmwareenginePrivateCloudManagementClusterOutputReference
type jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) InternalValue() *VmwareenginePrivateCloudManagementCluster {
	var returns *VmwareenginePrivateCloudManagementCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) NodeTypeConfigs() VmwareenginePrivateCloudManagementClusterNodeTypeConfigsList {
	var returns VmwareenginePrivateCloudManagementClusterNodeTypeConfigsList
	_jsii_.Get(
		j,
		"nodeTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) NodeTypeConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) StretchedClusterConfig() VmwareenginePrivateCloudManagementClusterStretchedClusterConfigOutputReference {
	var returns VmwareenginePrivateCloudManagementClusterStretchedClusterConfigOutputReference
	_jsii_.Get(
		j,
		"stretchedClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) StretchedClusterConfigInput() *VmwareenginePrivateCloudManagementClusterStretchedClusterConfig {
	var returns *VmwareenginePrivateCloudManagementClusterStretchedClusterConfig
	_jsii_.Get(
		j,
		"stretchedClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVmwareenginePrivateCloudManagementClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VmwareenginePrivateCloudManagementClusterOutputReference {
	_init_.Initialize()

	if err := validateNewVmwareenginePrivateCloudManagementClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vmwareenginePrivateCloud.VmwareenginePrivateCloudManagementClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVmwareenginePrivateCloudManagementClusterOutputReference_Override(v VmwareenginePrivateCloudManagementClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vmwareenginePrivateCloud.VmwareenginePrivateCloudManagementClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference)SetInternalValue(val *VmwareenginePrivateCloudManagementCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) PutNodeTypeConfigs(value interface{}) {
	if err := v.validatePutNodeTypeConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNodeTypeConfigs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) PutStretchedClusterConfig(value *VmwareenginePrivateCloudManagementClusterStretchedClusterConfig) {
	if err := v.validatePutStretchedClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStretchedClusterConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ResetNodeTypeConfigs() {
	_jsii_.InvokeVoid(
		v,
		"resetNodeTypeConfigs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ResetStretchedClusterConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetStretchedClusterConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

