variable "rabbitmq" {
  type = object({
    replicas = number
  })
}