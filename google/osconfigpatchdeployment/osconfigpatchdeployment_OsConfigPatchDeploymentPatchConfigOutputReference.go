package osconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v3/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v3/osconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigPatchDeploymentPatchConfigOutputReference interface {
	cdktf.ComplexObject
	Apt() OsConfigPatchDeploymentPatchConfigAptOutputReference
	AptInput() *OsConfigPatchDeploymentPatchConfigApt
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
	Goo() OsConfigPatchDeploymentPatchConfigGooOutputReference
	GooInput() *OsConfigPatchDeploymentPatchConfigGoo
	InternalValue() *OsConfigPatchDeploymentPatchConfig
	SetInternalValue(val *OsConfigPatchDeploymentPatchConfig)
	MigInstancesAllowed() interface{}
	SetMigInstancesAllowed(val interface{})
	MigInstancesAllowedInput() interface{}
	PostStep() OsConfigPatchDeploymentPatchConfigPostStepOutputReference
	PostStepInput() *OsConfigPatchDeploymentPatchConfigPostStep
	PreStep() OsConfigPatchDeploymentPatchConfigPreStepOutputReference
	PreStepInput() *OsConfigPatchDeploymentPatchConfigPreStep
	RebootConfig() *string
	SetRebootConfig(val *string)
	RebootConfigInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsUpdate() OsConfigPatchDeploymentPatchConfigWindowsUpdateOutputReference
	WindowsUpdateInput() *OsConfigPatchDeploymentPatchConfigWindowsUpdate
	Yum() OsConfigPatchDeploymentPatchConfigYumOutputReference
	YumInput() *OsConfigPatchDeploymentPatchConfigYum
	Zypper() OsConfigPatchDeploymentPatchConfigZypperOutputReference
	ZypperInput() *OsConfigPatchDeploymentPatchConfigZypper
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
	PutApt(value *OsConfigPatchDeploymentPatchConfigApt)
	PutGoo(value *OsConfigPatchDeploymentPatchConfigGoo)
	PutPostStep(value *OsConfigPatchDeploymentPatchConfigPostStep)
	PutPreStep(value *OsConfigPatchDeploymentPatchConfigPreStep)
	PutWindowsUpdate(value *OsConfigPatchDeploymentPatchConfigWindowsUpdate)
	PutYum(value *OsConfigPatchDeploymentPatchConfigYum)
	PutZypper(value *OsConfigPatchDeploymentPatchConfigZypper)
	ResetApt()
	ResetGoo()
	ResetMigInstancesAllowed()
	ResetPostStep()
	ResetPreStep()
	ResetRebootConfig()
	ResetWindowsUpdate()
	ResetYum()
	ResetZypper()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigPatchDeploymentPatchConfigOutputReference
type jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) Apt() OsConfigPatchDeploymentPatchConfigAptOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigAptOutputReference
	_jsii_.Get(
		j,
		"apt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) AptInput() *OsConfigPatchDeploymentPatchConfigApt {
	var returns *OsConfigPatchDeploymentPatchConfigApt
	_jsii_.Get(
		j,
		"aptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) Goo() OsConfigPatchDeploymentPatchConfigGooOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigGooOutputReference
	_jsii_.Get(
		j,
		"goo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GooInput() *OsConfigPatchDeploymentPatchConfigGoo {
	var returns *OsConfigPatchDeploymentPatchConfigGoo
	_jsii_.Get(
		j,
		"gooInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) InternalValue() *OsConfigPatchDeploymentPatchConfig {
	var returns *OsConfigPatchDeploymentPatchConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) MigInstancesAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"migInstancesAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) MigInstancesAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"migInstancesAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PostStep() OsConfigPatchDeploymentPatchConfigPostStepOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigPostStepOutputReference
	_jsii_.Get(
		j,
		"postStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PostStepInput() *OsConfigPatchDeploymentPatchConfigPostStep {
	var returns *OsConfigPatchDeploymentPatchConfigPostStep
	_jsii_.Get(
		j,
		"postStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PreStep() OsConfigPatchDeploymentPatchConfigPreStepOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigPreStepOutputReference
	_jsii_.Get(
		j,
		"preStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PreStepInput() *OsConfigPatchDeploymentPatchConfigPreStep {
	var returns *OsConfigPatchDeploymentPatchConfigPreStep
	_jsii_.Get(
		j,
		"preStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) RebootConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) RebootConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) WindowsUpdate() OsConfigPatchDeploymentPatchConfigWindowsUpdateOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigWindowsUpdateOutputReference
	_jsii_.Get(
		j,
		"windowsUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) WindowsUpdateInput() *OsConfigPatchDeploymentPatchConfigWindowsUpdate {
	var returns *OsConfigPatchDeploymentPatchConfigWindowsUpdate
	_jsii_.Get(
		j,
		"windowsUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) Yum() OsConfigPatchDeploymentPatchConfigYumOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigYumOutputReference
	_jsii_.Get(
		j,
		"yum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) YumInput() *OsConfigPatchDeploymentPatchConfigYum {
	var returns *OsConfigPatchDeploymentPatchConfigYum
	_jsii_.Get(
		j,
		"yumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) Zypper() OsConfigPatchDeploymentPatchConfigZypperOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigZypperOutputReference
	_jsii_.Get(
		j,
		"zypper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ZypperInput() *OsConfigPatchDeploymentPatchConfigZypper {
	var returns *OsConfigPatchDeploymentPatchConfigZypper
	_jsii_.Get(
		j,
		"zypperInput",
		&returns,
	)
	return returns
}


func NewOsConfigPatchDeploymentPatchConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigPatchDeploymentPatchConfigOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigPatchDeploymentPatchConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentPatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigPatchDeploymentPatchConfigOutputReference_Override(o OsConfigPatchDeploymentPatchConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentPatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference)SetInternalValue(val *OsConfigPatchDeploymentPatchConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference)SetMigInstancesAllowed(val interface{}) {
	if err := j.validateSetMigInstancesAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migInstancesAllowed",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference)SetRebootConfig(val *string) {
	if err := j.validateSetRebootConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rebootConfig",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PutApt(value *OsConfigPatchDeploymentPatchConfigApt) {
	if err := o.validatePutAptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putApt",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PutGoo(value *OsConfigPatchDeploymentPatchConfigGoo) {
	if err := o.validatePutGooParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoo",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PutPostStep(value *OsConfigPatchDeploymentPatchConfigPostStep) {
	if err := o.validatePutPostStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPostStep",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PutPreStep(value *OsConfigPatchDeploymentPatchConfigPreStep) {
	if err := o.validatePutPreStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPreStep",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PutWindowsUpdate(value *OsConfigPatchDeploymentPatchConfigWindowsUpdate) {
	if err := o.validatePutWindowsUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWindowsUpdate",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PutYum(value *OsConfigPatchDeploymentPatchConfigYum) {
	if err := o.validatePutYumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putYum",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) PutZypper(value *OsConfigPatchDeploymentPatchConfigZypper) {
	if err := o.validatePutZypperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putZypper",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetApt() {
	_jsii_.InvokeVoid(
		o,
		"resetApt",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetGoo() {
	_jsii_.InvokeVoid(
		o,
		"resetGoo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetMigInstancesAllowed() {
	_jsii_.InvokeVoid(
		o,
		"resetMigInstancesAllowed",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetPostStep() {
	_jsii_.InvokeVoid(
		o,
		"resetPostStep",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetPreStep() {
	_jsii_.InvokeVoid(
		o,
		"resetPreStep",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetRebootConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetRebootConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetWindowsUpdate() {
	_jsii_.InvokeVoid(
		o,
		"resetWindowsUpdate",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetYum() {
	_jsii_.InvokeVoid(
		o,
		"resetYum",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ResetZypper() {
	_jsii_.InvokeVoid(
		o,
		"resetZypper",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

