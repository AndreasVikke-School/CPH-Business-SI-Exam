resource "kubernetes_namespace" "postgres" {
  metadata {
    name = "postgres-si"
  }
}

resource "helm_release" "postgres" {
  name       = "postgres"
  namespace  = kubernetes_namespace.postgres.metadata.0.name

  repository = "https://charts.bitnami.com/bitnami"
  chart      = "postgresql-ha"
  timeout    = 900

  set {
    name  = "postgresql.password"
    value = "P@ssword!"
  }
   set {
    name  = "postgresql.repmgrPassword"
    value = "P@ssword!"
  }
  set {
    name  = "postgresql.replicaCount"
    value = var.postgres.replicas
  }
}

# ==== POSTGRES SERVICE ====
# resource "kubernetes_deployment" "postgres" {
#   metadata {
#     name      = "postgres"
#     namespace = kubernetes_namespace.postgres.metadata.0.name
#   }
#   spec {
#     replicas = 1
#     selector {
#       match_labels = {
#         app = "postgres"
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = "postgres"
#         }
#       }
#       spec {
#         container {
#           image = "postgres:10.5"
#           name  = "postgres"
#           port {
#             container_port = 5432
#           }
#           env {
#             name  = "POSTGRES_USER"
#             value = var.postgress_username
#           }
#           env {
#             name  = "POSTGRES_PASSWORD"
#             value = var.postgress_password
#           }
#         }
#       }
#     }
#   }
# }

# resource "kubernetes_service" "postgres" {
#   metadata {
#     name      = "postgres"
#     namespace = kubernetes_namespace.postgres.metadata.0.name
#   }
#   spec {
#     selector = {
#       app = "postgres"
#     }
#     type = "ClusterIP"
#     port {
#       port = 5432
#     }
#   }
# }
# ==== POSTGRES SERVICE END ====

# ==== PGADMIN SERVICE ====
resource "kubernetes_deployment" "pgadmin" {
  metadata {
    name      = "pgadmin"
    namespace = kubernetes_namespace.postgres.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "pgadmin"
      }
    }
    template {
      metadata {
        labels = {
          app = "pgadmin"
        }
      }
      spec {
        container {
          image = "dpage/pgadmin4"
          name  = "postgres"
          port {
            container_port = 80
          }
          env {
            name  = "PGADMIN_DEFAULT_EMAIL"
            value = var.postgress_email
          }
          env {
            name  = "PGADMIN_DEFAULT_PASSWORD"
            value = var.postgress_password
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "pgadmin" {
  metadata {
    name      = "pgadmin"
    namespace = kubernetes_namespace.postgres.metadata.0.name
  }
  spec {
    selector = {
      app = "pgadmin"
    }
    type = "LoadBalancer"
    port {
      port        = 5000
      target_port = 80
    }
  }
}
# ==== PGADMIN END ====
