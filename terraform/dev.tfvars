api_service_image_version      = "sha256-ac6be4a0ef4b"
postgres_service_image_version = "sha256-c7bfce869745"
redis_service_image_version    = "sha256-e760e8feac0b"
kafka_service_image_version    = "sha256-1164cad5aeb3"
rabbitmq_service_image_version = "sha256-de3ca0bee0cf"


neo4j = {
    core_replicas    = 3,
    replica_replicas = 1
}
postgres = {
    replicas = 1
}
rabbitmq = {
    replicas = 1
}
redis = {
    replicas: 6
}
kafka = {
    replicas = 1
}
