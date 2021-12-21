api_service_image_version      = "sha256-fbc9e8a2c1bc"
postgres_service_image_version = "sha256-5da3a4985f5e"
redis_service_image_version    = "sha256-e760e8feac0b"
kafka_service_image_version    = "sha256-1164cad5aeb3"
rabbitmq_service_image_version = "sha256-e47433cf09b3"
neo4j_service_image_version    = "sha256-f7ca345ebd26"

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
