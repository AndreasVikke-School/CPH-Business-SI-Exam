variable "neo4j" {
  type = object({
    core_replicas    = number,
    replica_replicas = number
  })
}