// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of routing controls for a control panel. A routing control is
// an Amazon Route 53 Application Recovery Controller construct that has one of two
// states: ON and OFF. You can map the routing control state to the state of an
// Amazon Route 53 health check, which can be used to control routing.
func (c *Client) ListRoutingControls(ctx context.Context, params *ListRoutingControlsInput, optFns ...func(*Options)) (*ListRoutingControlsOutput, error) {
	if params == nil {
		params = &ListRoutingControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoutingControls", params, optFns, c.addOperationListRoutingControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoutingControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoutingControlsInput struct {

	// The Amazon Resource Name (ARN) of the control panel.
	//
	// This member is required.
	ControlPanelArn *string

	// The number of objects that you want to return with this call.
	MaxResults int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRoutingControlsOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// An array of routing controls.
	RoutingControls []types.RoutingControl

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoutingControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRoutingControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRoutingControls{}, middleware.After)
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
	if err = addOpListRoutingControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoutingControls(options.Region), middleware.Before); err != nil {
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

// ListRoutingControlsAPIClient is a client that implements the
// ListRoutingControls operation.
type ListRoutingControlsAPIClient interface {
	ListRoutingControls(context.Context, *ListRoutingControlsInput, ...func(*Options)) (*ListRoutingControlsOutput, error)
}

var _ ListRoutingControlsAPIClient = (*Client)(nil)

// ListRoutingControlsPaginatorOptions is the paginator options for
// ListRoutingControls
type ListRoutingControlsPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoutingControlsPaginator is a paginator for ListRoutingControls
type ListRoutingControlsPaginator struct {
	options   ListRoutingControlsPaginatorOptions
	client    ListRoutingControlsAPIClient
	params    *ListRoutingControlsInput
	nextToken *string
	firstPage bool
}

// NewListRoutingControlsPaginator returns a new ListRoutingControlsPaginator
func NewListRoutingControlsPaginator(client ListRoutingControlsAPIClient, params *ListRoutingControlsInput, optFns ...func(*ListRoutingControlsPaginatorOptions)) *ListRoutingControlsPaginator {
	if params == nil {
		params = &ListRoutingControlsInput{}
	}

	options := ListRoutingControlsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoutingControlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoutingControlsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoutingControls page.
func (p *ListRoutingControlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoutingControlsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRoutingControls(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRoutingControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-control-config",
		OperationName: "ListRoutingControls",
	}
}
