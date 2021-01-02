---
title: Automatic TLS
subtitle: Go from HTTP to HTTPS in minutes and let Response Server manage, renew, and sync your certificates for you.
menuTitle: Automatic TLS
category: Guides
position: 1000
---

<alert type="info">

Heads up! You've stumbled upon a really cool feature and we're eager for everyone to use it. If you run into any issues or have questions about this guide please get in touch on Discord!

</alert>

## Prerequisites

Before continuing with the guide, please ensure your Response Server meets the following criteria. **If you're using Response Cloud you do not have to follow this guide, we've already taken care of configuring TLS for you!**

1. Response Server is routable on the public internet.
2. You have DNS records configured for the domain or sub-domain Response Server should be accessible from.

Not sure you meet these criteria? Reach out on Discord for help.

<alert type="warning">

If you're accessing Response from an IP address (like 127.0.0.1) or localhost your not quite ready for this guide. Be sure your Response Server is hosted on a VPS or dedicated server and that you can access it from the domain or sub-domain you intend for Response to be accessed from (like response.example.org).

</alert>
