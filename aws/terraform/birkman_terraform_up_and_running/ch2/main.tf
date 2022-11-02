resource "aws_instance" "example" {
	ami = "ami-0d593311db5abb72b"
	instance_type = "t2.micro"
	
	tags = {
		Name = "terraform-example"
	}
}

