api_service_image_version      = "sha256-1d2362b68ae4"
postgres_service_image_version = "sha256-5da3a4985f5e"
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
