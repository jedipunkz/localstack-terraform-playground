# Description: Create S3 bucket for Terraform state
# aws --endpoint-url=http://localhost:4566 s3 mb s3://terraform-tfstate-bucket
#

# resource "aws_s3_bucket" "terraform_state" {
#   bucket = "terraform-state-bucket"
# }
#
# resource "aws_s3_bucket_acl" "terraform_state" {
#   bucket = aws_s3_bucket.terraform_state.id
#
#   acl = "private"
# }
