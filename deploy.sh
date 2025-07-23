export AWS_PROFILE=quenby-admin
export AWS_REGION=us-east-1
aws_account_id=$(aws sts get-caller-identity --query Account --output text)
ecr_endpoint="$aws_account_id.dkr.ecr.us-east-1.amazonaws.com"
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin "$ecr_endpoint"

# export KO_DOCKER_REPO=localhost:5001
export KO_DOCKER_REPO="$ecr_endpoint/knative"
export KO_FLAGS="--base-import-paths --tags oauth"
KO_DEFAULTPLATFORMS=linux/amd64
cd data-plane
mvn clean
cd ..

rm eventing-kafka-*.yaml

export SKIP_CREATE_SECRETS=true
./hack/run.sh deploy
