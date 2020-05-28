data "sergiobuj_network" "nonprod" {
    namespace = "stage"
}

data "sergiobuj_network" "production" {
    namespace = "production"
}

output "vpc_non_prod" {
  value = data.sergiobuj_network.nonprod.vpc
}

output "vpc_prod" {
  value = data.sergiobuj_network.production.vpc
}

output "subnets_non_prod" {
  value = data.sergiobuj_network.nonprod.private_subnets.*.id
}

output "subnets_prod" {
  value = data.sergiobuj_network.production.private_subnets.*.id
}