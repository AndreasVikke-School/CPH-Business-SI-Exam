output "service_ip" {
  value = var.service_type == "LoadBalancer" ? "${kubernetes_service.service.status.0.load_balancer.0.ingress.0.ip}:${kubernetes_service.service.spec.0.port.0.port}" : ""
}