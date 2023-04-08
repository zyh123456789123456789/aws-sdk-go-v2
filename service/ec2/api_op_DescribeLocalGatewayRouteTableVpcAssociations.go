// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified associations between VPCs and local gateway route
// tables.
func (c *Client) DescribeLocalGatewayRouteTableVpcAssociations(ctx context.Context, params *DescribeLocalGatewayRouteTableVpcAssociationsInput, optFns ...func(*Options)) (*DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	if params == nil {
		params = &DescribeLocalGatewayRouteTableVpcAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocalGatewayRouteTableVpcAssociations", params, optFns, c.addOperationDescribeLocalGatewayRouteTableVpcAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocalGatewayRouteTableVpcAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewayRouteTableVpcAssociationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//   - local-gateway-id - The ID of a local gateway.
	//   - local-gateway-route-table-arn - The Amazon Resource Name (ARN) of the local
	//   gateway route table for the association.
	//   - local-gateway-route-table-id - The ID of the local gateway route table.
	//   - local-gateway-route-table-vpc-association-id - The ID of the association.
	//   - owner-id - The ID of the Amazon Web Services account that owns the local
	//   gateway route table for the association.
	//   - state - The state of the association.
	//   - vpc-id - The ID of the VPC.
	Filters []types.Filter

	// The IDs of the associations.
	LocalGatewayRouteTableVpcAssociationIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLocalGatewayRouteTableVpcAssociationsOutput struct {

	// Information about the associations.
	LocalGatewayRouteTableVpcAssociations []types.LocalGatewayRouteTableVpcAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocalGatewayRouteTableVpcAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGatewayRouteTableVpcAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGatewayRouteTableVpcAssociations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVpcAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeLocalGatewayRouteTableVpcAssociationsAPIClient is a client that
// implements the DescribeLocalGatewayRouteTableVpcAssociations operation.
type DescribeLocalGatewayRouteTableVpcAssociationsAPIClient interface {
	DescribeLocalGatewayRouteTableVpcAssociations(context.Context, *DescribeLocalGatewayRouteTableVpcAssociationsInput, ...func(*Options)) (*DescribeLocalGatewayRouteTableVpcAssociationsOutput, error)
}

var _ DescribeLocalGatewayRouteTableVpcAssociationsAPIClient = (*Client)(nil)

// DescribeLocalGatewayRouteTableVpcAssociationsPaginatorOptions is the paginator
// options for DescribeLocalGatewayRouteTableVpcAssociations
type DescribeLocalGatewayRouteTableVpcAssociationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLocalGatewayRouteTableVpcAssociationsPaginator is a paginator for
// DescribeLocalGatewayRouteTableVpcAssociations
type DescribeLocalGatewayRouteTableVpcAssociationsPaginator struct {
	options   DescribeLocalGatewayRouteTableVpcAssociationsPaginatorOptions
	client    DescribeLocalGatewayRouteTableVpcAssociationsAPIClient
	params    *DescribeLocalGatewayRouteTableVpcAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLocalGatewayRouteTableVpcAssociationsPaginator returns a new
// DescribeLocalGatewayRouteTableVpcAssociationsPaginator
func NewDescribeLocalGatewayRouteTableVpcAssociationsPaginator(client DescribeLocalGatewayRouteTableVpcAssociationsAPIClient, params *DescribeLocalGatewayRouteTableVpcAssociationsInput, optFns ...func(*DescribeLocalGatewayRouteTableVpcAssociationsPaginatorOptions)) *DescribeLocalGatewayRouteTableVpcAssociationsPaginator {
	if params == nil {
		params = &DescribeLocalGatewayRouteTableVpcAssociationsInput{}
	}

	options := DescribeLocalGatewayRouteTableVpcAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLocalGatewayRouteTableVpcAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLocalGatewayRouteTableVpcAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLocalGatewayRouteTableVpcAssociations page.
func (p *DescribeLocalGatewayRouteTableVpcAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
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

	result, err := p.client.DescribeLocalGatewayRouteTableVpcAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVpcAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLocalGatewayRouteTableVpcAssociations",
	}
}
