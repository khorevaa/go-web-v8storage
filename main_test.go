package main

import (
	"testing"

	"github.com/gocraft/web"
)

func TestContext_SayHello(t *testing.T) {
	type fields struct {
		HelloCount int
	}
	type args struct {
		rw  web.ResponseWriter
		req *web.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				HelloCount: tt.fields.HelloCount,
			}
			c.SayHello(tt.args.rw, tt.args.req)
		})
	}
}
