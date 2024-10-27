package s3bucket

import (
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateS3Bucket creates an S3 bucket with optional versioning
func CreateS3Bucket(ctx *pulumi.Context, bucketName string, versioningEnabled bool) (*s3.Bucket, error) {
    bucket, err := s3.NewBucket(ctx, bucketName, &s3.BucketArgs{
        Bucket: pulumi.String(bucketName),
        Versioning: &s3.BucketVersioningArgs{
            Enabled: pulumi.Bool(versioningEnabled),
        },
    })
    if err != nil {
        return nil, err
    }

    return bucket, nil
}
