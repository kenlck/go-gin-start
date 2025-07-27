// Define an environment named "local"
variable "envfile" {
    type    = string
    default = ".env"
}

locals {
    envfile = {
        for line in split("\n", file(var.envfile)): split("=", line)[0] => regex("=(.*)", line)[0]
        if !startswith(line, "#") && length(split("=", line)) > 1
    }
}

env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["file://multi.hcl", "file://schema.hcl"].
  src = "file://schema.pg.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = local.envfile["DATABASE_URL"]

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = local.envfile["DEV_DATABASE_URL"]
} 