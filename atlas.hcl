// Define an environment named "local"
env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["file://multi.hcl", "file://schema.hcl"].
  src = "file://schema.pg.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = "postgres://myuser:mypassword@localhost:5437/go_gin_start?sslmode=disable"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "postgres://myuser:mypassword@localhost:5437/go_gin_start_dev?sslmode=disable"
} 