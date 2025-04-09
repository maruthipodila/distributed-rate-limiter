# Distributed Rate Limiter

A cloud-native, distributed rate limiting service designed to manage API or service traffic across horizontally scaled systems.

## 🔥 Why This Project?

In modern architectures, services are often replicated across multiple nodes. To enforce consistent usage limits, a **distributed rate limiter** is essential.

## 📐 Architecture

- Multiple app nodes or services connect to a central **Redis** instance
- Redis acts as the **coordinator** for rate limits
- Supports algorithms like:
  - Token Bucket
  - Leaky Bucket
  - Sliding Window Log / Counter

![Architecture Diagram](./assets/distributed_rate_limiter_design.pdf)

## ✨ Features

- ✅ Global rate limit enforcement across distributed instances
- 🔄 Multiple algorithms supported
- 🚀 Pluggable backends (Redis, etcd planned)
- 📊 Prometheus metrics
- 🌐 REST and gRPC interfaces
- 🛠️ Configurable per-user, per-IP, per-endpoint limits

## 🚧 Roadmap

- [x] Redis-backed token bucket implementation
- [ ] gRPC API
- [ ] Admin dashboard
- [ ] Pluggable persistence (etcd, Postgres)
- [ ] Advanced rules (geo/tenant-specific throttling)

## 💡 Use Cases

- API Gateway throttling
- User/IP based rate limits
- SaaS multi-tenant abuse prevention

## 🛠️ Tech Stack

- Language: Go (Golang)
- Storage: Redis
- APIs: REST (Gin), gRPC (soon)
- Monitoring: Prometheus, Grafana (planned)

---

## 🤝 Contributing

We welcome contributions! Whether it's a bug fix, feature idea, or just improving docs — PRs are open.

## 📄 License
