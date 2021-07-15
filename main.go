package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var baseDir string

const timeout = 5 * time.Second

func main() {
	log.SetFlags(log.Lshortfile)
	var workNum int
	var postUrl string
	var command = &cobra.Command{
		Use:   "PicdownGo -u URL",
		Short: "Download pictures",
		Long:  `Download given url pictures by cmdline`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if postUrl == "" {
				return fmt.Errorf("please use 'PicdownGo -u URL'")
			}
			baseDir = "Pic"
			err = os.MkdirAll(baseDir, 0755)
			if err != nil {
				return fmt.Errorf("create dictionary failed,%w", err)
			}
			jsonFileAddr := fmt.Sprintf("./parser.json")
			jsonfile, err := reloadParser(jsonFileAddr)
			if err != nil {
				return fmt.Errorf("loading jsonfile failed,%w", err)
			}
			err = crawler(postUrl, workNum, jsonfile)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			log.Println("pictures have downloaded")
			return nil
		}}
	command.Flags().StringVarP(&postUrl, "URL", "u", "", "URL of post")
	command.Flags().IntVarP(&workNum, "workerNum", "w", 20, "number of workers")
	err := command.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fprintf:%v\n", err.Error())
		os.Exit(1)
	}
}
