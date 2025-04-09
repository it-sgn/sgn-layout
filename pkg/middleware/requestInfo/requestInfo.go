package requestInfo

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/v2/middleware"
)

const DomainKey = "x-md-global-domain"

// setRequestInfo Set Request information
func SetRequestInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			domain := ""
			// Meta information acquisition domain
			if md, ok := metadata.FromServerContext(ctx); ok {
				domain = md.Get(DomainKey)
			}
			// Request information to get domain
			if domain == "" {
				if tr, ok := transport.FromServerContext(ctx); ok {
					switch tpKind := tr.Kind(); tpKind {
					case transport.KindHTTP:
						trHttp := tr.(*http.Transport)
						if trHttp.RequestHeader().Get(DomainKey) != "" {
							domain = trHttp.RequestHeader().Get(DomainKey)
						}
					case transport.KindGRPC:
						trHttp := tr.(*grpc.Transport)
						if trHttp.RequestHeader().Get(DomainKey) != "" {
							domain = trHttp.RequestHeader().Get(DomainKey)
						}
					}
				}
			}

			// 获取请求域
			ctx = metadata.AppendToClientContext(ctx, DomainKey, domain)
			ctx = context.WithValue(ctx, DomainKey, domain)
			return handler(ctx, req)
		}
	}
}
