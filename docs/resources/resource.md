# `secret_resource` Resource

This is a special resource that can only be created on "import".

## Example Usage

```hcl
resource "secret_resource" "my_api_key" {
  lifecycle { prevent_destroy = true }
}
```

Then import the secret:

`terraform import secret_resource my_api_key "dfgdsgdsfgsdfgfd"`

To cycle a secret, use `terraform state rm` to forget the previous value.

## Argument Reference

None

## Attribute Reference

* `value` - (Computed, Sensitive) The stored secret
