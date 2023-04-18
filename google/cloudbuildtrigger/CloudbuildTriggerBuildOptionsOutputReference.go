package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/cloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildTriggerBuildOptionsOutputReference interface {
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
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	DynamicSubstitutions() interface{}
	SetDynamicSubstitutions(val interface{})
	DynamicSubstitutionsInput() interface{}
	Env() *[]*string
	SetEnv(val *[]*string)
	EnvInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudbuildTriggerBuildOptions
	SetInternalValue(val *CloudbuildTriggerBuildOptions)
	Logging() *string
	SetLogging(val *string)
	LoggingInput() *string
	LogStreamingOption() *string
	SetLogStreamingOption(val *string)
	LogStreamingOptionInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	RequestedVerifyOption() *string
	SetRequestedVerifyOption(val *string)
	RequestedVerifyOptionInput() *string
	SecretEnv() *[]*string
	SetSecretEnv(val *[]*string)
	SecretEnvInput() *[]*string
	SourceProvenanceHash() *[]*string
	SetSourceProvenanceHash(val *[]*string)
	SourceProvenanceHashInput() *[]*string
	SubstitutionOption() *string
	SetSubstitutionOption(val *string)
	SubstitutionOptionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() CloudbuildTriggerBuildOptionsVolumesList
	VolumesInput() interface{}
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
	PutVolumes(value interface{})
	ResetDiskSizeGb()
	ResetDynamicSubstitutions()
	ResetEnv()
	ResetLogging()
	ResetLogStreamingOption()
	ResetMachineType()
	ResetRequestedVerifyOption()
	ResetSecretEnv()
	ResetSourceProvenanceHash()
	ResetSubstitutionOption()
	ResetVolumes()
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

// The jsii proxy struct for CloudbuildTriggerBuildOptionsOutputReference
type jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) DynamicSubstitutions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) DynamicSubstitutionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicSubstitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) Env() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) EnvInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) InternalValue() *CloudbuildTriggerBuildOptions {
	var returns *CloudbuildTriggerBuildOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) Logging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) LoggingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) LogStreamingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) LogStreamingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) RequestedVerifyOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedVerifyOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) RequestedVerifyOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedVerifyOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) SecretEnv() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) SecretEnvInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) SourceProvenanceHash() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceProvenanceHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) SourceProvenanceHashInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceProvenanceHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) SubstitutionOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"substitutionOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) SubstitutionOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"substitutionOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) Volumes() CloudbuildTriggerBuildOptionsVolumesList {
	var returns CloudbuildTriggerBuildOptionsVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) WorkerPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) WorkerPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPoolInput",
		&returns,
	)
	return returns
}


func NewCloudbuildTriggerBuildOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudbuildTriggerBuildOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerBuildOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildTriggerBuildOptionsOutputReference_Override(c CloudbuildTriggerBuildOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetDynamicSubstitutions(val interface{}) {
	if err := j.validateSetDynamicSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicSubstitutions",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetEnv(val *[]*string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetInternalValue(val *CloudbuildTriggerBuildOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetLogging(val *string) {
	if err := j.validateSetLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetLogStreamingOption(val *string) {
	if err := j.validateSetLogStreamingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamingOption",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetRequestedVerifyOption(val *string) {
	if err := j.validateSetRequestedVerifyOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedVerifyOption",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetSecretEnv(val *[]*string) {
	if err := j.validateSetSecretEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretEnv",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetSourceProvenanceHash(val *[]*string) {
	if err := j.validateSetSourceProvenanceHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceProvenanceHash",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetSubstitutionOption(val *string) {
	if err := j.validateSetSubstitutionOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substitutionOption",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference)SetWorkerPool(val *string) {
	if err := j.validateSetWorkerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerPool",
		val,
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) PutVolumes(value interface{}) {
	if err := c.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetDynamicSubstitutions() {
	_jsii_.InvokeVoid(
		c,
		"resetDynamicSubstitutions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		c,
		"resetEnv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetLogging() {
	_jsii_.InvokeVoid(
		c,
		"resetLogging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetLogStreamingOption() {
	_jsii_.InvokeVoid(
		c,
		"resetLogStreamingOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetRequestedVerifyOption() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestedVerifyOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetSecretEnv() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretEnv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetSourceProvenanceHash() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceProvenanceHash",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetSubstitutionOption() {
	_jsii_.InvokeVoid(
		c,
		"resetSubstitutionOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ResetWorkerPool() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkerPool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

