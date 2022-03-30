package grpc

import (
	"appstruct/internal/adapters/app/api"
	"appstruct/internal/adapters/core/arithmetic"
	"appstruct/internal/adapters/framework/left/grpc/pb"
	"appstruct/internal/adapters/framework/right/db"
	"appstruct/internal/ports"
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"
)

const buffSize = 1 << 20

var (
	listener       *bufconn.Listener
	apiPort        ports.APIPort
	arithmeticPort ports.ArithmeticPort
	dbPort         ports.DBPort
	grpcPort       ports.GRPCPort
	dbDriver       string
	dbDSN          string
)

func init() {
	listener = bufconn.Listen(buffSize)
	grpcServer := grpc.NewServer()

	// framework - driven by application
	dbPort = getNewDBPortInitialized()

	// domain
	arithmeticPort = arithmetic.NewAdapter()

	// application
	apiPort = api.NewAdapter(dbPort, arithmeticPort)
	grpcPort = NewAdapter(apiPort)

	// framework - driver of the application
	pb.RegisterArithmeticServiceServer(grpcServer, grpcPort)
	go func() {
		defer dbPort.Close()
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func buffDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(buffDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc.DialContext() error: %v", err)
	}
	return conn
}

func getNewDBPortInitialized() ports.DBPort {
	dbDriver = os.Getenv("DB_DRIVER")
	dbDSN = os.Getenv("DB_DSN")
	adapter, err := db.NewAdapter(dbDriver, dbDSN)
	if err != nil {
		log.Fatalln("Error creating db adapter:", err)
	}

	return adapter
}

///TESTS
type calculateEvent func(context.Context, *pb.OperationParameters, ...grpc.CallOption) (*pb.Answer, error)
type calculateHandler func(client pb.ArithmeticServiceClient) calculateEvent

func executeCalculation(t *testing.T, calcHandler calculateHandler) (*pb.Answer, error) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Fatalf("grpc.ClientConn.Close() error: %v", err)
		}
	}(conn)

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 6,
		B: 2,
	}

	return calcHandler(client)(ctx, params)
}

func TestGetAdd(t *testing.T) {
	answer, err := executeCalculation(t, func(client pb.ArithmeticServiceClient) calculateEvent {
		return client.GetAdd
	})

	require.NoError(t, err)
	require.Equal(t, int32(8), answer.Value)
}

func TestGetSub(t *testing.T) {
	answer, err := executeCalculation(t, func(client pb.ArithmeticServiceClient) calculateEvent {
		return client.GetSub
	})

	require.NoError(t, err)
	require.Equal(t, int32(4), answer.Value)
}

func TestGetMulti(t *testing.T) {
	answer, err := executeCalculation(t, func(client pb.ArithmeticServiceClient) calculateEvent {
		return client.GetMulti
	})

	require.NoError(t, err)
	require.Equal(t, int32(12), answer.Value)
}

func TestGetDiv(t *testing.T) {
	answer, err := executeCalculation(t, func(client pb.ArithmeticServiceClient) calculateEvent {
		return client.GetDiv
	})

	require.NoError(t, err)
	require.Equal(t, int32(3), answer.Value)
}
