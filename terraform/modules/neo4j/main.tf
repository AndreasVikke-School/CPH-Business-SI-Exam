resource "kubernetes_namespace" "neo4j" {
  metadata {
    name = "neo4j-si"
  }
}

resource "helm_release" "neo4j" {
  name       = "neo4j"
  namespace  = kubernetes_namespace.neo4j.metadata.0.name

  repository = "https://helm.neo4j.com/neo4j"
  chart      = "neo4j-standalone"
  timeout    = 900

  values = [
    "${file("${path.module}/custom-values.yaml")}"
  ]

  set {
    name  = "neo4jPassword"
    value = "P@ssword!"
  }
}

data "kubernetes_service" "neo4j" {
  metadata {
    name      = "neo4j-neo4j"
    namespace = kubernetes_namespace.neo4j.metadata.0.name
  }

  depends_on = [
    helm_release.neo4j
  ]
}


# kubectl run --rm -it --image "neo4j:4.3.2-enterprise" cypher-shell -- cypher-shell -a "neo4j://neo4j.neo4j.svc.cluster.local:7687" -u neo4j -p "P@ssword\!"