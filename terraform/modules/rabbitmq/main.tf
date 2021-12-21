resource "kubernetes_namespace" "rabbitmq" {
  metadata {
    name = "rabbitmq-si"
  }
}


resource "kubernetes_deployment" "rabbitmq" {
  metadata {
    name = "rabbitmq"
    namespace = kubernetes_namespace.rabbitmq.metadata.0.name
  }
  
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "rabbitmq"
      }
    }
    template {
      metadata {
        labels = {
          app = "rabbitmq"
        }
      }
      spec {
        container {
          image = "rabbitmq:3-management-alpine"
          name  = "rabbitmq"
          port {
            container_port = 5672
          }

          env {
            name  = "RABBITMQ_DEFAULT_USER"
            value = "rabbitmq"
          }
          env {
            name  = "RABBITMQ_DEFAULT_PASS"
            value = "P@ssword!"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "rabbitmq" {
  metadata {
    name      = "rabbitmq"
    namespace = kubernetes_namespace.rabbitmq.metadata.0.name
  }
  spec {
    selector = {
      app = "rabbitmq"
    }
    type = "LoadBalancer"

    port {
      port        = 5672
      target_port = 5672
      name        = "server"
    }
    port {
      port        = 15672
      target_port = 15672
      name        = "browser"
    }
  }
}

# resource "helm_release" "rabbitmq" {
#   name       = "rabbitmq"
#   namespace  = kubernetes_namespace.rabbitmq.metadata.0.name

#   repository = "https://charts.bitnami.com/bitnami"
#   chart      = "rabbitmq"
#   timeout    = 1200

#   set {
#     name  = "service.type"
#     value = "LoadBalancer"
#   }
#   set {
#     name  = "replicaCount"
#     value = var.rabbitmq.replicas
#   }
#   set {
#     name  = "auth.username"
#     value = "rabbitmq"
#   }
#   set {
#     name  = "auth.password"
#     value = "P@ssword!"
#   }
# }

# data "kubernetes_service" "rabbitmq" {
#   metadata {
#     name      = "rabbitmq"
#     namespace = kubernetes_namespace.rabbitmq.metadata.0.name
#   }

#   depends_on = [
#     helm_release.rabbitmq
#   ]
# }