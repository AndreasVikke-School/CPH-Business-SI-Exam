locals {
  broker_env = {
    "BROKER_ID_COMMAND"                      = "hostname | awk -F'-' '{print $2}'"
    "KAFKA_ZOOKEEPER_CONNECT"                = "zookeeper:2181"
    "KAFKA_LISTENERS"                        = "INTERNAL://:9092,EXTERNAL://:9094"
    "KAFKA_ADVERTISED_LISTENERS"             = "INTERNAL://:9092,EXTERNAL://_{HOSTNAME_COMMAND}:9094"
    "KAFKA_LISTENER_SECURITY_PROTOCOL_MAP"   = "INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT"
    "KAFKA_INTER_BROKER_LISTENER_NAME"       = "INTERNAL"
    "KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR" = "1"
    "KAFKA_PORT"                             = "9092"
    "KAFKA_AUTO_CREATE_TOPICS_ENABLE"        = "true"
  }
}

resource "kubernetes_stateful_set" "kafka" {
  metadata {
    name      = "kafka"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }

  spec {
    replicas     = var.kafka.replicas
    service_name = "kafka"

    selector {
      match_labels = {
        app = "kafka"
      }
    }

    template {
      metadata {
        labels = {
          app = "kafka"
        }
      }
      spec {
        container {
          name    = "kafka"
          image   = "wurstmeister/kafka:2.11-1.1.0"
          port {
            container_port = 9094
            name           = "external"
          }
          port {
            container_port = 9092
            name           = "internal"
          }

          env {
            name  = "HOSTNAME_COMMAND"
            value = "echo ${kubernetes_service.kafka.status.0.load_balancer.0.ingress.0.ip}"
          }
          dynamic "env" {
            for_each = local.broker_env
            content {
              name  = env.key
              value = env.value
            }
          }
        }
      }
    }
  }

  depends_on = [
    kubernetes_service.kafka,
    kubernetes_deployment.zookeeper
  ]
}

resource "kubernetes_service" "kafka" {
  metadata {
    name      = "kafka"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "kafka"
    }
    type = "LoadBalancer"
    port {
      name = "internal"
      port = 9092
    }
    port {
      name = "external"
      port = 9094
    }
  }
}

resource "kubernetes_service" "zookeeper" {
  metadata {
    name      = "zookeeper"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "zookeeper"
    }
    type = "ClusterIP"
    port {
      name = "client"
      port = 2181
    }
    port {
      name = "server"
      port = 2888
    }
    port {
      name = "leader-election"
      port = 3888
    }
  }
}


resource "kubernetes_namespace" "kafka" {
  metadata {
    name = "kafka-si"
  }
}

# # ==== KAFKA ZOOKEEPER ====
resource "kubernetes_deployment" "zookeeper" {
  metadata {
    name      = "zookeeper"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "zookeeper"
      }
    }
    template {
      metadata {
        labels = {
          app = "zookeeper"
        }
      }
      spec {
        container {
          image = "zookeeper:3.4.12"
          name  = "zookeeper"
          command = [
            "/bin/sh",
            "-c",
            "export ZOO_MY_ID=$(expr $(hostname | grep -o \"[[:digit:]]*$\") + 1) && /docker-entrypoint.sh zkServer.sh start-foreground"
          ]
          port {
            container_port = 2181
            name           = "client"
          }
          port {
            container_port = 2888
            name           = "server"
          }
          port {
            container_port = 3888
            name           = "leader-election"
          }

          env {
            name  = "ZOO_SERVERS"
            value = "server.1=zookeeper.kafka:2888:3888"
          }
        }
      }
    }
  }
}

# ==== KAFDROP ====
resource "kubernetes_deployment" "kafka_kafdrop" {
  metadata {
    name      = "kafka-kafdrop"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "kafka-kafdrop"
      }
    }
    template {
      metadata {
        labels = {
          app = "kafka-kafdrop"
        }
      }
      spec {
        container {
          image = "obsidiandynamics/kafdrop"
          name  = "kafka-broker"
          port {
            container_port = 9000
          }
          env {
            name  = "KAFKA_BROKERCONNECT"
            value = "kafka:9092"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "kafka_kafdrop" {
  metadata {
    name      = "kafka-kafdrop"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "kafka-kafdrop"
    }
    type = "LoadBalancer"
    port {
      name = "kafdrop-port"
      port = 9000
    }
  }
}
# ==== KAFDROP END ====