// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used by activity workers and task states using the callback (https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token)
// pattern to report that the task identified by the taskToken completed
// successfully.
func (c *Client) SendTaskSuccess(ctx context.Context, params *SendTaskSuccessInput, optFns ...func(*Options)) (*SendTaskSuccessOutput, error) {
	if params == nil {
		params = &SendTaskSuccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendTaskSuccess", params, optFns, c.addOperationSendTaskSuccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendTaskSuccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendTaskSuccessInput struct {

	// The JSON output of the task. Length constraints apply to the payload size, and
	// are expressed as bytes in UTF-8 encoding.
	//
	// This member is required.
	Output *string

	// The token that represents this task. Task tokens are generated by Step
	// Functions when tasks are assigned to a worker, or in the context object (https://docs.aws.amazon.com/step-functions/latest/dg/input-output-contextobject.html)
	// when a workflow enters a task state. See GetActivityTaskOutput$taskToken .
	//
	// This member is required.
	TaskToken *string

	noSmithyDocumentSerde
}

type SendTaskSuccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendTaskSuccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendTaskSuccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendTaskSuccess{}, middleware.After)
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
	if err = addOpSendTaskSuccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendTaskSuccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendTaskSuccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "SendTaskSuccess",
	}
}
