## go modules replacement

In order to use this patch, if using Go modules

### Clone this repository
1. Clone this repository

```shell
git clone https://github.com/orijtech/google-cloud-go.git
cd google-cloud-go && git checkout storage-remove-writer-indefinite-retry
echo $(pwd)
```

### Replace in the go.mod directive
2. For the application using `cloud.google.com/go/storage`, please edit the `replace` clause in your `go.mod`
file as

```shell
replace cloud.google.com/go/storage => {{PWD}}
```
where `{{PWD}}` is the path that we obtained from [Step 1](#clone-this-repository)
