// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the configured alarms. Specify an alarm name in your
// request to return information about a specific alarm, or specify a monitored
// resource name to return information about all alarms for a specific resource. An
// alarm is used to monitor a single metric for one of your resources. When a
// metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see Alarms in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-alarms)
// .
func (c *Client) GetAlarms(ctx context.Context, params *GetAlarmsInput, optFns ...func(*Options)) (*GetAlarmsOutput, error) {
	if params == nil {
		params = &GetAlarmsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAlarms", params, optFns, c.addOperationGetAlarmsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAlarmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAlarmsInput struct {

	// The name of the alarm. Specify an alarm name to return information about a
	// specific alarm.
	AlarmName *string

	// The name of the Lightsail resource being monitored by the alarm. Specify a
	// monitored resource name to return information about all alarms for a specific
	// resource.
	MonitoredResourceName *string

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetAlarms request. If your results are
	// paginated, the response will return a next page token that you can specify as
	// the page token in a subsequent request.
	PageToken *string

	noSmithyDocumentSerde
}

type GetAlarmsOutput struct {

	// An array of objects that describe the alarms.
	Alarms []types.Alarm

	// The token to advance to the next page of results from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetAlarms request and specify the next page
	// token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAlarmsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAlarms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAlarms{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAlarms(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAlarms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetAlarms",
	}
}
