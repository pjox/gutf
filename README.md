# Gutf

Gutf is a terminal tool that converts files encoding to UTF-8.

For the moment it only converts Windows1254 to UTF-8 but more encodings will be added later. This tool will only convert files into UTF-8 because, let's be honest, using something different than UTF-8 today is just unnaceptable.

Different from iconv, gutf works with huge files as it does not load the whole file into memory.

## Usage

```text
gutf inputFile outputFile
```