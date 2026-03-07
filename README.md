# MY GO-RELOADED PROGECT DOCUMENTAION

## Project Definition

So basically, what this project is all about is receiving a file that contains a sentence of unrefined text as input file, then modify the contents in the file, but the modification isn't done in the input file, rather in another file, which is the output file.

## Cases/Instanses To Work With
- Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")

- Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")

- Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO!")

- Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

- Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")

- For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

- Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").

- Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".

- The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")

- If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")

- Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").

## How It Works

The entry point is `FileToFileHandler`, which reads the input file line by line, passes each line through `MyLogic`, and writes the result to the output file.

```
Input file → FileToFileHandler → MyLogic → Output file
```

`MyLogic` chains all transformation functions in order:

```
HexAndBin → PunctuationsHandler → LowUpAndCap → ArticleHandler
```

---

## What Each Function In My Code Does

### `FileToFileHandler(inputFile, outputFile string) error`
Handles all file I/O. Opens the input file, creates the output file, reads line by line, passes each line to `MyLogic`, and writes the result. Uses buffered reading and writing for efficiency. Returns an error if any file operation fails.

---

### `MyLogic(text string) string`
Orchestrates the transformation pipeline. Calls each handler in sequence on a single line of text and returns the fully transformed result.

---

### `HexAndBin(s string) string`
Converts hexadecimal or binary values to decimal when followed by a `(hex)` or `(bin)` tag.

| Input | Output |
|---|---|
| `"1E (hex) files were added"` | `"30 files were added"` |
| `"It has been 10 (bin) years"` | `"It has been 2 years"` |

---

### `PunctuationsHandler(s string) string`
Fixes common punctuation spacing mistakes.

| Fix | Input | Output |
|---|---|---|
| Space before punctuation | `"Hello !"` | `"Hello!"` |
| Space between punctuation | `"Wait . ."` | `"Wait.."` |
| Missing space after punctuation | `"Hello,world"` | `"Hello, world"` |
| Space after opening quote | `"' hello'"` | `"'hello'"` |
| Space before closing quote | `"'hello '"` | `"'hello'"` |
| Leading/trailing whitespace | `"  Hello  "` | `"Hello"` |

---

### `LowUpAndCap(s string) string`
Applies case transformations to words preceding inline tags. The optional number `N` controls how many words are affected (defaults to 1).

| Tag | Effect | Input | Output |
|---|---|---|---|
| `(up)` | UPPERCASE previous word | `"go (up)"` | `"GO"` |
| `(low)` | lowercase previous word | `"GO (low)"` | `"go"` |
| `(cap)` | Capitalize previous word | `"hello (cap)"` | `"Hello"` |
| `(up, N)` | UPPERCASE previous N words | `"so exciting (up, 2)"` | `"SO EXCITING"` |

---

### `ArticleHandler(s string) string`
Replaces the article `"a"`/`"A"` with `"an"`/`"An"` when followed by a word starting with a vowel or `h`. Preserves the original capitalisation of the article.

| Input | Output |
|---|---|
| `"a amazing rock"` | `"an amazing rock"` |
| `"A elephant"` | `"An elephant"` |

---

## Tests

Tests live in the `tests` package and cover each transformation function individually using table-driven test cases.

### `TestConvertToVowel`
Verifies that `ArticleHandler` correctly replaces `"a"`/`"A"` with `"an"`/`"An"` before vowel-starting words.

```
"There it was. A amazing rock!" → "There it was. An amazing rock!"
```

### `TestHexAndBin`
Verifies that `HexAndBin` correctly converts hex and binary values to decimal.

```
"1E (hex) files were added" → "30 files were added"
"It has been 10 (bin) years" → "It has been 2 years"
```

### `TestLowUpAndCap`
Verifies that `LowUpAndCap` applies case transformations to the correct number of preceding words.

```
"This is so exciting (up, 2)" → "This is SO EXCITING"
"Ready, set, go (up) !"       → "Ready, set, GO !"
```

### `TestPunctuations`
Verifies that `PunctuationsHandler` fixes spacing around punctuation and trims whitespace.

```
"  This is so EXCITING .  ...yes it is .   " → "This is so EXCITING.... yes it is."
"Ready , set , go  ! "                        → "Ready, set, go!"
```

### Running the tests

```bash
go test ./tests/...
```