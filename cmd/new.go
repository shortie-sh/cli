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

var ending string

// getCmd represents the get command
var newCmd = &cobra.Command{
	Use:   "new [url]",
	Short: "Create new redirect",
	Args: func(cmd *cobra.Command, args []string) error {
    	if len(args) < 1 {
      		return errors.New("requires url to be specified")
    	}
		return nil
  	},
	Run: func(cmd *cobra.Command, args []string) {
		var client = graphql.NewClient(viper.GetString("instance") + "/api/graphql", http.DefaultClient)
		if ending == "" {
			resp, err := createRedirectGenEnding(context.Background(), client, args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Created redirect at " + viper.GetString("instance") + "/"  + resp.CreateRedirect.Ending)
			return
		} else {
			resp, err := createRedirect(context.Background(), client, args[0], ending)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Created redirect at " + viper.GetString("instance") + "/"  + resp.CreateRedirect.Ending)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&ending, "ending", "e", "", "Custom ending for redirect")
}

//go:generate go run github.com/Khan/genqlient graphql/genqlient.yaml