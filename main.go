package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"pulumi-s3-go/s3bucket"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        // Create an S3 bucket
        _, err := s3bucket.CreateS3Bucket(ctx, "my-pulumi-bucket-20241027", true)
        if err != nil {
            return err
        }
        return nil
    })
}
