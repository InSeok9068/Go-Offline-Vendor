## Vendor Mod Setup

### download library (offline mode)

```shell
go get github.com/AlecAivazis/survey/v2
go get github.com/atotto/clipboard
go get github.com/briandowns/spinner
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/glamour
go get github.com/charmbracelet/lipgloss
go get github.com/fatih/color
go get github.com/gizak/termui/v3
go get github.com/go-rod/rod
go get github.com/gocolly/colly
go get github.com/jmoiron/sqlx
go get github.com/olekukonko/tablewriter
go get github.com/playwright-community/playwright-go
go get github.com/schollz/progressbar/v3
go get github.com/sijms/go-ora/v2
go get github.com/spf13/cobra
go get github.com/spf13/viper
go get github.com/stretchr/testify
go get github.com/xuri/excelize/v2
go get gopkg.in/gomail.v2
go get modernc.org/sqlite
```

```shell
go mod vendor
```

### vendor mode run

```shell
go run -mod vendor offline_modules.go
```

---

## Install Library

```shell
playwright install --with-deps
```

---

## Pocketbase

```go
_ "github.com/pocketbase/pocketbase/apis"
_ "github.com/pocketbase/pocketbase/cmd"
_ "github.com/pocketbase/pocketbase/core"
_ "github.com/pocketbase/pocketbase/core/validators"
_ "github.com/pocketbase/pocketbase/examples/base"
_ "github.com/pocketbase/pocketbase/forms"
_ "github.com/pocketbase/pocketbase/mails"
_ "github.com/pocketbase/pocketbase/mails/templates"
_ "github.com/pocketbase/pocketbase/migrations"
_ "github.com/pocketbase/pocketbase/plugins/ghupdate"
_ "github.com/pocketbase/pocketbase/plugins/jsvm"
_ "github.com/pocketbase/pocketbase/plugins/migratecmd"
_ "github.com/pocketbase/pocketbase/tests"
_ "github.com/pocketbase/pocketbase/tools/archive"
_ "github.com/pocketbase/pocketbase/tools/auth"
_ "github.com/pocketbase/pocketbase/tools/cron"
_ "github.com/pocketbase/pocketbase/tools/dbutils"
_ "github.com/pocketbase/pocketbase/tools/filesystem"
_ "github.com/pocketbase/pocketbase/tools/filesystem/blob"
_ "github.com/pocketbase/pocketbase/tools/hook"
_ "github.com/pocketbase/pocketbase/tools/inflector"
_ "github.com/pocketbase/pocketbase/tools/list"
_ "github.com/pocketbase/pocketbase/tools/logger"
_ "github.com/pocketbase/pocketbase/tools/mailer"
_ "github.com/pocketbase/pocketbase/tools/osutils"
_ "github.com/pocketbase/pocketbase/tools/picker"
_ "github.com/pocketbase/pocketbase/tools/router"
_ "github.com/pocketbase/pocketbase/tools/routine"
_ "github.com/pocketbase/pocketbase/tools/search"
_ "github.com/pocketbase/pocketbase/tools/security"
_ "github.com/pocketbase/pocketbase/tools/store"
_ "github.com/pocketbase/pocketbase/tools/subscriptions"
_ "github.com/pocketbase/pocketbase/tools/template"
_ "github.com/pocketbase/pocketbase/tools/tokenizer"
_ "github.com/pocketbase/pocketbase/tools/types"
_ "github.com/pocketbase/pocketbase/ui"
```
