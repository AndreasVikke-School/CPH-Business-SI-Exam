output "camunda_ip" {
    value = "${kubernetes_service.camunda.status.0.load_balancer.0.ingress.0.ip}:8082"
}