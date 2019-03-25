package filestore

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//S3Handler handles S3 sessions
type S3Handler struct {
	Session *session.Session
	Bucket  string
}

//GetHandler connects user to S3 session and creates Handler
func GetHandler() (S3Handler, error) {
	var handler S3Handler
	// get the S3 session
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		return handler, err
	}

	// create handler from session
	handler = S3Handler{
		Session: sess,
		Bucket:  os.Getenv("S3_BUCKET"),
	}
	return handler, nil
}

//UploadFile takes S3 session and key/value for file and stores in Bucket
func (h S3Handler) UploadFile(uuid string, body string) error {
	buffer := []byte(body)

	_, err := s3.New(h.Session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(h.Bucket),
		Key:                  aws.String(uuid),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(len(buffer))),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	// returns the error if there is one
	return err
}

//ReadFile takes S3 session and with key Reads file
func (h S3Handler) ReadFile(uuid string) (string, error) {
	// gets results from Bucket
	results, err := s3.New(h.Session).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String("scripts/" + uuid),
	})
	if err != nil {
		return "", err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return "", err
	}
	// returns string of file
	return string(buf.Bytes()), nil
}
