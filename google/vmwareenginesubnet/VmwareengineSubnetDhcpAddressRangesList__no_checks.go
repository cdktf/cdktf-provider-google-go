// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vmwareenginesubnet

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VmwareengineSubnetDhcpAddressRangesList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VmwareengineSubnetDhcpAddressRangesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VmwareengineSubnetDhcpAddressRangesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VmwareengineSubnetDhcpAddressRangesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VmwareengineSubnetDhcpAddressRangesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVmwareengineSubnetDhcpAddressRangesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

