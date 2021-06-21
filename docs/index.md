# secret Provider

Store secrets in the Terraform state

## Example

```tf
resource "secret_resource" "my_api_key" {
  lifecycle { prevent_destroy = true }
}

resource "some_other_resource" "foo" {
  attr = secret_resource.my_api_key.value
}
```

## Resources

* [secret_resource](resources/resource.md)
