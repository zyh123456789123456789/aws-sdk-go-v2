// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets an Amazon Web Services Migration Hub Refactor Spaces route.
func (c *Client) GetRoute(ctx context.Context, params *GetRouteInput, optFns ...func(*Options)) (*GetRouteOutput, error) {
	if params == nil {
		params = &GetRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRoute", params, optFns, c.addOperationGetRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRouteInput struct {

	// The ID of the application.
	//
	// This member is required.
	ApplicationIdentifier *string

	// The ID of the environment.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The ID of the route.
	//
	// This member is required.
	RouteIdentifier *string

	noSmithyDocumentSerde
}

type GetRouteOutput struct {

	// The ID of the application that the route belongs to.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the route.
	Arn *string

	// The Amazon Web Services account ID of the route creator.
	CreatedByAccountId *string

	// The timestamp of when the route is created.
	CreatedTime *time.Time

	// Unique identifier of the environment.
	EnvironmentId *string

	// Any error associated with the route resource.
	Error *types.ErrorResponse

	// Indicates whether to match all subpaths of the given source path. If this value
	// is false , requests must match the source path exactly before they are forwarded
	// to this route's service.
	IncludeChildPaths *bool

	// A timestamp that indicates when the route was last updated.
	LastUpdatedTime *time.Time

	// A list of HTTP methods to match. An empty list matches all values. If a method
	// is present, only HTTP requests using that method are forwarded to this route’s
	// service.
	Methods []types.HttpMethod

	// The Amazon Web Services account ID of the route owner.
	OwnerAccountId *string

	// A mapping of Amazon API Gateway path resources to resource IDs.
	PathResourceToId map[string]string

	// The unique identifier of the route. DEFAULT: All traffic that does not match
	// another route is forwarded to the default route. Applications must have a
	// default route before any other routes can be created. URI_PATH: A route that is
	// based on a URI path.
	RouteId *string

	// The type of route.
	RouteType types.RouteType

	// The unique identifier of the service.
	ServiceId *string

	// The path to use to match traffic. Paths must start with / and are relative to
	// the base of the application.
	SourcePath *string

	// The current state of the route.
	State types.RouteState

	// The tags assigned to the route. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key-value pair.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRoute{}, middleware.After)
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
	if err = addOpGetRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "refactor-spaces",
		OperationName: "GetRoute",
	}
}
