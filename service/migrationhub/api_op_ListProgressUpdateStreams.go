// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists progress update streams associated with the user account making this call.
func (c *Client) ListProgressUpdateStreams(ctx context.Context, params *ListProgressUpdateStreamsInput, optFns ...func(*Options)) (*ListProgressUpdateStreamsOutput, error) {
	if params == nil {
		params = &ListProgressUpdateStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProgressUpdateStreams", params, optFns, c.addOperationListProgressUpdateStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProgressUpdateStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProgressUpdateStreamsInput struct {

	// Filter to limit the maximum number of results to list per page.
	MaxResults *int32

	// If a NextToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in NextToken .
	NextToken *string

	noSmithyDocumentSerde
}

type ListProgressUpdateStreamsOutput struct {

	// If there are more streams created than the max result, return the next token to
	// be passed to the next call as a bookmark of where to start from.
	NextToken *string

	// List of progress update streams up to the max number of results passed in the
	// input.
	ProgressUpdateStreamSummaryList []types.ProgressUpdateStreamSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProgressUpdateStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProgressUpdateStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProgressUpdateStreams{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProgressUpdateStreams(options.Region), middleware.Before); err != nil {
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

// ListProgressUpdateStreamsAPIClient is a client that implements the
// ListProgressUpdateStreams operation.
type ListProgressUpdateStreamsAPIClient interface {
	ListProgressUpdateStreams(context.Context, *ListProgressUpdateStreamsInput, ...func(*Options)) (*ListProgressUpdateStreamsOutput, error)
}

var _ ListProgressUpdateStreamsAPIClient = (*Client)(nil)

// ListProgressUpdateStreamsPaginatorOptions is the paginator options for
// ListProgressUpdateStreams
type ListProgressUpdateStreamsPaginatorOptions struct {
	// Filter to limit the maximum number of results to list per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProgressUpdateStreamsPaginator is a paginator for ListProgressUpdateStreams
type ListProgressUpdateStreamsPaginator struct {
	options   ListProgressUpdateStreamsPaginatorOptions
	client    ListProgressUpdateStreamsAPIClient
	params    *ListProgressUpdateStreamsInput
	nextToken *string
	firstPage bool
}

// NewListProgressUpdateStreamsPaginator returns a new
// ListProgressUpdateStreamsPaginator
func NewListProgressUpdateStreamsPaginator(client ListProgressUpdateStreamsAPIClient, params *ListProgressUpdateStreamsInput, optFns ...func(*ListProgressUpdateStreamsPaginatorOptions)) *ListProgressUpdateStreamsPaginator {
	if params == nil {
		params = &ListProgressUpdateStreamsInput{}
	}

	options := ListProgressUpdateStreamsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProgressUpdateStreamsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProgressUpdateStreamsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProgressUpdateStreams page.
func (p *ListProgressUpdateStreamsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProgressUpdateStreamsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListProgressUpdateStreams(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProgressUpdateStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "ListProgressUpdateStreams",
	}
}
