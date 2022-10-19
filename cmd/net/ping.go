/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	fileOut string

	// Logic
	client = http.Client{
		Timeout: time.Second * 2,
	}
)

func ping(domain string, fileOut string) (int, error) {
	url := "http://" + domain

	fmt.Println("File OUT", len(fileOut))

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote URL and returns the response ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath, fileOut); err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(args)
			fmt.Println("RESPONSE", resp)

			if len(fileOut) > 0 {
				f, _ := os.Create(fileOut)
				defer f.Close()

				w := bufio.NewWriter(f)
				fmt.Fprintf(w, "Here is the file %v", fileOut)
				w.Flush()
				//f.WriteString("The file is")
			}
		}
	},
}

func init() {

	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to pin")
	pingCmd.Flags().StringVarP(&fileOut, "file", "F", "", "output to provided file name")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
