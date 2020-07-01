package external

import (
	pb "github.com/jon177/lky-api/go/v1/ns/external/api"

	"github.com/jon177/lky-network-server/internal/config"
	"google.golang.org/grpc"
	"github.com/jon177/lky-network-server/internal/api/helpers"
	"net/http"
	"strings"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/net/http2"

	log "github.com/sirupsen/logrus"
	"time"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"fmt"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
)

var (
	bind string
)

// Setup configures the API package.
func Setup(conf config.Config) error {

	bind = conf.NetworkServer.Api.Bind

	return setupApi(conf)
}

func setupApi(conf config.Config) error {
	grpcOpts := helpers.GetgRPCServerOptions()
	grpcServer := grpc.NewServer(grpcOpts...)
	pb.RegisterDeviceServiceServer(grpcServer, NewDeviceAPI())

	var clientHTTPHandler http.Handler

	// switch between gRPC and "plain" http handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			if clientHTTPHandler == nil {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}

			clientHTTPHandler.ServeHTTP(w, r)
		}
	})

	go func() {
		log.Fatal(http.ListenAndServe(bind, h2c.NewHandler(handler, &http2.Server{})))
	}()


	// give the http server some time to start
	time.Sleep(time.Millisecond * 100)

	// setup the HTTP handler
	var err error
	clientHTTPHandler, err = setupHTTPAPI(conf)
	if err != nil {
		return err
	}

	return nil
}

func setupHTTPAPI(conf config.Config) (http.Handler, error) {
	r := mux.NewRouter()

	// setup json api handler
	jsonHandler, err := getJSONGateway(context.Background())
	if err != nil {
		return nil, err
	}

	r.PathPrefix("/api").Handler(jsonHandler)

	return wsproxy.WebsocketProxy(r), nil
}

func getJSONGateway(ctx context.Context) (http.Handler, error) {
	// dial options for the grpc-gateway
	var grpcDialOpts []grpc.DialOption
	grpcDialOpts = append(grpcDialOpts, grpc.WithInsecure())

	bindParts := strings.SplitN(bind, ":", 2)
	if len(bindParts) != 2 {
		log.Fatal("get port from bind failed")
	}
	apiEndpoint := fmt.Sprintf("localhost:%s", bindParts[1])

	mux := runtime.NewServeMux()

	if err := pb.RegisterDeviceServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register application handler error")
	}

	return mux, nil
}