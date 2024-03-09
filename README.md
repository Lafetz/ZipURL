# ZipUrl

Let's face it, creating URL shorteners is about as common as cat videos on the internet.This isn't just about shortening URLs; it's about building a robust and maintainable system. The hexagonal architecture keeps the core logic of URL shortening clean and independent, surrounded by adaptable "adapters" that handle external concerns like databases and user interfaces. 

This means:
   
• Easy Testing: Isolate the core logic for unit tests, ensuring the heart of the system works flawlessly.
   
• Future-Proof Flexibility: Need to switch databases or implement a new API? No problem! Just swap out the relevant adapter without touching the core logic.
   
• Crystal-Clear Code: The hexagonal architecture promotes modularity, making the codebase easier to understand and maintain.

## Built with
- Go
- Gin
- Redis
- PostgreSQL

## ☑ Tasks
- [x] Jwt
- [x] Unit tests
- [ ] Integration tests
- [ ] Caching
- [ ] Logging
