resource "kubernetes_namespace" "camunda" {
  metadata {
    name = "camunda-si"
  }
}

# resource "kubernetes_config_map" "camunda" {
#   metadata {
#     name      = "camunda"
#     namespace = kubernetes_namespace.camunda.metadata.0.name
#   }

#   data = {
#     "test.form"  = <<-EOF
#     {
#       "schemaVersion": 2,
#       "components": [
#         {
#           "label": "Buy Amout",
#           "type": "number",
#           "id": "Field_1v9v7fy",
#           "key": "buyAmountField",
#           "description": "The amount of items to buy",
#           "validate": {
#             "required": true,
#             "min": 1,
#             "max": 10
#           }
#         }
#       ],
#       "type": "default",
#       "id": "buyForm"
#     }
#     EOF
#     # "bookContract_form.form"  = file("../applications/bpmn/bookContract_form.form")
#     # "buyItem_form.form"       = file("../applications/bpmn/buyItem_form.form")
#     # "restock_model.bpmn"      = file("../applications/bpmn/restock_model.bpmn")
#     # "restock_rules.dmn"       = file("../applications/bpmn/restock_rules.dmn")
#     # "vinylContract_form.form" = file("../applications/bpmn/vinylContract_form.form")
#   }
# }

resource "kubernetes_deployment" "camunda" {
  metadata {
    name      = "camunda"
    namespace = kubernetes_namespace.camunda.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "camunda"
      }
    }
    template {
      metadata {
        labels = {
          app = "camunda"
        }
      }
      spec {
        container {
          image = "camunda/camunda-bpm-platform:run-latest"
          name  = "camunda"
          port {
            container_port = 8080
          }

          # volume_mount {
          #   name       = "bpmn"
          #   mount_path = "camunda/configuration/resources"
          #   read_only  = false
          # }
        }
        # volume {
        #   name = "bpmn"

        #   config_map {
        #     name         = "camunda"
        #     default_mode = "0755"
        #   }
        # }
      }
    }
  }
}

resource "kubernetes_service" "camunda" {
  metadata {
    name      = "camunda"
    namespace = kubernetes_namespace.camunda.metadata.0.name
  }
  spec {
    selector = {
      app = "camunda"
    }
    type = "LoadBalancer"
    port {
      port = 8082
      target_port = 8080
    }
  }
}
