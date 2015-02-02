package secman

import "github.com/awslabs/aws-sdk-go/gen/s3"

type FakeS3 struct {
	ListRequests  []s3.ListObjectsRequest
	ListResponses []s3.ListObjectsOutput

	DeleteRequests  []s3.DeleteObjectRequest
	DeleteResponses []s3.DeleteObjectOutput

	PutRequests  []s3.PutObjectRequest
	PutResponses []s3.PutObjectOutput
}

func (f *FakeS3) ListObjects(req *s3.ListObjectsRequest) (*s3.ListObjectsOutput, error) {
	f.ListRequests = append(f.ListRequests, *req)
	resp := f.ListResponses[0]
	f.ListResponses = f.ListResponses[1:]
	return &resp, nil
}

func (f *FakeS3) DeleteObject(req *s3.DeleteObjectRequest) (*s3.DeleteObjectOutput, error) {
	f.DeleteRequests = append(f.DeleteRequests, *req)
	resp := f.DeleteResponses[0]
	f.DeleteResponses = f.DeleteResponses[1:]
	return &resp, nil
}

func (f *FakeS3) PutObject(req *s3.PutObjectRequest) (*s3.PutObjectOutput, error) {
	f.PutRequests = append(f.PutRequests, *req)
	resp := f.PutResponses[0]
	f.PutResponses = f.PutResponses[1:]
	return &resp, nil
}