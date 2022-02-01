/*
Copyright © 2022 Reese Armstrong <me@reeseric.ci>

*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [ending]",
	Short: "Get existing redirect",
	Args: func(cmd *cobra.Command, args []string) error {
    	if len(args) < 1 {
      		return errors.New("requires ending to be specified")
    	}
		return nil
  	},
	Run: func(cmd *cobra.Command, args []string) {
		var client = graphql.NewClient(viper.GetString("instance") + "/api/graphql", http.DefaultClient)
		resp, err := getRedirect(context.Background(), client, args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		if resp.GetRedirect.Url == "" {
			fmt.Println("no redirect found")
			return
		} else {
			fmt.Println(resp.GetRedirect.Url)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

//go:generate go run github.com/Khan/genqlient graphql/genqlient.yaml