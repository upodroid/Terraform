module "child" {
  source = "./child"
}

output "sens" {
    value = module.child.out
}