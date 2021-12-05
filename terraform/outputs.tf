output "api_ip" {
  value = module.api_service.service_ip
}
output "pgadmin_ip" {
  value = module.postgres_module.pgadmin_ip
}
output "kafdrop_ip" {
  value = module.kafka_module.kafdrop_ip
}
output "rabbitmq_ip" {
  value = module.rabbitmq_module.rabbitmq_ip
}
output "rabbitmq_web_ip" {
  value = module.rabbitmq_module.rabbitmq_web_ip
}