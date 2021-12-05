provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

provider "kubectl" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

provider "helm" {
  kubernetes {
    config_path    = "~/.kube/config"
    config_context = "minikube"
  }
}

module "kafka_module" {
  source = "./modules/kafka"
}
module "postgres_module" {
  source = "./modules/postgres"
}
module "redis_module" {
  source = "./modules/redis"
}
module "rabbitmq_module" {
  source = "./modules/rabbitmq"
}

resource "kubernetes_namespace" "services" {
  metadata {
    name = "services-si"
  }
}

module "api_service" {
  source = "./modules/service"

  name_prefix            = "api-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "api_service"
  image_version          = var.api_service_image_version
  container_port         = 8081
  container_replications = 2
  service_type           = "LoadBalancer"
  service_ports = {
    server = {
      port        = 8080,
      target_port = 8081
    }
  }
}
