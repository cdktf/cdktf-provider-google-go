package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobOutputReference interface {
	cdktf.ComplexObject
	Actions() DataLossPreventionJobTriggerInspectJobActionsList
	ActionsInput() interface{}
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
	InspectConfig() DataLossPreventionJobTriggerInspectJobInspectConfigOutputReference
	InspectConfigInput() *DataLossPreventionJobTriggerInspectJobInspectConfig
	InspectTemplateName() *string
	SetInspectTemplateName(val *string)
	InspectTemplateNameInput() *string
	InternalValue() *DataLossPreventionJobTriggerInspectJob
	SetInternalValue(val *DataLossPreventionJobTriggerInspectJob)
	StorageConfig() DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference
	StorageConfigInput() *DataLossPreventionJobTriggerInspectJobStorageConfig
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
	PutActions(value interface{})
	PutInspectConfig(value *DataLossPreventionJobTriggerInspectJobInspectConfig)
	PutStorageConfig(value *DataLossPreventionJobTriggerInspectJobStorageConfig)
	ResetActions()
	ResetInspectConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) Actions() DataLossPreventionJobTriggerInspectJobActionsList {
	var returns DataLossPreventionJobTriggerInspectJobActionsList
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) ActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) InspectConfig() DataLossPreventionJobTriggerInspectJobInspectConfigOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigOutputReference
	_jsii_.Get(
		j,
		"inspectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) InspectConfigInput() *DataLossPreventionJobTriggerInspectJobInspectConfig {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfig
	_jsii_.Get(
		j,
		"inspectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) InspectTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) InspectTemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) InternalValue() *DataLossPreventionJobTriggerInspectJob {
	var returns *DataLossPreventionJobTriggerInspectJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) StorageConfig() DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) StorageConfigInput() *DataLossPreventionJobTriggerInspectJobStorageConfig {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfig
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionJobTriggerInspectJobOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobOutputReference_Override(d DataLossPreventionJobTriggerInspectJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference)SetInspectTemplateName(val *string) {
	if err := j.validateSetInspectTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectTemplateName",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference)SetInternalValue(val *DataLossPreventionJobTriggerInspectJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) PutActions(value interface{}) {
	if err := d.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putActions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) PutInspectConfig(value *DataLossPreventionJobTriggerInspectJobInspectConfig) {
	if err := d.validatePutInspectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInspectConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) PutStorageConfig(value *DataLossPreventionJobTriggerInspectJobStorageConfig) {
	if err := d.validatePutStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		d,
		"resetActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) ResetInspectConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetInspectConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

