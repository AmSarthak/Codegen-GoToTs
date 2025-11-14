A tiny weekend project that parses Go structs and generates TypeScript interfaces automatically.
Zero dependencies. Pure Go.
Useful when your backend structs should match frontend TS models without copying manually.

🚀 Features

🔍 Parses Go structs using go/ast

🔄 Converts Go types → TypeScript types using simple mapping rules

📄 Generates .ts files with interfaces (not classes)

🧩 Supports:

basic types: string, int, float64, bool

slices: []string, []int, etc.