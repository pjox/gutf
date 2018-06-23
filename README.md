# Gutf

Gutf is a terminal tool that converts file encodings.

This tool will convert files from and into a variety of encodings (see the supported encodings section), but please try to use it responsibly and always convert you files into UTF-8. If by any event you need to convert your text files into something different than UTF-8, don't worry, Guft will help you do those unholy convertions.

Different from iconv, gutf works with huge files as it does not load the whole file into memory.

## Usage

```text
gutf [input] [encoding] [output] [encoding]
```

## Supported Encodings

```text
IBMCodePage037
IBMCodePage437
IBMCodePage850
IBMCodePage852
IBMCodePage855
WindowsCodePage858
IBMCodePage860
IBMCodePage862
IBMCodePage863
IBMCodePage865
IBMCodePage866
IBMCodePage1047
IBMCodePage1140
ISO8859-1
ISO8859-2
ISO8859-3
ISO8859-4
ISO8859-5
ISO8859-6
ISO-8859-6E
ISO-8859-6I
ISO8859-7
ISO8859-8
ISO-8859-8E
ISO-8859-8I
ISO8859-9
ISO8859-10
ISO8859-13
ISO8859-14
ISO8859-15
ISO8859-16
KOI8-R
KOI8-U
Macintosh
MacintoshCyrillic
Windows874
Windows1250
Windows1251
Windows1252
Windows1253
Windows1254
Windows1255
Windows1256
Windows1257
Windows1258
X-User-Defined
EUC-JP
ISO-2022-JP
ShiftJIS
EUC-KR
GB18030
GBK
HZ-GB2312
Big5
UTF-8
UTF-16BE(UseBOM)
UTF-16BE(IgnoreBOM)
UTF-16LE(IgnoreBOM)
UTF-32BE(UseBOM)
UTF-32BE(IgnoreBOM)
UTF-32LE(IgnoreBOM)
```
