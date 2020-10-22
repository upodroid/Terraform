variable "sensitive" {
    default   = "foo"
    sensitive = true
}

resource "aws_instance" "foo" {
    foo = var.sensitive
}

output "out" {
  value     = var.sensitive
}
