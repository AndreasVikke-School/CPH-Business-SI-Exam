resource "kubernetes_namespace" "rabbitmq" {
  metadata {
    name = "rabbitmq-si"
  }
}

resource "helm_release" "rabbitmq" {
  name       = "rabbitmq"
  namespace  = kubernetes_namespace.rabbitmq.metadata.0.name

  repository = "https://charts.bitnami.com/bitnami"
  chart      = "rabbitmq"
  timeout    = 900

  set {
    name  = "service.type"
    value = "LoadBalancer"
  }
  set {
    name  = "replicaCount"
    value = var.rabbitmq.replicas
  }
  set {
    name  = "auth.username"
    value = "rabbitmq"
  }
  set {
    name  = "auth.password"
    value = "P@ssword!"
  }
}

data "kubernetes_service" "rabbitmq" {
  metadata {
    name      = "rabbitmq"
    namespace = kubernetes_namespace.rabbitmq.metadata.0.name
  }

  depends_on = [
    helm_release.rabbitmq
  ]
}