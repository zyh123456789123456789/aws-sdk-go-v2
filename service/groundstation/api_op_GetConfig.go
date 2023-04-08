// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns Config information. Only one Config response can be returned.
func (c *Client) GetConfig(ctx context.Context, params *GetConfigInput, optFns ...func(*Options)) (*GetConfigOutput, error) {
	if params == nil {
		params = &GetConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfig", params, optFns, c.addOperationGetConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigInput struct {

	// UUID of a Config .
	//
	// This member is required.
	ConfigId *string

	// Type of a Config .
	//
	// This member is required.
	ConfigType types.ConfigCapabilityType

	noSmithyDocumentSerde
}

type GetConfigOutput struct {

	// ARN of a Config
	//
	// This member is required.
	ConfigArn *string

	// Data elements in a Config .
	//
	// This member is required.
	ConfigData types.ConfigTypeData

	// UUID of a Config .
	//
	// This member is required.
	ConfigId *string

	// Name of a Config .
	//
	// This member is required.
	Name *string

	// Type of a Config .
	ConfigType types.ConfigCapabilityType

	// Tags assigned to a Config .
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfig{}, middleware.After)
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
	if err = addOpGetConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "groundstation",
		OperationName: "GetConfig",
	}
}
