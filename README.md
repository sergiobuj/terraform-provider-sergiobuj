# Terraform provider Sample

Terraform provider sample. Proof of concept to create a terraform data resource that holds any kind of information.

```tree
.
├── README.md
├── example
│   ├── main.tf
│   └── terraform-provider-sergiobuj
├── data_source_network.go
├── main.go
└── provider.go
```

# Quickstart

```bash
$ go build -o example/terraform-provider-sergiobuj
$ cd example
$ terraform init
$ terraform apply
```

The Terraform configuration looks like this:

```terraform
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
```

Terraform output should show:

```bash
$ terraform output

vpc_non_prod = ARN:VPCstage
vpc_prod = ARN:VPCproduction
```

# Implementation

For the sample data resource see `data_source_network.go`.

- `func dataSourceNetwork() *schema.Resource`
  - To define the schema of the resource.
- `func dataSourcePreprodNetworkRead(d *schema.ResourceData, meta interface{}) error`.
  - To fill in all the data (make api requests or generate the data).
