output "neo4j_web_ip" {
    value = "${data.kubernetes_service.neo4j.status.0.load_balancer.0.ingress.0.ip}:7474"
}