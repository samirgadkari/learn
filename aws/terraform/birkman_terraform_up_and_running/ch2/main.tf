
/* Getting an error when trying to talk to this image using
*  http://PublicIPAddr:8080
*  Error is connection refused. Have to figure out how to SSH into the instance.
*/
resource "aws_instance" "example" {

	/* Found this AL2 (Amazon Linux 2) image by:
		Make sure you're in the right zone. Images are zone-based.
		EC2 Page -> Launch Instance -> select Amazon Linux
		Note down it's AMI ID. Go to:
		AMIs -> Paste the AMI ID there
		Check the columns to see if this is what you want.
		The AL2 images are based on Fedora. They also have
		SELinux (Security Enhanced Linux) layer enabled,
		although it is permissive, so make sure to lock it down and
		DO NOT use it for production as-is */
	ami = "ami-0d593311db5abb72b"
	instance_type = "t2.micro"

	/* Specify the ID of the VPC security group that should be associated
		with this ec2 instance */
	vpc_security_group_ids = [aws_security_group.instance.id]
	
	/* The user_data variable points to the commands to run on startup.
		In <<-EOF, the << is required, the - after is needed if your
		commands are indented. The EOF is a word that you choose.
		All lines from that word to the word appearing again are
		considered part of the list of commands */
	user_data = <<-EOF
				#!/bin/bash
				sudo yum update -y
				sudo yum install -y httpd
				echo "Hello, World" > index.html
				sudo systemctl start httpd
				EOF

	tags = {
		Name = "terraform-example"
	}
}

/* Since AWS does not allow traffic from/to an AWS instance,
	we have to create a security group and enable this traffic.
	When you specify ingress traffic, the outgoing traffic is
	automatically allowed. */
resource "aws_security_group" "instance" {
	name = "terraform-example-instance"

	ingress {
		from_port = 8080
		to_port = 8080
		protocol = "tcp"
		cidr_blocks = ["0.0.0.0/0"]
	}
}

