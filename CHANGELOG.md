## [unreleased]

### 🚀 Features

- Add local yolo mode with configurable exclusion paths
- Treat local yolo exclusion paths as extra trusted roots
- Add native engram memory and graph review builtins
- Add generic codegraph proxy builtin for the bcrg tool surface
- Merge codegraph MCP into a native CLI tool and rename bcrg proxy
- Replace codegraph tool with tokensave and add OpenCode provider support
- *(ui)* Show user@host:cwd in the header by default (#3583)
- Textarea selection (#3507)

### 🐛 Bug Fixes

- Separate esc to clear input from ctrl+esc to stop the task
- Recall the most recent prompt immediately instead of skipping it
- Keep bang shell commands recallable in prompt history
- Align mem scope between save/context and validate tool args
- Return full length from bounded writer to avoid short-write
- Clean codegraph tool dead code and friendly init error
- Register native codegraph/bcgraph/mem tools with agents
- *(bash)* Decide auto-approval from the parsed command, not its prefix
- *(bash)* Satisfy staticcheck QF1001 in isAssignment
- *(bash)* Keep read-only git filter and remote forms auto-approved
- *(ui)* Use dialog background for bash syntax highlighting in permissions (#3575)
- *(windows)* Scrolling up on modifier keys alone (#3598)

### 💼 Other

- Harden native engram/bcrg builtin tools
- Add timeout and output cap to bcrg subprocess calls
- *(bash)* Close four auto-approval gaps in the safe-command matcher
- Resolve conflicts from v0.91.0

### 📚 Documentation

- Add design for native engram + bcrg integration
- Trim session context (drop duplicate Beads block and CLAUDE.md)
- Streamline AGENTS.md and add development reference

### ⚡ Performance

- Resolve local yolo permission path only once per request
- Cut redundant whitespace-edit guidance from the system prompt

### 🧪 Testing

- Unset API keys during auth test

### ⚙️ Miscellaneous Tasks

- Remove .envrc file
- Set up Beads issue tracking and agent integration
- Ignore codegraph daemon pid file
- Update Beads integration hash
- Drop git fetch from install and release tasks
- Run dependabot monthly
- *(legal)* @teddytennant has signed the CLA
- *(legal)* @Lin-BoYuan has signed the CLA
- *(legal)* @lawrence3699 has signed the CLA
- Add clean and rebuild task targets for full recompiles
- Gofumpt
- Auto-update files
- Auto-update files
- Normalize line endings to `lf` via `.gitattributes` (#3555)
- *(legal)* @marwan562 has signed the CLA
- Re-record vcr cassettes
- Standardize go toolchain on v1.26.6 (#3592)
- *(legal)* @wangjiawen2013 has signed the CLA
- Auto-update files
- *(legal)* @Nayjest has signed the CLA
- *(legal)* @torrmal has signed the CLA
- *(legal)* @LinespottingPrivate has signed the CLA
- Update dependencies and model configs for deepseek-v4
- Update beads docs and suppress noisy git tag stderr
- Ignore codegraph lock file
