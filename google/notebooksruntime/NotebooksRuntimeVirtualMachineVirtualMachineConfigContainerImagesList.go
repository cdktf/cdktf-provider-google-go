package notebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/notebooksruntime/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList
type jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList {
	_init_.Initialize()

	if err := validateNewNotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList{}

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList_Override(n NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) Get(index *float64) NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

