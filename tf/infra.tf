#variable "global_access_key" {type=string}  
#variable "global_secret_key" {type=string}  
  
variable "env" {  
    default= "test"  
}  
provider "aws" {  
#  access_key = "${var.global_access_key}"  
#  secret_key = "${var.global_secret_key}"  
  region = "us-east-1"  
}  
  
resource "aws_s3_bucket" "bucket" {  
  bucket = "dmaic-appset-bucket"  
}  
  
