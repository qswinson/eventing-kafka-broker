package security

import (
	"context"
	"os"

	"github.com/IBM/sarama"
	"github.com/aws/aws-msk-iam-sasl-signer-go/signer"
)

type MSKAccessTokenProvider struct {
}

func (m *MSKAccessTokenProvider) Token() (*sarama.AccessToken, error) {
	signer.AwsDebugCreds = true
	// Load AWS region from environment variable
	awsRegion := os.Getenv("AWS_DEFAULT_REGION")
	if awsRegion == "" {
		awsRegion = "us-east-1"
	}
	token, _, err := signer.GenerateAuthToken(context.TODO(), awsRegion)
	return &sarama.AccessToken{Token: token}, err
}

func MSKAccessTokenProviderGeneratorFunc() *MSKAccessTokenProvider {
	return &MSKAccessTokenProvider{}
}
