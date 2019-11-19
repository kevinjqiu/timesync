package cmd

import (
	"context"
	"fmt"
	"github.com/kevinjqiu/timesync/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

var serverFlags struct {
	bind string
}

var clientFlags struct {
	serverAddr string
}

var rootCmd = &cobra.Command{
	Use:   "timesync",
	Short: "Sync time between client and server using Cristian's algorithm",
}

func newServerCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "Run server",
		Run: func(cmd *cobra.Command, args []string) {

			lis, err := net.Listen("tcp", serverFlags.bind)
			if err != nil {
				logrus.Fatalf("failed to listen: %v", err)
				return
			}

			grpcServer := grpc.NewServer()
			pkg.RegisterTimeSyncServer(grpcServer, pkg.NewServer())
			grpcServer.Serve(lis)
		},
	}

	cmd.Flags().StringVarP(&serverFlags.bind, "bind", "b", "localhost:8080", "Bind address")
	return cmd
}

func newClientCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "client",
		Aliases: []string{"c"},
		Short:   "Run client",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(clientFlags.serverAddr, grpc.WithInsecure())
			if err != nil {
				logrus.Fatalf("cannot connect to server: %v", err)
			}
			defer conn.Close()

			client := pkg.NewTimeSyncClient(conn)
			t1 := time.Now().UTC().UnixNano()
			serverTime, err := client.GetServerTime(context.Background(), &pkg.GetServerTimeParams{})
			t2 := time.Now().UTC().UnixNano()

			if err != nil {
				logrus.Fatal(err)
			}
			syncedTime := serverTime.Ts + (t2-t1)/2
			logrus.Infof("Server time:  %v", serverTime.Ts)
			logrus.Infof("RTT:          %v", time.Duration(t2-t1))
			logrus.Infof("Sync'ed time: %v", syncedTime)
		},
	}

	cmd.Flags().StringVarP(&clientFlags.serverAddr, "server", "s", "localhost:8080", "Server address")
	return cmd
}

func init() {
	rootCmd.AddCommand(newServerCobraCommand())
	rootCmd.AddCommand(newClientCobraCommand())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
