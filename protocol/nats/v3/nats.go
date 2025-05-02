package nats

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type cnn struct {
	conn              *nats.Conn
	Url               string
	SendWithJetStream bool
	js                jetstream.JetStream
}

func (c *cnn) SetConn(conn *nats.Conn) {
	c.conn = conn
}

func (c *cnn) GetConn() (*nats.Conn, error) {
	if c.conn != nil {
		return c.conn, nil
	}
	if c.Url != "" {
		conn, err := nats.Connect(c.Url)
		if err != nil {
			return nil, err
		}
		c.conn = conn
		return conn, nil
	}
	if c.conn == nil {
		return nil, errors.New("no existing connection and no URL provided")
	}
	return nil, nil
}

func (c *cnn) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *cnn) JetStream() (jetstream.JetStream, error) {
	if c.js != nil {
		return c.js, nil
	}
	if c.conn == nil {
		return nil, errors.New("no connection")
	}
	js, err := jetstream.New(c.conn)
	if err != nil {
		return nil, err
	}
	return js, nil
}

func (c *cnn) sendJS(ctx context.Context, m *nats.Msg) error {
	js, err := c.JetStream()
	if err != nil {
		return err
	}
	_, err = js.StreamNameBySubject(ctx, m.Subject)
	if err != nil {
		return err
	}
	_, err = js.PublishMsg(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (c *cnn) Send(ctx context.Context, subject string, data []byte, header nats.Header) error {
	if c.conn == nil {
		return errors.New("no connection")
	}
	m := &nats.Msg{
		Subject: subject,
		Data:    data,
		Header:  header,
	}
	if c.SendWithJetStream {
		return c.sendJS(ctx, m)
	}
	return c.conn.PublishMsg(m)
}
