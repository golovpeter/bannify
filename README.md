# 🎨 Bannify

A simple CLI tool to generate colorful ASCII art banners. Perfect for terminal splash screens, MOTDs, or just making your terminal output more fun!

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap username/bannify
brew install bannify
```

### From source

```bash
go install github.com/username/bannify@latest
```

### Manual build

```bash
git clone https://github.com/username/bannify.git
cd bannify
go build -o bannify
```

## Usage

```bash
# Basic usage
bannify HELLO

# Specify text with flag
bannify -t "HELLO WORLD"

# Choose a color
bannify -c rainbow "GO ROCKS"

# Add a border
bannify -b -c magenta "BREW"

# Combine options
bannify -b -c yellow "BUILD TOOLS"

# Show version
bannify -v
```

## Color Options

- `red` - Red text
- `green` - Green text
- `blue` - Blue text
- `yellow` - Yellow text
- `magenta` - Magenta text
- `cyan` - Cyan text (default)
- `rainbow` - Rainbow gradient effect 🌈

## Examples

### Simple banner
```bash
bannify HELLO
```

### Rainbow gradient with border
```bash
bannify -b -c rainbow "WELCOME"
```

### Perfect for MOTD
Add to your `~/.bashrc` or `~/.zshrc`:
```bash
bannify -c cyan "$(whoami)"
```

## Supported Characters

- A-Z (uppercase)
- 0-9 (numbers)
- Space and exclamation mark

Lowercase letters are automatically converted to uppercase.

## Use Cases

- 🖥️ Terminal splash screens
- 📜 Message of the day (MOTD)
- 🎬 Presentation demos
- 📦 Build script headers
- 🎉 Fun CLI output

## Contributing

Contributions are welcome! Feel free to:
- Add more ASCII fonts
- Add more color schemes
- Improve character support
- Fix bugs

## License

MIT License - feel free to use this however you want!

## Author

Created with ❤️ for the terminal enthusiasts

---

⭐ If you like this tool, give it a star on GitHub!