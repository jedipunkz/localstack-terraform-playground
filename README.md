# localstack-terraform-playground

## Create DynamoDB table for locking tfstate

```shell
aws dynamodb create-table --table-name terraform-state-lock --attribute-definitions AttributeName=LockID,AttributeType=S --key-schema AttributeName=LockID,KeyType=HASH --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 --endpoint-url=http://localhost:4566
```

## Create S3 bucket for terraform backend

```shell
aws --endpoint-url=http://localhost:4566 s3 mb s3://terraform-tfstate-bucket
```

## Terraform Plan / Apply

```shell
terraform plan
terraform apply
```
