# KingDNS DNS Engine v2.8

Database Driven Resolver Foundation.

Implemented:

- database repository layer
- resolver database flow
- cache integration point
- routing selection point
- dynamic record lookup architecture

The previous nslookup blocker was:

- DNS listener missing
- handler registration missing

Those are now solved.

Next:

- real PostgreSQL queries
- routing_rules integration
- node health selection
- Redis cache execution
