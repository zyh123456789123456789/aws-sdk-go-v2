// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of existing routes in a service mesh.
func (c *Client) ListRoutes(ctx context.Context, params *ListRoutesInput, optFns ...func(*Options)) (*ListRoutesOutput, error) {
	if params == nil {
		params = &ListRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoutes", params, optFns, c.addOperationListRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoutesInput struct {

	// The name of the service mesh to list routes in.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual router to list routes in.
	//
	// This member is required.
	VirtualRouterName *string

	// The maximum number of results returned by ListRoutes in paginated output. When
	// you use this parameter, ListRoutes returns only limit results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListRoutes request with the returned
	// nextToken value. This value can be between 1 and 100. If you don't use this
	// parameter, ListRoutes returns up to 100 results and a nextToken value if
	// applicable.
	Limit *int32

	// The Amazon Web Services IAM account ID of the service mesh owner. If the
	// account ID is not your own, then it's the ID of the account that shared the mesh
	// with your account. For more information about mesh sharing, see Working with
	// shared meshes (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html)
	// .
	MeshOwner *string

	// The nextToken value returned from a previous paginated ListRoutes request where
	// limit was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRoutesOutput struct {

	// The list of existing routes for the specified service mesh and virtual router.
	//
	// This member is required.
	Routes []types.RouteRef

	// The nextToken value to include in a future ListRoutes request. When the results
	// of a ListRoutes request exceed limit , you can use this value to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRoutes{}, middleware.After)
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
	if err = addOpListRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoutes(options.Region), middleware.Before); err != nil {
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

// ListRoutesAPIClient is a client that implements the ListRoutes operation.
type ListRoutesAPIClient interface {
	ListRoutes(context.Context, *ListRoutesInput, ...func(*Options)) (*ListRoutesOutput, error)
}

var _ ListRoutesAPIClient = (*Client)(nil)

// ListRoutesPaginatorOptions is the paginator options for ListRoutes
type ListRoutesPaginatorOptions struct {
	// The maximum number of results returned by ListRoutes in paginated output. When
	// you use this parameter, ListRoutes returns only limit results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListRoutes request with the returned
	// nextToken value. This value can be between 1 and 100. If you don't use this
	// parameter, ListRoutes returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoutesPaginator is a paginator for ListRoutes
type ListRoutesPaginator struct {
	options   ListRoutesPaginatorOptions
	client    ListRoutesAPIClient
	params    *ListRoutesInput
	nextToken *string
	firstPage bool
}

// NewListRoutesPaginator returns a new ListRoutesPaginator
func NewListRoutesPaginator(client ListRoutesAPIClient, params *ListRoutesInput, optFns ...func(*ListRoutesPaginatorOptions)) *ListRoutesPaginator {
	if params == nil {
		params = &ListRoutesInput{}
	}

	options := ListRoutesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoutesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoutesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoutes page.
func (p *ListRoutesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoutesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListRoutes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "ListRoutes",
	}
}
