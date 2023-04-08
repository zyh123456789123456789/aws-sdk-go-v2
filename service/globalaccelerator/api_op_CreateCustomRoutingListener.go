// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a listener to process inbound connections from clients to a custom
// routing accelerator. Connections arrive to assigned static IP addresses on the
// port range that you specify.
func (c *Client) CreateCustomRoutingListener(ctx context.Context, params *CreateCustomRoutingListenerInput, optFns ...func(*Options)) (*CreateCustomRoutingListenerOutput, error) {
	if params == nil {
		params = &CreateCustomRoutingListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomRoutingListener", params, optFns, c.addOperationCreateCustomRoutingListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomRoutingListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomRoutingListenerInput struct {

	// The Amazon Resource Name (ARN) of the accelerator for a custom routing listener.
	//
	// This member is required.
	AcceleratorArn *string

	// A unique, case-sensitive identifier that you provide to ensure the
	// idempotency—that is, the uniqueness—of the request.
	//
	// This member is required.
	IdempotencyToken *string

	// The port range to support for connections from clients to your accelerator.
	// Separately, you set port ranges for endpoints. For more information, see About
	// endpoints for custom routing accelerators (https://docs.aws.amazon.com/global-accelerator/latest/dg/about-custom-routing-endpoints.html)
	// .
	//
	// This member is required.
	PortRanges []types.PortRange

	noSmithyDocumentSerde
}

type CreateCustomRoutingListenerOutput struct {

	// The listener that you've created for a custom routing accelerator.
	Listener *types.CustomRoutingListener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomRoutingListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCustomRoutingListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCustomRoutingListener{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateCustomRoutingListenerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCustomRoutingListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomRoutingListener(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomRoutingListener struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomRoutingListener) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomRoutingListener) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCustomRoutingListenerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomRoutingListenerInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCustomRoutingListenerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomRoutingListener{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomRoutingListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "CreateCustomRoutingListener",
	}
}
