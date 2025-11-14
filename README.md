Go Struct → TypeScript Interfaces + REST API Backend

✅ 1. Code Generator

Automatically parses Go structs using go/ast and generates matching TypeScript interfaces used by the frontend.

✅ 2. REST API Backend (Go)

A small API that serves inventory data (fake 60-server dataset), supports:

Filtering (e.g., GPU model)
/getServerByGPU?gpu=NVIDIA

Pagination

/getAIServers?limit=20

Utility endpoints for inventory queries

✅ 3. Fake Inventory Data

A seeded JSON dataset that simulates servers, CPU/GPU models, manufacturers, etc.

Stored in:

backend/data/inventory.json

✅ 4. Upcoming

React and Typescript based UI for inventory management.
