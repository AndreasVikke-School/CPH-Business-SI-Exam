variable "api_service_image_version" {
  type    = string
  default = "latest"
}

variable "postgres_service_image_version" {
  type    = string
  default = "latest"
}

variable "redis_service_image_version" {
  type    = string
  default = "latest"
}

variable "kafka_service_image_version" {
  type    = string
  default = "latest"
}

variable "neo4j_service_image_version" {
  type    = string
  default = "latest"
}

variable "neo4j" {
  type = object({
    core_replicas    = number,
    replica_replicas = number
  })
  default = {
    core_replicas = 3
    replica_replicas = 1
  }
}
variable "postgres" {
  type = object({
    replicas = number
  })
  default = {
    replicas = 1
  }
}
variable "rabbitmq" {
  type = object({
    replicas = number
  })
  default = {
    replicas = 1
  }
}
variable "redis" {
  type = object({
    replicas = number
  })
  default = {
    replicas = 1
  }
}
variable "kafka" {
  type = object({
    replicas = number
  })
  default = {
    replicas = 1
  }
}