// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accepts an invitation for the member account to contribute data to a behavior
// graph. This operation can only be called by an invited member account. The
// request provides the ARN of behavior graph. The member account status in the
// graph must be INVITED .
func (c *Client) AcceptInvitation(ctx context.Context, params *AcceptInvitationInput, optFns ...func(*Options)) (*AcceptInvitationOutput, error) {
	if params == nil {
		params = &AcceptInvitationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptInvitation", params, optFns, c.addOperationAcceptInvitationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptInvitationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptInvitationInput struct {

	// The ARN of the behavior graph that the member account is accepting the
	// invitation for. The member account status in the behavior graph must be INVITED .
	//
	// This member is required.
	GraphArn *string

	noSmithyDocumentSerde
}

type AcceptInvitationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptInvitationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAcceptInvitation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAcceptInvitation{}, middleware.After)
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
	if err = addOpAcceptInvitationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptInvitation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptInvitation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "AcceptInvitation",
	}
}
