api_service_image_version      = "sha256-eb6b5ce43531"
postgres_service_image_version = "sha256-5bc7fef84c4f"
redis_service_image_version    = "sha256-e760e8feac0b"
kafka_service_image_version    = "sha256-1164cad5aeb3"
rabbitmq_service_image_version = "sha256-87d93d716bbd"
neo4j_service_image_version    = "sha256-84e34fab8804"

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
    replicas = 6
}
kafka = {
    replicas = 1
}
