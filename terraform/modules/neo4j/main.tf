resource "kubernetes_namespace" "neo4j" {
  metadata {
    name = "neo4j-si"
  }
}

resource "kubernetes_deployment" "neo4j" {
  metadata {
    name = "neo4j"
    namespace = kubernetes_namespace.neo4j.metadata.0.name
  }
  
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "neo4j"
      }
    }
    template {
      metadata {
        labels = {
          app = "neo4j"
        }
      }
      spec {
        container {
          image = "neo4j:4.4.2"
          name  = "neo4j"
          port {
            container_port = 7687
          }

          env {
            name  = "NEO4J_AUTH"
            value = "neo4j/P@ssword!"
          }
        }
      }
    }
  }
}

# resource "helm_release" "neo4j" {
#   name       = "neo4j"
#   namespace  = kubernetes_namespace.neo4j.metadata.0.name

#   chart      = "${path.module}/chart/"
#   timeout    = 1200

#   set {
#     name  = "neo4jPassword"
#     value = "P@ssword!"
#   }
#   set {
#     name  = "acceptLicenseAgreement"
#     value = "yes"
#   }
#   set {
#     name  = "core.service.type"
#     value = "LoadBalancer"
#   }

#   set {
#     name  = "core.numberOfServers"
#     value = var.neo4j.core_replicas
#   }
#   set {
#     name  = "readReplica.numberOfServers"
#     value = var.neo4j.replica_replicas
#   }
# }

resource "kubernetes_service" "neo4j" {
  metadata {
    name      = "neo4j"
    namespace = kubernetes_namespace.neo4j.metadata.0.name
  }
  spec {
    selector = {
      app = "neo4j"
    }
    type = "LoadBalancer"

    port {
      port        = 7687
      target_port = 7687
      name        = "server"
    }
    port {
      port        = 7474
      target_port = 7474
      name        = "browser"
    }
  }
}
