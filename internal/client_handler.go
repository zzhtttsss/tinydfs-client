package internal

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"tinydfs-base/common"
	"tinydfs-base/protocol/pb"
)

var (
	GlobalClientHandler *ClientHandler
	once                = &sync.Once{}
)

type ClientHandler struct {
	pb.UnimplementedMasterAddServiceServer
}

func init() {
	if GlobalClientHandler == nil {

		once.Do(func() {
			GlobalClientHandler = &ClientHandler{}
		})
	}
}

func (c *ClientHandler) Check4Add(args *pb.CheckArgs4AddArgs) (*pb.CheckArgs4AddReply, error) {
	addr := viper.GetString(common.MasterAddr) + viper.GetString(common.MasterPort)
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//TODO 能否复用？
	client := pb.NewMasterAddServiceClient(conn)
	ctx := context.Background()
	reply, err := client.CheckArgs4Add(ctx, args)
	return reply, err
}

func (c *ClientHandler) CheckAndStat(args *pb.CheckAndStatArgs) (*pb.CheckAndStatReply, error) {
	addr := viper.GetString(common.MasterAddr) + viper.GetString(common.MasterPort)
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewMasterStatServiceClient(conn)
	ctx := context.Background()
	reply, err := client.CheckAndStat(ctx, args)
	return reply, err
}

func (c *ClientHandler) CheckAndList(args *pb.CheckAndListArgs) (*pb.CheckAndListReply, error) {
	addr := viper.GetString(common.MasterAddr) + viper.GetString(common.MasterPort)
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewMasterListServiceClient(conn)
	ctx := context.Background()
	reply, err := client.CheckAndList(ctx, args)
	return reply, err
}

func (c *ClientHandler) GetDataNodes4Add(args *pb.GetDataNodes4AddArgs) (*pb.GetDataNodes4AddReply, error) {
	addr := viper.GetString(common.MasterAddr) + viper.GetString(common.MasterPort)
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewMasterAddServiceClient(conn)
	ctx := context.Background()
	reply, err := client.GetDataNodes4Add(ctx, args)
	return reply, err
}

func (c *ClientHandler) UnlockDic4Add(args *pb.UnlockDic4AddArgs) (*pb.UnlockDic4AddReply, error) {
	addr := viper.GetString(common.MasterAddr) + viper.GetString(common.MasterPort)
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewMasterAddServiceClient(conn)
	ctx := context.Background()
	reply, err := client.UnlockDic4Add(ctx, args)
	return reply, err
}

func (c *ClientHandler) ReleaseLease4Add(args *pb.ReleaseLease4AddArgs) (*pb.ReleaseLease4AddReply, error) {
	addr := viper.GetString(common.MasterAddr) + viper.GetString(common.MasterPort)
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewMasterAddServiceClient(conn)
	ctx := context.Background()
	reply, err := client.ReleaseLease4Add(ctx, args)
	return reply, err
}
