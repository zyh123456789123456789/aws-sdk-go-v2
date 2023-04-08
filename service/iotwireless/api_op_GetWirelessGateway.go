// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a wireless gateway.
func (c *Client) GetWirelessGateway(ctx context.Context, params *GetWirelessGatewayInput, optFns ...func(*Options)) (*GetWirelessGatewayOutput, error) {
	if params == nil {
		params = &GetWirelessGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWirelessGateway", params, optFns, c.addOperationGetWirelessGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWirelessGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWirelessGatewayInput struct {

	// The identifier of the wireless gateway to get.
	//
	// This member is required.
	Identifier *string

	// The type of identifier used in identifier .
	//
	// This member is required.
	IdentifierType types.WirelessGatewayIdType

	noSmithyDocumentSerde
}

type GetWirelessGatewayOutput struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The description of the resource.
	Description *string

	// The ID of the wireless gateway.
	Id *string

	// Information about the wireless gateway.
	LoRaWAN *types.LoRaWANGateway

	// The name of the resource.
	Name *string

	// The ARN of the thing associated with the wireless gateway.
	ThingArn *string

	// The name of the thing associated with the wireless gateway. The value is empty
	// if a thing isn't associated with the gateway.
	ThingName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWirelessGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWirelessGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWirelessGateway{}, middleware.After)
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
	if err = addOpGetWirelessGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWirelessGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWirelessGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetWirelessGateway",
	}
}
