# Snippet Manager

A command-line tool for managing and organizing code snippets. Built with Go and Cobra, this tool allows you to save, retrieve, and organize your frequently used code snippets.

## Features

- Save code snippets with descriptions and tags
- Auto-generated unique IDs for each snippet
- JSON-based storage in your home directory
- Tag-based organization
- Simple and intuitive command-line interface

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/snippet-manager.git

# Navigate to the project directory
cd snippet-manager

# Build the project
go build
```

## Usage

### Adding a Snippet

```bash
snippet-manager add -n "name" -c "content" -d "description" -t tag1,tag2
```

Options:
- `-n, --name`: Name of the snippet (required)
- `-c, --content`: Content of the snippet (required)
- `-d, --description`: Description of the snippet
- `-t, --tag`: Tags for the snippet (comma-separated)

Example:
```bash
snippet-manager add -n "http-server" -c "http.ListenAndServe(\":8080\", nil)" -d "Basic HTTP server" -t go,server
```

### Storage

Snippets are stored in JSON format at `~/.snippet-manager/snippets.json`.

## Project Structure

```
snippet-manager/
├── cmd/           # Command implementations
├── internal/      # Internal packages
│   ├── snippet/   # Snippet type and operations
│   └── store/     # Storage implementation
└── main.go        # Entry point
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.