package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var EamilServiceSess *session.Session

// NewSession makes a new aws session
func NewSession(id, key, region string) (*session.Session, error) {
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(id, key, ""),
	})

	if err != nil {
		return nil, err
	}
	return s, nil
}
