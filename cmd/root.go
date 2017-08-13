package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ghodss/yaml"
)

var cfgFile string
var jsonFile string

func isJSON(b []byte) bool {
	var js interface{}
	return json.Unmarshal(b, &js) == nil 
}

func json2yaml(json []byte) (s string) {
	if ! isJSON(json) {
		fmt.Println("The json argument is not valid!")
		os.Exit(1)
	}

	if y, err := yaml.JSONToYAML(json); err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		s = string(y)
	}

	return
}

var RootCmd = &cobra.Command{
	Use:   "json2yaml",
	Short: "Convert json to yaml then prints to stdout",
	Long: `json2yaml is a command line tool that convert json to yaml
either from stdin or from a file then prints the result to stdout 

examples:

	$ json2yaml --file=foo.json
	$ json2yaml < foo.json
	$ echo '{"foo": "bar"}' | json2yaml
	
`,
	Run: func(cmd *cobra.Command, args []string) {
	
		stdin := ""
		fi, _ := os.Stdin.Stat()

		if (fi.Mode() & os.ModeNamedPipe != 0) || (fi.Size() != 0) {
			bytes, _ := ioutil.ReadAll(os.Stdin)
			stdin = string(bytes)
		}

		if (jsonFile == "" && stdin == "") || (jsonFile != "" && stdin != "") {
			cmd.Help()

		} else if jsonFile != "" {

			b, err := ioutil.ReadFile(jsonFile)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}

			fmt.Print(json2yaml(b))

		} else {
			fmt.Print(json2yaml([]byte(stdin)))
		}
		
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.json2yaml.yaml)")
	RootCmd.Flags().StringVarP(&jsonFile, "file", "f", "", "json file")
}


func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".json2yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
