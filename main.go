package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/luisnquin/s3-edit/cli"
	"github.com/luisnquin/s3-edit/cli/s3"
	"github.com/luisnquin/s3-edit/config"
)

func main() {
	awsProfile := flag.String("profile", "", "AWS profile name")
	awsRegion := flag.String("region", "", "AWS region")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s <S3-FILE-PATH> [flags]\n", os.Args[0])

		fmt.Fprintf(os.Stderr, "\nFlags: \n")
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "\nExamples: \n")
		fmt.Fprintf(os.Stderr, "  %s s3://bucket/path/to/file\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s s3://bucket/path/to/file -profile=my-profile\n", os.Args[0])
	}

	flag.Parse()

	if args := flag.Args(); len(args) == 1 {
		path, err := s3.ParsePath(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var awsConfig aws.Config

		if *awsRegion != "" {
			awsConfig.Region = awsRegion
		}

		params, err := config.NewAWSParams(*awsProfile, awsConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cli.Edit(path, params)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
