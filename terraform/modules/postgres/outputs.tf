output "pgadmin_ip" {
  value = "${kubernetes_service.pgadmin.status.0.load_balancer.0.ingress.0.ip}:${kubernetes_service.pgadmin.spec.0.port.0.port}"
}