# ascii-art-color

## Description
Ascii-art-color is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. 

The output should manipulate colors using the flag` --color=<color> <substring to be colored>`, in which --color is the flag and `<color>` is the color desired by the user and `<substring to be colored> `is the substring that you can chose to be colored. These colors can be achieved using different notations (color code systems, like `RGB`, `hsl`, `ANSI`...).

---
### Requirement:
+ The substring to be colored can be a single letter or more.
+ If the substring is not specified, the whole `string` should be colored.
+ The flag must have exactly the same format as above, any other formats must return the following usage message:
```text
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

---
## How It Works:
+ This project gets its ascii character from some banner files with a specific graphical template representation. The files are formatted in a way that is not necessary to change them.

+ In the banner files, each character has a height of 8 lines.

+ Characters are separated by a new line `\n`.

+ `ANSI escape codes` are standardized, in-band signaling sequences—typically starting with an Escape character (
ESC or ASCII 27)—embedded within text streams to instruct terminal emulators to format text, manipulate colors, move the cursor, and clear the screen.

+ Position markers for substring coloring checks for the presence of a substring in the input and colors the substring instead of the full input, and in the absence of the substring, all the full string must be colored.

---

## Usage
```bash
# Normal run
go run . "Hello"

# Color whole string
go run . --color=red "Hello"

# Color substring only
go run . --color=red kit "a king kitten have kit"

# With newlines
go run . --color=blue "Hello\nThere"

# Invalid flag format
go run . --colour=red "Hello"
# Output: Usage: go run . [OPTION] [STRING]...
```
---
## Supported Colors
+ red, green, yellow, blue, magenta, cyan, white

---
## Supported Input
+ Letters upper and lower case
+ Numbers
+ Special characters
+ `\n` for new lines

---
## Banner Files
+ standard — classic style
+ shadow — shadow effect style 
+ thinkertoy — playful style

---
## Requirements 
+ Go 1.22 or higher
+ No external dependencies

---
## Author
**Samuel Ojetunde || Michael Akinwale**