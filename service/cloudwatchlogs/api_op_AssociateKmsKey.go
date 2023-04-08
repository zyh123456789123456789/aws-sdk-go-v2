// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified KMS key with the specified log group. Associating a
// KMS key with a log group overrides any existing associations between the log
// group and a KMS key. After a KMS key is associated with a log group, all newly
// ingested data for the log group is encrypted using the KMS key. This association
// is stored as long as the data encrypted with the KMS keyis still within
// CloudWatch Logs. This enables CloudWatch Logs to decrypt this data whenever it
// is requested. CloudWatch Logs supports only symmetric KMS keys. Do not use an
// associate an asymmetric KMS key with your log group. For more information, see
// Using Symmetric and Asymmetric Keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// . It can take up to 5 minutes for this operation to take effect. If you attempt
// to associate a KMS key with a log group but the KMS key does not exist or the
// KMS key is disabled, you receive an InvalidParameterException error.
func (c *Client) AssociateKmsKey(ctx context.Context, params *AssociateKmsKeyInput, optFns ...func(*Options)) (*AssociateKmsKeyOutput, error) {
	if params == nil {
		params = &AssociateKmsKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateKmsKey", params, optFns, c.addOperationAssociateKmsKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateKmsKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateKmsKeyInput struct {

	// The Amazon Resource Name (ARN) of the KMS key to use when encrypting log data.
	// This must be a symmetric KMS key. For more information, see Amazon Resource
	// Names (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms)
	// and Using Symmetric and Asymmetric Keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
	// .
	//
	// This member is required.
	KmsKeyId *string

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	noSmithyDocumentSerde
}

type AssociateKmsKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateKmsKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateKmsKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateKmsKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateKmsKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateKmsKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateKmsKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "AssociateKmsKey",
	}
}
