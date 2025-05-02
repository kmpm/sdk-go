package nats

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/nats-io/nats.go"
)

type Sender struct {
	cnn            *cnn
	hasConnControl bool
	sendSubject    string
}

type SenderOption func(*Sender) error

func WithConn(conn *nats.Conn) SenderOption {
	return func(s *Sender) error {
		if conn == nil {
			return errors.New("nats connection is required")
		}
		s.cnn.SetConn(conn)
		s.hasConnControl = false
		return nil
	}
}

func WithURL(url string) SenderOption {
	return func(s *Sender) error {
		if url == "" {
			return errors.New("nats url is required")
		}
		s.cnn.Url = url
		s.hasConnControl = true
		return nil
	}
}

func NewSender(ctx context.Context, opts ...SenderOption) (s *Sender, err error) {
	s = &Sender{
		cnn: &cnn{},
	}
	defer func() {
		if err != nil {
			s.Close(ctx)
		}
	}()

	for _, fn := range opts {
		if err := fn(s); err != nil {
			return nil, err
		}
	}
	_, err = s.cnn.GetConn()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Sender) getSendSubject(ctx context.Context) string {
	subject, ok := ctx.Value(ctxKeySubject).(string)
	if ok {
		return subject
	}
	return s.sendSubject
}

func (s *Sender) Send(ctx context.Context, in binding.Message, transformers ...binding.Transformer) (err error) {
	subject := s.getSendSubject(ctx)
	if subject == "" {
		return newValidationError(fieldSendSubject, messageNoSendSubject)
	}
	defer func() {
		if err2 := in.Finish(err); err2 != nil {
			if err == nil {
				err = err2
			} else {
				err = fmt.Errorf("failed to call in.Finish() when error already occurred: %s: %w", err2.Error(), err)
			}
		}
	}()

	writer := new(bytes.Buffer)
	header, err := WriteMsg(ctx, in, writer, transformers...)
	if err != nil {
		return err
	}

	return s.cnn.Send(ctx, subject, writer.Bytes(), header)
}

func (s *Sender) Close(ctx context.Context) error {
	if s.hasConnControl {
		s.cnn.Close()
		return nil
	}
	return nil
}

var _ protocol.Sender = (*Sender)(nil)
var _ protocol.Closer = (*Sender)(nil)
