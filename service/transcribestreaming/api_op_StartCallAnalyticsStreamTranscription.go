// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribestreaming

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
	"time"
)

// Starts a bidirectional HTTP/2 or WebSocket stream where audio is streamed to
// Amazon Transcribe and the transcription results are streamed to your
// application. Use this operation for Call Analytics (https://docs.aws.amazon.com/transcribe/latest/dg/call-analytics.html)
// transcriptions. The following parameters are required:
//   - language-code
//   - media-encoding
//   - sample-rate
//
// For more information on streaming with Amazon Transcribe, see Transcribing
// streaming audio (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html)
// .
func (c *Client) StartCallAnalyticsStreamTranscription(ctx context.Context, params *StartCallAnalyticsStreamTranscriptionInput, optFns ...func(*Options)) (*StartCallAnalyticsStreamTranscriptionOutput, error) {
	if params == nil {
		params = &StartCallAnalyticsStreamTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartCallAnalyticsStreamTranscription", params, optFns, c.addOperationStartCallAnalyticsStreamTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartCallAnalyticsStreamTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCallAnalyticsStreamTranscriptionInput struct {

	// Specify the language code that represents the language spoken in your audio. If
	// you're unsure of the language spoken in your audio, consider using
	// IdentifyLanguage to enable automatic language identification. For a list of
	// languages supported with streaming Call Analytics, refer to the Supported
	// languages (https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html)
	// table.
	//
	// This member is required.
	LanguageCode types.CallAnalyticsLanguageCode

	// Specify the encoding of your input audio. Supported formats are:
	//   - FLAC
	//   - OPUS-encoded audio in an Ogg container
	//   - PCM (only signed 16-bit little-endian audio formats, which does not include
	//   WAV)
	// For more information, see Media formats (https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio)
	// .
	//
	// This member is required.
	MediaEncoding types.MediaEncoding

	// The sample rate of the input audio (in hertz). Low-quality audio, such as
	// telephone audio, is typically around 8,000 Hz. High-quality audio typically
	// ranges from 16,000 Hz to 48,000 Hz. Note that the sample rate you specify must
	// match that of your audio.
	//
	// This member is required.
	MediaSampleRateHertz *int32

	// Labels all personally identifiable information (PII) identified in your
	// transcript. Content identification is performed at the segment level; PII
	// specified in PiiEntityTypes is flagged upon complete transcription of an audio
	// segment. You can’t set ContentIdentificationType and ContentRedactionType in
	// the same request. If you set both, your request returns a BadRequestException .
	// For more information, see Redacting or identifying personally identifiable
	// information (https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html)
	// .
	ContentIdentificationType types.ContentIdentificationType

	// Redacts all personally identifiable information (PII) identified in your
	// transcript. Content redaction is performed at the segment level; PII specified
	// in PiiEntityTypes is redacted upon complete transcription of an audio segment.
	// You can’t set ContentRedactionType and ContentIdentificationType in the same
	// request. If you set both, your request returns a BadRequestException . For more
	// information, see Redacting or identifying personally identifiable information (https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html)
	// .
	ContentRedactionType types.ContentRedactionType

	// Enables partial result stabilization for your transcription. Partial result
	// stabilization can reduce latency in your output, but may impact accuracy. For
	// more information, see Partial-result stabilization (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization)
	// .
	EnablePartialResultsStabilization bool

	// Specify the name of the custom language model that you want to use when
	// processing your transcription. Note that language model names are case
	// sensitive. The language of the specified language model must match the language
	// code you specify in your transcription request. If the languages don't match,
	// the custom language model isn't applied. There are no errors or warnings
	// associated with a language mismatch. For more information, see Custom language
	// models (https://docs.aws.amazon.com/transcribe/latest/dg/custom-language-models.html)
	// .
	LanguageModelName *string

	// Specify the level of stability to use when you enable partial results
	// stabilization ( EnablePartialResultsStabilization ). Low stability provides the
	// highest accuracy. High stability transcribes faster, but with slightly lower
	// accuracy. For more information, see Partial-result stabilization (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization)
	// .
	PartialResultsStability types.PartialResultsStability

	// Specify which types of personally identifiable information (PII) you want to
	// redact in your transcript. You can include as many types as you'd like, or you
	// can select ALL . To include PiiEntityTypes in your Call Analytics request, you
	// must also include either ContentIdentificationType or ContentRedactionType .
	// Values must be comma-separated and can include: BANK_ACCOUNT_NUMBER ,
	// BANK_ROUTING , CREDIT_DEBIT_NUMBER , CREDIT_DEBIT_CVV , CREDIT_DEBIT_EXPIRY ,
	// PIN , EMAIL , ADDRESS , NAME , PHONE , SSN , or ALL .
	PiiEntityTypes *string

	// Specify a name for your Call Analytics transcription session. If you don't
	// include this parameter in your request, Amazon Transcribe generates an ID and
	// returns it in the response. You can use a session ID to retry a streaming
	// session.
	SessionId *string

	// Specify how you want your vocabulary filter applied to your transcript. To
	// replace words with *** , choose mask . To delete words, choose remove . To flag
	// words without changing them, choose tag .
	VocabularyFilterMethod types.VocabularyFilterMethod

	// Specify the name of the custom vocabulary filter that you want to use when
	// processing your transcription. Note that vocabulary filter names are case
	// sensitive. If the language of the specified custom vocabulary filter doesn't
	// match the language identified in your media, the vocabulary filter is not
	// applied to your transcription. For more information, see Using vocabulary
	// filtering with unwanted words (https://docs.aws.amazon.com/transcribe/latest/dg/vocabulary-filtering.html)
	// .
	VocabularyFilterName *string

	// Specify the name of the custom vocabulary that you want to use when processing
	// your transcription. Note that vocabulary names are case sensitive. If the
	// language of the specified custom vocabulary doesn't match the language
	// identified in your media, the custom vocabulary is not applied to your
	// transcription. For more information, see Custom vocabularies (https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html)
	// .
	VocabularyName *string

	noSmithyDocumentSerde
}

type StartCallAnalyticsStreamTranscriptionOutput struct {

	// Shows whether content identification was enabled for your Call Analytics
	// transcription.
	ContentIdentificationType types.ContentIdentificationType

	// Shows whether content redaction was enabled for your Call Analytics
	// transcription.
	ContentRedactionType types.ContentRedactionType

	// Shows whether partial results stabilization was enabled for your Call Analytics
	// transcription.
	EnablePartialResultsStabilization bool

	// Provides the language code that you specified in your Call Analytics request.
	LanguageCode types.CallAnalyticsLanguageCode

	// Provides the name of the custom language model that you specified in your Call
	// Analytics request.
	LanguageModelName *string

	// Provides the media encoding you specified in your Call Analytics request.
	MediaEncoding types.MediaEncoding

	// Provides the sample rate that you specified in your Call Analytics request.
	MediaSampleRateHertz *int32

	// Provides the stabilization level used for your transcription.
	PartialResultsStability types.PartialResultsStability

	// Lists the PII entity types you specified in your Call Analytics request.
	PiiEntityTypes *string

	// Provides the identifier for your Call Analytics streaming request.
	RequestId *string

	// Provides the identifier for your Call Analytics transcription session.
	SessionId *string

	// Provides the vocabulary filtering method used in your Call Analytics
	// transcription.
	VocabularyFilterMethod types.VocabularyFilterMethod

	// Provides the name of the custom vocabulary filter that you specified in your
	// Call Analytics request.
	VocabularyFilterName *string

	// Provides the name of the custom vocabulary that you specified in your Call
	// Analytics request.
	VocabularyName *string

	eventStream *StartCallAnalyticsStreamTranscriptionEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *StartCallAnalyticsStreamTranscriptionOutput) GetStream() *StartCallAnalyticsStreamTranscriptionEventStream {
	return o.eventStream
}

func (c *Client) addOperationStartCallAnalyticsStreamTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartCallAnalyticsStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartCallAnalyticsStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addEventStreamStartCallAnalyticsStreamTranscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddRequireMinimumProtocol(stack, 2, 0); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddStreamingEventsPayload(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = eventstreamapi.AddInitializeStreamWriter(stack); err != nil {
		return err
	}
	if err = addOpStartCallAnalyticsStreamTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartCallAnalyticsStreamTranscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartCallAnalyticsStreamTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartCallAnalyticsStreamTranscription",
	}
}

