package slacker

import (
	"context"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// ClientOption an option for client values
type ClientOption func(*ClientDefaults)
type InteractiveActionHandlerFunc func(ctx context.Context, api *slack.Client, client *socketmode.Client, callback slack.InteractionCallback)

// WithDebug sets debug toggle
func WithDebug(debug bool) ClientOption {
	return func(defaults *ClientDefaults) {
		defaults.Debug = debug
	}
}

func WithInteractiveActionHandler(f InteractiveActionHandlerFunc) ClientOption {
	return func(defaults *ClientDefaults) {
		defaults.InteractiveActionHandler = f
	}
}

// ClientDefaults configuration
type ClientDefaults struct {
	Debug bool
	InteractiveActionHandler InteractiveActionHandlerFunc
}

func newClientDefaults(options ...ClientOption) *ClientDefaults {
	config := &ClientDefaults{
		Debug: false,
		InteractiveActionHandler: func(ctx context.Context, api *slack.Client, client *socketmode.Client, callback slack.InteractionCallback) {},
	}

	for _, option := range options {
		option(config)
	}
	return config
}

// ReplyOption an option for reply values
type ReplyOption func(*ReplyDefaults)

// WithAttachments sets message attachments
func WithAttachments(attachments []slack.Attachment) ReplyOption {
	return func(defaults *ReplyDefaults) {
		defaults.Attachments = attachments
	}
}

// WithBlocks sets message blocks
func WithBlocks(blocks []slack.Block) ReplyOption {
	return func(defaults *ReplyDefaults) {
		defaults.Blocks = blocks
	}
}

// WithThreadReply specifies the reply to be inside a thread of the original message
func WithThreadReply(useThread bool) ReplyOption {
	return func(defaults *ReplyDefaults) {
		defaults.ThreadResponse = useThread
	}
}

// ReplyDefaults configuration
type ReplyDefaults struct {
	Attachments    []slack.Attachment
	Blocks         []slack.Block
	ThreadResponse bool
}

// NewReplyDefaults builds our ReplyDefaults from zero or more ReplyOption.
func NewReplyDefaults(options ...ReplyOption) *ReplyDefaults {
	config := &ReplyDefaults{
		Attachments:    []slack.Attachment{},
		Blocks:         []slack.Block{},
		ThreadResponse: false,
	}

	for _, option := range options {
		option(config)
	}
	return config
}

// ReportErrorOption an option for report error values
type ReportErrorOption func(*ReportErrorDefaults)

// ReportErrorDefaults configuration
type ReportErrorDefaults struct {
	ThreadResponse bool
}

// WithThreadError specifies the reply to be inside a thread of the original message
func WithThreadError(useThread bool) ReportErrorOption {
	return func(defaults *ReportErrorDefaults) {
		defaults.ThreadResponse = useThread
	}
}

// NewReportErrorDefaults builds our ReportErrorDefaults from zero or more
// ReportErrorOption.
func NewReportErrorDefaults(options ...ReportErrorOption) *ReportErrorDefaults {
	config := &ReportErrorDefaults{
		ThreadResponse: false,
	}

	for _, option := range options {
		option(config)
	}
	return config
}
