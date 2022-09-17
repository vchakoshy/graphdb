/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"math/rand"

	"github.com/spf13/cobra"
	"github.com/vchakoshy/graphdb/graph"
)

// cliTestCmd represents the cliTest command
var cliTestCmd = &cobra.Command{
	Use:   "cliTest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		g := graph.NewGraph()
		var i int64
		for i = 1; i <= 100; i++ {
			for j := 0; j < 2; j++ {
				g.AddFollow(i, rand.Int63n(100))
			}
		}

		for i = 1; i <= 100; i++ {
			f, err := g.GetFollows(i)
			if err != nil {
				continue
			}

			for _, j := range f {
				log.Printf("user: %d follows %d", i, j)
			}

		}

		f, _ := g.GetFriendsOfFriends(1)
		for _, j := range f {
			log.Printf("user: %d fof: %d", 1, j)
		}

	},
}

func init() {
	rootCmd.AddCommand(cliTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
