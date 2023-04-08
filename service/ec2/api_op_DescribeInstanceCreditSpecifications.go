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

// Describes the credit option for CPU usage of the specified burstable
// performance instances. The credit options are standard and unlimited . If you do
// not specify an instance ID, Amazon EC2 returns burstable performance instances
// with the unlimited credit option, as well as instances that were previously
// configured as T2, T3, and T3a with the unlimited credit option. For example, if
// you resize a T2 instance, while it is configured as unlimited , to an M4
// instance, Amazon EC2 returns the M4 instance. If you specify one or more
// instance IDs, Amazon EC2 returns the credit option ( standard or unlimited ) of
// those instances. If you specify an instance ID that is not valid, such as an
// instance that is not a burstable performance instance, an error is returned.
// Recently terminated instances might appear in the returned results. This
// interval is usually less than one hour. If an Availability Zone is experiencing
// a service disruption and you specify instance IDs in the affected zone, or do
// not specify any instance IDs at all, the call fails. If you specify only
// instance IDs in an unaffected zone, the call works normally. For more
// information, see Burstable performance instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
// in the Amazon EC2 User Guide.
func (c *Client) DescribeInstanceCreditSpecifications(ctx context.Context, params *DescribeInstanceCreditSpecificationsInput, optFns ...func(*Options)) (*DescribeInstanceCreditSpecificationsOutput, error) {
	if params == nil {
		params = &DescribeInstanceCreditSpecificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceCreditSpecifications", params, optFns, c.addOperationDescribeInstanceCreditSpecificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceCreditSpecificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceCreditSpecificationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - instance-id - The ID of the instance.
	Filters []types.Filter

	// The instance IDs. Default: Describes all your instances. Constraints: Maximum
	// 1000 explicitly specified instance IDs.
	InstanceIds []string

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// . You cannot specify this parameter and the instance IDs parameter in the same
	// call.
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstanceCreditSpecificationsOutput struct {

	// Information about the credit option for CPU usage of an instance.
	InstanceCreditSpecifications []types.InstanceCreditSpecification

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceCreditSpecificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceCreditSpecifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceCreditSpecifications{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceCreditSpecifications(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceCreditSpecificationsAPIClient is a client that implements the
// DescribeInstanceCreditSpecifications operation.
type DescribeInstanceCreditSpecificationsAPIClient interface {
	DescribeInstanceCreditSpecifications(context.Context, *DescribeInstanceCreditSpecificationsInput, ...func(*Options)) (*DescribeInstanceCreditSpecificationsOutput, error)
}

var _ DescribeInstanceCreditSpecificationsAPIClient = (*Client)(nil)

// DescribeInstanceCreditSpecificationsPaginatorOptions is the paginator options
// for DescribeInstanceCreditSpecifications
type DescribeInstanceCreditSpecificationsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// . You cannot specify this parameter and the instance IDs parameter in the same
	// call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceCreditSpecificationsPaginator is a paginator for
// DescribeInstanceCreditSpecifications
type DescribeInstanceCreditSpecificationsPaginator struct {
	options   DescribeInstanceCreditSpecificationsPaginatorOptions
	client    DescribeInstanceCreditSpecificationsAPIClient
	params    *DescribeInstanceCreditSpecificationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstanceCreditSpecificationsPaginator returns a new
// DescribeInstanceCreditSpecificationsPaginator
func NewDescribeInstanceCreditSpecificationsPaginator(client DescribeInstanceCreditSpecificationsAPIClient, params *DescribeInstanceCreditSpecificationsInput, optFns ...func(*DescribeInstanceCreditSpecificationsPaginatorOptions)) *DescribeInstanceCreditSpecificationsPaginator {
	if params == nil {
		params = &DescribeInstanceCreditSpecificationsInput{}
	}

	options := DescribeInstanceCreditSpecificationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstanceCreditSpecificationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceCreditSpecificationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstanceCreditSpecifications page.
func (p *DescribeInstanceCreditSpecificationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceCreditSpecificationsOutput, error) {
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

	result, err := p.client.DescribeInstanceCreditSpecifications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstanceCreditSpecifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceCreditSpecifications",
	}
}
