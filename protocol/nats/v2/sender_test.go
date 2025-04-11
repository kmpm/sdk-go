package nats

import (
	"context"
	"testing"
)

func TestGetSendSubject(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		sendSubject string
		wantSubject string
	}{
		{
			name: "using default subject",
			args: args{
				ctx: context.Background(),
			},
			sendSubject: "default",
			wantSubject: "default",
		},
		{
			name: "using context subject",
			args: args{

				ctx: WithSubject(context.Background(), "context"),
			},
			sendSubject: "default",
			wantSubject: "context",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			sender := &Sender{Subject: tt.sendSubject}
			subject := sender.getSubject(tt.args.ctx)
			if subject != tt.wantSubject {
				t.Errorf("Sender.getSubject() = %v, want %v", subject, tt.wantSubject)
			}
		})
	}
}
