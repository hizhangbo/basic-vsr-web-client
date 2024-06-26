// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.6
// source: basicvsr/v1/basicvsr.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBasicVSRExecBasicVsr = "/basicvsr.v1.BasicVSR/ExecBasicVsr"
const OperationBasicVSRGetStatus = "/basicvsr.v1.BasicVSR/GetStatus"

type BasicVSRHTTPServer interface {
	ExecBasicVsr(context.Context, *GPURequest) (*ExecReply, error)
	// GetStatus Sends a greeting
	GetStatus(context.Context, *GPURequest) (*GPUReply, error)
}

func RegisterBasicVSRHTTPServer(s *http.Server, srv BasicVSRHTTPServer) {
	r := s.Route("/")
	r.GET("/basicvsr", _BasicVSR_GetStatus0_HTTP_Handler(srv))
	r.GET("/basicvsr/exec/{name}", _BasicVSR_ExecBasicVsr0_HTTP_Handler(srv))
}

func _BasicVSR_GetStatus0_HTTP_Handler(srv BasicVSRHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GPURequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBasicVSRGetStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStatus(ctx, req.(*GPURequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GPUReply)
		return ctx.Result(200, reply)
	}
}

func _BasicVSR_ExecBasicVsr0_HTTP_Handler(srv BasicVSRHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GPURequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBasicVSRExecBasicVsr)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExecBasicVsr(ctx, req.(*GPURequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExecReply)
		return ctx.Result(200, reply)
	}
}

type BasicVSRHTTPClient interface {
	ExecBasicVsr(ctx context.Context, req *GPURequest, opts ...http.CallOption) (rsp *ExecReply, err error)
	GetStatus(ctx context.Context, req *GPURequest, opts ...http.CallOption) (rsp *GPUReply, err error)
}

type BasicVSRHTTPClientImpl struct {
	cc *http.Client
}

func NewBasicVSRHTTPClient(client *http.Client) BasicVSRHTTPClient {
	return &BasicVSRHTTPClientImpl{client}
}

func (c *BasicVSRHTTPClientImpl) ExecBasicVsr(ctx context.Context, in *GPURequest, opts ...http.CallOption) (*ExecReply, error) {
	var out ExecReply
	pattern := "/basicvsr/exec/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBasicVSRExecBasicVsr))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BasicVSRHTTPClientImpl) GetStatus(ctx context.Context, in *GPURequest, opts ...http.CallOption) (*GPUReply, error) {
	var out GPUReply
	pattern := "/basicvsr"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBasicVSRGetStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
