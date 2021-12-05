variable "name_prefix" {
  type = string
}

variable "namespace" {
  type    = string
  default = "default"
}

variable "image_name" {
  type = string
}

variable "image_version" {
  type    = string
  default = "latest"
}

variable "container_port" {
  type    = string
  default = "latest"
}

variable "container_replications" {
  type    = number
  default = 80
}

variable "service_type" {
  type    = string
  default = "ClusterIP"
}

variable "service_ports" {
  type = map(object({
    port        = string
    target_port = string
  }))
}

variable "container_env" {
  type = map(string)
  default = {}
}