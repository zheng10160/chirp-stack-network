package external

import (
	"golang.org/x/net/context"

	pb "github.com/jon177/lky-api/go/v1/ns/external/api"
	"github.com/jon177/lky-network-server/internal/api/helpers"
	"github.com/jon177/lky-network-server/internal/storage"
)

type DeviceAPI struct{}

// NewDeviceAPI creates a new NodeAPI.
func NewDeviceAPI() *DeviceAPI {
	return &DeviceAPI{}
}

func (a *DeviceAPI) Get(ctx context.Context, req *pb.GetDeviceRequest) (*pb.GetDeviceResponse, error) {
	d, err := storage.GetDevice(ctx, storage.DB(), req.DevId)

	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	resp := pb.GetDeviceResponse{
		Device: &pb.Device{
			DevId:       d.ID,
			Reserved:    d.Reserved,
			KeyMain:     d.KeyMain,
			KeyExt:      d.KeyExt,
			BatchId:     d.BatchId,
			BatchSerial: d.BatchSerial,
		},
	}

	return &resp, nil
}
