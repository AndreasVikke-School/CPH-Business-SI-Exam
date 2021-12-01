output "kafdrop_ip" {
  value = "${kubernetes_service.kafka_kafdrop.status.0.load_balancer.0.ingress.0.ip}:${kubernetes_service.kafka_kafdrop.spec.0.port.0.port}"
}