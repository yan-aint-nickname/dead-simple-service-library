# Dead Simple Service Library

TODO:

- components:
  - redis (opt-out)
  - postgres (opt-out)
    - driver pgx/pgxpool
    - query builder ?? buildsqlx/sqlx (opt-out) ?? I prefer raw sql
    - migration managment ?? (opt-out)
  - kafka (opt-out)
  - mongodb (opt-out)
  - [x] http framework - gin (obligatory) https://github.com/gin-gonic/gin ~> feature-rich + easy to setup
    - [x] use viper as config manager
    - [x] sentry integration
    - [ ] logrus
    - [ ] dbpool integration
    - [ ] kafka integration (pub)
    - [ ] gql/apollo client/server
    - [ ] more examples?
  - worker/job scheduler (opt-in/opt-out?) https://github.com/go-co-op/gocron?tab=readme-ov-file
    - [ ] db integration
    - [ ] kafka integration (pub/sub)
    - [ ] gql client
  - graphql ?? (opt-in) https://github.com/graph-gophers/graphql-go

- telemetry:
  - opentelemetry (obligatory) https://opentelemetry.io/docs/languages/go/getting-started/
  - [x] sentry-sdk (obligatory) https://github.com/getsentry/sentry-go

- json (std encoding/json)

- logging https://github.com/sirupsen/logrus ~> loguru-like + std compatible

- static analysis:
  - golangci-lint (obligatory/opt-out)
  - gocritic (obligatory/opt-out)
  - gosec

- test
  - go test -v -timeout 30s -coverprofile=cover.out -cover (obligatory/opt-out)
  - go tool cover -func=cover.out (obligatory/opt-out)

- pre-commit:
  - pre-commit (global/opt-out)

- local development:
  - replace github.com/user/library => /local/path/library (replace directive works)
- styleguide: https://github.com/uber-go/guide/blob/master/style.md
