package grpc

import (
	"context"
	"os"
	"time"

	pb "github.com/juanjoss/url-service/handlers/grpc/qrgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = os.Getenv("GRPC_SERVER_HOST") + ":" + os.Getenv("GRPC_SERVER_PORT")
)

func GenerateQR(source string) ([]byte, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return []byte{}, err
	}
	defer conn.Close()

	client := pb.NewQrGeneratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GenerateQR(ctx, &pb.QrGenRequest{Source: source})
	if err != nil {
		return []byte{}, err
	}

	return res.Barcode, nil
}
