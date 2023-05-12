# 
# aws dynamodb create-table --table-name terraform-state-lock --attribute-definitions AttributeName=LockID,AttributeType=S --key-schema AttributeName=LockID,KeyType=HASH --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 --endpoint-url=http://localhost:4566
#

# resource "aws_dynamodb_table" "terraform_state_lock" {
#   name           = "terraform-state-lock"
#   hash_key       = "LockID"
#   read_capacity  = 1
#   write_capacity = 1
#
#   attribute {
#     name = "LockID"
#     type = "S"
#   }
# }
