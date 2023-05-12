# terraform {
#   backend "local" {}
# }

terraform {
  backend "s3" {
    endpoint          = "http://localhost:4566"
    dynamodb_endpoint = "http://localhost:4566"

    skip_credentials_validation = true
    skip_metadata_api_check     = true
    force_path_style            = true

    encrypt        = true
    bucket         = "terraform-state-bucket"
    dynamodb_table = "terraform-state-lock"
    key            = "terraform.tfstate"
    region         = "us-east-1"
  }
}

provider "aws" {
  region = "us-east-1"

  access_key = "mock_access_key"
  secret_key = "mock_secret_key"

  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    iam      = "http://localhost:4566"
    kinesis  = "http://localhost:4566"
    lambda   = "http://localhost:4566"
    s3       = "http://localhost:4566"
    ec2      = "http://localhost:4566"
    dynamodb = "http://localhost:4566"
  }
}
