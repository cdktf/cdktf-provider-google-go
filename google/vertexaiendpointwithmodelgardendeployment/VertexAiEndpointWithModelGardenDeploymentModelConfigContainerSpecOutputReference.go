// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/vertexaiendpointwithmodelgardendeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	DeploymentTimeout() *string
	SetDeploymentTimeout(val *string)
	DeploymentTimeoutInput() *string
	Env() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnvList
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	GrpcPorts() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPortsList
	GrpcPortsInput() interface{}
	HealthProbe() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeOutputReference
	HealthProbeInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe
	HealthRoute() *string
	SetHealthRoute(val *string)
	HealthRouteInput() *string
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InternalValue() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
	SetInternalValue(val *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec)
	LivenessProbe() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbeOutputReference
	LivenessProbeInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe
	Ports() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPortsList
	PortsInput() interface{}
	PredictRoute() *string
	SetPredictRoute(val *string)
	PredictRouteInput() *string
	SharedMemorySizeMb() *string
	SetSharedMemorySizeMb(val *string)
	SharedMemorySizeMbInput() *string
	StartupProbe() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeOutputReference
	StartupProbeInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe
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
	PutEnv(value interface{})
	PutGrpcPorts(value interface{})
	PutHealthProbe(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe)
	PutLivenessProbe(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe)
	PutPorts(value interface{})
	PutStartupProbe(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe)
	ResetArgs()
	ResetCommand()
	ResetDeploymentTimeout()
	ResetEnv()
	ResetGrpcPorts()
	ResetHealthProbe()
	ResetHealthRoute()
	ResetLivenessProbe()
	ResetPorts()
	ResetPredictRoute()
	ResetSharedMemorySizeMb()
	ResetStartupProbe()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference
type jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) DeploymentTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) DeploymentTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Env() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnvList {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GrpcPorts() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPortsList {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPortsList
	_jsii_.Get(
		j,
		"grpcPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GrpcPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grpcPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthProbe() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeOutputReference
	_jsii_.Get(
		j,
		"healthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthProbeInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe {
	var returns *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe
	_jsii_.Get(
		j,
		"healthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthRoute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthRouteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) InternalValue() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec {
	var returns *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) LivenessProbe() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbeOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) LivenessProbeInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe {
	var returns *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Ports() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPortsList {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPortsList
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PredictRoute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PredictRouteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) SharedMemorySizeMb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedMemorySizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) SharedMemorySizeMbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedMemorySizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) StartupProbe() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) StartupProbeInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe {
	var returns *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference_Override(v VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetDeploymentTimeout(val *string) {
	if err := j.validateSetDeploymentTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentTimeout",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetHealthRoute(val *string) {
	if err := j.validateSetHealthRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthRoute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetInternalValue(val *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetPredictRoute(val *string) {
	if err := j.validateSetPredictRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictRoute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetSharedMemorySizeMb(val *string) {
	if err := j.validateSetSharedMemorySizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedMemorySizeMb",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutEnv(value interface{}) {
	if err := v.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEnv",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutGrpcPorts(value interface{}) {
	if err := v.validatePutGrpcPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGrpcPorts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutHealthProbe(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe) {
	if err := v.validatePutHealthProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putHealthProbe",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutLivenessProbe(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe) {
	if err := v.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutPorts(value interface{}) {
	if err := v.validatePutPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPorts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutStartupProbe(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe) {
	if err := v.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		v,
		"resetArgs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		v,
		"resetCommand",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetDeploymentTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetDeploymentTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		v,
		"resetEnv",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetGrpcPorts() {
	_jsii_.InvokeVoid(
		v,
		"resetGrpcPorts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetHealthProbe() {
	_jsii_.InvokeVoid(
		v,
		"resetHealthProbe",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetHealthRoute() {
	_jsii_.InvokeVoid(
		v,
		"resetHealthRoute",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		v,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		v,
		"resetPorts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetPredictRoute() {
	_jsii_.InvokeVoid(
		v,
		"resetPredictRoute",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetSharedMemorySizeMb() {
	_jsii_.InvokeVoid(
		v,
		"resetSharedMemorySizeMb",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		v,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

