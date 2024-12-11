# md2cb - Markdown to Clipboard with Rich Text support

A simple command line tool that converts markdown to rich text and copies it to the clipboard. Perfect for converting markdown tables to rich text that can be pasted into applications that support formatted text.

## BEWARE

> This is MAC OS only!

## Requirements

- macOS (uses native Cocoa framework for clipboard operations)
- Go 1.17 or later
- pandoc (`brew install pandoc`)

## Installation

```bash
go install github.com/oderwat/md2cb@latest
```

Or download the latest release from the [releases page](https://github.com/oderwat/md2cb/releases).

## Usage

The program reads markdown from stdin and places the converted rich text on the clipboard:

```bash
echo "| Header | Content |
|-|-|
| Cell 1 | Cell 2 |" | md2cb
```

Or read from a file:

```bash
cat your_table.md | md2cb
```

## How it Works

1. Uses pandoc to convert markdown to HTML
2. Converts the HTML to rich text using macOS's native NSAttributedString
3. Places the result on the clipboard in both HTML and rich text formats
4. Ready to paste into any application that supports formatted text

## License

MIT License - see [LICENSE.txt](LICENSE.txt) file