// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets an infrastructure configuration.
func (c *Client) GetInfrastructureConfiguration(ctx context.Context, params *GetInfrastructureConfigurationInput, optFns ...func(*Options)) (*GetInfrastructureConfigurationOutput, error) {
	if params == nil {
		params = &GetInfrastructureConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInfrastructureConfiguration", params, optFns, c.addOperationGetInfrastructureConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInfrastructureConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// GetInfrastructureConfiguration request object.
type GetInfrastructureConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the infrastructure configuration that you
	// want to retrieve.
	//
	// This member is required.
	InfrastructureConfigurationArn *string

	noSmithyDocumentSerde
}

// GetInfrastructureConfiguration response object.
type GetInfrastructureConfigurationOutput struct {

	// The infrastructure configuration object.
	InfrastructureConfiguration *types.InfrastructureConfiguration

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInfrastructureConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInfrastructureConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInfrastructureConfiguration{}, middleware.After)
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
	if err = addOpGetInfrastructureConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInfrastructureConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInfrastructureConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "GetInfrastructureConfiguration",
	}
}
