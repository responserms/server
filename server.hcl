encryption_key = env("ENCRYPTION_KEY", "none")

events {
  type = "nats"
  url = "nats://host.docker.internal:4222"
}

// cluster {
//   member_port = 3320
//   discovery_port = 3322

//   autojoin "nats" {
//     url = "nats://host.docker.internal:4222"
//     subject = "response-cluster"
//   }

//   autojoin "cloud" {
//     provider = "digitalocean"
//     args = {
//       region = "nyc3"
//       tag_name = "response-server"
//       api_token = env("DIGITALOCEAN_API_TOKEN")
//     }
//   }
// }

developer {
  profiling = true
}