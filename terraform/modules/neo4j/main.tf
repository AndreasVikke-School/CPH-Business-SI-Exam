resource "kubernetes_namespace" "neo4j" {
  metadata {
    name = "neo4j-si"
  }
}

resource "helm_release" "neo4j" {
  name       = "neo4j"
  namespace  = kubernetes_namespace.neo4j.metadata.0.name

  chart      = "${path.module}/chart/"
  timeout    = 1200

  set {
    name  = "neo4jPassword"
    value = "P@ssword!"
  }
  set {
    name  = "acceptLicenseAgreement"
    value = "yes"
  }
  set {
    name  = "core.service.type"
    value = "LoadBalancer"
  }

  set {
    name  = "core.numberOfServers"
    value = var.neo4j.core_replicas
  }
  set {
    name  = "readReplica.numberOfServers"
    value = var.neo4j.replica_replicas
  }
}

data "kubernetes_service" "neo4j" {
  metadata {
    name      = "neo4j-neo4j"
    namespace = kubernetes_namespace.neo4j.metadata.0.name
  }

  depends_on = [
    helm_release.neo4j
  ]
}
