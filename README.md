## Vendor Mod Setup

### download library (offline mode)

```shell
go mod vendor
```

### vendor mode run

```shell
go run -mod vendor offline_modules.go
```

---

## Install Library

### cobra-cli

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

### sqlc

```shell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### ALL

```shell
go install github.com/spf13/cobra-cli@latest
go install golang.org/x/tools/cmd/godoc@latest
go install github.com/playwright-community/playwright-go/cmd/playwright@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
playwright install --with-deps
```

---

## Pocketbase

###

```shell
go get github.com/pocketbase/pocketbase@latest
```

```go
_ "github.com/pocketbase/pocketbase/apis"
_ "github.com/pocketbase/pocketbase/cmd"
_ "github.com/pocketbase/pocketbase/core"
_ "github.com/pocketbase/pocketbase/daos"
_ "github.com/pocketbase/pocketbase/examples/base"
_ "github.com/pocketbase/pocketbase/forms/validators"
_ "github.com/pocketbase/pocketbase/mails/templates"
_ "github.com/pocketbase/pocketbase/migrations/logs"
_ "github.com/pocketbase/pocketbase/models/schema"
_ "github.com/pocketbase/pocketbase/models/settings"
_ "github.com/pocketbase/pocketbase/plugins/ghupdate"
_ "github.com/pocketbase/pocketbase/plugins/jsvm"
_ "github.com/pocketbase/pocketbase/plugins/migratecmd"
_ "github.com/pocketbase/pocketbase/resolvers"
_ "github.com/pocketbase/pocketbase/tests"
_ "github.com/pocketbase/pocketbase/tokens"
_ "github.com/pocketbase/pocketbase/tools/archive"
_ "github.com/pocketbase/pocketbase/tools/auth"
_ "github.com/pocketbase/pocketbase/tools/cron"
_ "github.com/pocketbase/pocketbase/tools/dbutils"
_ "github.com/pocketbase/pocketbase/tools/filesystem"
_ "github.com/pocketbase/pocketbase/tools/hook"
_ "github.com/pocketbase/pocketbase/tools/inflector"
_ "github.com/pocketbase/pocketbase/tools/list"
_ "github.com/pocketbase/pocketbase/tools/logger"
_ "github.com/pocketbase/pocketbase/tools/mailer"
_ "github.com/pocketbase/pocketbase/tools/migrate"
_ "github.com/pocketbase/pocketbase/tools/osutils"
_ "github.com/pocketbase/pocketbase/tools/rest"
_ "github.com/pocketbase/pocketbase/tools/routine"
_ "github.com/pocketbase/pocketbase/tools/search"
_ "github.com/pocketbase/pocketbase/tools/security"
_ "github.com/pocketbase/pocketbase/tools/store"
_ "github.com/pocketbase/pocketbase/tools/subscriptions"
_ "github.com/pocketbase/pocketbase/tools/template"
_ "github.com/pocketbase/pocketbase/tools/tokenizer"
_ "github.com/pocketbase/pocketbase/tools/types"
_ "github.com/pocketbase/pocketbase/tools/ui"
```
