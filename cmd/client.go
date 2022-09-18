/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vchakoshy/graphdb/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial("127.0.0.1:8080", opts...)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		client := service.NewGraphdbClient(conn)
		ctx := context.Background()
		client.AddFollow(ctx, &service.Follow{From: 1, To: 2})
		client.AddFollow(ctx, &service.Follow{From: 2, To: 3})
		client.AddFollow(ctx, &service.Follow{From: 2, To: 4})
		client.AddFollow(ctx, &service.Follow{From: 2, To: 5})

		res, err := client.GetFriendsOfFriends(ctx, &service.User{Id: 1})
		if err != nil {
			panic(err)
		}
		for _, i := range res.GetUsers() {
			log.Println("fof of ", 1, i)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
