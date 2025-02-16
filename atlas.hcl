env "local" {
  url = "sqlite://local.db"
}

env "prod" {
  url = "libsql://your-db-url"
} 