resource "kubernetes_namespace" "neo4j" {
  metadata {
    name = "neo4j-si"
  }
}

resource "helm_release" "neo4j" {
  name       = "neo4j"
  namespace  = kubernetes_namespace.neo4j.metadata.0.name

  chart      = "${path.module}/chart/"
  timeout    = 900

#   values = [
#     "${file("${path.module}/custom-values.yaml")}"
#   ]
  set {
    name  = "neo4jPassword"
    value = "P@ssword!"
  }
  set {
    name  = "acceptLicenseAgreement"
    value = "yes"
  }
  set {
    name  = "core.service.type"
    value = "LoadBalancer"
  }
}

# data "kubernetes_service" "neo4j" {
#   metadata {
#     name      = "neo4j-neo4j"
#     namespace = kubernetes_namespace.neo4j.metadata.0.name
#   }

#   depends_on = [
#     helm_release.neo4j
#   ]
# }


# kubectl run --rm -it --image "neo4j:4.3.2-enterprise" cypher-shell -- cypher-shell -a "neo4j://neo4j.neo4j-si.svc.cluster.local:7687" -u neo4j -p "P@ssword\!"