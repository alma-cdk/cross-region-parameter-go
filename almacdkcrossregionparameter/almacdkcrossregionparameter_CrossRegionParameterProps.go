// Store AWS SSM Parameter Store Parameters into another AWS Region with AWS CDK
package almacdkcrossregionparameter

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Experimental.
type CrossRegionParameterProps struct {
	// Information about the SSM Parameter that you want to add.
	//
	// Required by this construct (AWS considers it as optional).
	//
	// Example:
	//   'Some message for the Swedes'
	//
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// SSM Parameter name.
	//
	// Example:
	//   '/parameter/path/message'
	//
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Target region for the parameter.
	//
	// Must be some other region than the current Stack's region.
	//
	// Example:
	//   'eu-north-1'
	//
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The SSM Parameter value that you want to add.
	//
	// Limits:
	// - Standard parameters have a value limit of 4 KB.
	// - Advanced parameters have a value limit of 8 KB.
	//
	// Example:
	//   'Hej då!'
	//
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// A regular expression used to validate the SSM Parameter Value.
	//
	// For example, for String types with values restricted to numbers,
	// you can specify the following: `^\d+$`.
	//
	// Example:
	//   '^\d+$'
	//
	// See: https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PutParameter.html#systemsmanager-PutParameter-request-AllowedPattern
	//
	// Experimental.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// The AWS Key Management Service (AWS KMS) ID that you want to use to encrypt a parameter.
	//
	// Either the default AWS KMS key automatically assigned to your AWS account or a custom key. Required for parameters that use the SecureString data type.
	//
	// The KMS Key must exists in the target region.
	//
	// If you don't specify a key ID, the system uses the default key associated with your AWS account.
	//
	// Example:
	//   '1234abcd-12ab-34cd-56ef-1234567890ab'
	//
	// See: https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PutParameter.html#systemsmanager-PutParameter-request-KeyId
	//
	// Experimental.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// The SSM Parameter Tier to assign to a parameter.
	//
	// - Parameter Store offers a standard tier and an advanced tier for parameters. Standard parameters have a content size limit of 4 KB and can't be configured to use parameter policies. You can create a maximum of 10,000 standard parameters for each Region in an AWS account. Standard parameters are offered at no additional cost.
	// - Advanced parameters have a content size limit of 8 KB and can be configured to use parameter policies. You can create a maximum of 100,000 advanced parameters for each Region in an AWS account. Advanced parameters incur a charge. For more information, see Standard and advanced parameter tiers in the AWS Systems Manager User Guide.
	// - You can change a standard parameter to an advanced parameter any time. But you can't revert an advanced parameter to a standard parameter. Reverting an advanced parameter to a standard parameter would result in data loss because the system would truncate the size of the parameter from 8 KB to 4 KB. Reverting would also remove any policies attached to the parameter. Lastly, advanced parameters use a different form of encryption than standard parameters.
	// - If you no longer need an advanced parameter, or if you no longer want to incur charges for an advanced parameter, you must delete it and recreate it as a new standard parameter.
	//
	// Example:
	//   ParameterTier.INTELLIGENT_TIERING
	//
	// Experimental.
	ParameterTier awsssm.ParameterTier `field:"optional" json:"parameterTier" yaml:"parameterTier"`
	// The type of SSM Parameter that you want to add.
	//
	// Example:
	//   ParameterType.STRING_LIST
	//
	// Experimental.
	ParameterType awsssm.ParameterType `field:"optional" json:"parameterType" yaml:"parameterType"`
	// One or more policies to apply to a SSM Parameter.
	// See: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html
	//
	// Experimental.
	Policies *string `field:"optional" json:"policies" yaml:"policies"`
	// Tags to add into the SSM Paramater that you want to add.
	//
	// Example:
	//   [
	//     {
	//       Key: 'STRING_VALUE',
	//       Value: 'STRING_VALUE'
	//     },
	//   ]
	//
	// Experimental.
	Tags *[]*TagProp `field:"optional" json:"tags" yaml:"tags"`
}

