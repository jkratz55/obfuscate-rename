# Obfuscate Rename

Obfuscate Rename is a tool to rename files to random strings using UUIDs. It is useful to obfuscate the names of files before sharing, protecting test/production data, etc.

Obfuscate Rename supports recursively traversing and renaming files recursively and can filter files by extension.

## Installation

```bash
go install github.com/jkratz55/obfuscate-rename
```

## Usage

Example:
```bash
obfuscate-rename -d . --extension .mp4
```