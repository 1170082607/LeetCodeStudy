# Repository Guidelines

## Project Structure & Module Organization
This Go 1.21 module (`go.mod` -> `LeetCodeStudy`) collects LeetCode solutions under `top_interview_150/array_and_strings`. Each solution file follows `<difficulty>_<id>_<slug>.go` where difficulty is `eazy`, `medium`, or `hard`, paired with a sibling `_test.go`. Keep problem-specific helpers unexported and colocated. IDE settings in `.idea/` support GoLand; leave them untouched unless syncing config changes.

## Build, Test, and Development Commands
- `go test ./...` - runs every solution test suite; use before pushing.
- `go test ./top_interview_150/array_and_strings -run Test_canJump` - tight loop on one case while iterating.
- `gofumpt -w .` - repository-standard formatter; run from the repo root after edits.

## Coding Style & Naming Conventions
Follow Go defaults: tabs for indentation, gofumpt-ordered imports, and lowerCamelCase identifiers for functions like `canJump`. File names keep the difficulty prefix, zero-padded LeetCode id, and descriptive slug (`medium_122_max_profit.go`). Keep comments bilingual-friendly, matching existing Chinese problem statements. Avoid introducing packages beyond `array_and_strings` unless expanding the directory tree.

## Testing Guidelines
Tests rely on the standard `testing` package with table-driven cases (`tests := []struct {...}`). Name test functions `Test_<function>` and store them in sibling `_test.go` files. When adding new solutions, mirror the filename and cover at least the official example inputs. Prefer descriptive `tt.name` values over empty strings for new tables.

## Commit & Pull Request Guidelines
Recent history uses short, imperative summaries referencing problem ids (`solve 13 and 42`, `use "gofumpt -w ." in root path`). Keep the same voice, grouping related problems per commit when practical. Pull requests should list the solved problems, cite any LeetCode discussions referenced, and note the test command you ran. Include screenshots only if comparing performance metrics; otherwise a short outcomes bullet is enough.

## Agent Workflow Tips
Document new subdirectories in this guide when expanding beyond `array_and_strings`, and update examples accordingly. If you introduce shared utilities, add a README stub beside them and explain how they stay LeetCode-compliant.
