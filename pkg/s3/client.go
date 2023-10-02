package s3

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Client represents the S3 client.
type S3Client struct {
	client       *s3.S3
	uploadBucket *string
}

// NewS3Client creates a new S3 client.
func NewS3Client() (*S3Client, error) {
	// Create a new AWS session using the default credentials and configuration
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your desired AWS region
	})
	if err != nil {
		return nil, err
	}

	// Create an S3 client
	svc := s3.New(sess)

	return &S3Client{
		client: svc,
	}, nil
}

// UploadFile uploads a file to the specified S3 bucket and key.
func (c *S3Client) UploadFile(bucket, key, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	_, err = c.client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(key),
		ACL:                  aws.String("private"), // Set ACL (Access Control List) to private, change as needed.
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"), // You can enable server-side encryption by setting this to "AES256".
	})
	if err != nil {
		return err
	}

	fmt.Printf("File uploaded successfully to S3: %s\n", key)
	return nil
}

func main() {
	// Example usage
	s3Client, err := NewS3Client()
	if err != nil {
		log.Fatal("Failed to create S3 client:", err)
	}

	bucketName := "your-bucket-name"
	fileKey := "path/to/destination/folder/filename.ext"
	localFilePath := "path/to/local/file/filename.ext"

	err = s3Client.UploadFile(bucketName, fileKey, localFilePath)
	if err != nil {
		log.Fatal("Failed to upload file to S3:", err)
	}
}
