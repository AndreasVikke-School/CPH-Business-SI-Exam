resource "kubernetes_deployment" "service" {
  metadata {
    name = "${var.name_prefix}service"
    # namespace = kubernetes_namespace.test.metadata.0.name
    namespace = var.namespace
  }
  spec {
    replicas = var.container_replications
    selector {
      match_labels = {
        app = "${var.name_prefix}service"
      }
    }
    template {
      metadata {
        labels = {
          app = "${var.name_prefix}service"
        }
      }
      spec {
        container {
          image = "ghcr.io/andreasvikke/cph-business-ls-exam/${var.image_name}:${var.image_version}"
          name  = "${var.name_prefix}service-container"
          port {
            container_port = var.container_port
          }

          dynamic "env" {
            for_each = var.container_env
            content {
              name  = env.key
              value = env.value
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "service" {
  metadata {
    name      = "${var.name_prefix}service"
    namespace = var.namespace
  }
  spec {
    selector = {
      app = "${var.name_prefix}service"
    }
    type = var.service_type

    dynamic "port" {
      for_each = var.service_ports

      content {
        port        = port.value.port
        target_port = port.value.target_port
        name        = port.key
      }
    }
  }
}
