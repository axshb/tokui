```
 _____     _             _
|_   _|__ | | ___ _   _ (_)
  | |/ _ \| |/ / | | | || |
  | | (_) |   <| |_| | || |
  |_|\___/|_|\_\\__,_|_||_|
  ```
A blazing fast TUI scratchpad for real-time LLM token counting, built on Go with Bubble Tea and Lipgloss. 

### Features
- Real-time token counting using the `cl100k_base` encoding (GPT-4).
- File browser to check the token counts of files in the directory.

### Showcase
![Tokui Screenshot](https://github.com/axshb/tokui/blob/main/DOCS/screenshot.png)

### Keybinds
- Tab / Shift+Arrows: Cycle focus between Sidebar and Editor.
- Enter: Select file (Sidebar) or New line (Editor).
- Ctrl+X: Clear editor buffer.
- Ctrl+C: Quit.

### Installation & Development

Using Go:
1. `git clone https://github.com/yourusername/tokui.git`
2. `cd tokui`
3. `go mod tidy`
4. `go run .`

Using Nix:
1. `nix develop` (enters a shell with dependencies)
2. `go run .`

### Commit Rules

- chore: for maintenance, build process changes, tooling, etc.
- docs: for documentation changes
- feat: for features
- fix: for bugfixes/optimizations.
