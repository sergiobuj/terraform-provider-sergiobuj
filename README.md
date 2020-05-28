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

Terraform output should show:

```bash
$ terraform output

subnets_non_prod = [
  "ARN::subnet:default",
]
subnets_prod = [
  "ARN::subnet:production",
]
vpc_non_prod = ARN:VPCstage
vpc_prod = ARN:VPCproduction
```

# Implementation

For the sample data resource see `data_source_network.go`.

- `func dataSourceNetwork() *schema.Resource`
  - To define the schema of the resource.
- `func dataSourcePreprodNetworkRead(d *schema.ResourceData, meta interface{}) error`.
  - To fill in all the data (make api requests or generate the data).