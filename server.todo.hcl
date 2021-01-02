encryption_key = "/60w9G2zIfa6d47evaCcxl9E1L/k7OwWhzyMLA+YfCA="

variables {
  nats_url = "nats://host.docker.internal:4222"
}

events {
  type = "nats"
  url = nats_url
}

cluster {
  member_port = 3320
  discovery_port = 3322

  // members = [
  //   "10.3.0.2:3322"
  // ]

  autojoin "nats" {
    url = nats_url
    subject = "response-cluster"
  }

  // autojoin "cloud" {
  //   provider = "digitalocean"
  //   args = {
  //     region = "nyc3"
  //     tag_name = "response-server"
  //     api_token = env("DIGITALOCEAN_API_TOKEN")
  //   }
  // }
}

developer {
  profiling = true
}