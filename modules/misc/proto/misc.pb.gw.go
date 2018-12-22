// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: modules/misc/proto/misc.proto

/*
Package miscpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package miscpb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_MiscSystem_Version_0(ctx context.Context, marshaler runtime.Marshaler, client MiscSystemClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VersionRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Version(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterMiscSystemHandlerFromEndpoint is same as RegisterMiscSystemHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMiscSystemHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMiscSystemHandler(ctx, mux, conn)
}

// RegisterMiscSystemHandler registers the http handlers for service MiscSystem to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMiscSystemHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMiscSystemHandlerClient(ctx, mux, NewMiscSystemClient(conn))
}

// RegisterMiscSystemHandlerClient registers the http handlers for service MiscSystem
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MiscSystemClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MiscSystemClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MiscSystemClient" to call the correct interceptors.
func RegisterMiscSystemHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MiscSystemClient) error {

	mux.Handle("GET", pattern_MiscSystem_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MiscSystem_Version_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MiscSystem_Version_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MiscSystem_Version_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "misc", "version"}, ""))
)

var (
	forward_MiscSystem_Version_0 = runtime.ForwardResponseMessage
)