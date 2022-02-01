/*
Copyright © 2022 Reese Armstrong <me@reeseric.ci>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var instance string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shortie",
	Short: "A cli for the shortie.sh url shortener",
	Long: `This is the command line companion for shortie.sh.
	shortie is a URL shortener that is similiar in featureset to bit.ly or tinyurl. 
	The main differences are that it is much simpler to use, open-source, and privacy respecting!
	
	The source code is available at https://github.com/shortie-sh/server
	The source code for this CLI is available at https://github.com/shortie-sh/cli
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(defCmd string) {
	var cmdFound bool
	cmd :=rootCmd.Commands()
  
	for _,a:=range cmd{
	  for _,b:=range os.Args[1:] {
		if a.Name()==b {
		 cmdFound=true
		  break
		}
	  }
	}
	if !cmdFound {
	  args:=append([]string{defCmd}, os.Args[1:]...)
	  rootCmd.SetArgs(args)
	}
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }

func init() {
	cobra.OnInitialize(initInstance)
	rootCmd.PersistentFlags().StringVar(&instance, "instance", "https://server.shortie.sh", "shortie-server instance")

}

func initInstance() {
	viper.BindEnv("instance", "SHORTIE_INSTANCE")
	viper.SetDefault("instance", "https://server.shortie.sh")
	if instance != "" && instance != "https://server.shortie.sh"  {
		viper.Set("instance", instance)
	}
	viper.SetEnvPrefix("shortie")
	viper.AutomaticEnv()
}
