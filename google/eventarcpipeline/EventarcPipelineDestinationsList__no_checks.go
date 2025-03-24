// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventarcpipeline

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventarcPipelineDestinationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EventarcPipelineDestinationsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventarcPipelineDestinationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventarcPipelineDestinationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventarcPipelineDestinationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventarcPipelineDestinationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventarcPipelineDestinationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventarcPipelineDestinationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

