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
- Route CLI output through rtk to compress tokens when available
- Strengthen system prompt file-read and search tool routing
- Mark view as last-resort reader and drop Claude config
- Add ctrl+/ editor keys help dialog

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
- Change select-all keybinding from ctrl+shift+a to ctrl+a
- Remove vim-style single-letter keybindings from chat view
- Clarify that file reads use view, not read; skip VCR tests without API key
- Use keyed literals for HelpRow to satisfy go vet
- Clean rebase conflict leftovers in go.sum and editor keys

### 💼 Other

- Harden native engram/bcrg builtin tools
- Add timeout and output cap to bcrg subprocess calls
- *(bash)* Close four auto-approval gaps in the safe-command matcher
- Tighten coder system prompt ~30% while preserving all rules

### 📚 Documentation

- Add design for native engram + bcrg integration
- Trim session context (drop duplicate Beads block and CLAUDE.md)
- Streamline AGENTS.md and add development reference

### ⚡ Performance

- Resolve local yolo permission path only once per request
- Cut redundant whitespace-edit guidance from the system prompt

### 🎨 Styling

- Apply gofumpt formatting (final newlines, alignment)

### 🧪 Testing

- Cover ctrl+/ editor help dialog open, close, and key isolation
- Tighten editor help dialog tests per review

### ⚙️ Miscellaneous Tasks

- Remove .envrc file
- Set up Beads issue tracking and agent integration
- Ignore codegraph daemon pid file
- Update Beads integration hash
- Drop git fetch from install and release tasks
- Add clean and rebuild task targets for full recompiles
- Auto-update files
- Auto-update files
- *(legal)* @marwan562 has signed the CLA
- Auto-update files
- Update dependencies and model configs for deepseek-v4
- Update beads docs and suppress noisy git tag stderr
- Ignore codegraph lock file
- Add git-cliff config for changelog generation
## [0.91.1] - 2026-08-25

### 🐛 Bug Fixes

