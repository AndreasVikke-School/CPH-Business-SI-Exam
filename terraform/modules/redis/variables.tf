variable "redis" {
  type = object({
    replicas = number
  })
}