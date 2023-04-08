// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specified task set within a service. This is used when a service uses
// the EXTERNAL deployment controller type. For more information, see Amazon ECS
// deployment types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) DeleteTaskSet(ctx context.Context, params *DeleteTaskSetInput, optFns ...func(*Options)) (*DeleteTaskSetOutput, error) {
	if params == nil {
		params = &DeleteTaskSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTaskSet", params, optFns, c.addOperationDeleteTaskSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTaskSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTaskSetInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task set found in to delete.
	//
	// This member is required.
	Cluster *string

	// The short name or full Amazon Resource Name (ARN) of the service that hosts the
	// task set to delete.
	//
	// This member is required.
	Service *string

	// The task set ID or full Amazon Resource Name (ARN) of the task set to delete.
	//
	// This member is required.
	TaskSet *string

	// If true , you can delete a task set even if it hasn't been scaled down to zero.
	Force *bool

	noSmithyDocumentSerde
}

type DeleteTaskSetOutput struct {

	// Details about the task set.
	TaskSet *types.TaskSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTaskSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTaskSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTaskSet{}, middleware.After)
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
	if err = addOpDeleteTaskSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTaskSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTaskSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeleteTaskSet",
	}
}
