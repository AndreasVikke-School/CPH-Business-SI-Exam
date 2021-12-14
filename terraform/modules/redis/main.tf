locals {
  # Directories start with "C:..." on Windows; All other OSs use "/" for root.
  is_windows = substr(pathexpand("~"), 0, 1) == "/" ? false : true
}

resource "kubernetes_namespace" "redis" {
  metadata {
    name = "redis-si"
  }
}

resource "kubernetes_config_map" "redis_cluster" {
  metadata {
    name      = "redis-cluster"
    namespace = kubernetes_namespace.redis.metadata.0.name
  }

  data = {
    "update-node.sh" = <<-EOF
        #!/bin/sh
        REDIS_NODES="/data/nodes.conf"
        sed -i -e "/myself/ s/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/$${POD_IP}/" $${REDIS_NODES}
        exec "$@"
    EOF
    "redis.conf"     = <<-EOF
        cluster-enabled yes
        cluster-require-full-coverage no
        cluster-node-timeout 5000
        cluster-config-file /data/nodes.conf
        cluster-migration-barrier 1
        appendonly yes
        protected-mode no
    EOF
  }
}

resource "kubernetes_stateful_set" "redis_cluster" {
  metadata {
    name      = "redis-cluster"
    namespace = kubernetes_namespace.redis.metadata.0.name
  }

  spec {
    replicas     = var.redis.replicas
    service_name = "redis-cluster"

    selector {
      match_labels = {
        app = "redis-cluster"
      }
    }

    template {
      metadata {
        labels = {
          app = "redis-cluster"
        }
      }
      spec {
        container {
          name    = "redis"
          image   = "redis:5.0.5-alpine"
          command = ["/conf/update-node.sh", "redis-server", "/conf/redis.conf"]

          port {
            container_port = 6379
            name           = "client"
          }
          port {
            container_port = 16379
            name           = "gossip"
          }

          env {
            name = "POD_IP"
            value_from {
              field_ref {
                field_path = "status.podIP"
              }
            }
          }

          volume_mount {
            name       = "conf"
            mount_path = "/conf"
            read_only  = false
          }
          volume_mount {
            name       = "data"
            mount_path = "/data"
            read_only  = false
          }
        }
        volume {
          name = "conf"

          config_map {
            name         = "redis-cluster"
            default_mode = "0755"
          }
        }
      }
    }

    volume_claim_template {
      metadata {
        name = "data"
      }

      spec {
        access_modes = ["ReadWriteOnce"]

        resources {
          requests = {
            storage = "1Gi"
          }
        }
      }
    }
  }

  provisioner "local-exec" {
    command     = "kubectl exec --namespace=${kubernetes_namespace.redis.metadata.0.name} -it redis-cluster-0 -- sh -c \"redis-cli --cluster create --cluster-replicas 1 $(kubectl get pods -l app=redis-cluster --namespace=${kubernetes_namespace.redis.metadata.0.name} -o jsonpath=\"{range.items[?(@.kind=='Pod')]}{.status.podIP}:6379 {end}\") --cluster-yes\" || echo Cluster Already SetUp"
    interpreter = local.is_windows ? ["pwsh", "-Command"] : []
  }

  depends_on = [
    kubernetes_config_map.redis_cluster
  ]
}

resource "kubernetes_service" "redis_cluster" {
  metadata {
    name      = "redis-cluster"
    namespace = kubernetes_namespace.redis.metadata.0.name
  }
  spec {
    selector = {
      app = "redis-cluster"
    }
    type = "ClusterIP"
    port {
      port = 6379
      name = "client"
    }
    port {
      port = 16379
      name = "gossip"
    }
  }
}
