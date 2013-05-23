http://www.joelonsoftware.com/articles/Unicode.html
characters, character sets, encodings, and Unicode.

**ASCII** - 
* 7 bits
* Represent every character using a number between 32 and 127
* Space = 32, 'A' = 65
* Codes below 32 were called unprintable, used for control characters.
* lots of people had their own ideas of what should go where in the space from 128 to 255.
* IBM-PC - OEM character set which provided some accented characters for European languages and a bunch of line drawing characters.
* In the ANSI standard, everybody agreed on what to do below 128.
* Asia - DBCS, the "double byte character set", some letters were stored in one byte and others took two. 

**Unicode**
- a letter maps to something called a code point which is still just a theoretical concept. How that code point is represented in memory or on disk is a whole nuther story. .e.g "U+0639" Unicode + Hexadecimal.

**Encodings**
- UTF-8 - stores Unicode code points, in memory using 8 bit bytes, variable byte-length up to 6 bytes
- Unicode Byte Order Mark at beginning of string for high-endian or low-endian mode.
- UTF-7, UTF-16 etc etc.

**UTF-8**

Let us consider how to encode the Euro sign, €.
* The Unicode code point for "€" is U+20AC.
* According to the scheme table above, this will take three bytes to encode, since it is between U+07FF and U+FFFF.
* Hexadecimal 20AC is binary 0010000010101100. The two leading zeros are added because, as the scheme table shows, a three-byte encoding needs exactly sixteen bits from the code point.
* Because it is a three-byte encoding, the leading byte starts with three 1s, then a 0 (1110...)
The remaining bits of this byte are taken from the code point (11100010), leaving ...000010101100.
Each of the continuation bytes starts with 10 and takes six bits of the code point (so 10000010, then 10101100).
* The three bytes 11100010 10000010 10101100 can be more concisely written in hexadecimal, as E2 82 AC.

The following table summarises this conversion, as well as others with different lengths in UTF-8. The colors indicate how bits from the code point are distributed among the UTF-8 bytes. Additional bits added by the UTF-8 encoding process are shown in black.

```
Character  Binary code point	Binary UTF-8	Hexadecimal UTF-8
$	U+0024	0100100	00100100	24
¢	U+00A2	000 10100010	11000010 10100010	C2 A2
€	U+20AC	00100000 10101100	11100010 10000010 10101100	E2 82 AC
𤭢	U+24B62	00010 01001011 01100010	11110000 10100100 10101101 10100010	F0 A4 AD A2
```

**HTML**

```html
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
```
But that meta tag really has to be the very first thing in the <head> section because as soon as the web browser sees this tag it's going to stop parsing the page and start over after reinterpreting the whole page using the encoding you specified.
