package aws

import (
	"bytes"
	"fmt"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Aws struct {
	region_name string
	session     *session.Session
}

func NewAws(region_name string) (*Aws, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region_name),
	})
	if err != nil {
		fmt.Print("Unabale to initilase aws")
	}
	return &Aws{
		region_name: region_name,
		session:     sess,
	}, nil
}

type Boto3Wrapper struct {
	s3_client *s3.S3
}

func NewBoto3Wrapper(accessKey, secretKey, region string) (*Boto3Wrapper, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}
	return &Boto3Wrapper{
		s3_client: s3.New(sess),
	}, nil
}

func (bw *Boto3Wrapper) GetSignedUrl(method, bucket, key string, expiry time.Duration) (string, error) {
	req, _ := bw.s3_client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	req.HTTPRequest.Method = method
	url, err := req.Presign(expiry)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (bw *Boto3Wrapper) UploadObjectWithPublicAccess(bucket, key, attachmentType string, file []byte) (string, error) {
	reader := bytes.NewReader(file)
	_, err := bw.s3_client.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(bucket),
		Key:                aws.String(key),
		ACL:                aws.String("public-read"),
		ContentType:        aws.String(attachmentType),
		ContentDisposition: aws.String("attachment"),
		Body:               reader,
	})
	if err != nil {
		return "", err
	}

	encodedKey := url.PathEscape(key)
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, bw.s3_client.Config.Region, encodedKey)
	return url, nil
}

func (bw *Boto3Wrapper) DeleteObject(bucket, key string) error {
	_, err := bw.s3_client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err
}

func (bw *Boto3Wrapper) GetS3Clinet() (*s3.S3, error) {
	return bw.s3_client, nil
}
