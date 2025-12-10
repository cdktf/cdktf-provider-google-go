// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package netappvolume

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetappVolumeMountOptionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeMountOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeMountOptionsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeMountOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeMountOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeMountOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetappVolumeMountOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

