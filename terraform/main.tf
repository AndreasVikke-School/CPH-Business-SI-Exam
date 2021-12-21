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
  kafka  = var.kafka
}
module "postgres_module" {
  source   = "./modules/postgres"
  postgres = var.postgres
}
module "redis_module" {
  source = "./modules/redis"
  redis  = var.redis
}
module "rabbitmq_module" {
  source   = "./modules/rabbitmq"
  rabbitmq = var.rabbitmq
}
module "neo4j_module" {
  source = "./modules/neo4j"
  neo4j  = var.neo4j
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

module "postgres_service" {
  source = "./modules/service"

  name_prefix            = "postgres-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "postgres_service"
  image_version          = var.postgres_service_image_version
  container_port         = 50051
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 50051,
      target_port = 50051
    }
  }
}

module "redis_service" {
  source = "./modules/service"

  name_prefix            = "redis-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "redis_service"
  image_version          = var.redis_service_image_version
  container_port         = 50051
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 50051,
      target_port = 50051
    }
  }
}

module "kafka_service" {
  source = "./modules/service"

  name_prefix            = "kafka-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "kafka_service"
  image_version          = var.kafka_service_image_version
  container_port         = 5000
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 5000,
      target_port = 5000
    }
  }

  depends_on = [
    module.kafka_module
  ]
}

module "rabbitmq_service" {
  source = "./modules/service"

  name_prefix            = "rabbitmq-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "rabbitmq_service"
  image_version          = var.rabbitmq_service_image_version
  container_port         = 50001
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 50001,
      target_port = 50001
    }
  }
}

module "neo4j_service" {
  source = "./modules/service"

  name_prefix            = "neo4j-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "neo4j_service"
  image_version          = var.neo4j_service_image_version
  container_port         = 50051
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 50051,
      target_port = 50051
    }
  }
}