// StartCallAnalyticsStreamTranscriptionEventStream provides the event stream handling for the StartCallAnalyticsStreamTranscription operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewStartCallAnalyticsStreamTranscriptionEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type StartCallAnalyticsStreamTranscriptionEventStream struct {
	// AudioStreamWriter is the EventStream writer for the AudioStream events. This
	// value is automatically set by the SDK when the API call is made Use this member
	// when unit testing your code with the SDK to mock out the EventStream Writer.
	//
	// Must not be nil.
	Writer AudioStreamWriter

	// CallAnalyticsTranscriptResultStreamReader is the EventStream reader for the
	// CallAnalyticsTranscriptResultStream events. This value is automatically set by
	// the SDK when the API call is made Use this member when unit testing your code
	// with the SDK to mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader CallAnalyticsTranscriptResultStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewStartCallAnalyticsStreamTranscriptionEventStream initializes an StartCallAnalyticsStreamTranscriptionEventStream.
// This function should only be used for testing and mocking the StartCallAnalyticsStreamTranscriptionEventStream
// stream within your application.
//
// The Writer member must be set before writing events to the stream.
//
// The Reader member must be set before reading events from the stream.
func NewStartCallAnalyticsStreamTranscriptionEventStream(optFns ...func(*StartCallAnalyticsStreamTranscriptionEventStream)) *StartCallAnalyticsStreamTranscriptionEventStream {
	es := &StartCallAnalyticsStreamTranscriptionEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Send writes the event to the stream blocking until the event is written.
// Returns an error if the event was not written.
func (es *StartCallAnalyticsStreamTranscriptionEventStream) Send(ctx context.Context, event types.AudioStream) error {
	return es.Writer.Send(ctx, event)
}

// Events returns a channel to read events from.
func (es *StartCallAnalyticsStreamTranscriptionEventStream) Events() <-chan types.CallAnalyticsTranscriptResultStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *StartCallAnalyticsStreamTranscriptionEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *StartCallAnalyticsStreamTranscriptionEventStream) safeClose() {
	close(es.done)

	t := time.NewTicker(time.Second)
	defer t.Stop()
	writeCloseDone := make(chan error)
	go func() {
		if err := es.Writer.Close(); err != nil {
			es.err.SetError(err)
		}
		close(writeCloseDone)
	}()
	select {
	case <-t.C:
	case <-writeCloseDone:
	}

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *StartCallAnalyticsStreamTranscriptionEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Writer.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *StartCallAnalyticsStreamTranscriptionEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var inputErrCh <-chan struct{}
	if v, ok := es.Writer.(errorSet); ok {
		inputErrCh = v.ErrorSet()
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-inputErrCh:
		es.err.SetError(es.Writer.Err())
		es.Close()

	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
