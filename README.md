## Vendor Mod Setup

### install

```shell
go mod vendor
```

### run

```shell
go run -mod vendor offline_modules.go
```

---

## Install Library

### godoc

```shell
go install github.com/spf13/cobra-cli@latest
```

### godoc

```shell
go install golang.org/x/tools/cmd/godoc@latest
```

### playwright-go

```shell
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps
# Or
go install github.com/playwright-community/playwright-go/cmd/playwright@latest
playwright install --with-deps
```

### sqlite

```shell
go install github.com/mattn/go-sqlite3@latest
```

### sqlc

```shell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
