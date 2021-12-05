output "rabbitmq_ip" {
    value = "${data.kubernetes_service.rabbitmq.status.0.load_balancer.0.ingress.0.ip}:${data.kubernetes_service.rabbitmq.spec.0.port.0.port}"
}
output "rabbitmq_web_ip" {
    value = "${data.kubernetes_service.rabbitmq.status.0.load_balancer.0.ingress.0.ip}:15672"
}