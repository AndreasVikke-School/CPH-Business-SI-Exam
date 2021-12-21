output "rabbitmq_ip" {
    value = "${kubernetes_service.rabbitmq.status.0.load_balancer.0.ingress.0.ip}:${kubernetes_service.rabbitmq.spec.0.port.0.port}"
}
output "rabbitmq_web_ip" {
    value = "${kubernetes_service.rabbitmq.status.0.load_balancer.0.ingress.0.ip}:${kubernetes_service.rabbitmq.spec.0.port.1.port}"
}