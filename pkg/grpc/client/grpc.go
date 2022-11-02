package grpc

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/juanjoss/shorturl/pkg/grpc/qrgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GenerateQR(source string) ([]byte, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("qrgen:%s", os.Getenv("QRGEN_GRPC_SERVER_PORT")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
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
