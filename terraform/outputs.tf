output "api_ip" {
  value = module.api_service.service_ip
}
output "pgadmin_ip" {
  value = module.postgres_module.pgadmin_ip
}
output "kafdrop_ip" {
  value = module.kafka_module.kafdrop_ip
}