- *(mcp)* Repair stdio startup diagnostics (#3634)
- Return to the start of the line (#3635)
- Wrapped messages counted as newlines (#3627)

### 📚 Documentation

- Add MindsHub custom provider example (#3618)
- Add documentation on how to override catwalk URL (#3585)

### ⚙️ Miscellaneous Tasks

- Bump golangci version to 2.13.1 (#3628)
- Bump bubbles (#3636)
- *(legal)* @Nikhils-G has signed the CLA
## [nightly] - 2026-08-24

### ⚙️ Miscellaneous Tasks

- *(legal)* @Nayjest has signed the CLA
- Merge main after v0.91.0
- *(legal)* @torrmal has signed the CLA
- *(legal)* @LinespottingPrivate has signed the CLA
- *(legal)* @4RH1T3CT0R7 has signed the CLA
## [0.91.0] - 2026-08-22

### 🚀 Features

- Textarea selection (#3507)

### 🐛 Bug Fixes

- *(windows)* Scrolling up on modifier keys alone (#3598)

### 🧪 Testing

- Unset API keys during auth test

### ⚙️ Miscellaneous Tasks

- Standardize go toolchain on v1.26.6 (#3592)
- *(legal)* @wangjiawen2013 has signed the CLA
- Auto-update files
## [0.90.0] - 2026-08-19

### 🚀 Features

- *(ui)* Show user@host:cwd in the header by default (#3583)

### 🐛 Bug Fixes

- *(bash)* Decide auto-approval from the parsed command, not its prefix
- *(bash)* Satisfy staticcheck QF1001 in isAssignment
- *(bash)* Keep read-only git filter and remote forms auto-approved
- *(ui)* Use dialog background for bash syntax highlighting in permissions (#3575)

### 💼 Other

- *(bash)* Close four auto-approval gaps in the safe-command matcher

### ⚙️ Miscellaneous Tasks

- Run dependabot monthly
- *(legal)* @teddytennant has signed the CLA
- *(legal)* @Lin-BoYuan has signed the CLA
- *(legal)* @lawrence3699 has signed the CLA
- Gofumpt
- Auto-update files
- Auto-update files
- Normalize line endings to `lf` via `.gitattributes` (#3555)
- *(legal)* @marwan562 has signed the CLA
- Re-record vcr cassettes
## [0.89.0] - 2026-08-12

### 🚀 Features

- *(ux)* Add bash syntax highlighting to bash tool
- *(ux)* Strip redundant cd-to-project prefix from bash tool display
- Use anthropic sdk from upstream (#3544)
- *(session)* Restore last used provider/model from session
- *(session)* Print session id on exit (#3398)
- Ctrl+end to go to the bottom and follow (#3535)

### 🐛 Bug Fixes

- *(ui)* Don't block interactive prompts on MCP initialization
- *(mcp)* Only wait for slow MCP servers in non-interactive runs
- Send hyper api key when fetching provider catalog (#3546)

### 💼 Other

- Differentiate invocation types (#3533)

### ⚙️ Miscellaneous Tasks

- *(legal)* @olifarhaan has signed the CLA
- *(legal)* @ach3rry has signed the CLA
- *(legal)* @ktsoator has signed the CLA
- *(mcp)* Log how long each MCP server takes to connect
- Auto-update files
## [0.88.1] - 2026-08-07

### 🚀 Features

- *(mcp)* Don't hold callback port open permanently (#3481)

### 🐛 Bug Fixes

- *(dialog)* Add clear alias to summarize (#3464)
- *(lsp)* Handle window/workDoneProgress/create to prevent server crash (#3445)
- *(providers)* Swap the provider cache instead of rewriting it
- *(hyper)* Keep working cache when provider cache can't be written
- *(dialog)* Raise permission-guard quiet period to survive natural typing pauses (#3393)
- *(server)* Typing lag over TCP in client/server mode
- *(server)* Keep concurrent sessions from killing each other
- *(server)* Release test DB before Windows temp-dir cleanup
- *(workspace)* Wait for recovery notice before asserting resync
- *(workspace)* Retire integration-test clients so the DB closes
- *(server)* Widen permission cross-client test timeout to stop CI flake
- *(server)* Mcp auth in client-server mode
- *(server)* Replace stale servers that predate shutdown_if_idle
- Reduce default and suggested mcp timeout (#3509)
- *(hyper)* Render error messages from back-end (#3513)

### 💼 Other

- Hide hypercredits if the server sends USD as the monetary unit

### 🚜 Refactor

- *(ui)* Route all agent-model rebuilds through one helper
- *(server)* Tidy mcp client-server auth wiring

### 📚 Documentation

- Add missing OAuth flags to mcp add in crush-config skill
- *(agents)* Mark crushrc as primary config format over crush.json

### ⚙️ Miscellaneous Tasks

- *(legal)* @traitimtrongvag has signed the CLA
- *(legal)* @Lee-Si-Yoon has signed the CLA
- *(legal)* @piakdev has signed the CLA
- *(legal)* @namore has signed the CLA
- *(mcp,ui)* Make 'needs auth' color tang, consolidate theme naming
- Auto-update files
- *(legal)* @nosey-dewdrop has signed the CLA
- Auto-update files
- Gufumpt
- Auto-update files
- Auto-update files
## [0.88.0] - 2026-07-31

### 🚀 Features

- *(oauth)* Give the browser redirect a real landing page
- Add shell config builtins and ConfigBuilder infrastructure
- Add crush.sh discovery and loading
- Add structured logging to shell config builtins
- Add provider-model builtin for defining provider models
- Expose CRUSH_VERSION to shell config scripts
- Add "option reset" to wipe shell config list options
- Verb-first provider and model shell config commands
- Add verb-first hook add/remove shell config commands
- Add verb-first lsp add/remove shell config commands
- Add verb-first mcp add/remove shell config commands
- Make permissions use verb-first allow subcommand
- Rename shell config to crushrc with local .crushrc override
- Add permissions deny to hide tools from the agent
- Configure attribution from crushrc
- Clarify model pricing flags in crushrc
- Expose remaining advanced config through crushrc
- *(config)* Apply top-level env vars on startup
- *(bedrock)* Retry the turn automatically after AWS SSO re-auth

### 🐛 Bug Fixes

- *(mcp)* Normalize oauth metadata redirects (#3415)
- *(mcp)* Remove orphaned tokens from oauths MCP (#3418)
- *(mcp)* Pin go-sdk to main for protocol version header fix (#3421)
- *(illumos)* Support building and running crush" (#3422)
- *(auth)* Stop parallel sessions from invalidating each other's login
- Stop long thinking blocks from re-rendering the entire document every frame (#3454)
- *(ui)* Prevent double spinner on session reload after kill (#3457)
- Resolve short session IDs in local 'crush run --session' (#3460)
- *(mcp)* MCP server loading in server-client mode
- *(edit)* Auto-correct whitespace mismatches in edit tool
- *(server)* Deliver permission and question events reliably to clients
- Accept mixed-case booleans in shell config
- Discover global crush.sh, not just crush.json
- Never execute crushrc from machine data directories
- Make crushrc source test pass on Windows
- *(shellconfig)* Bound crushrc execution with context and timeout
- *(config)* Only warn when JSON and crushrc keys actually conflict
- *(shellconfig)* Dedup model add, document optionSpecs, pin deny target
- *(config)* Track all discovered config paths for staleness
- *(config)* Capture rollback snapshot before configureProviders
- *(config)* Validate crushrc JSON output, skip empty config paths
- *(shellconfig)* Assert section/childMap don't overwrite non-map values
- *(config)* Retry transient Windows rename failures in atomicWriteFi… (#3469)
- *(tools)* Surface DuckDuckGo rate limiting instead of empty results

### 💼 Other

- *(config)* Add crushrc load benchmark

### 🚜 Refactor

- Use positional provider ID for provider-model, rename options to option
- Simplify shell config option booleans
- Make list option keys singular
- Build shell config imperatively instead of via JSON fragments
- Rename provider/model unset to remove (alias rm)
- Make permissions deny the only tool-blocking command
- Unify notification config under a single notifications key
- *(shellconfig)* Collapse option type system into one table
- *(shellconfig)* Replace copy-pasted flag loops with declarative engine
- *(shell)* Unify jq into the builtin registry
- *(shellconfig)* Unify flagInt/flagInt64, make boolTrue a flagKind, checked assertions

### 📚 Documentation

- Update AGENTS.md with Bash config format
- Default crush-config skill to the crush.sh format
- Add config guide and future-work notes
- Use a local Ollama example in the config quick-start
- Lead README configuration with the crush.sh format
- *(config)* Make documentation more human
- Note the deferred visible-but-always-denied tool state
- *(readme)* Update readme per bash-based config
- Present crushrc commands as CLI-style help
- Finish crushrc examples and option help
- *(config)* Add windows paths, copyedits
- *(config)* Add notes on future state migration, update readme/docs
- Fix version typo and clarify config trust wording
- *(crushrc)* Note that empty headers are dropped from requests
- Fix typos, grammar, and document deny-wins precedence
- *(readme)* Improve flow of configuration section a bit
- Document aws_auth_refresh and top-level env config
- *(ui)* Capture dialog rendering and chat perf rules

### 🧪 Testing

- *(shell)* Fix binary passthrough test on multi-call coreutils
- Add comprehensive tests for shell config builtins
- Add black-box tests for the permissions command
- Verify permissions command effect via real config load
- Verify option command effect via real config load
- Verify mcp, lsp, hook, provider, and model commands via config load
- Use crushrc filename in shell config tests
- *(crushrc)* Pin deny-wins-over-allow precedence for permissions
- *(crushrc)* Isolate CRUSH_GLOBAL env vars in permissions tests
- *(shell)* Add builtin dispatch integration tests
- *(tools)* Cover DuckDuckGo rate-limit detection

### ⚙️ Miscellaneous Tasks

- Auto-update files
- *(legal)* @timotheosh has signed the CLA
- *(legal)* @pshickeydev has signed the CLA
- Gufumpt
- Auto-update files
- *(legal)* @kesku has signed the CLA
- Auto-update files
- *(legal)* @Qalipso has signed the CLA
- Bump mcp sdk to 1.7.0 (#3447)
- Regen golden files
- *(legal)* @nobilelucifero has signed the CLA
- Bump fantasy to v0.39.0 (#3470)
- *(crushrc)* Log conflicting keys when json and crushrc compete
- *(crushrc)* Drop deadcode
- Merge main into crushrc
- Auto-update files
- Format
- Regenerate swagger docs
- Auto-update files
## [0.87.0] - 2026-07-24

### 🚀 Features

- *(mcp)* Add OAuth support for HTTP MCP servers
- *(mcp)* Detect claude/channel capability and receive channel messages
- *(cmd)* Add --channels opt-in flag for MCP channel servers
- *(client)* Reconnect the event stream after it drops
- *(ui)* Tell a lost connection apart from an uninitialized agent
- *(mcp)* OAuth 2.1 authorization for HTTP MCP servers
- *(mcp)* SSE OAuth, metadata fixups, and resource param stripping

### 🐛 Bug Fixes

- *(client)* Add MCP prompts to ui in client-server mode
- *(mcp)* Address OAuth review findings (timeout, refresh, concurrency)
- *(agent)* Serialize in-process run dispatch to prevent concurrent turns
- *(mcp)* Tighten channel meta key validation for XML name safety
- *(cmd)* Make --channels a persistent flag so `crush run` inherits it
- *(backend)* Log channel flag mismatch on duplicate workspace creation
- *(mcp)* Route channel messages through must-deliver broker path
- *(server)* Stop mapping channel events to spurious state changes
- *(agent)* Close dispatch completion-boundary cancel race
- *(mcp)* Buffer channel notifications during capability negotiation
- *(mcp)* Filter channel events from the shared app event stream
- *(mcp)* Prevent clients from inheriting channel opt-ins
- *(mcp)* Prevent panic when auth signal fires twice (#3403)
- *(server)* Don't shut down while a workspace is being created
- *(server)* Wait before shutting down so back-to-back sessions don't race it
- *(test)* Fix nil pointer crash in Windows CI tests
- *(config)* Load system-wide config from /etc/crush/crush.json (#2984)
- *(copilot)* Add additional responses models (#3416)
- *(ui)* Keep shell progress output from corrupting the TUI

### 📚 Documentation

- *(mcp)* Add OAuth configuration guide and regenerate schema

### ⚡ Performance

- *(ui)* Avoid quadratic shell output rendering (#3381)
- *(lsp)* Filter servers before searching `$PATH` (#3370)

### 🎨 Styling

- Gofmt formatting cleanup

### 🧪 Testing

- *(mcp)* Cover MCP prompt error paths and document endpoint
- *(mcp)* Add regression tests for OAuth review fixes
- *(server)* Cover the in-flight create counter that guards shutdown

### ⚙️ Miscellaneous Tasks

- *(mcp)* Merge main; fix createTransport arity; inject browser opener
- Auto-update files
- *(cli)* Hide --channels
- Auto-update files
- Auto-update files
- *(legal)* @kkmkoi has signed the CLA

### ◀️ Revert

- Remove #3348 OAuth implementation (superseded)
## [0.86.0] - 2026-07-20

### 🚀 Features

- Add --all and --crawl-dir modes to stats subcommand (#2811)
- Add keybinding and logic to copy verification URL in OAuth dialog #3323 (#3324)
- Recover cleanly from mid-stream provider connection resets
- *(ui)* Add scrollable sidebar with focus-based navigation
- *(ui)* Widen sidebar to 32 cells and reserve scrollbar column
- *(ui)* Scroll sidebar with mouse wheel when focused
- *(ui)* Scroll sidebar with mouse wheel when focused
- *(ui)* Show g/G shortcuts in sidebar full help
- *(ui)* Auto expand reasoning dialog based on count (#3332)
- *(lsp)* Add LSP superpowers tools
- *(lsp)* Update system prompt to recommend lsp tools
- *(lsp)* Add 4 new lsp tools
- *(lsp)* Fancy diff view renders
- *(ui)* Clickable ✕ remove button on attachment chips
- On quit dialog, add hint on how to skip confirmation (#3356)

### 🐛 Bug Fixes

- *(tools)* Report every matching line per file in internal grep (#2994)
- *(queue)* Fix pill border on queued messages (#3333)
- Keep the spinner animating when a response restarts
- Enable thinking blocks for minimax m3 on opencode providers (#3335)
- Keep chat pinned to bottom after a resize (#3336)
- *(agent)* Preserve attachment chips in user messages after sending
- *(mcp)* Clear tools and close the session on MCP error; reap stdio process groups
- *(mcp)* Wait for MCP init before building the tool list
- *(mcp)* Serialize renewals, restore all registries, arm init gate
- *(ui)* Prevent sidebar text clipping and restore sidebar focus
- *(ui)* Leave an empty column between sidebar content and scrollbar
- *(ui)* Only allow sidebar focus when content overflows
- *(ui)* Keep sidebar layout while scrolling full content
- *(ui)* Scroll sidebar details below the logo
- *(ui)* Align the compact sidebar logo
- *(ui)* Resolve timer display conflicts in thinking animation and tool spinners (#3353)
- Prevent new sessions from hanging in client-server mode
- *(ui)* Use fgmostsubtle for canceled text (#3360)
- *(lsp)* Fix TrackConfigured startup race and relative path resolution
- *(lsp/ui)* Fix stale sidebar state on render
- *(lsp)* Add word boundaries to symbol grep and validate against LSP
- *(lsp)* Lsp restart failing
- *(config)* Prevent data race when reading config during reload (#3362)
- *(commands)* Scope logout command to oauth providers
- *(ui)* Keep chips in place when toggling attachment delete-mode
- *(ui)* Restore right padding on the attachment remove button
- *(refactor)* Refactor the edit tool to not duplicate find and replace
- *(tui)* Prevent stale background refreshes from overwriting UI state (#169)
- *(baseten)* Fix "none" reasoning level (#3386)

### 🚜 Refactor

- *(ui)* Move sidebar scroll state mutation out of draw function
- Simplify edit tools and enforce Sourcegraph result limits

### ⚡ Performance

- *(tui)* Keep synchronous workspace probes off the per-message Update path

### ⚙️ Miscellaneous Tasks

- *(legal)* @santhreal has signed the CLA
- Bump modernc.org/sqlite (#3334)
- Auto-update files
- *(legal)* @joestump has signed the CLA
- *(legal)* @H-TTTTT has signed the CLA
- *(ui)* Remove dead Editor.AddFile keybinding
- Generate golden files
- Auto-update files
- *(legal)* @TheJhyeFactor has signed the CLA
## [0.85.0] - 2026-07-16

### 🚀 Features

- *(question)* Add question tool with structured UI
- *(question)* Add client server integration
- *(question)* Add mouse support
- *(question)* Add paste support in text areas
- *(question)* Redo tab resizing and mouse -> keyboard transition
- *(question)* Adjust question prompts and error messages
- *(question)* Add mouse scrolling
- *(question)* Make escape cancel the question instead of submitting empty answers
- *(question)* Extend length limits on question tool
- Integrate fantasy OnAuthRefresh for transparent auth retry
- *(question)* Allow newlines and make free text like pop
- *(question)* Tweak the confirmation tab
- *(question)* Tweak yes/no to add shortcuts
- *(dialog)* Hide list info column when it would crowd item names
- Wire LM Studio vision capabilities to SupportsImages (#3280)
- Elapsed seconds timer (#3223)

### 🐛 Bug Fixes

- *(question)* Fix scrollbar disappearing in single-select
- *(question)* Address PR review feedback
- *(question)* Make scroll containers consistent
- *(question)* Fix a bug with the hover state coming deselected
- *(errors)* Properly supress context cancelation error
- *(oauth)* Fix the ui freeze on loading model endpoints
- *(events)* Restore file picker event (#3314)
- Respect client-server mode during login (#3315)
- *(tests)* Bail on precancelled ctx in handleJQ (#3058)
- Match question tool cursor color to main editor
- *(dialog)* Make OAuth dialog width responsive to terminal size
- *(dialog)* Make APIKeyInput width responsive to terminal size
- *(dialog)* Constrain Quit dialog padding on narrow screens
- *(dialog)* Scale FilePicker down on small screens
- *(dialog)* Fix OAuth content wrapping and full-width layout
- *(dialog)* Truncate dialog titles instead of wrapping on small screens
- *(dialog)* Subtract border size in Reasoning and Notifications width
- *(dialog)* Truncate TitleInfo when it exceeds dialog width
- *(dialog)* Clamp DrawCenter and DrawOnboarding content to screen bounds
- *(dialog)* Centralize keybind hint rendering and stop overflow
- *(dialog)* Mute command list shortcut hints
- *(dialog)* Keep dialogs stable on small screens
- *(dialog)* Stop list item names from wrapping past the list width
- *(dialog)* Remove dead space beside the list scrollbar
- *(dialog)* Account for the input prompt width so long values don't wrap
- *(dialog)* Fix model provider label truncation at low widths
- *(dialog)* Center permission buttons in fullscreen

### 💼 Other

- *(client/server)* Fix non interactive init

### 🚜 Refactor

- Make the coordinator use a struct
- *(dialog)* Dedupe scrollbar joins and tame the permission Draw

### ⚡ Performance

- *(ui)* Memoize the chroma syntax-highlight style
- *(ui)* Memoize chroma lexer lookups by filename
- *(ui)* Keep chat resize smooth on large conversations

### ⚙️ Miscellaneous Tasks

- *(legal)* @pranavthakur0-0 has signed the CLA
- *(legal)* @albatrossflyon-coder has signed the CLA
- *(legal)* @B-A-M-N has signed the CLA
- Auto-update files
- Auto-update files
- *(legal)* @seven7763 has signed the CLA
- *(legal)* @joestump-agent has signed the CLA
- Switch back to upstream openai sdk (#3308)
## [0.84.1] - 2026-07-11

### ⚙️ Miscellaneous Tasks

- *(legal)* @coloryourlife has signed the CLA
## [0.84.0] - 2026-07-10

### 🚀 Features

- *(agent)* Send session hash in header for cache affinity

### 🐛 Bug Fixes

- *(release)* Drop broken .termux.deb artifact (#2848)
- *(server)* Show update notices in client-server mode

### ⚙️ Miscellaneous Tasks

- *(legal)* @josh-socium has signed the CLA
- Gofmt
## [0.83.0] - 2026-07-08

### 🚀 Features

- Prepare for gpt-5.6 (#3270)

### 🐛 Bug Fixes

- *(bangmode)* Don't add extra ! when browsing history

### ⚙️ Miscellaneous Tasks

- *(legal)* @ychampion has signed the CLA
- *(legal)* @retakt has signed the CLA
## [0.82.0] - 2026-07-07

### 🚀 Features

- *(ui)* Preserve newlines in expanded tool content (#3239)

### 🐛 Bug Fixes

- *(scrollbar)* Only show on human scroll
- *(ui)* Make chat follow-scroll reliable when content grows (#3240)
- *(copilot)* Guard nil request body in initiator transport (#3246)
- *(baseten)* Make the thinking on/off toggle work for baseten
- *(alibaba)* Fix missing thinking traces on deepseek via aliababa (#3259)

### 💼 Other

- Add hyper

### 📚 Documentation

- *(readme)* Mention `MOONSHOT_API_KEY` (#3251)

### ⚙️ Miscellaneous Tasks

- *(legal)* @sablea has signed the CLA
- *(legal)* @Osamaali313 has signed the CLA
- *(legal)* @lockp111 has signed the CLA
- Add alibaba us (#3249)
- *(legal)* @TheRAWagent has signed the CLA
## [0.81.0] - 2026-06-29

### 🚀 Features

- *(ui)* Add scrollbar to chat view (#3018)
- *(herdr)* Add herdr socket integration
- *(ui)* Allow expanding toolcall names

### 🐛 Bug Fixes

- *(config)* Prevent stale ReasoningEffort from leaking across providers (#3209)
- *(hyper)* Refresh hyper oauth token before fetching credits (#3212)
- *(config)* Prevent startup deadlock when configured model ID is invalid
- Cache streaming thinking renders to avoid CPU burn during long reasoning traces
- *(prompts)* Unindent heredoc
- *(agent)* Fall back to first reasoning level when default is unset (#3218)
- *(ui)* Ansi.truncate instead of byte slicing in todos
- *(ui)* Use ansi.truncate everywhere
- *(bash)* Keep TruncateOutput from splitting UTF-8 characters
- *(agent)* Prevent session bricking when non-vision models receive tool result media
- *(fsext)* DirTrim renders wrong character for non-ASCII directory names (#3214)
- *(shell)* Skip persistence when session no longer exists

### ⚙️ Miscellaneous Tasks

- *(legal)* @feizhuzheng has signed the CLA
- Auto-update files
- *(legal)* @syf2211 has signed the CLA
- *(legal)* @EduardF1 has signed the CLA
- Auto-update files
- Regen golden files
- Auto-update files
## [0.80.0] - 2026-06-25

### 🚀 Features

- Log provider warnings from fantasy step results
- *(providers)* Add llamacpp enricher
- *(ui)* Optimize model ui rendering

### 🐛 Bug Fixes

- *(hooks)* Bridge Claude Code additionalContext onto HookResult.Context
- Correct model discovery enrichment for local providers
- *(docs)* Set proper auto discovery models flag
- *(scrolling)* Stale scrolling acceleration (#3197)
- Prevent flaky diffview golden tests on Windows CI
- *(glob)* Keep file search fast and bounded on large directories
- Validate tool call JSON before storing to prevent stuck conversations
- *(config)* Make concurrent config access race-free via copy-on-write
- *(auth)* Stop two Crush instances from invalidating each other's login
- Render pills box reliably when todos or queue appear mid-session

### ⚡ Performance

- *(ui)* Skip theme rebuild when provider keeps the same theme
- *(config)* Make model selection and config reload fast

### ⚙️ Miscellaneous Tasks

- *(legal)* @xorangekiller has signed the CLA
- *(legal)* @sithglan has signed the CLA
- *(legal)* @macro-ss has signed the CLA
- *(legal)* @andrinoff has signed the CLA
- *(legal)* @calyptobai has signed the CLA
- *(legal)* @truffle-dev has signed the CLA
- *(legal)* @chardoncs has signed the CLA
- Bump fantasy to 0.33.2
- Bump golangci version to fix lint issues in ci
- Fix lint on interp.ExecHandler (#3200)
- *(legal)* @notno has signed the CLA
- *(legal)* @vismaytiwari has signed the CLA
- Auto-update files
## [0.79.1] - 2026-06-20

### 🐛 Bug Fixes

- *(bangmode)* Activate bang mode when ! is preceded by whitespace
- *(bangmode)* Engage bang mode when pasting text starting with !
- *(bangmode)* Include bang commands in history
- *(bangmode)* Sync bang mode with external editor

### ⚙️ Miscellaneous Tasks

- *(legal)* @ken-jo has signed the CLA
## [0.79.0] - 2026-06-19

### 🚀 Features

- *(schema)* Fix the reflection on provider options
- *(bang)* Show pending spinner
- *(bang)* Stream results in
- *(bang)* Remap ansi 16 colors
- *(bangmode)* Cancel command execution
- *(bangmode)* Copy message result
- *(bangmode)* Properly strip ansi in copy and context
- *(bangmode)* Adjust command output color
- *(bangmode)* Allow prefixing a string with bangmode

### 🐛 Bug Fixes

- *(bang)* Set title in bang mode
- *(bangmode)* Interleave stderr and stdout
- *(bangmode)* Fix duplicate command message race condition

### ⚙️ Miscellaneous Tasks

- Auto-update files
- *(bangmode)* Adjust ANSI16 colors
- *(bangmode)* Adjust working indicator text color
## [0.78.0] - 2026-06-19

### 🚀 Features

- *(scrolling)* Move to delta coalescing filter (#3135)
- Auto-discover models from openai-compat providers
- *(enricher)* Add litellm enricher
- *(enricher)* Add ollama enricher
- *(enricher)* Add omlx enricher
- *(enricher)* Add lmstudio enricher
- *(dialog)* Track last closed dialog in open with grace
- Add support for fireworks provider
- Add bang mode for direct shell command execution (#3013)
- *(ui)* Extract hypercredits from fantasy
- *(clipboard)* Migrate clipboard to golang.design/x/clipboard

### 🐛 Bug Fixes

- *(subagents)* Fixed subagents returning empty responses (#3125)
- *(shell)* Isolate child processes from Crush's session (#3097)
- *(title)* Fallback title generation
- *(server)* Show LSP status in client-server mode
- *(clipboard)* Support platforms without clipboard access

### 📚 Documentation

- Add auto discovery to readme

### ⚙️ Miscellaneous Tasks

- *(legal)* @warmjademe has signed the CLA
- *(legal)* @jialudev has signed the CLA
- *(legal)* @knight110001 has signed the CLA
- Auto-update files
## [0.77.0] - 2026-06-14

### 🚀 Features

- *(server)* Add helper to detect stale unix sockets
- *(server)* Store runtime socket in per-user runtime directory
- *(hooks)* Add name field to hooks
- *(config)* Load user-level context files

### 🐛 Bug Fixes

- *(server)* Log earlier so socket cleanup is always recorded
- *(server)* Clear leftover sockets so server can always start
- *(server)* Detect and remove a dead socket before starting server
- *(server)* Correct socket location test on macOS
- Address potential indentation on commit messages trailings (#3106)

### 📚 Documentation

- *(hooks)* Add name field

### 🧪 Testing

- Update VCR cassettes

### ⚙️ Miscellaneous Tasks

- *(legal)* @yuseferi has signed the CLA
- *(legal)* @gmit3 has signed the CLA
- *(legal)* @brianjlandau has signed the CLA
- *(test,server)* Cover stale socket cleanup and socket location
- Auto-update files
- *(legal)* @kypkk has signed the CLA
- Auto-update files
- Bump lipgloss to v2.0.4 and glamour to v2.0.1
- *(bash/git)* Improve git commit and PR message standards (#3052)
## [0.76.0] - 2026-06-05

### 🚀 Features

- *(server)* Make server prompts independent of client connections
- Discover skills from git root in monorepos (#3078)

### 🐛 Bug Fixes

- *(ui)* Fix session rename rendering (#3071)
- *(server)* Prevent cancels from affecting future prompts
- *(server)* Ignore background errors from unrelated prompts
- *(server)* Close a queued-prompt cancel race
- *(server)* Complete queued prompts independently
- *(server)* Complete background prompts that fail early
- *(permissions)* Ignore all keys briefly after dialog opens (#3055)
- *(cli)* Show all providers in crush models, not just configured ones

### 🚜 Refactor

- *(commands)* Centralise skill discovery
- *(server)* Share prompt validation before background dispatch

### ⚙️ Miscellaneous Tasks

- *(server)* Honor cancels immediately after prompt acceptance
- *(server)* Report background prompt failures via events
- *(server)* Run accepted server prompts in the background
- *(server)* Acknowledge accepted prompts with HTTP 202
- *(server)* Apply asynchronous prompt contract to clients
- *(server,tesst)* Cover multi-client prompt cancellation flows
- *(server,tests)* Cover async cancellation cleanup behavior
- *(server,tests)* Cover the accepted-prompt cancel race end to end
- *(legal)* @GroovyCarrot has signed the CLA
- Auto-update files
- *(legal)* @CnsMaple has signed the CLA
- Update fantasy and catwalk (#3087)
## [0.75.0] - 2026-06-02

### 🚀 Features

- Prepare alibaba migration to `/messages` (#3067)

### 🐛 Bug Fixes

- *(agent)* Move reauth notification from agent to coordinator
- *(oauth)* Harden IsExpired with minimum buffer and ExpiresIn guard
- *(agent)* Add 401 retry and reauth notification to sub-agent runs
- *(agent)* Centralize 401 retry logic and fix notify-on-success bug
- *(oauth)* Stop fabricating token lifetime when expires_in is missing
- *(config)* Sort SetConfigFields keys, remove redundant MkdirAll, rename test
- Avoid startup crash if unable to find default models (#3066)

### 🚜 Refactor

- *(lock)* Add canonical internal/lock package, migrate db and cmd callers
- *(config)* Replace reload booleans with reloadMu mutex

### ⚙️ Miscellaneous Tasks

- *(legal)* @ardi1s has signed the CLA
- Auto-update files
- *(legal)* @cak3ninja has signed the CLA
## [0.74.1] - 2026-05-29

### 🐛 Bug Fixes

- *(noninteractive)* Crush run reliability in client/server mode
- *(scrollbar)* Fix track position calculation
- *(opencode)* Fix qwen3.7-max on opencode (#3040)
- Set non-interactive env vars in shell to prevent editor hangs (#3025)

### ⚙️ Miscellaneous Tasks

- *(legal)* @xulongzhe has signed the CLA
## [0.74.0] - 2026-05-28

### 🚀 Features

- Support notifications for ssh terminal
- *(notifications)* Add configurable backend and bell support
- *(notifications)* Migrate disabled and add picker
- *(tools)* Add diff view for denied tools

### 🐛 Bug Fixes

- *(ui)* Properly render model name in summarize
- *(client)* Prevent event subscription panic on cancellation
- *(server)* One client's cancel should not send a 500 to others
- *(copilot)* Add additional responses models
- *(bedrock)* Load aws credentials for bedrock europe
- *(bedrock)* Improve detection of pre-existing aws credentials

### ⚙️ Miscellaneous Tasks

- Auto-update files
- Gofumpt
- Fix(bedrock): apply region if given for aws config as well
- Auto-update files
## [0.73.0] - 2026-05-26

### 🚀 Features

- Add aws bedrock europe (#3016)

### 🐛 Bug Fixes

- *(backend)* Fix data race in tests using captureDebugLogs
- *(server)* Release pooled DB on test shutdown so Windows can clean temp dir
- *(ui)* Improve model changed notification (#3015)

### ⚙️ Miscellaneous Tasks

- *(labeler)* Do not use custom token
- *(labeler)* Add missing permission
## [0.72.0] - 2026-05-25

### 🚀 Features

- *(skills)* Add descriptions to skill picker and use attachements
- *(ui)* Add scrollbar to sessions dialog (#3005)
- *(ui)* Auto-expand pills when terminal height is sufficient
- *(ui)* Add ctrl+y keybinding to toggle yolo mode
- *(ui)* Show notification when toggling yolo mode (#3008)

### 🐛 Bug Fixes

- *(tests)* Fix flaky async windows test
- *(ui)* Only auto-expand pills once per session lifecycle

### ⚙️ Miscellaneous Tasks

- *(legal)* @officialasishkumar has signed the CLA
- *(legal)* @yhyu13 has signed the CLA
## [0.71.0] - 2026-05-22

### 🚀 Features

- *(permissions)* Require a permission prompt for chained commands
- *(db)* Refuse to open a data directory in use by another crush
- *(server)* Share one workspace per directory across clients
- *(server)* Broadcast config changes to all connected clients
- *(tui)* Auto close permission prompt when another client responds
- *(api)* Expose in progress flag on session responses
- *(server)* Track which session each client is currently viewing
- *(api)* Report how many clients are watching each session
- *(skills)* Add support for user invocable skills
- Render scrollbar for model list (dialog and onboarding) (#2978)

### 🐛 Bug Fixes

- *(prompts)* Tweak file reads to encourage more targetted reads
- *(permissions)* Make permission resolution idempotent across clients
- *(db)* Only enforce the data directory lock in client server mode
- *(server)* Support attachments in client-server mode
- *(server)* Display available skills in client
- Potential data race on `permissionService` (#2964)
- Update fantasy with stream fixes (#2968)
- *(agent)* Estimate missing streamed usage
- *(agent)* Correct fallback usage accounting
- *(agent)* Harden fallback usage accounting
- *(agent)* Clear stale summary token counts
- *(ui)* Mark estimated context usage
- *(session)* Preserve estimated usage marker
- *(ui)* Preserve estimated usage percentage color
- *(ui)* Add locking around markdown rendering
- *(ui)* Guard divide-by-zero display error
- Fix sometimes sending reasoning effort when it shouldn't (#2982)
- *(bedrock)* Enforce `us-east-1` as region for bedrock (#2985)
- *(db)* Keep SQLite temp files in memory
- *(models)* Fix sorting of hyper
- *(ui)* Scroll to the properly select model

### 📚 Documentation

- Describe how Crush shares a workspace across clients

### 🧪 Testing

- *(server)* Cover the multi client flows end to end

### ⚙️ Miscellaneous Tasks

- *(legal)* @dcu has signed the CLA
- *(tests)* Update golden files
- *(db)* Log lock metadata write failures and explain lock file lifetime
- *(legal)* @mei2jun1 has signed the CLA
- *(legal)* @g2mt has signed the CLA
- *(legal)* @Ricardo-M-L has signed the CLA
- *(legal)* @Muttaqin86 has signed the CLA
- Auto-update files
- Isolate store test
## [0.70.0] - 2026-05-18

### 🚀 Features

- Add aliababa (singapore) provider (#2949)

### 🐛 Bug Fixes

- Retry on network errors (#2945)
- *(auth)* Support oauth in client-server mode
- *(server)* Recover from handler panics + return 500
- *(ui)* Restore pills to-do in client/server mode

### ⚡ Performance

- Remove hot-path slog.Info from SSE event delivery (#2929)

### ⚙️ Miscellaneous Tasks

- Auto-update files
- Auto-update files
- *(legal)* @Broderick-Westrope has signed the CLA
- *(legal)* @processtrader has signed the CLA
- Auto-update files
- Gofumpt project
## [0.69.1] - 2026-05-15

### 🐛 Bug Fixes

- *(ui)* Keep tool spinners animating during long-running tasks

### ⚡ Performance

- *(chat)* Cache the prefixed render of chat messages
- Batch streaming message updates
- *(chat)* Cache the parts of an assistant message separately
- *(chat)* Show only the tail of long reasoning blocks when expanded
- *(chat)* Skip re-rendering chat list items that have not changed
- *(chat)* Only render the chat lines that fit on screen
- *(chat)* Reuse the rendered prefix of a streaming reply
- *(chat)* Skip re-parsing the rendered chat when nothing has changed

### ⚙️ Miscellaneous Tasks

- *(legal)* @CatMe0w has signed the CLA
- *(ui)* Make spinner deterministic for testing
## [0.69.0] - 2026-05-15

### 🚀 Features

- *(dirs)* Add some styling to the dirs command

### 🐛 Bug Fixes

- *(deepseek)* Fix 400 bad request for deepseek provider (#2923)

### ⚙️ Miscellaneous Tasks

- *(legal)* @yookibooki has signed the CLA
## [0.68.0] - 2026-05-14

### 🚀 Features

- *(prompts)* Remove long prompt option
- *(fsext)* Stop upward lookup at a boundary directory
- *(prompts)* Template prompts and add github and ripgrep info
- *(prompts)* Extend templating system to more prompts
- *(logs)* Add a log line for dropped events
- *(tools)* Create an allow list for MCP tools (#2800)
- *(oauth)* Add logout command (#2838)
- *(responses)* Opt in specific copilot models to responses

### 🐛 Bug Fixes

- *(paste)* Normalize windows newlines
- *(config)* Always resolve the data directory to an absolute path (#2883)
- *(config)* Scope .crush discovery to the current repo
- *(config)* Scope crush.json discovery to the current repo
- *(prompts)* Don't include ripgrep and gh prompts in testing
- *(permission)* Fix publish-before-lock race and use O(1) session permission lookups
- *(db)* Use connection pool to avoid corrupted writes
- *(dns)* Fix tmux dns resolver
- *(pubsub)* Raise default per-subscriber buffer (64 -> 4096)
- *(config)* Use large model for small if not configured (#2873)
- *(ui)* Calculation bug that would cause modified files to wrap
- *(bedrock)* Honor reasoning_effort for Anthropic-on-Bedrock models (#2887)
- *(tools)* Switch bufio scanner for reader (#2884)
- Make thinking on/off toggle work for deepseek provider
- *(ui)* Regression in tool output in client/server mode (#2878)
- Render permission dialog content in client/server mode (#2877)
- *(server)* Probe readiness over HTTP instead of statting the socket
- *(server)* Keep the spawned server alive after the parent exits
- *(server)* Spawn after stale shutdown + handle socket errors
- *(server)* Serialize concurrent server spawns with a per-host lock
- *(lint)* Require `exec.CommanadContext` over `exec.Command`
- *(test,race)* Probe server health during the run, not after
- *(ui)* Show custom skills reliably on startup
- *(auth)* Add better atomic refresh for hyper
- Address thinking on/off toggle for provider (#2916)
- *(reasoning)* Enforce can reason from catwalk before sending reasoning_effor
- Detect stale server during development with BuildID

### 🚜 Refactor

- *(server)* Derive per-host cache dir from parsed host URL
- *(skills)* Make coordinator the sole skill discovery publisher

### 🧪 Testing

- *(config)* Tests for the data directory paths
- *(server)* Regression for the client/server spawn race

### ⚙️ Miscellaneous Tasks

- Auto-update files
- *(golden)* Rerecord vhs
- *(legal)* @jan-xyz has signed the CLA
- *(tests)* Update golden files
- *(legal)* @johnjansen has signed the CLA
- *(legal)* @pablodz has signed the CLA
- Allow manual trigger of schema update
- Fetch updated info from hyper
- Fix indentation
- *(legal)* @akhenakh has signed the CLA
- *(legal)* @taciturnaxolotl has signed the CLA
- Auto-update files
## [0.67.0] - 2026-05-11

### 🚀 Features

- *(shell)* Add ExpandValue for config value shell expansion
- *(config)* Resolve MCP args and thread resolver through env/headers/args
- *(config)* Resolve MCP url through shell expansion
- *(shell)* Shebang/binary/in-process dispatch handler
- *(hooks)* Run via shell.Run instead of sh -c
- *(hooks)* Propagate CRUSH/AGENT env vars to builtin shell
- Add touch tool for empty files

### 🐛 Bug Fixes

- *(ui/chat)* Make keyboard expand work for assistant thinking blocks (#2791)
- *(tools/view)* Detect image mime type; don't rely on extension (#2757)
- *(shell)* Ctx-aware jq builtin
- *(shell)* Convert path to posix path in tests
- *(shell)* Fix build error post-refactor
- *(tools/touch)* Gate outside-workingDir paths via permission prompt
- *(lsp)* Update powernap with fix for lsps windows (#2862)
- *(agent)* Release activeRequests before publishing TypeAgentFinished
- *(schema)* Fix schema descriptions being cut off
- *(config)* Individual errors on json parse
- Properly follow the `Assisted-by` header spec (#2871)
- Limit view size checks to returned content (#2785)

### 💼 Other

- Switch config value expansion to lenient by default
- Fail provider header expansion loudly and drop empty values
- Expand shell variables in args and env
- Remove unused environment variable resolver
- Resolve conflicts with main

### 🚜 Refactor

- *(config)* Resolve shell vars via shell.ExpandValue
- *(config)* Make Resolved{Env,Headers} pure and error-returning
- *(shell)* Extract stateless run entrypoint
- *(tools)* Remove touch tool; allow empty write content

### 📚 Documentation

- *(README)* Add note about shell expansion in MCP config
- Document lenient shell expansion and security model
- Update resolver godoc to match lenient default
- *(skill)* Document shell expansion in crush-config skill
- *(hooks)* Document new embedded shell model
- *(hooks)* Clarify relative paths

### 🧪 Testing

- Cover shell expansion in MCP config resolution
- Use forward slashes in shell commands for Windows compat
- Update config tests to match lenient expansion
- Pin provider skip behavior for api_key and endpoint expansion
- Realign MCP init tests with lenient shell expansion

### ⚙️ Miscellaneous Tasks

- *(legal)* @sven2718 has signed the CLA
- *(legal)* @acheong08 has signed the CLA
- *(legal)* @smeinecke has signed the CLA
- Auto-update files
- Modernize errors.As to errors.AsType
- Auto-update files
## [0.66.1] - 2026-05-08

### 🐛 Bug Fixes

- *(config)* Atomically update multiple fields during oauth
- *(tools)* Fix a potential nill crash in cached glob results
- *(posthog)* Do not discard custom properties of an error (#2829)
- *(errors)* Surface errors in subagents
- *(ui)* Add exit alias to the quit command
- *(agent)* Support flat_rate cost handling (#2116)
- *(tui)* Show initialization mark errors in status footer (#2825)
- *(ui)* Allow oauth modals to consume enter
- *(tools)* Don't return a go error on glob tool failure
- *(pubsub)* Respect channelBufferSize parameter in Subsribe
- *(tools)* Truncate long running background commands to 30k chars
- Update fantasy with tool call fixes (#2839)

### 📚 Documentation

- *(readme)* Document `HYPER_API_KEY`
- *(readme)* Fixed typo in hooks (#2801)

### ⚙️ Miscellaneous Tasks

- *(legal)* @yeonuk-hwang has signed the CLA
## [0.66.0] - 2026-05-06

### 🚀 Features

- Add Nix flake for development environment (#2512)

### 🐛 Bug Fixes

- *(ui)* Prevent duplicate custom skills from rendering
- *(agent)* Drain queued messages after manual session summarize
- Skip image attachments in history when model doesn't support them (#2818)
- *(summarize)* Reauthenticate oauth tokens when used to summarize
- *(ui)* Display error on summarization instead of leaving spinning

### 💼 Other

- Yollo mode via flag doesn't activate prompt

### 🚜 Refactor

- *(coordinator)* Extract token refresh helpers to reduce duplication

### ⚙️ Miscellaneous Tasks

- *(legal)* @ardevd has signed the CLA
## [0.65.3] - 2026-05-04

### 🐛 Bug Fixes

- *(db)* Prevent SQLITE_NOTADB corruption under concurrent sub-agents (#2690)
- *(ui)* Cache glamour renderers
- *(config)* Check config file for newer token before OAuth refresh

### 🚜 Refactor

- *(ui)* Pair markdown cache invalidation with the styles mutation

### ⚙️ Miscellaneous Tasks

- *(legal)* @somjik-api has signed the CLA
- *(legal)* @ilgax has signed the CLA
- Auto-update files
## [0.65.1] - 2026-05-01

### ⚙️ Miscellaneous Tasks

- Remove snapcraft token
## [0.65.0] - 2026-05-01

### 🚀 Features

- *(hyper)* Show remaining hypercredits in the sidebar (#2766)
- Launch hyper beta (#2768)
- *(ui)* Add hypercredit readout to small top header

### 🐛 Bug Fixes

- Fix thinking on/off toggle for certain openai-compat providers

### 🚜 Refactor

- *(hyper)* Simplify by removing old code

### 🧪 Testing

- Re-record test fixtures

### ⚙️ Miscellaneous Tasks

- *(legal)* @vorticalbox has signed the CLA
- *(legal)* @lloydzhou has signed the CLA
- *(ui)* Change wording: rewrote input to rewrote output (#2742)
- *(legal)* @carlosgrillet has signed the CLA
- Update `view` tool limit to 200KB
- Auto-update files
- *(legal)* @SAY-5 has signed the CLA
- *(legal)* @mkaaad has signed the CLA
- *(legal)* @pragneshbagary has signed the CLA
- *(ui)* Hypercrush small type treatment
## [0.64.0] - 2026-04-29

### 🚀 Features

- *(ui)* Switch to hyper theme when provider hyper is chosen

### 🐛 Bug Fixes

- *(hyper)* Re-auth at selection time to ensure provider availability
- Remove unused build (#2734)
- *(ui)* Restore 'update available' notification
- *(ui)* Notification width and text truncation
- *(styles)* Fix some regressions where colors were incorrect
- *(hooks)* Recompile matchers after config reload
- *(app)* Replace single events channel with pubsub.Broker for fan-out (#2663)
- *(tools/job_kill)* Use longer job_kill desc to improve reliability (#2747)

### 🚜 Refactor

- *(hooks)* Move matcher compilation into the runner

### ⚙️ Miscellaneous Tasks

- *(styles)* Use hypercrush theme when hyper is selected
- *(styles)* Clean up theme definitions
- *(ui,styles)* Color edits and copyedits for the oAuth view
- *(hyper)* Update endpoint to new one
- Auto-update files
- *(legal)* @georgeglarson has signed the CLA
## [0.63.0] - 2026-04-27

### 🚀 Features

- PreToolUse hook (#2598)

### 🐛 Bug Fixes

- *(ui)* Don't show disabled skills
- *(ui)* Fix dialog box shift when session rename is active
- *(ui/hooks)* Restore hook styling

### 📚 Documentation

- *(hooks)* Improve hook documentation

### ⚙️ Miscellaneous Tasks

- *(tests)* Regression test for (lack of) disabled skills in the ui
- Auto-update files
- *(ui)* Theme prep
- *(hooks,skills)* Update crush-hooks skill per recent changes
## [0.62.1] - 2026-04-24

### 🐛 Bug Fixes

- Remove minimax api key validate (#2688)
- *(lsp)* Replace sticky unavailable cache with retry backoff (#2498)
- *(agent)* Implement OnRetry logging with structured retry fields (#2700)
- *(ui)* Logo and grad arguments from earlier refactor
- *(styles)* Use semantic names in styles + drop deadcode
- *(hyper)* Fix re-authorization flow not triggering on certain conditions (#2703)

### 📚 Documentation

- *(readme)* Tiny updates

### 🧪 Testing

- Re-record vcr cassettes

### ⚙️ Miscellaneous Tasks

- *(legal)* @ne275 has signed the CLA
- *(legal)* @flynn-eye has signed the CLA
- Auto-update files
- *(ui)* Add new letterforms: h, y, p, e, with alts
- *(ui)* Use Lip Gloss for color blends
- *(ui)* Formal hypercrush type treatment
## [0.62.0] - 2026-04-22

### 🚀 Features

- Generally render output that looks like a diff as a diff (#2607)

### 🐛 Bug Fixes

- *(lsp)* Mitigate stale diagnostics
- Reduce `fetch` and `view` tools truncation size to 100KB
- Reduce token usage, use short tool descriptions by default (#2679)
- Silence unless warning about non-existent skill paths

### 🧪 Testing

- Re-record vcr cassettes

### ⚙️ Miscellaneous Tasks

- *(legal)* @axeprpr has signed the CLA
- Remove CODEOWNERS
- *(legal)* @iuga has signed the CLA
- Update catwalk to new domain (#2680)
- Auto-update files
## [0.61.1] - 2026-04-21

### ⚙️ Miscellaneous Tasks

- *(legal)* @gavmor has signed the CLA
- Use stable goreleaser action
## [0.61.0] - 2026-04-20

### 🚀 Features

- Enable progress bar on iterm2 (#2641)
- *(skills)* Log skill activation to help diagnose load failures
- *(prompt)* Require loading appropriate skills before acting

### 🐛 Bug Fixes

- *(oauth)* Fix copy to clipboard on terminals that don't support osc52 (#2642)
- Use same chroma formatter as diffview for markdown (#2656)

### 🚜 Refactor

- *(skills)* Strip implementation hints from crush-owned descriptions

### ⚙️ Miscellaneous Tasks

- *(legal)* @pi128 has signed the CLA
- *(legal)* @enrell has signed the CLA
## [0.60.0] - 2026-04-17

### 🚀 Features

- *(hyper)* Use openai-compatible endpoint for hyper (#2640)
## [0.59.0] - 2026-04-16

### 🚀 Features

- Show progress bar on rio terminal (#2624)
- *(ui)* Add skills discovery status to sidebar and landing page (#2384)
- *(shell)* Add jq as a bash builtin + a jq skill for it

### 🐛 Bug Fixes

- Inject synthetic tool_result for orphaned tool_use on session resume (#2622)
- *(prompt)* Remind coder to follow `<git_commits>` format on commit

### 🧪 Testing

- Reduce amount of logs print on `go test -v`
- Simplify integration tests, run for a single provider / model
- Re-record vcr cassettes
- Run migration setup once to avoid race condition
- Re-record vcr cassettes
- Re-record vcr cassettes

### ⚙️ Miscellaneous Tasks

- *(legal)* @taigong12 has signed the CLA
- Bump bubbletea to v2.0.6 to fix wide char issue
- `task fmt`
- Only run `build` and `lint` workflows once in pull requests
- Fix test race condition
## [0.58.0] - 2026-04-15

### 🚀 Features

- Add opencode zen and opencode go support

### 🐛 Bug Fixes

- *(system-prompt)* Remove "portuguese" as example

### 🧪 Testing

- Re-record vcr cassettes

### ⚙️ Miscellaneous Tasks

- Bump bubbletea to v2.0.5 to fix tea_debug.log file issue
- *(legal)* @KimBioInfoStudio has signed the CLA
- Auto-update files
## [0.57.0] - 2026-04-13

### 🚀 Features

- *(config)* Support `HYPER_API_KEY` for hyper auth (#2583)

### 🐛 Bug Fixes

- Show attribution setting on `crush_info` tool (#2594)
- *(tools)* Drastically reduce tool call description lengths
- *(tools)* Modernize string split
- Use proper bool check
- *(tools/bash)* Restore cross-platform instructions in bash tool
- *(tools)* Add missing strconv import
- *(agent)* Prevent session corruption due to malformed image data (#2597)
- *(events)* Prevent early events from being dropped before init (#2611)
- *(ci)* Use stable Go version in security workflow
- *(ui)* Format code in ui.go
- *(agent)* Validate tool call/results + strip tags from titles

### 🚜 Refactor

- Simplify skills parsing and improve discovery visibility (#2350)

### 🧪 Testing

- Always use short tool descriptions for tests

### ⚙️ Miscellaneous Tasks

- *(legal)* @gurnben has signed the CLA
- *(legal)* @dshumphr has signed the CLA
- *(legal)* @bensantora has signed the CLA
- *(tools/agent)* Improve and correct tool description
- *(tools/view)* Improve tool description
- *(tools/edit)* Improve tool description
- *(tools/ls)* Improve tool description
- *(tools/glob)* Improve tool description
- *(tools/write)* Improve tool description
- *(tools/fetch,web_fetch,web_search)* Improve tool descriptions
- *(tools/download)* Improve download tool description
- *(tools/multiedit)* Improve tool description
- *(tools/grep)* Improve tool description
- *(tools/sourcegraph)* Improve tool description
- *(tools/lsp)* Improve LSP tool descriptions
- *(tools/mcp)* Generally improve mcp tool descriptions
- *(tools/agentic_fetch)* Improve tool description
- *(tools/jobs)* Improve job tool descriptions
- *(tools/todos)* Improve todo tool description
- *(tools/crush_info)* Improve tool description
- *(tools/crush_logs)* Improve tool description
- *(tools/bash)* Improve tool description
- *(tools)* Gate short descs with CRUSH_SHORT_TOOL_DESCRIPTIONS
- *(tests)* Re-record vcr cassettes
- *(legal)* @ahostbr has signed the CLA
- *(agent)* Move filter logic into a function
## [0.56.0] - 2026-04-08

### 🚀 Features

- Open Hyper auth dialog automatically on unauthorized error
- *(tools)* Crush_info tool for readling live config
- *(tools)* Add crush_info tool implementation files
- *(tools)* Add crush_logs tool for reading application logs
- *(tools/crush_info)* Handle config staleness
- *(tools/crush_info)* Staleness detection and auto-reload
- *(tools/crush_info)* Add skill status

### 🐛 Bug Fixes

- *(schema)* Fix `crush.json` schema generation (#2574)
- *(ui)* Subscribe to `app.LSPEvent` instead of `workspace.LSPEvent` (#2565)
- *(app)* Derive shutdown context from `context.Background()` instead of cancelled `globalCtx` (#2242)
- *(lsp)* Clone slice to avoid mutation

### 📚 Documentation

- *(readme)* Mention `AVIAN_API_KEY` (#2557)

### 🎨 Styling

- Standardize log capitalization

### ⚙️ Miscellaneous Tasks

- *(skills/crush-config)* Improve crush config skill (#2556)
- *(legal)* @smlx has signed the CLA
- Add custom error and message for hyper unauthorized (#2577)
- Auto-update files
- Update unauthorize message
## [0.55.1] - 2026-04-06

### 🐛 Bug Fixes

- Support local models with unknown max_tokens and context window (#2554)

### ⚙️ Miscellaneous Tasks

- *(legal)* @Kartik33 has signed the CLA
- *(legal)* @talha7k has signed the CLA
- *(legal)* @avianion has signed the CLA
- *(legal)* @shahidshabbir-se has signed the CLA
## [0.55.0] - 2026-04-02

### 🚀 Features

- *(server)* Initial implementation of Crush RPC server/client
- Send server client version info
- Add swagger api docs generation
- Add generated swagger docs and serve them in the server
- Add AppWorkspace implementation of Workspace interface gated behind CRUSH_CLIENT_SERVER env var
- *(skills)* Builtin skills + skill disabling + crush-config builtin (#2466)

### 🐛 Bug Fixes

- Ensure proper resource cleanup and add retry logic for workspace creation
- Setup logging in server and client commands
- *(server)* Encode sessions in proto format
- *(client)* Add ListWorkspaces method to retrieve all workspaces from the server
- *(server)* Log to stderr if it's a terminal
- Update UI tests to use test workspace

### 🚜 Refactor

- Centralize user config path in a single func `home.Config()` (#2542)
- Modernize (#2548)
- *(server)* Move agent, session, permission, and event logic to backend package
- Rename Permissions.SkipRequests to Overrides().SkipPermissionRequests
- Use client and server for workspace access and management
- Migrate run and login commands to use client API instead of direct app access
- Update client methods to return proto types
- Simplify LSP diagnostic counts retrieval

### ⚙️ Miscellaneous Tasks

- *(legal)* @MrRolie has signed the CLA
- Auto-update files
- Add a workaround for slog calls in config.Load leaking to stderr
- Make lint green
- *(lint)* Fix yet another lint issue
- Auto-update files
## [0.54.0] - 2026-03-31

### 🚀 Features

- *(init-cmd)* Mention progressive disclosure (#1786)
- *(init)* Elicit control/data flow, arch descs (#1790)
- *(cli/session)* Show skill metadata (#2541)

### 🐛 Bug Fixes

- *(event)* Prevent panic on non-string telemetry keys (#2502)
- *(lsp)* Respect lsp auto start config (#2487)
- *(csync)* Use pointer receiver for JSONSchemaAlias (#2521)
- Do not commit `.crush/.gitignore` (#2531)
- Skip non-existent command dirs instead of creating them (#2534)
- *(taskfile)* Fix syntax error on task `run:onboarding`
- Conditionally show image keybindings based on model support (#2522)
- *(commands)* Add timeout context for MCP prompt retrieval (#2517)
- *(ls)* Respect git's `core.excludesfile` config if set (#2314)

### 🚜 Refactor

- *(ui)* Replace hardcoded cursor offset with style-based calculation (#2530)

### ⚙️ Miscellaneous Tasks

- *(lint)* Fix some small lint warnings (#2488)
- Auto-update files
- *(legal)* @Dexterity104 has signed the CLA
- *(legal)* @majiayu000 has signed the CLA
- *(legal)* @afsuyadi has signed the CLA
- *(legal)* @owldev127 has signed the CLA
## [0.53.0] - 2026-03-27

### 🚀 Features

- *(ui)* Variable height prompt input field (#2468)

### 🐛 Bug Fixes

- Enhance session title prompt to fix language issue (#2497)
- *(dialog)* Use valid key binding name for OAuth success state (#2491)
- Append error tool message when tool call is cancelled (#2492)
- *(diffview)* Use `udiff.Lines` and not `udiff.Strings`
- *(ui)* Always clean up external editor temp file (#2503)

### 📚 Documentation

- *(readme)* Add q&a section and note about clipboard support (#2509)

### 🧪 Testing

- *(diffview)* Update test fixtures

### ⚙️ Miscellaneous Tasks

- *(legal)* @fuleinist has signed the CLA
- *(legal)* @iceymoss has signed the CLA
- *(taskfile)* Add one more file to dependencies
- *(legal)* @srivilliamsai has signed the CLA
- Update catwalk
## [0.52.0] - 2026-03-25

### 🚀 Features

- Load project skills automatically from `.crush/skills` and more
- *(ui)* Improve file completion ranking algorithm
- *(ui)* Prioritize filename-based completion ranking
- *(bash)* Set `CRUSH=1`, `AGENT=crush` and `AI_AGENT=crush` (#2484)

### 🐛 Bug Fixes

- *(log)* Don't conflate body drain errors with HTTP errors
- On windows, also load skills from `$HOME/.config/crush/skills`
- Complete file picker dialog action (#2483)
- *(hyper)* Ensure it's possible to override with `HYPER_URL`
- *(system-prompt)* Tell the model to respond in the prompt spoken language
- Session titles should keep the same spoken language

### 📚 Documentation

- *(readme)* Update skill docs with updates and new paths
- *(ui)* Add comprehensive comments to completion ranking algorithm

### 🧪 Testing

- Re-record vcr cassettes after system prompt change

### ⚙️ Miscellaneous Tasks

- *(lint)* Modernize interfaces in the db package
- Update generated `.crush/.gitignore` to not ignore skills
## [0.51.3] - 2026-03-24

### 🐛 Bug Fixes

- Improve long text detection to account for long text in a single line (#2442)
- Reduce max read size from 5mb to 1mb for view and fetch tools (#2447)
- *(tui)* Remove duplicate ctrl+g help binding in FullHelp (#2465)

### ⚙️ Miscellaneous Tasks

- Attempt to fix dependabot
- *(legal)* @whatnick has signed the CLA
- *(legal)* @Alex-wuhu has signed the CLA
- Update cla action and bring back pr number
- *(legal)* @hongquan has signed the CLA in $pullRequestNo
- Update `google.golang.org/grpc` with security fix (#2464)
- *(labeler)* Add avian label
- *(legal)* @faelis has signed the CLA in $pullRequestNo
- Update ncruces/go-sqlite3 to v0.33.0 (#2461)
- *(legal)* @UnderLotus has signed the CLA in $pullRequestNo
- *(legal)* @malikwirin has signed the CLA in $pullRequestNo
- *(cla)* Revert commit title change
## [0.51.2] - 2026-03-20

### 🐛 Bug Fixes

- *(fantasy)* Fix copilot tool calls, fix ollama compatibility (#2444)
## [0.51.1] - 2026-03-19

### 🐛 Bug Fixes

- Update fantasy with a fix for avian (#2438)
- Remove deadcode
- *(mcp)* Handle raw/whitespace base64
- *(mcp/docker)* Harden tests
- *(ui/docker)* Stable sort mcp parameters
- *(mcp/docker)* Only write config after docker startup succeeds
- *(ui/docker)* Don't block ui when checking for docker desktop

### 📚 Documentation

- *(mcp/docker)* Comments

### ⚙️ Miscellaneous Tasks

- Gofumpt
## [0.51.0] - 2026-03-19

### 🚀 Features

- Add support for `--session` and `--continue` for the tui (#2422)
- Add aliases: `crush r` -> `crush run`, `crush s` -> `crush session` (#2424)
- Docker mcp integration (#2026)

### ⚙️ Miscellaneous Tasks

- *(legal)* @nghiant03 has signed the CLA
- *(taskfile)* Enhance `deps` task to work for recent releases
- Auto-update files
## [0.50.1] - 2026-03-17

### ⚙️ Miscellaneous Tasks

- Add support for gpt 5.4 mini and nano (#2419)
## [0.50.0] - 2026-03-16

### 🚀 Features

- Be able to continue non-interactive sessions (#2401)

### 🐛 Bug Fixes

- *(vercel)* Fix validation of api keys for vercel (#2415)
- *(ui)* Render nested tools consistently across states
- Keep original order when filtering models (#2416)

### 💼 Other

- Modernize new pointer creation (e.g. new(true))

### ⚙️ Miscellaneous Tasks

- *(legal)* @xulongwu4 has signed the CLA
- *(labeler)* Add vercel label
- *(events)* Add events for the CLI session series of commands (#2408)
## [0.49.0] - 2026-03-13

### 🚀 Features

- Add configurable auto-background threshold for bash tool (#2183)
- Allow to paste bedrock api keys (#2407)

### ⚙️ Miscellaneous Tasks

- Update openai sdk to v3 (#2406)
- *(ui)* Add toggle for transparent bg to command palette (#2405)
## [0.48.0] - 2026-03-12

### 🚀 Features

- Set user-agent to Charm Crush/<version> (#2357)
- *(notification)* Alert on turn completion and permission request (#1356)
- CLI-based session access and management (#2373)

### 🐛 Bug Fixes

- *(ui)* Format xhigh as XHigh (not Xhigh) (#2369)
- *(ui)* Properly truncate info message (#2379)
- *(events)* Panic when metrics are disabled
- *(noninteractive)* Actually use models to generate titles (#2372)
- *(events)* Remove redundant posthog exit event (#2371)

### 🚜 Refactor

- *(config)* Introduce ConfigStore and Scope for better config m… (#2395)

### 🧪 Testing

- Re-record test fixtures

### ⚙️ Miscellaneous Tasks

- Improve examples in `crush --help`
- *(agent)* Cleanup logic
- *(agent)* Allocate errors once, reuse errors
- *(legal)* @ZeitbyteRepo has signed the CLA
- Update AGENTS.md (#2388)
- Fix govulncheck (#2399)
- *(legal)* @seroperson has signed the CLA
- Update and improve ui/AGENTS.md
- Update user-agent string to better follow the convention
- *(ui)* Add command palette entry for toggling notifications (#2402)
## [0.47.2] - 2026-03-05

### 🐛 Bug Fixes

- Suppress message when clipboard is empty (#2361)

### ⚙️ Miscellaneous Tasks

- Add gpt 5.4 and gpt 5.4 pro (#2363)
## [0.47.1] - 2026-03-05

### 🐛 Bug Fixes

- *(lsp/edit)* Properly handle non-ascii chars (e.g. CJK) (#2325)

### ⚙️ Miscellaneous Tasks

- Add hyper to labeler
- Update hyper (#2354)
- Update golangci-lint to v2.10 and fix new issues (#2355)
- Update bubble tea to v2.0.1 (#2360)
## [0.47.0] - 2026-03-04

### 🚀 Features

- *(shell)* Add blocking wait option to job_output tool (#2189)

### 🐛 Bug Fixes

- *(lsp)* Fallback to Kill() on timeout (#2349)

### ⚙️ Miscellaneous Tasks

- Bump powernap to v0.1.2
- Update catwalk
## [0.46.2] - 2026-03-02

### 🐛 Bug Fixes

- *(lsp)* Treat adjacent ranges as non-overlapping per LSP spec (#2322)
- *(ui)* Follow scroll when at bottom (#2336)

### ⚙️ Miscellaneous Tasks

- *(legal)* @detro has signed the CLA
- *(legal)* @taoeffect has signed the CLA
- *(legal)* @vmfu has signed the CLA
## [0.46.1] - 2026-02-27

### 🐛 Bug Fixes

- Use typed context keys in tests to satisfy staticcheck
- *(tools/view)* Perform UTF-8 validity check only if read succeeds
- *(mcp)* Restore handling for unsupported resources/list method
- *(tools/view)* Fix view paging, test for edge cases

### 💼 Other

- Find references, double timeout

### 🚜 Refactor

- Simplify context value retrieval using generics
- Extract common sub-agent execution logic
- Use params struct for runSubAgent and add unit tests

### ⚡ Performance

- *(tools/view)* Avoid scanning entire file to count lines
- *(lsp)* Don't watch for changes when simply reading files
- *(tools/view)* Pause briefly for LSP diagnostics when viewing a file
- *(lsp)* Use shared timeout for parallel diagnostics collection

### ⚙️ Miscellaneous Tasks

- *(lint)* Don't shadow err vars
## [0.46.0] - 2026-02-26

### 🚀 Features

- Add minimax china provider (#2315)
- Add support for anthropic thinking effort (#2318)

### 🐛 Bug Fixes

- *(lsp)* Replace recursive fastwalk with filepath.Glob in root marker detection (#2316)

### ⚙️ Miscellaneous Tasks

- *(legal)* @mavaa has signed the CLA
- *(legal)* @aisk has signed the CLA
## [0.45.1] - 2026-02-25

### 🐛 Bug Fixes

- Initialize lsp manager callback to prevent nil pointer panic (#2307)
- *(lsp)* Fix multiple bugs in lsp client lifecycle and handlers (#2305)
- *(ui)* Truncate status messages that would otherwise wrap (#2306)
- *(event)* Guard against panic (#2310)

### ⚙️ Miscellaneous Tasks

- *(events)* Log when crush stats is called
- Update catwalk
- Notify me on winget prs
## [0.45.0] - 2026-02-24

### 🚀 Features

- Add support or gemini 3+ thinking levels

### 🐛 Bug Fixes

- *(agent)* Fix minor bugs in coordinator and view tool (#2276)
- Wrap correct error (#2296)
- *(app)* Fix goroutine leak, shutdown context, and model matching (#2298)
- *(mcp)* Gracefully handle Method not found for resources/list (#2239)

### ⚙️ Miscellaneous Tasks

- *(legal)* @Jaylonnet has signed the CLA
- Update `x/powernap` (#2295)
- Allow `MIT-0` dependencies (#2297)
- Update lip gloss, bubble tea and bubbles to v2.0.0 (#2299)
- Update fantasy to v0.10.0
- Update catwalk to v0.22.0
## [0.44.0] - 2026-02-22

### 🚀 Features

- *(ui)* Indicate when skills are loaded

### 🐛 Bug Fixes

- *(ui)* Apply the message style to each line of the message items
- *(ui)* Fix cases where compact header and footer bleed off screen (#2279)

### ⚙️ Miscellaneous Tasks

- *(legal)* @erikstmartin has signed the CLA
- *(styles)* Create specific styles for image loading indicators
## [0.43.4] - 2026-02-20

### ⚙️ Miscellaneous Tasks

- Fix release action, update go to 1.26
## [0.43.3] - 2026-02-20

### 🐛 Bug Fixes

- *(ci)* Allow Unlicense in dependency review action (#2259)
- *(lsp)* Prevent nil client from being stored in clients map (#2262)
- Use `Authorization` header for MiniMax (#2269)
- Proper validate io.net api keys (#2272)
- *(agent)* Pass correct model config to small provider builder (#2236)
- More reliably detect windows drive (#2273)

### 🚜 Refactor

- Simply code

### ⚙️ Miscellaneous Tasks

- Update to go 1.26 stable for govulncheck
## [0.43.2] - 2026-02-19

### 🐛 Bug Fixes

- Address `nil` pointer dereference panics on lsp client methods (#2256)
- *(ui)* Optimize assistant message rendering to improve performance (#2258)

### ⚙️ Miscellaneous Tasks

- *(cla)* Allow prs from github copilot
- *(cla)* Attempt without `[bot]` suffix
- *(cla)* Attempt `copilot`
- *(cla)* Try capital letter
## [0.43.1] - 2026-02-18

### 🐛 Bug Fixes

- Detect and stop tool call infinite loops (#2130) (#2214)
- *(ui)* Early exit AtBottom() when totalHeight exceeds viewport height
- *(ui)* Toggle pills to follow scroll (#2218)
- *(lsp)* Properly remove clients from map on stop/kill
- Make reasoning effort dialog smaller (#2247)
- *(deps)* Update `go-nativeclipboard` version to compile to RISC-V. (#2216)
- *(ui)* Cache blurred and focused renderings separately for assistant messages (#2252)

### 🚜 Refactor

- Simplify some code by using `cmp.Or` (#2253)

### ⚡ Performance

- Remove mutex from lsp manager

### ⚙️ Miscellaneous Tasks

- Update fantasy with fix for json schema in openai (#2221)
- *(legal)* @maxbrunet has signed the CLA
- *(legal)* @0xarcher has signed the CLA
- *(legal)* @julienrbrt has signed the CLA
- *(labeler)* Add minimax label
- Skip intermittent test on windows
## [0.43.0] - 2026-02-13

### 🚀 Features

- *(pills)* Add toggle todos/pills menu item (#2202)

### 🐛 Bug Fixes

- *(ui)* Completions offset for attachments row (#2208)
- Add io.net api key validation fallback

### 📚 Documentation

- *(readme)* Mention io.net provider

### ⚡ Performance

- Replace regex-based gitignore with glob-based matching (#2199)

### ⚙️ Miscellaneous Tasks

- *(legal)* @wallacegibbon has signed the CLA
- *(legal)* @PHPCraftdream has signed the CLA
- Update catwalk to v0.19.0
- Auto-update files
## [0.42.0] - 2026-02-11

### 🚀 Features

- *(lsp)* Show user-configured LSPs in the UI (#2192)

### 🐛 Bug Fixes

- *(mcp)* Cancel context on MCP session close to prevent leak (#2157)
- *(sqlite)* Increase busy timeout (#2181)
- *(lsp)* Files outside cwd (#2180)
- Clear regex cache on new session to prevent unbounded growth (#2161)
- *(config)* Correct Task agent ID in SetupAgents (#2101)
- Respect disable_default_providers (#2177)
- *(ui)* Correctly position cursor when attachments are present (#2190)
- *(ui)* Adjust sessions dialog size
- *(ui)* Truncate dialog titles with ellipsis
- *(ui)* Dialogs: loop around and scroll list when navigating with up/down keys
- *(ui)* Ensure the min size accounts for the dialog border
- *(grep)* Do not go outside cwd, add timeout (#2188)
- Ensure all providers are shown unless `disable_default_providers` is set (#2197)
- Address potential panic on shell command execution (#2200)
- Change binding to open/close todo list from `ctrl+space` to `ctrl+t` (#2201)

### 🚜 Refactor

- Use csync.Map for regex caches (#2187)

### 📚 Documentation

- Update LICENSE copyright
- *(readme)* Mention subscriptions (#2184)
- *(ui)* Comment typo

### ⚙️ Miscellaneous Tasks

- *(legal)* @portertech has signed the CLA
- Update fantasy (#2186)
- *(taskfile)* Avoid compiling when not needed
- *(taskfile)* Add `run:catwalk` task to run with local catwalk
- *(taskfile)* Add `modernize` task
- Run `modernize`
- *(taskfile)* Add `run:onboarding` to test onboarding flow
- Update `AGENTS.md`, mention `x/ansi` package
- Add `omitempty` back as `omitzero`
- Auto-update files
- Golangci-lint 2.9 (#2193)
- *(taskfile)* Add `-v` to `go build`
- *(taskfile)* Run binary with extension on windows
- Auto-update files
## [0.41.0] - 2026-02-09

### 🚀 Features

- Add ability to re-authenticate / edit api key

### 🐛 Bug Fixes

- *(ui)* Fix help wrapping on dialogs

### ⚡ Performance

- Track and start lsp on command (#2176)
## [0.40.0] - 2026-02-09

### 🚀 Features

- *(mcp)* Resources support (#2123)
- Add clipboard image paste support (ctrl+v) (#2148)
- *(lsp)* Start LSPs on demand, improve auto-start (#2103)

### 🐛 Bug Fixes

- *(ui)* Use plain letters for lsp status (#2121)
- Cap posthog shutdown timeout (#2138)
- *(mcp)* Race condition, logs (#2145)
- Make it possible to add api key for minimax
- Build linux/386 (#2153)
- Thinking sidebar (#2151)
- *(ui)* `highlighter` now modifies the cell directly. (#2171)
- *(ui)* Clear image cache when FilePicker closes to prevent unbounded memory growth (#2158)
- Prevent goroutine orphaning in mcp.Close() and shell.KillAll() (#2159)
- *(ui)* Prevent nil pointer in completions size update (#2162)
- Improving shutdown (#2175)

### 🚜 Refactor

- Remove global config (#2132)
- Remove empty slice declarations (#2150)
- Cleanup the code a bit
- *(lsp)* Use same handle file in client and manager (#2168)

### 📚 Documentation

- *(readme)* Add mention to minimax

### ⚡ Performance

- Timer leak in setupSubscriber (#2147)

### ⚙️ Miscellaneous Tasks

- Update fantasy to v0.7.1 (#2139)
- *(legal)* @francescoalemanno has signed the CLA
- Auto-update files
- Update `charm.land/catwalk` with minimax support
- Fix pure go build
- *(legal)* @biisal has signed the CLA
- *(legal)* @mishudark has signed the CLA
- Bump bubbletea and ultraviolet dependencies
## [0.39.3] - 2026-02-05

### 🐛 Bug Fixes

- *(ui)* List: ensure the offset line does not go negative when scrolling up
## [0.39.2] - 2026-02-05

### 🐛 Bug Fixes

- Change hyper url (#2120)
- *(ui)* Completions popup gets too narrow on single item (#2125)
- *(ui)* Fix bug preventing pasting text on windows (#2126)
- *(ui)* Api key dialog typo (#2131)
- *(ui)* Consistent box sizing (#2127)
- Hyper provider cancel (#2133)
- Realtime session file changes (#2134)

### 🚜 Refactor

- Remove old tui (#2008)

### ⚙️ Miscellaneous Tasks

- Update ui/agents.md (#2122)
- *(legal)* @inquam has signed the CLA
- *(legal)* @nickgrim has signed the CLA

### ◀️ Revert

- The width changes in #2127 (#2135)
## [0.39.1] - 2026-02-04

### 🐛 Bug Fixes

- *(ui)* Padding in the view (#2107)
- *(ui)* Cursor mispositioned when pasting large blocks of text in textarea (#2113)
- *(ui)* Context percentage updates (#2115)
- *(ui)* Ensure we anchor the chat view to the bottom when toggling (#2117)
- *(ui)* Only scroll to selected item if item collapsed

### ⚙️ Miscellaneous Tasks

- *(legal)* @zhiquanchi has signed the CLA
- Use OIDC for npm login (#2094)
## [0.39.0] - 2026-02-03

### 🚀 Features

- Add configurable timeout for LSP initialization (#2075)
- Release new ui refactor (#2105)
- *(ui)* Transparent mode (#2087)

### 🐛 Bug Fixes

- *(ui)* Fix permissions dialog rendering on small windows (#2093)
- *(ui)* Scroll to expanded item (#2088)
- *(ui)* Ensure `%d Queued` text is visible (#2096)
- *(styles)* Increase text contrast in active session deletion item
- Fix pasting files on some terminal emulators (#2106)

### ⚙️ Miscellaneous Tasks

- *(legal)* @acmacalister has signed the CLA
- *(style)* Add specific style for session rename placeholder
- *(styles)* Make rename style definitions match UI language
- Auto-update files
## [0.38.1] - 2026-02-02

### 🐛 Bug Fixes

- Address potential panic on initialization (#2092)
## [0.38.0] - 2026-02-02

### 🚀 Features

- Add support for vercel provider (#2090)

### 🐛 Bug Fixes

- *(ui)* Update layout and size after session switch
- *(ui)* Show auto-discovered LSPs (#2077)
- *(lsp)* Improve auto discovery (#2086)
- Ensure the commands and models dialogs render with borders (#2068)
- Ensure all tools work when behind a http proxy (#2065)
- *(lsp)* Improve lsp tools (#2089)

### 🚜 Refactor

- *(chat)* Handle double click & triple click (#1959)

### ⚙️ Miscellaneous Tasks

- *(legal)* @bittoby has signed the CLA
- *(legal)* @ijt has signed the CLA
- *(legal)* @khalilgharbaoui has signed the CLA
- Handle hyper config correctly (#2027)
## [0.37.0] - 2026-01-30

### 🚀 Features

- Implement prompt history (#2005)
- *(lsp)* Auto-discover LSPs (#1834)
- *(mcp)* Support server side instructions (#2015)
- Filetracker per session (#2033)
- Open commands dialog on pressing `/` (#2034)
- Allow to disable indeterminate progress bar (#2048)

### 🐛 Bug Fixes

- *(ui)* Ensure the message list does not scroll beyond the last item (#1993)
- Layout calculations when editor has attachments (#2012)
- Make the check for sidebar toggle inclusive (#2013)
- *(ui)* Use setState method to change UI state and focus (#1994)
- *(lsp)* Scope client to working directory (#1792)
- Schema incorrectly marks optional fields as required (#1996)
- Decouple thinking/reasoning from provider type (#2032)
- *(stats)* Resizing breaks pie charts (#2030)
- Make the commands dialog less taller (#2035)
- *(ui)* Fix selection of code blocks with tabs inside markdown (#2039)
- *(ui)* Fix wrong color on selected item info on dialogs (#2041)
- Respect disabled indeterminate progress bar setting on app start (#2054)
- Improve logs, standardize capitalized (#2047)
- Allow HYPER_URL with embedded provider (#2031)
- Panic when matching titles in session dialogue
- Slice string at the grapheme level, not byte level
- *(ui)* Typo in ListItemStyles type name
- Do not scroll to bottom if user has scrolled up (#2049)
- *(ui)* Switch focus on click (#2055)
- *(ui)* Arrow navigation wasnt working when todo view is open (#2052)
- *(posthog)* Check correct error; prevent panic (#2036)

### 🚜 Refactor

- Terminal capability handling (#2014)

### 📚 Documentation

- Improve clarity and fluency of mandarin tagline (#2022)

### ⚙️ Miscellaneous Tasks

- Fix typo on `crush stats` html page
- Auto-update files
- Format nix (#2009)
- *(legal)* @oug-t has signed the CLA
- *(legal)* @liannnix has signed the CLA
- Use goreleaser nightly on snapshot build
- Auto-update files
- *(ui)* String efficiency
- `chmod +x scripts/check_log_capitalization.sh`
- Update catwalk and its import paths to `charm.land/catwalk`
## [0.36.0] - 2026-01-27

### 🚀 Features

- Delete sessions (#1963)
- Crush stats (#1920)
- Update session title (#1988)
- Add ability to drag & drop multiple file at once + support windows (#1992)

### 🐛 Bug Fixes

- Should also copy on `y` (additionally to `c`) (#1989)
- Enable left/right scrolling of diff (#1984)
- *(agent)* Read step data for summarization check (#1787)
- Stats chart don't account for cached tokens
- Token calculation (#2004)

### 🚜 Refactor

- Use `ContainsAnyOf` from `x`

### 📚 Documentation

- *(readme)* Add SYNTHETIC_API_KEY (#1971)
- *(readme)* Add Z.ai API key info

### ⚙️ Miscellaneous Tasks

- Bump glamour to v2.0.0-20260123212943-6014aa153a9b
- Fix snapshot goreleaser dist
- *(legal)* @billycao has signed the CLA
- *(legal)* @gdamjan has signed the CLA
- *(labeler)* Add automation for `area: crush run` label
- Use goreleaser nightly (#1987)
- *(goreleaser)* Fix aur_sorces build to properly set the version (#1978)
## [0.35.0] - 2026-01-23

### 🚀 Features

- Lsp_restart (#1930)
- *(ui)* Add keybinding to copy chat message content to clipboard (#1947)
- Implement onboarding flow on the new ui codebase

### 🐛 Bug Fixes

- Lsp sort
- Tab to chat only when in chat
- Handle new session when focused on the list
- Make sure we have a fresh model/tools on each call
- Add back suspend
- *(ui)* Models: ensure select loop breaks correctly and scroll to top on filter
- *(ui)* Only copy chat highlight when we have highlighted content
- Route mouse events to the dialog if its showing (#1953)
- Commands height (#1954)
- Permission notification (#1955)
- Completions width (#1956)
- New/update message behavior (#1958)
- *(ui)* Rework cursor can appear out of place on multi-line (#1948)
- *(ui)* Prevent AAAA probe bleed in terminals without Kitty graphics support (#1967)
- *(list)* Prevent panic due to negative index
- *(dialogs)* Prevent panic due to negative index
- Ensure hyper is the first provider in the list

### 🚜 Refactor

- *(ui)* Enable initialize
- *(ui)* Add references tool (#1940)
- Use different ansi image library (#1964)
- Rename `uiConfigure` to `uiOnboarding`

### 📚 Documentation

- Add vercel ai gateway to readme (#1951)

### ⚙️ Miscellaneous Tasks

- Do not scroll sessions if not neccessary
- Change the dialog sizes a bit
- *(legal)* @huaiyuWangh has signed the CLA
- *(legal)* @akitaonrails has signed the CLA
- *(legal)* @mcowger has signed the CLA
- *(legal)* @jerilynzheng has signed the CLA
- Goreleaser build --snapshot on every commit to main (#1910)
- *(legal)* @AnyCPU has signed the CLA
## [0.34.0] - 2026-01-21

### 🚀 Features

- *(ui)* Filepicker: support kitty graphics in tmux
- *(ui)* Filepicker: support kitty graphics in tmux (#1884)
- *(ui)* Use the new UI behind a feature flag
- Increase paste lines as attachment threshold (#1936)
- *(ui)* Copy chat highlighted content to clipboard
- Crush run --model, and crush models (#1889)
- *(ui)* Show working directory in window title

### 🐛 Bug Fixes

- *(ui)* Filepicker: simplify tmux kitty image encoding
- Correct spelling from marshall to marshal
- Correct typo in log message
- Recent models dont go into the schema (#1892)
- Use strconv.ParseBool
- *(sec)* Do not output resolved command (#1934)
- *(ui)* Increase paste lines threshold (#1937)
- *(ui)* Implement Highlightable interface for message items
- *(ui)* List: move focused logic to render callback

### 🚜 Refactor

- Add assistant info item (#1881)
- Pills section (#1916)
- Move domain-specific type out of generic pubsub package
- Mcp tool item (#1923)
- *(ui)* Reimplement UI components (#1652)

### 📚 Documentation

- *(readme)* Remove link to site that is offline (#1926)

### ⚙️ Miscellaneous Tasks

- *(prompt)* Small edits (#1915)
- *(tools/edit)* Various fixes (#1921)
- Remove dead (duplicate) code (#1913)
- Remove unnecessary testing concerns from env.New
- Return empty slices instead of nil for safety
- Remove redundant zero value initialization
- Simplify struct initialization to direct return
- Add log when new UI is enabled
## [0.33.3] - 2026-01-18

### 🐛 Bug Fixes

- Don't build native clipboard for ios
- Do not wait for MCP on interactive mode
## [0.33.2] - 2026-01-16

### 🐛 Bug Fixes

- Add freebsd support for clipboard on 386 architecture
- Skip native clipboard support on unsupported platforms

### ⚙️ Miscellaneous Tasks

- Bump purego to v0.10.0-alpha.3.0.20260115160133-57859678ab72
## [0.33.1] - 2026-01-16

### 🐛 Bug Fixes

- Editor: exclude native clipboard support from linux/386 builds (#1903)
## [0.33.0] - 2026-01-16

### 🚀 Features

- Add clipboard image paste functionality to chat editor (#181) (#1151)
- Implement hyper oauth flow in the new ui codebase
- Implement github copilot oauth flow in the new ui codebase

### 🐛 Bug Fixes

- *(ci)* Security: allow Google Patent License for Go modules
- *(ci)* Update security workflow to use setup-go and install govulncheck
- Race in agent.go (#1853)
- Ensure that hyper and copilot models show up even if not configured
- Address "verifying..." now showing on side of spinner
- Align "authentication successful" on the left
- Address double vertical margins between sections
- *(ui)* Filepicker: remove redundant Init method and Action type
- *(ui)* Filepicker: simplify cmd return
- Resolve extra headers for providers (#1764)
- Make hyper and copilot link styled on ui (#1872)
- Try to make the search tool more reliable (#1779)
- Mcps loading in non interactive mode (#1894)

### 🚜 Refactor

- *(ui)* Dialog: cleanup render logic and use RenderContext (#1871)
- Make oauth dialog generic and move provider logic to interface
- Remove init in favor of returning cmd on new
- Remove duplication on functions to open dialogs
- Reasoning dialog (#1880)

### 📚 Documentation

- *(readme)* Update features section
- *(readme)* Remove extra comma

### ⚙️ Miscellaneous Tasks

- *(README)* Update crush art (#1861)
- *(legal)* @kuxoapp has signed the CLA
- *(legal)* @mhpenta has signed the CLA
- Implement missing command and fix summarize (#1882)
- Tidy agent package (#1857)
- Fix more typos (#1863)
- Remove duplicated log
- Update `posthog-go` to latest version
- Use posthog's default exception reporting
## [0.32.1] - 2026-01-13

### 🚀 Features

- Implement api key input dialog on new ui codebase (#1836)
- *(ui)* Dialog: add file picker dialog with image preview
- *(ui)* Filepicker: add image attachment support with preview

### 🐛 Bug Fixes

- Update fantasy with panic fix for google gemini (#1840)
- *(ui)* Dialog: saveKeyAndContinue should return Action
- *(ui)* Filepicker: defer image preview until after transmission
- Race condition where title might not be generated (#1844)

### 🚜 Refactor

- Compact mode (#1850)

### ⚙️ Miscellaneous Tasks

- *(legal)* @jeis4wpi has signed the CLA
- *(legal)* @uppet has signed the CLA
- Fix some typos
- *(legal)* @andreasdotorg has signed the CLA
- *(sec)* Add more security jobs, improve build, enable race detector (#1849)
- Fix govulncheck
## [0.32.0] - 2026-01-12

### 🚀 Features

- Completions menu (#1781)
- Attachments (#1797)
- Open editor in the right position (#1803)
- Paste as file (#1800)
- Allow to send the prompt if its empty but has text attachments (#1809)
- Open editor in the right position (#1804)
- Allow to send the prompt if its empty but has text attachments (#1806)
- Add `disable_default_providers` option (#1675)

### 🐛 Bug Fixes

- *(ui)* Adjust app and help area margins
- *(models)* Ensure that we show unknown providers on the list
- *(ui)* Dialogs: ensure returned commands are executed
- *(ui)* Completions: simplify completions popup message handling
- *(ui)* Completions: simplify Close method
- *(ui)* Simplify suffix handling in message editor
- *(ui)* Dry up up/down key binding in dialog commands
- *(ui)* Completions: load files asynchronously
- *(ui)* Dry up session dialog key bindings
- *(sqlite)* Busy timeout (#1815)
- Use cached lsp.DiagnosticCounts (#1814)
- Make sure to unlock in goroutine (#1820)

### 🚜 Refactor

- Use `cmp.Or`
- *(ui)* Dialog: message and draw handling (#1822)

### 📚 Documentation

- *(README)* Add FreeBSD installation instructions

### ⚡ Performance

- Reduce memory usage (#1812)
- *(shell)* Reduce allocations in updateShellFromRunner (#1817)
- Use strings.Builder for string concatenation in loops (#1819)
- Fix possibly unclosed resp.body (#1818)
- *(config)* Simplify loadFromConfigPaths (#1821)
- Improve startup and shutdown speed (#1829)

### ⚙️ Miscellaneous Tasks

- Fix const type
- *(legal)* @mohaanymo has signed the CLA
- *(legal)* @zyriab has signed the CLA
- *(legal)* @aleksclark has signed the CLA
- Auto-update files
## [0.31.0] - 2026-01-07

### 🚀 Features

- *(ui)* Status: add status bar with info messages and help toggle
- *(skills)* Also load from .config/agents (#1755)
- Remove claude code support (#1783)

### 🐛 Bug Fixes

- *(posthog)* Normalize `interactive` prop case
- *(posthog)* Correct bool prop name for non-interactive mode (#1771)
- Mark files that are attched as read (#1777)
- *(mcp)* Centrally filter disabled tools (#1622)

### 🚜 Refactor

- Add more tools
- Add agent tools and rename simple->compact

### ⚙️ Miscellaneous Tasks

- Some cleanup
- Rename isSpinning to spinning
- Fix typo in const name
## [0.30.3] - 2026-01-05

### 🐛 Bug Fixes

- Quota for subagents in copilot
- Enable responses api for github copilot

### ⚙️ Miscellaneous Tasks

- Make it so the small model also is always considered a subagent
- *(copilot)* Update message: "wait a minute" -> "wait 5 minutes"
- Update fantasy to v0.6.0
## [0.30.2] - 2026-01-04

### 🐛 Bug Fixes

- *(sessions)* Nil pointer dereference (#1759)
## [0.30.1] - 2026-01-03

### 🚀 Features

- Agent skills (#1690)
- *(ui)* Model dialog: implement model selection handling
- *(ui)* Models dialog: filter by provider and model name
- *(ui)* Models dialog: improve group filtering by ignoring spaces

### 🐛 Bug Fixes

- *(tui)* Guard model selection when list is empty (#1715)
- *(sessions)* Select the current session in dialog
- *(ui)* Do not allow summarizing if agent is busy
- *(ui)* Model dialog: skip non-model items when navigating selection
- *(ui)* List: countLines should return 1 for empty strings
- *(ui)* Models dialog: filter each group separately
- *(sessions)* Tag removal, handle multibyte
- *(sessions)* Generate title with large model if small model fails
- Copilot quota handling (#1738)

### 🚜 Refactor

- *(skills)* Use fastwalk to resolve symlinks (#1732)

### ⚙️ Miscellaneous Tasks

- *(legal)* @yuguorui has signed the CLA
- *(legal)* @aeroxy has signed the CLA
- Auto-update files
- Minor internal/app improvements (#1696)
- *(ui)* Add TODO for model API and auth validation
- *(legal)* @nikolayk812 has signed the CLA
- *(cli)* Simplify help text (#1752)
- Fix aur on arm64 (#1739)
## [0.29.1] - 2025-12-22

### 🐛 Bug Fixes

- Prevent filename insertion when dragging attachments (#1683)

### ⚙️ Miscellaneous Tasks

- *(legal)* @Mr777x-enf has signed the CLA
## [0.29.0] - 2025-12-19

### 🚀 Features

- Paste long content as an attachment (#1634)

### 🐛 Bug Fixes

- Remove unsupported image types
- Race condition (#1649)
- Initial api key load (#1672)
- Splash padding y (#1680)
- *(aws-bedrock)* Update fantasy with `panic` fix (#1681)

### ⚙️ Miscellaneous Tasks

- *(legal)* @flatsponge has signed the CLA
- *(legal)* @jonhoo has signed the CLA
- Auto-update files
- Auto-update files
## [0.28.0] - 2025-12-18

### 🚀 Features

- Add ability to set envs to override config directories (#1661)
- *(ui)* Model selection dialog
- Make hyper and copilot login work on onboarding

### 🐛 Bug Fixes

- *(ui)* Dialog: clarify session age display
- *(tui/mcp)* Singularize tool/prompt count when 1 (#1623)
- *(ui)* Remove redundant check in thinking rendering
- *(ui)* Improve thinking message truncation display
- *(ui)* Properly align session item age text
- *(ui)* Dialog: show shortcut/info in list items
- *(ui)* Dialog: show provider name for recent models
- *(ui)* Dialog: no need for groupItems map in ModelsList.VisibleItems
- *(ui)* Handle model selection in models dialog
- *(ui)* Dialog sessions refactor and close commands dialog on select
- Handle `ctrl+c` in hypercrush auth (#1665)
- *(onboarding)* Address `c` key press not working on onboarding (#1663)
- *(copilot)* Change import to happen on demand, and also on onboarding

### 🚜 Refactor

- *(config)* Add ToProvider method to ProviderConfig
- *(chat)* Rename GetMessageItems to ExtractMessageItems and GetToolRenderer to ToolRenderer
- *(ui)* Chat: abstract tool message rendering
- *(ui)* Dialog: unify session and model switching dialogs
- *(chat)* Add bash related tools
- *(chat)* Add file tools
- *(chat)* Add search tools

### ⚡ Performance

- *(ui)* Dialog: preallocate slice for filterable items in ModelsList

### ⚙️ Miscellaneous Tasks

- *(chat)* Rename funcs
- Auto-update files
- Auto-update files
- Remove cspell.json (#1664)
- *(chat)* Remove empty assistant messages
- *(ui)* Standardize truncation message format
- *(ux)* Remove extra uneeded line
## [0.27.0] - 2025-12-17

### 🚀 Features

- Add github copilot support
- Add github copilot auth flow via tui

### 🚜 Refactor

- *(hyper)* A couple of small adjustments

### ⚙️ Miscellaneous Tasks

- *(cmd/projects)* Improve help text (#1647)
- Auto-update files
- Auto-update files
## [0.26.0] - 2025-12-16

### 🚀 Features

- *(ui)* Wip: basic chat message sending
- *(ui)* Dialog: wrap navigation from last to first and vice versa
- *(chat)* Expandable thinking for assistant
- *(ui)* List: expose filterable items source type and return values for selection methods
- Hyper integration (#1642)

### 🐛 Bug Fixes

- *(dialog)* Commands: execute command handlers properly
- *(ui)* Editor: show yolo prompt correctly
- Check if token is expired before sending the request (#1641)
- *(ui)* OpenEditor and handle pasted files in editor
- *(ui)* Improve key handling and keybindings for chat/editor
- *(ui)* Dialog: align radio buttons with checkboxes
- *(ui)* Adjust dialog sizing to account for dynamic title and help heights
- *(chat)* Only spin when there is no and no tool calls
- *(chat)* Race condition
- *(chat)* Reset index and paused animations
- *(chat)* Do not mark tools with results as canceled
- *(ui)* Ensure MCPs are displayed in configured order

### 🚜 Refactor

- *(ui)* Dialog: improve command and session dialogs
- *(ui)* Dialog: rename Add/Remove to Open/Close
- *(ui)* Rename files.go to session.go and update session loading logic
- *(tui)* Move ExecShell to uiutil package
- *(chat)* Implement user message (#1644)
- *(chat)* Simple assistant message
- *(chat)* Only show animations for items that are visible
- *(chat)* Initial setup for tool calls
- *(chat)* Hook up events for tools

### ⚙️ Miscellaneous Tasks

- Bump bubbletea to fix moving the cursor to the bottom of screen
- Update bubbletea to latest rc2 version
- Bump bubbletea to latest rc2 version
- *(chat)* Add some missing docs
- *(chat)* Remove unused style
- *(chat)* Some more docs missing
- *(chat)* Small improvements
## [0.25.0] - 2025-12-16

### 🚀 Features

- Add centralized project tracking (for Splitrail) (#1553)
- *(ui)* Wip: add commands dialog to show available commands

### 🐛 Bug Fixes

- Typo in Subscriber type name (#1616)
- *(deps)* Use charm.land glamour and log (#1469)
- *(ui)* Simplify QuitDialogKeyMap by embedding key bindings directly
- Remove 100MB file size limit from download tool (#1631)
- *(uiutil)* Add Cursor interface
- Update fantasy with fix for empty messages bug (#1639)

### 🚜 Refactor

- *(ui)* Simplify dialog model and rendering
- *(list)* Simplify focus and highlight interfaces
- *(ui)* List: remove mouse highlighting state, add render callbacks
- *(tui)* Move command loading to uicmd package
- *(tui)* Unify commandType definition
- *(tui)* Move UI message handling to internal/uiutil
- *(ui)* Use uiutil for command handling and error reporting

### ⚙️ Miscellaneous Tasks

- *(legal)* @strawberry-code has signed the CLA
- *(ui)* Standardize dialog identifiers
- *(ui)* Remove left padding in thinking blocks (#1629)
- Send `interactive` attribute (for `crush run`) (#1635)
- Search improvements (#1632)
## [0.24.0] - 2025-12-12

### 🚀 Features

- Todo tool
- *(ui)* New session selector dialog

### 🐛 Bug Fixes

- Edit tool error on new file creation due to logical fallthrough (#1566)
- *(tui)* Fix list wrap behave when it has unfocusable items (#1312)
- *(noninteractive)* Support output redirection (aka, pipes) (#1594)
- *(ui)* List highlight selection scrolling (#1575)

### 🚜 Refactor

- *(ui)* Dialog: use Action pattern and lipgloss layers

### ⚙️ Miscellaneous Tasks

- *(legal)* @nonsleepr has signed the CLA
- Update bubbletea to latest rc2 version
- Bump bubbletea and ultraviolet to fix non en lang rendering
## [0.23.0] - 2025-12-11

### 🚀 Features

- *(tui)* Show / hint when editor is empty (#1512)
- *(mcp-config)* Add list of disabled_tools (#1533)
- Update fantasy and catwalk for support to gpt 5.2 and 5.1 codex max

### 🐛 Bug Fixes

- *(ui)* Use new lipgloss compositor
- *(ui)* List: prevent negative offset in list rendering

### 🧪 Testing

- Fix local test suite, ignore machine environ (#1605)

### ⚙️ Miscellaneous Tasks

- Bump lipgloss to the latest v2 version
- Bump bubbletea and ultraviolet to fix kitty keyboard exiting
- Bump bubbletea to module replace
- Bump bubbletea and ultraviolet to fix cursor related rendering artifacts
- Bump bt and uv to fix scroll empty line artifact bug
- Skip two flacky tests depending on os (#1606)
- Bump dependencies
- *(taskfile)* Add task to update fantasy and catwalk
## [0.22.2] - 2025-12-10

### 🚀 Features

- *(ui)* Initial chat ui implementation
- *(ui)* Chat: add navigation and keybindings
- *(ui)* Add common utils and refactor chat message items and tools
- *(ui)* Initial lazy-loaded list implementation
- *(ui)* Optimize ScrollToBottom in lazylist
- *(ui)* Invalidate rendered items on focus/blur
- *(ui)* Add text highlighting support to lazylist
- *(ui)* Add mouse click handling to lazy list items

### 🐛 Bug Fixes

- Typo choise -> choice (#1579)
- *(ui)* Move canvas initialization to View method
- *(ui)* List: scroll to bottom after session load
- *(ui)* Fix scrolling up last item
- *(ui)* Scroll down off by one
- *(ui)* Correct scrolling up behavior in lazy list
- *(ui)* Dry highlighting items in lazylist
- *(noninteractive)* Cancel on signal (#1584)
- *(claude)* Add authentication refresh on 401 errors (#1581)

### 🚜 Refactor

- *(list)* Remove item IDs for simpler API
- *(ui)* Remove dirty tracking from LazyList
- *(ui)* Cleanup and remove unused list code
- *(ui)* Rename lazylist package to list and update imports

### ⚡ Performance

- Improve startup, specifically provider updates (#1577)

### ⚙️ Miscellaneous Tasks

- *(legal)* @mengwong has signed the CLA
## [0.22.1] - 2025-12-08

### 🐛 Bug Fixes

- Update fantasy to fix image loading
## [0.22.0] - 2025-12-08

### 🚀 Features

- Support image results from tools (#1549)
- Add web search (#1565)

### 🐛 Bug Fixes

- *(list)* Cap rendered filterable list while keeping full search set (#1492)
- Faster shutdown (#1570)
- Ignore mouse clicks when `isProjectInit` is active (#1561)
- Use visual width instead of byte length for text truncation (#1562)
- Update MultiEdit permission desc to reflect actual applied edits (#1564)
- Prevent crash when pressing ctrl+f on model selector (#1573)
- Prevent nil pointer dereference when updating agent model  (#1560)

### ⚙️ Miscellaneous Tasks

- *(legal)* @mike1858 has signed the CLA
- Fix fantasy version
- *(legal)* @Guxinpei has signed the CLA
- *(legal)* @mikluko has signed the CLA
- *(legal)* @Gustave-241021 has signed the CLA
## [0.21.0] - 2025-12-04

### 🚀 Features

- Show progress bar on iterm2 and rio
- Allow tools to run in parallel (#1543)

### 🐛 Bug Fixes

- Add ValidArgs and Args to login cmd (#1540)
- *(ui)* Hide cursor when editor is not visible
- *(ui)* Change pointer receivers to value receivers
- *(ui)* Change UI model receiver to pointer back

### 🚜 Refactor

- Use contains any

### ⚙️ Miscellaneous Tasks

- Bump ultraviolet to fix rendering issues with unnecessary cursor movements and erasing line background color
- *(labeler)* Add copilot label
- Go mod tidy
- *(legal)* @marifcelik has signed the CLA
- Improve prompt (#1502)
- Improve zai tool calls (#1551)
## [0.20.1] - 2025-12-01

### ⚙️ Miscellaneous Tasks

- Fix cosign (#1538)
## [0.20.0] - 2025-12-01

### 🚀 Features

- *(ui)* Add optimized list component with focus navigation
- Implement text highlighting in list items (#1536)
- Add `crush login claude` command (#1537)

### 🐛 Bug Fixes

- Remove `max_tokens` minimum requirement to fix json schema issue (#1532)
## [0.19.4] - 2025-12-01

### 🐛 Bug Fixes

- Fix `c` key not working in model filter (#1534)
- *(editor)* Fix opening `$EDITOR` w/ and w/o args (#1520)

### ⚙️ Miscellaneous Tasks

- *(legal)* @masroor-ahmad has signed the CLA
- *(legal)* @thezbm has signed the CLA
## [0.19.3] - 2025-11-29

### 🐛 Bug Fixes

- Refresh oauth token in the background
## [0.19.2] - 2025-11-26

### ⚙️ Miscellaneous Tasks

- Update bubbletea to fix restoring terminal first render
## [0.19.1] - 2025-11-26

### 🐛 Bug Fixes

- *(ui)* Prevent panic when session is nil
- *(editor)* Fix opening `$EDITOR` when it contains arguments (#1481)
- Fix `h` and `l` keys not working on models filter
- *(claude)* Simplify code and fix potential unauthorized error
## [0.19.0] - 2025-11-26

### 🚀 Features

- Add support for claude code max (#1514)

### 🐛 Bug Fixes

- Use canvas to render UI view

### ⚙️ Miscellaneous Tasks

- Initialize functionality & landing page
- Add chat sidebar (#1510)
- Auto-update generated files
## [0.18.6] - 2025-11-24

### ⚙️ Miscellaneous Tasks

- *(legal)* @heimoshuiyu has signed the CLA
- Swap error related logs to use error log level (#1505)
- Update bubbletea to use string view content
## [0.18.5] - 2025-11-21

### 🚀 Features

- *(ui)* Implement tea.Layer Draw for UI model

### 📚 Documentation

- *(readme)* Update "aws bedrock" to "amazon bedrock" (#1478)

### ⚙️ Miscellaneous Tasks

- Add gopls settings from Carlos's dotfiles (#1424)
- *(ui)* Adjust status notification details (#1490)
## [0.18.4] - 2025-11-19

### 🐛 Bug Fixes

- Detect version for `go install ...@main` (#1476)
- Kimi coding api key validation (#1477)

### ⚙️ Miscellaneous Tasks

- *(legal)* @micahwalter has signed the CLA
## [0.18.3] - 2025-11-19

### 🚀 Features

- *(ui)* Simplify editor and embed into main UI model
- Notify about new crush versions (#361)
- Show a different message if crush is built from source

### 🐛 Bug Fixes

- *(ui)* Use Content instead of Layer for main view
- Add missing openai-compat in schema (#1461)
- Don't notify update available when running local build (#1465)
- Handle google reasoning (#1474)

### 🚜 Refactor

- *(ui)* Pass app.App to common.Common and access config via common
- *(ui)* Rework init and support different layouts (#1463)

### 🎨 Styling

- Small code style updates

### ⚙️ Miscellaneous Tasks

- Bump bubbles to v2.0.0-rc.1 and update textarea and help width usage
- Auto-update generated files
## [0.18.2] - 2025-11-17

### 🐛 Bug Fixes

- *(mcp)* Always call `ListTools` to discover available tools (#1447)

### ⚙️ Miscellaneous Tasks

- Bump bubbletea and ultraviolet to support mode 2026
- *(legal)* @Iflgit has signed the CLA
- *(legal)* @iainlane has signed the CLA
- *(attribution)* Default to assisted-by + email (#1444)
- *(setup)* Model chooser language copyedit
- Bump bubbletea and ultraviolet to reduce resize tearing (#1460)
## [0.18.1] - 2025-11-13

### 📚 Documentation

- *(readme)* Add initialize_as (#1438)

### 🧪 Testing

- Record vcr cassettes

### ⚙️ Miscellaneous Tasks

- Update fantasy & add support for gpt-5.1 (#1439)
## [0.18.0] - 2025-11-13

### 🚀 Features

- *(config)* Default to AGENTS.md w/ new setting (#1403)

### 🐛 Bug Fixes

- *(test)* Set a fixed attribution to avoid system prompt mismatch

### 🧪 Testing

- Migrate tests to `charm.land/x/vcr`

### ⚙️ Miscellaneous Tasks

- Auto-update generated files
## [0.17.0] - 2025-11-12

### 🚀 Features

- Recent models section in picker (#1374)
- *(config)* Add trailer_style option
- *(config)* Migrate deprecated co_authored_by

### 🐛 Bug Fixes

- Append `ImageURLContent` part in `unmarshallParts` (#1387)
- *(mcp)* Client being killed (#1419)
- Recent models test
- *(bash)* Use model name instead of ID

### 🚜 Refactor

- *(bash)* Make whitespace conform to spec

### 📚 Documentation

- *(readme)* Move "local models" section to be under "custom providers" (#1402)

### 🧪 Testing

- Record vcr cassettes for #1379

### ⚙️ Miscellaneous Tasks

- *(legal)* @novalis78 has signed the CLA
- *(task)* Fetch tags before installing
- *(legal)* @alewtschuk has signed the CLA
- Update fantasy to v0.2.1
- *(taskfile)* Add task to record all cassettes
- Record all cassettes after fantasy update
## [0.16.1] - 2025-11-07

### 🐛 Bug Fixes

- Improve ux for presenting errors from providers (#1388)

### ⚙️ Miscellaneous Tasks

- Upgrade dependencies and fix related code (#1404)
## [0.16.0] - 2025-11-07

### 🚀 Features

- Background jobs & remove persistent shell (#1328)

### ⚙️ Miscellaneous Tasks

- Remove flaky tests
## [0.15.2] - 2025-11-06

### 🚀 Features

- *(ui)* Ctrl+l/ctrl+m to open model switcher (#1274)

### 🐛 Bug Fixes

- *(sqlite)* Remove WAL pragma (#1280)

### ⚙️ Miscellaneous Tasks

- *(legal)* @nanvenomous has signed the CLA
- Upgrade bubbletea to fix rendering issue #1389
## [0.15.1] - 2025-11-05

### ⚙️ Miscellaneous Tasks

- *(legal)* @danielmerja has signed the CLA
- Fix: upgrade dependencies and fix rendering on non-tc terminals
## [0.15.0] - 2025-11-05

### 🚀 Features

- *(tui/chat)* Use @ for files, / for commands (#1377)
- Agentic fetch tool (#1315)

### ⚙️ Miscellaneous Tasks

- *(legal)* @LarsArtmann has signed the CLA
## [0.14.0] - 2025-11-04

### 🚀 Features

- Show progress bar on boot for feedback (#1371)
- *(mcp)* Refactor, support prompts

### 🐛 Bug Fixes

- *(lint)* Don't shadow 'env' variable
- Rethink global config path on windows: use `$HOME/.config` (#1352)
- *(noninteractive)* Always print newline after output
- *(noninteractive)* Strip leading newline from assistant responses
- *(noninteractive)* Spinner text on light backgrounds
- *(catwalk)* Improve fetch logging message
- *(mcp)* Tool/prompt list update
- Missing handle call
- Improve code
- Handle delayed mcp init

### 🚜 Refactor

- Simplify `home.Dir()` (#1353)
- *(agent)* Clarify user-initiated cancellations and denials (#1368)
- *(catwalk)* Remove stale helper to avoid confusion
- *(mcp)* Some more decoupling

### 📚 Documentation

- Add comment and commit notes to CRUSH.md
- Add some package-level GoDoc comments
- *(readme)* Improve mcp example to have the real github url (#1378)

### ⚙️ Miscellaneous Tasks

- On schema update commits, use @charmcli as the author (#1344)
- *(legal)* @heynemann has signed the CLA
- *(legal)* @niklasschaeffer has signed the CLA
- Switch from bar cursor to blocky cursor
- *(noninteractive)* Accept a writer for output
## [0.13.7] - 2025-10-31

### 🚀 Features

- Keep successful edits even if some fail (#1327)
- *(ui)* Add initial sidebar component with logo

### 🐛 Bug Fixes

- Address panic due to possible `nil` map (#1348)

### ⚙️ Miscellaneous Tasks

- *(.gitignore)* Ignore `/tmp` dir
- Gramatical edit (#1350)
- Improve azure support (#1351)
## [0.13.6] - 2025-10-30

### 🐛 Bug Fixes

- Fix broken tools on `crush run` (#1333)
- *(ux)* Fix crush logo flicker on window resizes (#1338)
- Upgrade bubbletea to fix rendering issues (#1340)
- Only enable built-in core utils by default on windows

### 🚜 Refactor

- *(shell)* Some small code adjustments

### ⚙️ Miscellaneous Tasks

- *(legal)* @teras has signed the CLA
## [0.13.5] - 2025-10-29

### 🐛 Bug Fixes

- *(tui)* Make permissions dialog height responsive to resize (#877)
- *(tui)* Disable progress bar for unsupported terminals (#1329)

### ⚙️ Miscellaneous Tasks

- Add vcr cassette files to `.gitignore` (#1332)
## [0.13.4] - 2025-10-29

### 🐛 Bug Fixes

- Openrouter key validation

### ⚙️ Miscellaneous Tasks

- *(legal)* @bradflaugher has signed the CLA
## [0.13.3] - 2025-10-28

### 🚀 Features

- *(ui)* Add basic dialog and new UI structure
- *(ui)* Restructure UI package into subpackages
- *(ui)* Add chat and editor models with keybindings

### 🐛 Bug Fixes

- *(test)* Normalize file separator to make vcr cassette match
- *(tools)* Fix handling of abs paths on windows
- Add prefix if exists in all calls (#1311)
- *(ui)* Lint: use consistent receiver names in EditorModel methods
- *(ui)* Accurately calculate help and main height in UI layout
- *(ui)* Don't use separate showFullHelp field
- Update fantasy and catwalk with fixes
- *(agent-test)* Instruct the model to not output the file timestamp
- *(test)* Normalize output of download tool to make tests pass on windows

### 🚜 Refactor

- Make test code a little bit cleaner
- *(ui)* Reorganize ui components into flat structure

### 🧪 Testing

- Re-record vcr cassettes

### ⚙️ Miscellaneous Tasks

- *(legal)* @dawndiy has signed the CLA
- Remove skip for `TestCoderAgent` on windows
- Update fantasy (#1316)
- *(ui)* Add comments to internal/ui/model/ui.go
- Skipping agent tests on windows temporarily
## [0.13.2] - 2025-10-28

### ⚙️ Miscellaneous Tasks

- Go mod tidy
## [0.13.1] - 2025-10-28

### 🐛 Bug Fixes

- Fantasy nil panic
## [0.13.0] - 2025-10-27

### 🐛 Bug Fixes

- *(tests)* Remove embedded providers, regenerate fixtures

### 📚 Documentation

- *(readme)* Add note section about `openai` vs. `openai-compat`

### ⚡ Performance

- Init coder agent in a goroutine (#1289)

### ⚙️ Miscellaneous Tasks

- Small fixes and catwalk update
- More prompt improvements
- Do not override old headers
- Clone headers
- Do not respond with error
- Improve initialize
- Improve summarize logic
- Improve edit instructions
- Do not prompt for init when empty
- Go mod tidy
- *(legal)* @Supratim69 has signed the CLA
- *(legal)* @plandem has signed the CLA
- Change providers url in the log
- Update docs
- Make thinking mode persistent
- Small fix
- Auto-update generated files
## [0.13.0-beta.2] - 2025-10-24

### 🐛 Bug Fixes

- *(test)* Improve go-vcr request matching logic

### 🚜 Refactor

- Simplify code with `cmp.Or`

### ⚙️ Miscellaneous Tasks

- *(legal)* @blouflab has signed the CLA
- Skip `TestCoderAgent` for now
## [0.12.3] - 2025-10-24

### 🐛 Bug Fixes

- *(bedrock)* Update anthropic sdk with fix for aws sso (#1297)

### ⚙️ Miscellaneous Tasks

- *(legal)* @Jesssullivan has signed the CLA
- *(legal)* @mmangkad has signed the CLA
## [0.12.2] - 2025-10-23

### 🐛 Bug Fixes

- *(test)* Close sqlite db on test finish
- Bring back metrics on fantasy
- Prompt prefix
- Cost calculation when using openrouter
- *(test)* Try to have the correct dir for windows
- Diagnostics log double quotting

### 🧪 Testing

- Try fix windows tests

### ⚙️ Miscellaneous Tasks

- Use :exacto in openrouter for supported models
- Add yolo back
- Add reasoning start
- Support provider extra body and provider options
- Fix bedrock
- Remove empty thinking
- Do not push prereleases to npm, brew, etc (#1291)
- Fix parameter update

### ◀️ Revert

- "test: try fix windows tests"
## [0.13.0-beta.1] - 2025-10-23

### 🚀 Features

- *(fantasy)* Add support for bedrock

### 🐛 Bug Fixes

- *(tests)* Regenerate tests and skip some
- Add detail to error message to allow better debugging
- Accept both keys for google vertex
- *(logs)* Disable color output when stdout is not a tty (#1286)

### 💼 Other

- Fix linting

### 🚜 Refactor

- Use v2 declarative API (#1229)

### 🧪 Testing

- Re-record 2 tests to make suite green

### ⚙️ Miscellaneous Tasks

- Load embedded version of providers always
- Small update
- Change references description to match others
- Small system prompt improvemets & fantasy
## [0.12.1] - 2025-10-22

### 🐛 Bug Fixes

- Always fetch providers live and not in background (#1281)
- Only debug if enabled (#1279)

### 📚 Documentation

- *(readme)* Move slack url

### ⚡ Performance

- *(list)* Optimize filter performance and limit results (#1193)

### ⚙️ Miscellaneous Tasks

- Remove vx.y.z from the release notes (#1276)
- Embed version in build/install tasks (#1278)
## [0.12.0] - 2025-10-21

### 🚀 Features

- *(lsp)* Find references tool (#1233)
- *(bedrock)* Add support for `AWS_BEARER_TOKEN_BEDROCK` for bedrock

### 🐛 Bug Fixes

- *(mcp)* Make sure to cancel context on error (#1246)
- *(mcp)* Improve cache hits when using MCPs (#1271)
- *(tui)* Remove ctrl+d deny keybind (#1269)

### 📚 Documentation

- Update aws bedrock docs on readme

### ⚙️ Miscellaneous Tasks

- *(legal)* @dpolishuk has signed the CLA
## [0.11.2] - 2025-10-16

### 🚀 Features

- Paste/close bindings in user cmd dialog (#1221)

### 🐛 Bug Fixes

- *(bedrock)* Detect credentials set by `aws configure` (#1232)
- *(grep)* Check mime type (#1228)
- *(mcp)* Add type assertion guards (#1239)
- *(tui)* Paste on arguments input (#1240)
- *(mcp)* Append to os.Environ() (#1242)
- *(mcp)* Avoid nil errors for tool parameters (#1245)
- *(mcp)* Improve STDIO error handling (#1244)

### 🚜 Refactor

- Use clamp from /x/exp/ordered (#1236)

### 🧪 Testing

- Add tests for the dirs cmd (#1243)

### ⚙️ Miscellaneous Tasks

- Go mod tidy
- *(legal)* @BrunoKrugel has signed the CLA
## [0.11.1] - 2025-10-14

### 🐛 Bug Fixes

- *(ls)* Properly handle limits (#1230)

### ⚙️ Miscellaneous Tasks

- *(task)* Also push empty named commit in release (#1231)
- *(task)* Set commit desc automatically
## [0.11.0] - 2025-10-13

### 🚀 Features

- *(mcp)* Notifications support - tools/list_changed (#967)
- Limit filepath walk, automatic low limits when not git repo (#1052)
- *(agent)* Add support for azure provider
- *(tui)* Progress bar (#1162)
- Add responses API support
- Allow users to disable cache
- Add support for google reasoning

### 🐛 Bug Fixes

- *(test)* Make prompt deterministic
- Add title lock
- *(agent)* Cancelation logic
- *(template)* For coder so we handle memory files correctly
- *(summary)* Handle cancel and other cases
- *(agent)* Tool results
- *(agent)* Save assistant message
- *(tests)* Regenerate tests
- Don't supress application level panics
- *(mcp)* Fix ui description, double spaces (#1210)
- *(lsp)* Small UI improvements (#1211)
- Move some logs to debug
- *(vertex)* Small fix for anthropic models via google vertex (#1214)
- *(tui)* Panic (#1220)
- *(tui)* Fix progress not cleaning up some times (#1219)
- *(agent)* Queued messages issues
- *(tools)* Some errors need to be sent to the LLM
- *(agent)* Make sure we finish thinking on error
- Handle OpenAI compat type correctly in TestConnection
- Fix google vertex key

### 💼 Other

- Initial setup
- Remove old implementation

### 🚜 Refactor

- Agent tool
- *(mcp)* Use the new mcp library (#1208)
- Adjustment based on fantasy api changes

### 🧪 Testing

- Initial test setup
- Add coder agent
- Add more tests and make it easy to add tests for multiple models
- Test all the tool calls

### ⚙️ Miscellaneous Tasks

- *(task)* Annotate tags during release
- Small test updates
- Merge master
- Fix tool calls
- Fix dependencies
- Change reasoning effort
- Merge catwalk and user options
- Print a bug reporting notice when crush crashes
- *(lint)* Ignore staticcheck in helpful crash error
- *(task)* Add helper for fetching tags
- *(task)* Fetch tags before releasing
- *(task)* Just use svu if it's already installed
- *(legal)* @kucukkanat has signed the CLA
- *(legal)* @thuggys has signed the CLA
- *(legal)* @nikaro has signed the CLA
- Allow to pass args to task run
- *(legal)* @daps94 has signed the CLA
- Small fixes
- Update catwalk
- Improve provider config, handle reasoning, add vertex
- Upgrade to latest fantasy
- Improve summary
- Improve coder prompt
- Less text
- Small fix
- Fix summary by adding options
- Prompt improvements
- Persist reasoning across sessions
- `go mod tidy`
- Update anthropic's sdk with a panic fix
## [0.10.4] - 2025-09-30

### 🚀 Features

- Crush dirs (#551)

### 🐛 Bug Fixes

- *(mcp)* SSE MCPs not working (#1157)
- *(style)* Heartbit in --version

### ⚙️ Miscellaneous Tasks

- *(taskfile)* Change `release` task to add a commit for the tag (#1159)
- Update session chooser key help text
- Update model chooser key help text
- Update changelog group names
## [0.10.3] - 2025-09-29

### 🐛 Bug Fixes

- *(openai)* 429 insuffice not retry (#546)
- *(mcp)* Do not eat list tools errors (#1138)
- *(agent)* Remove timout for now (#1158)
## [0.10.2] - 2025-09-29

### 🐛 Bug Fixes

- *(gemini)* Use full MIME type for binary content in message conversion (fixes charmbracelet/crush#995)
- *(agent)* Timer should reset after each chunk
- *(lint)* Remove empty line

### ⚙️ Miscellaneous Tasks

- *(legal)* @Wangch29 has signed the CLA
- Small fixes
- Increase timeout a bit
## [0.10.1] - 2025-09-29

### 🚀 Features

- If agent has been disabled do not set the agent fn

### 🐛 Bug Fixes

- Strip path from `$SHELL` (#1119)
- *(mcp/lsp)* Expand variable in commands (#1116)
- *(fsext)* Panic on fastwalk (#1122)
- *(provider)* Do not retry auth errors
- Improve retry
- Improve retry
- Improve shutdown (#1133)
- *(lsp)* Allow directories as root markers (#1117)
- *(stream)* Stream hang, add stream timeout (#1070)
- *(gemini)* Add baseURL resolution and conditional HTTPOptions configuration (#1144)
- *(gemini)* Add missing newline at end of file

### 🚜 Refactor

- Use http.Status... consts
- Use http.Status... consts

### 🧪 Testing

- Ensure agent tool name is on list of tool names

### ⚙️ Miscellaneous Tasks

- Comment
- Add task release
- Fix version
- Add name for helper tool name resolver
- *(metrics)* Have a better identifier fallback (#1130)
- *(legal)* @Kaneki-x has signed the CLA
- Pin actions (#1132)
- *(legal)* @maxious has signed the CLA
## [0.9.3] - 2025-09-24

### 🚀 Features

- *(config)* Search`crush.json` recursively up from the working directory (#898)
- *(lsp)* Load defaults by either name or command (#1109)
- Add alt/option+esc binding to current esc key behavior
- *(config)* Allow custom providers of type gemini (#585)
- *(permissions)* Pretty-print MCP JSON

### 🐛 Bug Fixes

- *(mcp)* Pass down mcp name to logger (#1078)
- Fix a typo in README.md (#1084)
- *(deps)* Update powernap for zig
- Session summarization dialog hanging indefinitely (#528)
- *(tui)* Yes/no init selection (#1074)
- Remove <think></think> from title
- Lsp/mcp command expand ~ (#1105)
- *(grep)* Resolve Windows path parsing with null separation (#1095)
- *(mcp)* Improve timeout errors (#1108)
- *(lsp)* Command
- *(lsp)* Improve error messages
- Disable providers (#1087)

### 🚜 Refactor

- Put tool descriptions in markdown files (#1077)
- Remove unused prompt (#1083)

### 📚 Documentation

- Add bit about nixos module (#606)
- *(readme)* Add mini Chinese description (thanks @ohjia)
- Aws profile/region envs (#1104)
- Document more mcp options
- Add huggingface inference

### 🧪 Testing

- Fix
- Verify tools are taken from agent when disabled (#1103)

### ⚙️ Miscellaneous Tasks

- *(legal)* @msteinert has signed the CLA
- *(legal)* @zoete has signed the CLA
- Task fmt (#1098)
- -trimpath, remove broken target
- Home pkg godoc
- Add metrics and error tracking
## [0.9.2] - 2025-09-19

### 🐛 Bug Fixes

- Handle z.ai key validation differently

### 🚜 Refactor

- *(tidy)* Remove nested if and else block
## [0.9.1] - 2025-09-19

### 🚀 Features

- *(lsp)* Remove internal watcher (#1062)

### 🐛 Bug Fixes

- *(lsp)* Use csync for lsp clients (#1073)
## [0.9.0] - 2025-09-17

### 🚀 Features

- LSP implementation using x/powernap (#1011)
- Add attribution settings to config and bash tool (#1025)

### 🐛 Bug Fixes

- *(config)* Look for more than just crush.md

### 📚 Documentation

- *(readme)* Mention contributing guide (#1067)
- Add missing anchor
- *(readme)* Tidy attribution section

### ⚙️ Miscellaneous Tasks

- Task run
- *(legal)* @khushveer007 has signed the CLA
- *(issue-labeler)* Add azure
- Fix labeler script
- Various attention to detail edits via @andreynering
## [0.8.3] - 2025-09-16

### 🐛 Bug Fixes

- Only enable watcher for git repos (#1060)

### 💼 Other

- *(gemini)* Ensure tool responses have the user role

### 📚 Documentation

- *(readme)* Add cerebras API key to table

### ⚙️ Miscellaneous Tasks

- Bump fang to v0.4.1 to fix #1041
- *(legal)* @dvcrn has signed the CLA
## [0.8.2] - 2025-09-15

### 🐛 Bug Fixes

- Remove ulimt as go 1.19 automatically raises file descriptors
- Request MaximizeOpenFileLimit for unix
- Make the limit really high on non-unix
- Windows lint for number
- Introduce notify ignore files

### 🚜 Refactor

- Make func unexported

### ⚙️ Miscellaneous Tasks

- *(legal)* @WhiskeyJack96 has signed the CLA
- *(legal)* @Grin1024 has signed the CLA
- Remove anim example (#1045)
- Update `github.com/raphamorim/notify` to v0.9.4
## [0.8.1] - 2025-09-13

### 🚀 Features

- Fix too many open files issue (#1033)

### 📚 Documentation

- *(readme)* Copyedits to provider disabling section
## [0.8.0] - 2025-09-12

### 🚀 Features

- *(config)* Define disabled tools option which filters out agent tools access (#1016)
- Add reasoning dialog
- Add ability to disable providers auto-update from catwalk
- Add `crush update-providers` command

### 🐛 Bug Fixes

- Allow searching by provider name
- Allow multi word search
- Fix the group search
- Improve group filtering (#1024)
- Handle providers that do not send unique tool call IDs
- Add else
- Agent tool not working when switching models
- Fix agent
- Set reasoning param when selecting model
- Add mcps after the filter

### 📚 Documentation

- *(readme)* Add bluesky to socials
- *(readme)* Document how to disable providers auto-update

### ⚙️ Miscellaneous Tasks

- Auto-update generated files
- Revert  ctx change
- Do not hide the reasoning content
- Bump ultraviolet to fix double encoding keys on windows
- *(legal)* @Amolith has signed the CLA
- Lint
- Refactor
- Remove duplicate tools
- Provider error message copyedit (#1029)
## [0.7.10] - 2025-09-11

### 🐛 Bug Fixes

- Ensure it's possible to quit (`ctrl+c`) even when a dialog is open (#1007)

### ⚙️ Miscellaneous Tasks

- *(legal)* @tauraamui has signed the CLA
- *(readme)* Fix build badge url
- *(legal)* @kim0 has signed the CLA
## [0.7.9] - 2025-09-10

### 🚀 Features

- Optimize LSP file watcher and ignore files (#959)

### 🐛 Bug Fixes

- Esc key not being recognized
- Add back init setting ulimit

### 🚜 Refactor

- *(fsext)* Improve hierarchical ignore handling and consolidate file exclusion logic (#999)

### ⚙️ Miscellaneous Tasks

- *(legal)* @SubodhSenpai has signed the CLA
## [0.7.8] - 2025-09-09

### 💼 Other

- *(sqlite)* Enable `secure_delete` pragma (#966)

### ⚙️ Miscellaneous Tasks

- *(legal)* @adriens has signed the CLA
- Bump bubbletea/ultraviolet to enable bracketed paste on windows (#1003)
## [0.7.7] - 2025-09-08

### 🐛 Bug Fixes

- *(openrouter)* Fix api key validation for openrouter (#997)

### 🚜 Refactor

- Check for id instead of name

### ⚙️ Miscellaneous Tasks

- Disable `lint-sync` for now
- *(legal)* @vadiminshakov has signed the CLA
## [0.7.6] - 2025-09-07

### 🐛 Bug Fixes

- Handle no content for gemini provider

### ⚙️ Miscellaneous Tasks

- *(legal)* @shaitanu has signed the CLA
## [0.7.5] - 2025-09-02

### 🐛 Bug Fixes

- Tool calls break the converation if interrupted.

### 💼 Other

- Show persistent shell path in permission dialog (#916)

### ⚙️ Miscellaneous Tasks

- *(labeler)* Adjust xai grok label
- Rename "issue labeler" to just "labeler" (because it also labels prs)
- *(legal)* @undo76 has signed the CLA
- `IsSubset` was moved to `x/exp/slice` (#923)
- *(legal)* @andersonjoseph has signed the CLA
- *(legal)* @tisDDM has signed the CLA
## [0.7.4] - 2025-08-28

### 🐛 Bug Fixes

- Handle providers that do not send the right index

### ⚙️ Miscellaneous Tasks

- Lint
## [0.7.3] - 2025-08-28

### 🐛 Bug Fixes

- Openai provider tool calls
- Assistant message

### ⚙️ Miscellaneous Tasks

- *(legal)* @negz has signed the CLA
- *(labeler)* Add shell label
## [0.7.2] - 2025-08-27

### 🚀 Features

- Support Anthropic base url option (#702)
- Show path when asking if the user wants to initialize project (#867)

### 🐛 Bug Fixes

- Return nil for empty tools slice to handle omitzero properly (#861)
- Prevent nil pointer dereference in mcp tools parameters (#850)
- Panic at mcp config (#860)
- Improve provider cache logs (#885)
- Resolve the baseurl for anthropic
- Validate resolved url
- *(tui)* Further guard against type accessor and map panics (#783)
- Fix `panic` that happens on `crush run` with tool calls
- Fix panic with gemini via litellm
- Fix goroutine panic due to waiting for channel for too long
- *(filepicker)* General fixes to size and position
- *(lsp)* Simplify init/ping, store capabilities (#713)

### 🚜 Refactor

- Home.Dir, home.Short, home.Long (#884)

### 📚 Documentation

- Update license link to point to `LICENSE.md` (#892)
- *(readme)* Fix typo (#900)

### ⚙️ Miscellaneous Tasks

- *(legal)* @khareyash05 has signed the CLA
- *(legal)* @mpj has signed the CLA
- *(legal)* @xPrimeTime has signed the CLA
- *(legal)* @mercmobily has signed the CLA
- Fix termux /usr/etc instead of /etc
- *(legal)* @xhos has signed the CLA
- Use charmcli account for dependabot merge (#553)
- Fix `gh` login before merge attempt (#895)
- Wait for lint before trying to merge
- *(legal)* @henrebotha has signed the CLA
- Remove `dependabot` step
- *(issue-labeler)* Remove `arch` work from `os: linux` label
- *(dependabot)* Remove `docker` as we don't have docker for now
- Group dependabot updates to reduce noise
- Fix typo in constant name
- Run `modernize` (#906)
## [0.7.1] - 2025-08-21

### 🐛 Bug Fixes

- *(windows)* Downgrade ultraviolet to fix non-win32 terminals on windows (#856)

### ⚙️ Miscellaneous Tasks

- *(legal)* @linw1995 has signed the CLA
- Fix termux dirs
## [0.7.0] - 2025-08-20

### 🚀 Features

- *(fsext)* Add function to search for something in parent directories
- If a `.crush` directory is present in a parent dir, use that
- Add `fsext` function to get owner of directory
- *(cmd)* Support overriding the data directory

### 🐛 Bug Fixes

- *(onboarding)* Fix onboarding screen freezing when on click on it
- *(fsext)* Stop traversing if the directory owner changes
- Treat data directory the same in logs as elsewhere
- Scrolling the editor after paste (#466)

### 🚜 Refactor

- Move `HomeDir()` to `fsext` package
- Improve check a little bit

### ⚙️ Miscellaneous Tasks

- *(goreleaser)* `go_version` should be an input and not a secret
- *(issue-labeler)* Add `area: onboarding`
- *(legal)* @ericcoleta has signed the CLA
- Update error message, this is not user fault
- Disable `dependabot-sync`
- *(issue-labeler)* Add more words for some labels
- Bump bubbletea and ultraviolet to support win32 input mode (#838)
## [0.6.3] - 2025-08-18

### 🚀 Features

- Create `.crush/.gitignore` automatically
- *(lsp)* Allow to set custom env to lsp servers via config (#778)
- Add `IsSubset` helper

### 🐛 Bug Fixes

- *(tui)* Prevent dialog keymaps from being swallowed (#782)
- *(tui)* Fix model filter placeholder text (#790)
- Update erro msg for bash tool (#803)
- *(llm)* Log error when retrying (#781)
- Expose Required params to Anthropic API (#752)
- *(shell)* Refactor arguments blocker to check for flags in any position
- *(shell)* Block `go test -exec` and ensure it works using equals
- Fix panic that can happen on sending a message (#817)

### 💼 Other

- Remove 'go' from safe programs (#820)

### 🚜 Refactor

- Migrate bool to empty struct for lower memory usage
- Have shorter argument name
- Remove duplicated `emerge` entry

### 🧪 Testing

- Use synctest
- Simplify

### ⚙️ Miscellaneous Tasks

- *(go)* Upgrade to go 1.25 and enable `GOEXPERIMENT=greenteagc`
- Update golangci-lint version to v2.4 which supports go 1.25
- *(legal)* @lpmitchell has signed the CLA
- *(taskfile)* Set `CGO_ENABLED=0` by default
- Set commit authors to @charmcli (#614)
- Update winget configs, LICENSE.md (#597)
- Revert #25, do not ask llm update `.gitignore`
- Auto-update generated files
- *(legal)* @marcosktsz has signed the CLA
- *(legal)* @sainadh-d has signed the CLA
- Fix `golangci-lint` locally (#815)
- *(issue-labeler)* Add `security` label
## [0.6.2] - 2025-08-15

### 🐛 Bug Fixes

- *(fsext)* Prevent `.*` on gitignore from ignoring entire root dir (#766)
- *(tui)* Guard against panics in map member access
- *(keyboard-input)* Operate on characters, not bytes
- *(lint)* Correct shadowed variable
- Remove timout

### ⚙️ Miscellaneous Tasks

- Update codeowners (#760)
- *(issue-labeler)* Add android label
- *(legal)* @neomantra has signed the CLA
- *(goreleaser)* Start releasing for android / termux (#780)
- Remove extraneous comment
- *(legal)* @uri has signed the CLA
## [0.6.1] - 2025-08-14

### 🚀 Features

- *(mcp)* Configurable MCP timeout
- *(lsp)* Add filetypes configuration (#666)
- *(mcp)* Ping and recreate mcp client if needed (#772)

### 🐛 Bug Fixes

- *(lint)* Check length before slicing to avoid a panic in list
- *(llm)* Set request timeout (#736)
- Truncate long paths in compact header (#773)
- *(tui)* Underline quit dialog buttons (#548)
- *(lsp)* Return a copy of lsp diagnostics to avoid data race (#681)
- *(gemini)* Fix tool calls for google gemini (#779)

### ⚙️ Miscellaneous Tasks

- Bump bubbletea to v2.0.0-beta.4.0.20250813201422-d4d69f63338d
- Bump dependencies to fix linux console perf
- Auto-update generated files
- *(legal)* @samiulsami has signed the CLA
## [0.6.0] - 2025-08-13

### 🚀 Features

- Steering (#605)

### 🐛 Bug Fixes

- *(editor)* Better sorting of files when completing with / (#733)
- Stdio mcp startups to match mark3labs upgrade (#742)
- Make the queue push the messages above
- Add top padding
- *(mcp)* Tool output join with new line (#686)

### ⚙️ Miscellaneous Tasks

- *(legal)* @orospakr has signed the CLA
- Bump bubbletea to v2.0.0-beta.4.0.20250813191918-4ea1703d4181
## [0.5.0] - 2025-08-11

### 🚀 Features

- *(config)* Allow configure the default diff mode (#454)

### 🐛 Bug Fixes

- *(gemini)* Retry at rate limit
- *(tool)* Fix `edit` and `multi-edit` tools on windows
- *(sidebar)* Compute the right line count even on windows / crlf
- *(sidebar)* Fix full path appearing on sidebar on windows

### ⚡ Performance

- Reduce GC pressure in rendering pipeline (#687)

### ⚙️ Miscellaneous Tasks

- *(legal)* @pwnintended has signed the CLA
- *(legal)* @tazjin has signed the CLA
- *(legal)* @liznear has signed the CLA
- *(legal)* @jamestrew has signed the CLA
- *(legal)* @wwwjfy has signed the CLA
- Sync dependabot config (#308)
- Bump glamour/ultraviolet/cellbuf for performance improvements
- Fix dependabot label
## [0.4.0] - 2025-08-08

### 🚀 Features

- *(csync.Map)* Added GetOrSet
- Improve .crushignore and .gitignore
- Allow for using CRUSH_ prefixed env-vars without clobbering default env vars (#391)
- *(diffview)* Add support for mouse scrolling
- *(diffview)* Only process scroll events when pointer is over dialog
- Add yolo mode command (#654)

### 🐛 Bug Fixes

- *(selection)* Fix panic that can happen on selection
- Improve ignore
- Config.HomeDir sync.OnceValue, use user.Current().HomeDir
- Fix panic: start can't be greater than end
- Fix panic: `fuzzy` library doesn't like empty lists
- *(tools)* Do not truncate output (#657)
- *(mcp)* Give a time limit for startup (#660)
- *(mcp)* Append to output instead of replacing (#658)

### 🚜 Refactor

- Use core utils middleware from `mvdan/sh` (#323)

### 📚 Documentation

- *(readme)* Add winget to installation instructions
- *(readme)* Break out windows package managers

### ⚡ Performance

- Make it a tad faster
- Bump bubble tea for improved CPU usage

### ⚙️ Miscellaneous Tasks

- *(legal)* @CyrusZei has signed the CLA
- *(legal)* @maxjustus has signed the CLA
- *(legal)* @akaytatsu has signed the CLA
- *(legal)* @theguy000 has signed the CLA
- *(issue-labeler)* Add panic / crash label
## [0.3.0] - 2025-08-07

### 🚀 Features

- *(list)* Add HasSelection method and selectionView for text-only output
- *(chat)* Copy selected text in chat messages via shared key binding
- *(tui)* Chat: clear selection on esc
- *(tui)* Chat: add copy and clear selection key bindings to help

### 🐛 Bug Fixes

- *(tui)* Editor: sanitize pasted paths
- Add $schema to jsonschema
- *(chat)* Focus chat and editor on mouse click
- *(chat)* Expose copy key binding
- *(list)* Include inbetween empty lines when selecting text
- Handle compact mode

### 💼 Other

- Selection
- Implement selection

### 📚 Documentation

- *(readme)* Add scoop (#550)
- *(readme)* Add missing `}`
- Fix path (#628)
- *(readme)* Add note on claude max and github copilot

### ⚙️ Miscellaneous Tasks

- Rename default theme + infer default theme name (#570)
- *(legal)* @0xWelt has signed the CLA
- *(legal)* @kslamph has signed the CLA
- *(legal)* @Sunsvea has signed the CLA
- Remove cache from prefix
- *(legal)* @taciturnaxolotl has signed the CLA
- *(legal)* @bashbunni has signed the CLA
- *(issue-labeler)* Add docs label
- *(issue-labeler)* Add comment explaining why single entries
- *(legal)* @edafonseca has signed the CLA
- *(legal)* @smores56 has signed the CLA
- *(legal)* @danielsz has signed the CLA
- Bump bubbles to fix textarea cursor issue
- *(legal)* @pavelzw has signed the CLA
- *(list)* Set selection colors
- Handle smaller text than screen size
- Support reasoning model minimal
- Update openai prompt
## [0.2.2] - 2025-08-06

### 🐛 Bug Fixes

- *(mcp)* Update lib, set transport logger (#510)
- Correctly show mcp and lsp states
- Summarize provider should be large
- Sidebar files jumping
- Width related commands not showing
- Make sure to make rendered string concurrency sage
- *(mcp)* Set logger on all mcp clients (#554)
- Fix openai compatible provider
- Handle nil body in http log
- Some openai providers

### 📚 Documentation

- Improve wording
- *(readme)* Add Ollama and LMStudio config (#538)
- *(readme)* Add bedrock and vertex info

### ⚙️ Miscellaneous Tasks

- *(legal)* @bold84 has signed the CLA
- *(cla)* Add bot to allowlist
- *(legal)* @jooray has signed the CLA
- *(legal)* @Ed4ward has signed the CLA
- *(legal)* @ngnhng has signed the CLA
- Add scoop (#545)
- Fix dependabot dependencies label
- Print action
- Remove unused varibale
- Small fix
- Limit tool output
- Make consts private
- Set notices size
- *(legal)* @zloeber has signed the CLA
- *(legal)* @nelsenm2 has signed the CLA
- *(legal)* @mohseenrm has signed the CLA
- Fix models
- Simplify model filtering
- Fix missing assignment
- Reasoning for openai providers that support it
- Update catwalk
- Small fix
- Change LSP/MCP loading icon color + refactor icons + theme prep
- *(tests)* Update golden files for status
- Remove commented code
- Update catwalk
- Bump ultraviolet to v0.0.0-20250805154935-01be9d7ef65d
- Bump bubbletea to v2.0.0-beta.4.0.20250805190305
- Bump charmbracelet/x/ansi to v0.10.0
- *(tui/status)* Remove parens around MCP detail (#580)
## [0.2.1] - 2025-08-02

### 🐛 Bug Fixes

- Http log

### ⚙️ Miscellaneous Tasks

- *(legal)* @yumosx has signed the CLA
- *(issue-labeler)* `pull_request` -> `pull_request_target`
- *(issue-labeler)* Add qwen
- *(issue-labeler)* Add ci label
## [0.2.0] - 2025-08-02

### 🚀 Features

- Support .crushignore as well as .gitignore
- Debug logs request response details (#407)
- Grep should support gitignore/crushignore (#428)

### 🐛 Bug Fixes

- Lint noctx issues
- *(tui)* Fix background color of file view component (#349)
- Permissions path
- Improve path prefix checking reliability
- Fix mcp clients
- Handle code agent not initialized
- Do not init MCP client on every tool request (#459)
- Mcp client must be started (#474)

### 🚜 Refactor

- Simplify
- Move run to its own file

### 📚 Documentation

- Update
- *(readme)* Add openrouter environment variable (#451)

### ⚙️ Miscellaneous Tasks

- Fmt
- *(legal)* @alvaro17f has signed the CLA
- *(legal)* @bbrodriges has signed the CLA
- Ignore chore(legal) in release notes
- Refactor implementation
- *(legal)* @SyedaAnshrahGillani has signed the CLA
- *(legal)* @spachava753 has signed the CLA
- Add action to automate issue / pr labeling (#470)
- *(issue-label)* Allow manual trigger
- *(issue-labeler)* Fix workflow dispatch
- *(issue-labeler)* Check only title to be less error-prone
- *(issue-labeler)* Add os labels
- *(issue-labeler)* Remove slash from labels
- *(issue-labeler)* Temporarily enable issue body
- *(issue-labeler)* Have a single regexp per label to make it work
- Add script o manually run issue labeler
- *(legal)* @tabletcorry has signed the CLA
## [0.1.11] - 2025-07-31

### 🐛 Bug Fixes

- Use anthropic provider for vertexAI (#398)
- Nil pointer in view (#403)

### ⚙️ Miscellaneous Tasks

- *(legal)* @jedisct1 has signed the CLA
- *(legal)* @Djiit has signed the CLA
- *(legal)* @steipete has signed the CLA
- *(legal)* @lmn451 has signed the CLA
- *(legal)* @petersanchez has signed the CLA
- Bump charmbracelet/ultraviolet to v0.0.0-20250731212901-76da584cc9a5
## [0.1.10] - 2025-07-31

### 🐛 Bug Fixes

- Fix vertex provider
- *(onboarding)* Fix y/n key press on model list and api key field (#402)

### ⚙️ Miscellaneous Tasks

- *(legal)* @fluffypony has signed the CLA
- Fix tests
- Aur sources (#378)
- Better changelog
## [0.1.9] - 2025-07-31

### 🐛 Bug Fixes

- Panic in non-interactive mode when no providers set up
- Fix openrouter api

### 📚 Documentation

- *(readme)* More BSDs + nix run (#387)

### ⚙️ Miscellaneous Tasks

- Bump bubbletea/v2
- *(legal)* @taigrr has signed the CLA
## [0.1.8] - 2025-07-30

### 🐛 Bug Fixes

- Fix region detection for bedrock

### 📚 Documentation

- *(readme)* Un-bury the environment variable listing
## [0.1.7] - 2025-07-30

### 🚀 Features

- Make "list" tool preview cleaner (#335)

### 🐛 Bug Fixes

- Lint for var blocks
- *(tui)* Apply border styling to error messages in selected state

### 📚 Documentation

- *(cla)* Update cla comment format (#345)
- *(readme)* Update nix instructions (#359)

### ⚡ Performance

- Avoid regex compilation in runtime

### ⚙️ Miscellaneous Tasks

- *(legal)* @kujtimiihoxha has signed the CLA in $pullRequestNo
- *(cla)* Temporarily remove pr number from commit message
- Fix word case
- Bump ultraviolet
- *(legal)* @rio has signed the CLA
- Add agents to the default context
- Fix openrouter issues
- Small fix
- Prevent switching the model while agent is working
- Fix custom headers and system prefix
- Remove log
- Nur license (#347)
- Openbsd and netbsd (#344)
- Fix bedrock and azure
- *(legal)* @douglarek has signed the CLA
## [0.1.6] - 2025-07-29

### 🐛 Bug Fixes

- *(ci)* Update license to FSL-1.1-MIT
- *(ci)* Release: sign deb and rpm packages with GPG key (#346)

### 📚 Documentation

- *(readme)* Add catwalk callout
- *(readme)* Signed and notarized on macOS

### ⚙️ Miscellaneous Tasks

- Fix nur manpage
## [0.1.5] - 2025-07-29

### ⚙️ Miscellaneous Tasks

- Go mod tidy
- *(legal)* @meowgorithm has signed the CLA in $pullRequestNo
- Entitlements
## [0.1.4] - 2025-07-29

### 📚 Documentation

- *(readme)* Hold brew for signature

### ⚙️ Miscellaneous Tasks

- *(brew)* I give up for now (#342)
## [0.1.3] - 2025-07-29

### 📚 Documentation

- *(readme)* Pull winget while we wait for approval

### ⚙️ Miscellaneous Tasks

- Update GoReleaser config (#338)
## [0.1.1] - 2025-07-29

### ⚙️ Miscellaneous Tasks

- Fix missing secrets for goreleaser
## [0.1.0] - 2025-07-29

### 🚀 Features

- *(groq)* Add support for Groq using the OpenAI provider
- Simpler diff implementation
- Configure context paths (#86)
- Add azure openai models (#74)
- Model selection for given provider (#57)
- Add support for OpenRouter  (#92)
- Themes (#113)
- Test for getContextFromPaths (#105)
- Custom commands (#133)
- Support named arguments in custom commands (#158)
- Add configuration persistence for model selections (#154)
- Support VertexAI provider (#153)
- Non-interactive mode
- Title block package
- Integrate new title block
- Add new filepicker
- Add file picker help
- Profiling
- Add attachments to messages
- Normalize `CRUSH.md` as the default name
- Open exit dialog on typing `exit` or `quit`
- Advanced keyboard
- Implement show all help
- Update implementation to main
- Implement lsp errors
- Add `.crush` automatically to `.gitignore` on initializing
- Add file changes
- Use range
- *(diffview)* Getting started with the api design
- *(diffview)* Basic working functionality
- *(diffview)* Show line numbers on the left
- *(diffview)* Show hunk lines and test for multiple hunks
- *(diffview)* Implement split / side-by-side view
- *(diffview)* Implement resizer to make diff wider or shorter as asked
- *(diffview)* Show ellipsis if we reached the height limit
- *(diffview)* Implement ability to set horizontal offset
- *(diffview)* Add support for setting vertical offset
- *(diffview)* Add syntax highlighting
- Use the new diff formatter
- Remove test file
- Responsive layout
- Add details overlay
- Use fang (#39)
- Add keybinding to permissions dialog to navigate the diffview
- *(diffview)* Prevent infinite vertical scroll by default
- Use mvdan.cc/sh instead of a the user shell (#45)
- *(diffview)* Set tab width to 4
- Full windows support
- Add support for context from gemini.md
- Add config validation and provider system with mock support
- Add provider-specific prompts for OpenAI, Gemini, and Anthropic
- Add dynamic model switching with agent provider updates
- Add reasoning effort support and new providers
- Make it possible for api key to be programatic
- Add github copilot provider (#230)
- Add Arch Linux package support to nfpms
- Alternate S character for logo
- Read stdin (#101)
- *(log)* Add atomic check for initialization
- *(tui)* Support adding newlines in chat editor
- *(tui)* Display newline shitf+enter in help when supported
- *(tui)* Set the textarea value back after closing the editor
- Stream content in non-interactive mode (#133)
- *(tui)* Use lipgloss on escape sequences
- *(permissions)* Add `--yolo` flag for auto-accepting all permissions
- Implement thinking mode for anthropic models
- *(mouse)* Delegate mouse actions to terminal emulator (#218)
- Verify the API key when setting provider
- Add a download tool
- Improve providers startup (#287)
- *(shell)* Use coreutils from u-root
- *(tui)* Completions: add select and insert keybinds
- Use new catwalk
- *(tui)* Completions: dynamically adjust width based on items
- Slices as well
- Crush run (#322)
- Add schema command
- *(tui)* Editor: change textarea placeholder based on CoderAgent state (#263)

### 🐛 Bug Fixes

- *(anthropic)* Skip empty messages
- *(anthropic)* Increase max retries
- Handle anthropic 429s
- Status messages
- Set provider defaults correctly in AWS projects
- Gemini tool calling
- *(openrouter)* Set api key from env (#129)
- Allow text selection (#127)
- More intuitive keybinds (#121)
- Tweak the logic in config to ensure that env vs file configurations merge properly (#115)
- *(provider/gemini)* Prevent empty parts in assistant messages
- Tui: properly calculate quit dialog size
- Tui: use KeyPressMsg instead of KeyMsg
- *(lint)* Correct typo
- Reference Charmtone colors directly in theme
- Reference Charmtone colors directly in theme
- *(logo)* Possible division by zero
- *(logo)* Un-expose the renderer for the letter S
- *(ui)* Small ui fixes
- Re-reference Charmtone colors directly in theme
- *(lint)* Remove extraneous newline
- *(chat)* Fix focus selection
- Crush
- Regex pattern caching to eliminate repeated compilation overhead
- Update code with gofmt
- Fmt files
- Format files
- Fix "failed to generate title: context deadline exceeded" error
- Format internal/llm/tools/sourcegraph.go
- Remove pool logic
- *(diffview)* Fix left pad for ansi stuff
- *(diffview)* Fix diffview width when code is too narrow
- *(diffview)* Improve code width and make it half space on each side
- *(diffview)* Add fixes + more extensive testing for width handling
- *(diffview)* Fix small bug when width is very small
- *(diffview)* Add fixes + more extensive testing for height handling
- *(diffview)* Respect line style even if it's the last one with ellipsis
- *(diffview)* Do not print ellipsis after content has being fully printed
- *(diffview)* Fix behavior of tabs and add tab width setting
- *(diffview)* Fix debug tasks after some test changes
- Remove broken patch tool
- Format internal/lsp/watcher/watcher_performance_test.go
- Remove watcher_performance_test.go
- Fix the file history table and implement realtime file updates
- Fix title generation
- Anim fmt files
- Do not call `SyntaxHighlight`. `ChromaStyle` is enough
- Remove unused func
- Fix setting sidebar
- Prevent header text jumping by aligning ctrl+d help text
- Small improvements around command execution
- Improve killing shell children processes
- Improve shell
- Remove
- Email
- Add update listener
- Fix the agent tool
- Fix sidebar init, fix focus
- Truncate the diff on the messages
- Windows basic issues
- Added index on created_at for sessions, messages, and files (#53)
- Schema, change theme (#73)
- Add parens for correct OOO (#72)
- Formatting in validation test
- Correct model selection index in models dialog
- Resolve golangci-lint issues
- *(mcp)* Ensure required field if nil (#278)
- *(tool/grep)* Always show file names with rg (#271)
- Secrets
- Fix config
- Don't render periods of ellipsis when there's no label
- Fix logo
- Fix config paths
- Fix the path
- Add panic recovery to TUI subscription handler
- Splash page view return type
- *(ci)* Disable gomoddirectives linter
- Shifting LSPs and MCPs in sidebar
- Increase padding in quit dialog
- *(fsext)* Only log warnings if log is initialized
- Format files
- Remove debouncer
- *(tui)* Editor: ensure the textarea scrolls when inserting newlines
- *(tui)* Permission: make sure the fetch content has the correct background color
- *(tui)* Editor: change open editor key binding from ctrl+e to ctrl+v
- *(tui)* Logo: make sure we truncate the logo if it exceeds the width
- *(tui)* Copy textarea value to editor on open
- *(tui)* Editor: make sure we update the textarea after closing the editor
- *(tui)* Editor: make sure we move the cursor to the end of input
- Use the same spinner in non-interactive mode (#131)
- Ensure new line at end of non-interactive output (#132)
- *(tui)* Chat: out of bounds error in param list rendering
- *(spinner)* Fix animation not ticking sometimes
- *(logs)* Typo
- *(tui)* Editor: position completions correctly and account for padding
- *(tui)* Completions: close when no items match query
- Non-interactive and context cancellation when SIGINT (#143)
- Signal
- *(tui)* Chat: remove paste key binding from onboarding
- Use fur production (#177)
- Agent init
- Remove unused key release and uniform key layout options
- Address panic when deciding which model to use
- Timeout and context cancel (#180)
- *(tui)* Make sure we treat \r\n as \n
- Fix permission cancel logic
- *(windows)* Use `mvdan/sh` + general fixes
- *(windows)* Remove powershell/cmd builtins and keep only programs
- *(tui)* Completions should not close on no results (#198)
- Do not ignore all dot files (#197)
- Fix empty basepath
- *(tui)* Permissions: default to diff split mode when dialog is wide enough
- Fix openai provider
- *(tui)* Chat: properly align and pad the version text
- *(tui)* Chat: change version text color
- Imports
- *(tui)* Permissions: ensure content viewport has a minimum height
- *(tui)* Permissions: properly pad command block
- *(tui)* Chat: fix compact mode details toggle
- *(tui)* Permissions: set a maximum width for the dialog
- *(tui)* Completions: close and reset completions on cancel key
- *(tui)* Adjust completions popup to fit within window width
- *(tui)* Completions: don't export variable
- *(tui)* Escape control characters in tool call content
- *(logs)* Limit the logs output
- Properly capitalize reasoning and thinking status text in sidebar
- Hide completions tui when no results (#206)
- *(tui)* Renderer: replace tabs with spaces in plain content
- Start filepicker in cwd instead of os.Homedir
- Fix fetch and view tool
- Azure provider
- *(main)* Don't use JoinVertical to avoid inserting unnecessary spaces
- *(lint)* Variable shadowing
- *(tools)* Bash: add scp and ssh to banned commands
- *(messages)* Truncate attachment paths by rune, not by byte
- *(messages)* Properly measure width when rendering attachment paths
- *(tui)* Render a smaller logo on splash screen
- *(tui)* Logo: simplify logo word stretching
- *(tui)* Splash: cache logo rendering on resize
- *(tui)* Splash: re-render logo when screen size changes
- *(tui)* Chat: adjust padding for message list
- *(tui)* Splash: trigger smaller logo on 55 columns
- *(tui)* Renderer: line wrapping in chat
- Improvements
- Improv diff
- Improvements
- Pkg
- Todo
- Test
- Method name
- Load providers in background
- *(tui)* Status: properly truncate info messages
- Anim out of bounds (#283)
- *(diffview)* Fix rendering issue caused by line breaks added by chroma
- *(diffview)* Escape content to avoid surprises
- Structure logging
- *(tui)* Cursor position in the textarea
- *(non-interactive)* Check bounds when reading message bytes
- Typo
- *(tui)* App: show error message when window is too small
- Cache update logic
- Improve lazy slice
- Prevent nil ptr
- Csync.Map
- Sync
- Fix panic on onboarding when no model is selected on the list
- Hide info box when the initialization message is shown
- *(tui)* Completions: keep track of the popup position
- *(tui)* Completions: don't set initial width
- *(tui)* Completions: readjust position on filter change
- *(tui)* Completions: reposition popup on window resize
- *(tui)* Completions: ensure minimum height for completions list
- *(tui)* Completions: improve positioning and handling completions
- Don’t panic when fetching document URI path fails
- Allow to override catwalk url
- *(tui)* Typo
- Breaks
- Unused slice methods
- Small fixes in csync
- *(onboarding)* Fix "no model selected" error showing after onboarding (#316)
- Remove unused constants, fix inefficient assign
- Simplify stuff
- Simpler
- Backtick
- *(ls)* Fix path expand on ls tool

### 💼 Other

- Fix whitespace issues
- Fix unneeded lint comment
- Temporarily disable the broken rules
- Setup to find issues on tests as well
- Remove old docs
- Integrate to existing app
- Improved list
- Oboarding splashscreen
- Api key
- Init screen and simple flow
- Endpoint resolution for openai client
- Fixup panic for oob range
- Add to messages list
- Initial rework

### 🚜 Refactor

- Use context for automatic lsp process cleanup
- Upgrade Anthropic SDK to v1.4.0 and adapt provider code
- Initial permissions refactor
- New init dialog
- New compact dialog
- Rename package `diff` to `diffview`
- *(diffview)* Convert styles into a reusable `LineStyle` type
- *(diffview)* Ask hunk by parameter instead of accessing by index
- *(diffview)* Refactor tests to be table driven
- *(diffview)* Cache num digits on a struct variable
- *(diffview)* Move padding out of `styles.go`
- *(diffview)* Refactor order of functions to match call order
- *(diffview)* Move `setPadding` to be inside `adjustStyles`
- *(diffview)* Reduce number of golden files
- *(diffview)* Make code a little bit shorter
- *(diffview)* Wrap default styles into functions
- Rename `internal/fileutil` into `internal/fsutil`
- Replace magic numbers with named constants
- Use uv bubbletea
- *(log)* Rename log initialization function for clarity
- Remove weird context value usage (#153)
- Improve command blocking system and fix test isolation
- *(tools)* Replace custom LineScanner with bufio.Scanner
- Use sync primitives in GetMcpTools
- Move `cmd` package inside `internal` as `internal/cmd`
- *(diffview)* Simplify how we handle line endings
- Move ansi escape function to the `ansiext` package
- Use csync.Map instead of sync.Map
- Use embed for prompts

### 📚 Documentation

- *(readme)* Add demo GIF (#160)
- *(crush.md)* Add commenting directive
- *(readme)* Empty out readme + add a few essential elements
- *(readme)* Add (very) basic getting started instuctions
- *(crush.md)* Add more specific Go formatting instructions
- Add testing guidelines for mock providers
- *(readme)* Various small edits
- Add targeted documentation for state management
- *(readme)* Various readme improvements
- *(readme)* Simplify and cleanup readme
- *(readme)* Add MCP info
- *(readme)* Add (temporary) installation instructions for non-developers
- Godoc
- Todo
- *(readme)* Now supporting windows!
- *(readme)* Add  notes about APKs and PKGs
- *(readme)* Genericize MCP examples
- *(readme)* Add header
- *(legal)* Update license (#318)
- Add $schema to crush.json
- *(readme)* General update
- *(readme)* Shave down GIF
- *(readme)* Update header and layout
- *(readme)* Additional note in contact info
- Update

### ⚡ Performance

- *(anim)* Use faster PRNG
- *(anim)* Eliminate string concatenation
- *(anim)* Reduce pointer dereferences
- *(anim)* Preallocate strings.Builder capacities
- Drop reflection usage and use string builder
- Shell commands using exponential backoff
- Optimize HTTP client pooling and binary file detection
- Lsp batching
- Make spinner less cpu demanding
- *(diffview)* Only compute syntax highlight for printed lines
- Concurrency improvements
- Reduce lock contention
- Attempt to improve permissions dialog performance
- *(diffview)* Use faster checksum algorithm
- Init tools in background
- *(ansiext)* Grow string builder to optimize allocations

### 🧪 Testing

- *(diffview)* Add test for custom context lines
- *(diffview)* Fix file name in test
- *(diffview)* Update golden files to remove final line breaks
- *(diffview)* Add golden files for width testing
- *(diffview)* Add golden files for height testing
- *(diffview)* Record golden files for xoffset test
- *(diffview)* Record golden files for yoffset test
- *(diffview)* Update golden files for syntax highlighting
- *(diffview)* Allow tests to run in parallel
- *(diffview)* Update golden files with a small bug fix
- *(diffview)* Update golden files for diffview tests
- Improve tests (#315)

### ⚙️ Miscellaneous Tasks

- Add a `.editorconfig` file
- Add a taskfile and default golangci-lint config
- Run `gofumpt`
- Run linting on ci
- Truncate version number if too long
- Rename 'title' package to 'logo'
- Fix typo in comment
- Align 'Charm' to edge of C curve
- Cleanup image component
- Change the module
- Rename opencode -> crush
- Rename OpenCode -> Crush
- Rename OPENCODE -> CRUSH
- Rename .opencode.json -> .crush.json
- Initial logs changes
- Bump bubbletea to latest v2-viewable commit
- Add a bench test for grep
- Update hard coded domains and author names
- Change the layout package location
- Rename active dialog id method
- Remove test keymap
- Small fix
- Small ux tweeks
- Having some fun with the `go-udiff` package
- *(diffview)* Update `go-udiff`
- Add taskfile to easily run tests
- *(diffview)* Fix lint
- *(diffview)* Add missing doc comments
- Fix broken performance tests:
- Small sidebar fix
- Fix linting
- *(diffview)* Remove unused `SyntaxHighlight` option
- Small fixes
- Add version
- Change the min width to 120
- Remove unused diff funcs
- Fix list focus, fix files with no changes
- Add dedicates diff styles
- Small ui fixes
- Dynamic change the diff format on size
- Add assistant section
- Small spelling stuff
- Remove unused func
- Change cancel logic
- Cancel requests when user closes
- Small fix
- Small fixes
- Small fix
- Implement agent tool
- Improve error messages
- Update goreleaser config
- Tool border, navigation help
- The diamond as the model icon
- Remove extra gap between model icon and name after messages
- If we're in the home dir top level show the full path
- *(spinner)* Set spinner to its original speed
- Typo
- *(permissions)* No need to generate a cache key
- Delete uneeded `cachedContentKey`
- *(lint)* Gofmt
- *(theme)* Use charmtone in glamour theme
- Move shell to its own package
- Lint fixes
- Fix linting
- Initial setup for the new config
- Move fur structs, small provider changes
- Add name to tools
- Move to the new config
- Change to the small/large model conceptg
- Tests
- Clean up unused dependencies
- Migrate config file from .crush.json to crush.json
- Add new fields to provider model
- Fix lint
- Change how max tokens works
- Make it possible to override maxTokens
- Enable thinking for anthropic models
- Generate schema, fix tests
- Change how we generate the json schema
- Implement large/small model switching
- Adjust wording of small/large model labels
- Change how summarize sessions gets triggered
- Fix lint
- Rebase
- Fix shell
- Small fixes to the config
- Remove debug log
- Small fixes
- Fix lint
- Remove unnecessary commnet
- Release and nightly jobs (#90)
- Fix token
- Fix username
- Revert, fix made on meta
- Nightly
- Fixes
- *(gitignore)* Ignore all .crush directories no matter where they are
- *(lint)* Normalize receiver var names
- *(lint)* Pull out magic number
- Overhaul spinner
- Integrate overhauled spinner
- Match initialchar colors in label position to label colors
- Drop installation script
- Fix openrouter
- Fix mcp servers
- Change how we handle max tokens for anthropic
- Fix tests
- Fix build workflow (#94)
- Disable changelog on nightly
- Update lipgloss
- *(spinner)* Re-implement color cycling
- Initial config tests
- Implement new logs
- Sort keys
- Small fix
- Change config signature
- Add more tests and handle more cases
- Add default models handling and tests
- Lint
- Add helper funcs
- Remove logs
- Lint
- Remove mcp server
- Small fixes
- Go mod tidy
- Pass datadir to db
- Make diff config independent
- Make permissions config independent
- Make getContextFromPaths config independent
- Make tools config independent
- Pass config to app
- Fix error output
- Fix max tokens limit
- Allow no config
- Fix logs
- Move highlight inside of tui
- Move diffview inside of tui
- Fix tests
- Rewrite the chat page
- Changes after rebase
- Show configured indicator
- Add a way to persist configurations
- Go mod tidy
- Make agent optional
- Allow custom providers without api key
- Try fix windows tests
- Go mod tidy
- Use internal versions of bubbletea and lipgloss v2
- Bump uv dependencies
- Bump bubbletea v2 uv
- Bump dependencies
- Go mod vendor
- Bump ultraviolet
- Update vendor
- Update dependencies
- Update lipgloss
- Update dependencies
- Update vendor dependencies
- Small update
- Fix
- Handle disabled state for mcps
- Chore: add mcp disabled state
- Bump dependencies
- Update vendors
- Bump dependencies
- Update vendors
- Lint
- Add api key step
- Remove unused fzf
- Small message change
- Fix openai provider for multi tool call
- Add info to the splash screen
- Improve help
- Add completions help and small improvements
- Improve agent tool ui
- Fix cursor
- Fix bash output
- Fix cursor
- Small ux changes
- Fix initialize padding
- Fix padding for info
- Disable go mod proxying for now
- Fix newline with \
- Update vendors
- *(tui)* Editor: return error on write failure
- Only build nightlies if there are something to build (#126)
- *(pro)* Nightly
- Remove background color on 'waiting for tool to start'
- *(lint)* Add missing GoDoc
- *(theme)* Use new BBQ tone
- Simplify 'canceled' styling + use american spelling
- Use american spelling for 'canceled' in all user-facing copy
- Small padding change
- Improve initalize text
- Bump dependencies
- Go mod vendor
- Upgrade bubbles v2 to latest v2-exp commit
- Add instruction for `.golden` files on `.gitattributes`
- Fix provider
- Fix bash commands with tabs
- Fix extra space
- Remove extra help
- Disable commands when other dialogs are open
- Fix missing keys
- Fix diffview tests on ci on windows
- Implement correct banned commands
- Fix errors not showing in bash tool
- Better naming
- Small change
- *(tui/completions)* Pull out magic number
- *(logs)* Copyedit
- Correctly mark project as initialized
- Fix panic on message send
- Center permissions
- Prevent multiple dialogs from opening
- Fix cancellation
- Handle cancel when panic
- Handle tools cancelled
- Show API errors
- Lint
- Add some status tests
- Move the permissions dialog a bit up
- Make the sidebar a bit more responsive
- Try to make the number of items in the sidebar dynamic
- Small fixe details screen
- *(tui)* Add tests for escapeLine
- *(goreleaser)* Disable windows for now
- Improve wording
- *(lint)* Enable staticcheck and fix its issues
- Small fix
- Cleanup the UI
- Remove empty line
- Small editor fix
- Improve thinking
- Lint
- Bump ultraviolet to latest main
- Go mod vendor
- Bump dependencies
- Update vendors
- Update `Co-Authored-By` address to `crush@charm.land`
- Some mcp improvements
- Add github mcp
- Remove unnecessary config
- Fix permissions dialog
- Small fix
- Add thinking/reasoning indicator
- Small fixes
- Unvendor crush
- Capitalize 'Go' in LSP config for this project
- Add extra body & header to openai and anthropic
- *(sidebar)* Remove space after periods of ellipsis; make darker
- *(sidebar)* Increase logo height breakpoint threshold
- Go mod tidy
- Run `modernize`
- *(windows)* Show warning on windows for now
- Remove box, have less styling
- Add CODEOWNERS file
- Bump ultraviolet to latest
- Bump ultraviolet
- Tidy various comments
- Dependabot, sync and lint jobs
- *(messages)* Pre-allocate user message attachment slice
- *(messages)* No need to append to a pre-allocated slice
- *(messages)* Greatly simplify filename truncation with ansi.Truncate
- Conditionally add thinking header
- Remove thinking when done
- Improve diff view performance
- Improve shell output and summary to include cwd
- Fix `build` github action
- Publish to npm (#266)
- Build apk for alpine
- *(windows)* Remove warning
- *(goreleaser)* Release for windows
- Bump charmbracelet/ultraviolet to v0.0.0-20250723145313-809e6f5b43a1
- *(diffview)* `TrimSuffix` -> `TrimRight`
- Update example
- *(non-interactive)* Log read error
- *(tui)* Small window design adjustments
- Be specific about what is too small
- Trigger
- Update readme
- Small status update
- Initial implementation
- More tests
- Test selection
- Implement new filterable list
- Improve anim performance
- Small improvements
- Change checksum algo
- Fix viewport offset
- Grouped list
- Use new list in completions
- Rebase fix + sync map
- Safe slice
- Lint
- Add mouse support
- Remove provider tests on startup
- Fix selected item
- Update golden
- Fix tests
- Fix completions
- Small fix
- Rename seq2
- Address review requests
- Publish to winget, nur; fixes brew token (#297)
- Make anim concurrent safe
- Fix tool spinning
- Fix sidebar
- Fix lint issue
- Run `modernize`
- Adjust codeowners to not prevent dependabot prs being merged (#319)
- Update dependencies
- Auto generate on changes
- Build more targets (#327)
- Improve permissions & edit tool
- Experimental new coder prompt
- Remove log
- Prevent new sessions when agent busy
- Fix editor
- Add suspend
- Add copy
- Small fix
- Add copy message
- Fix truncation line
- Default to the new prompt
- Small scroll fix
- Support default headers
- Proper caching for anthropic models with openrouter
- Use the catwalk name
- Support clipboard past image
- Lint
- Add more commands
- Fix gemini validation
- Small fix
- Prompt improvements
- Add new session
- Remove log
- Change open editor
- Lint
- Update ultraviolet dependency
- *(lint)* Change timeout to 10 minutes
- Remove old scripts
- Fix schema command
- Remove log
- Improve permissions
- Fix homedir
- Rever prompts
- Fix enter
- Fix model selector
- Toss in some more placeholders
- Small fix
- *(legal)* Add cla bot (#320)
- Fix token perms
- Bot PAT
- *(legal)* @caarlos0 has signed the CLA in charmbracelet/crush#$pullRequestNo
- *(legal)* Fix signatures
- *(legal)* Fix cla job
- *(legal)* Cla cleanup
- *(legal)* @caarlos0 has signed the CLA in $pullRequestNo
- *(legal)* @raphamorim has signed the CLA in $pullRequestNo
- *(legal)* @aymanbagabas has signed the CLA in $pullRequestNo
- *(legal)* @andreynering has signed the CLA in $pullRequestNo
- *(legal)* @ras0q has signed the CLA in $pullRequestNo
