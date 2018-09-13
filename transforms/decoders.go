package transforms

//Combines two hex chars into a single byte
func x2c(c1 byte, c2 byte) byte {
	var char byte
	if c1 >= 'A' {
		char = ((c1 & 0xdf) - 'A') + 10
	} else {
		char = c1 - '0'
	}
	char *= 16
	if c2 >= 'A' {
		char += ((c2 & 0xdf) - 'A') + 10
	} else {
		char += c2 - '0'
	}
	return char
}

//Converts a hex char into a byte
func xsingle2c(c byte) byte {
	if c >= 'A' {
		return ((c & 0xdf) - 'A') + 10
	} else {
		return c - '0'
	}
}

func validHex(X byte) bool {
	return ((X >= '0') && (X <= '9')) || ((X >= 'a') && (X <= 'f')) || ((X >= 'A') && (X <= 'F'))
}

func isDigit(X byte) bool {
	return (X >= '0') && (X <= '9')
}

//checks if the byte is a octo decimal digit
func isODidit(X byte) bool {
	return (X >= '0') && (X <= '7')
}

/*checks for white-space  characters.   In  the  "C"  and  "POSIX"
locales,  these  are:  space,  form-feed ('\f'), newline ('\n'),
carriage return ('\r'), horizontal tab ('\t'), and vertical  tab
('\v'). */
func isspace(X byte) bool {
	return X == ' ' || X == '\n' ||
		X == '\r' || X == '\t' ||
		X == '\f' || X == '\v'
}

//URLDecodeUni converts a url encoded string into a UTF-8 encoded string
//with support for unicode characters.
//
//`input` is the string that should be transformed.
//
//`unicodeCodePage is the unicode code page that should be used to decode.
//A recommended default is 20127 which is the us ASCII table
func URLDecodeUni(
	input string,
	unicodeCodePage int,
) string {

	if input == "" {
		return ""
	}

	newString := make([]byte, len(input))

	index, newIndex, xv, code, fact := 0, 0, 0, 0, 0
	hmap := byte(0)
	hmapFound := false

	inputLength := len(input)

	for index < inputLength {
		if input[index] == '%' {
			// Character is a percent sign.

			if (index+1) < inputLength && ((input[index+1] == 'u') || (input[index+1] == 'U')) {
				// IIS-specific %u encoding.

				if index+5 < inputLength {
					// We have at least 4 data bytes.
					if validHex(input[index+2]) &&
						validHex(input[index+3]) &&
						validHex(input[index+4]) &&
						validHex(input[index+5]) {

						code = 0
						fact = 1
						hmapFound = false

						if len(unicodemap) > 0 && unicodeCodePage > 0 {

							for i := 5; i >= 2; i-- {
								if validHex(input[index+i]) {
									if input[index+i] >= 97 {
										xv = int(input[index+i]) - 97 + 10
									} else if input[index+i] >= 65 {
										xv = int(input[index+i]) - 65 + 10
									} else {
										xv = int(input[index+i]) - 48
									}
									code += xv * fact
									fact *= 16
								}
							}

							if code >= 0 && code <= 65535 {
								hmap, hmapFound = unicodemap[unicodeCodePage][code]
							}
						}

						if hmapFound {
							newString[newIndex] = hmap
						} else {
							// We first make use of the lower byte here, ignoring the higher byte.
							newString[newIndex] = x2c(input[index+4], input[index+5])

							// Full width ASCII (ff01 - ff5e) needs 0x20 added
							if (newString[newIndex] > 0x00) &&
								(newString[newIndex] < 0x5f) &&
								(input[index+2] == 'f' || input[index+2] == 'F' &&
									(input[index+3] == 'f' || input[index+3] == 'F')) {
								newString[newIndex] += 0x20

							}
						}
						newIndex++
						index += 6
					} else {
						// Invalid data, skip %u.
						newString[newIndex] = input[index]
						newIndex++
						newString[newIndex] = input[index+1]
						newIndex++
						index += 2
					}
				} else {
					// Not enough bytes (4 data bytes), skip %u.
					newString[newIndex] = input[index]
					newIndex++
					newString[newIndex] = input[index+1]
					newIndex++
					index += 2
				}
			} else {
				// Standard URL encoding.

				// Are there enough bytes available?
				if index+2 < inputLength {
					// Yes

					// Decode a %xx combo only if it is valid.
					c1, c2 := input[index+1], input[index+2]

					if validHex(c1) && validHex(c2) {
						newString[newIndex] = x2c(c1, c2)
						newIndex++
						index += 3
					} else {
						// Not a valid encoding, skip this %
						newString[newIndex] = input[index]
						newIndex++
						index++
					}
				} else {
					// Not enough bytes available, skip this %
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			}
		} else {

			// Character is not a percent sign.
			if input[index] == '+' {
				newString[newIndex] = ' '
			} else {
				newString[newIndex] = input[index]
			}
			newIndex++
			index++
		}
	}

	return string(newString[:newIndex])
}

//URLDecode converts a url encoded string into a UTF-8 encoded string
func URLDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		if input[index] == '%' {
			/* Character is a percent sign. */

			/* Are there enough bytes available? */
			if index+2 < inputLength {
				c1 := input[index+1]
				c2 := input[index+2]

				if validHex(c1) && validHex(c2) {
					/* Valid encoding - decode it. */
					newString[newIndex] = x2c(input[index+1], input[index+2])
					newIndex++
					index += 3
				} else {
					/* Not a valid encoding, skip this % */
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			} else {
				/* Not enough bytes available, copy the raw bytes. */
				newString[newIndex] = input[index]
				newIndex++
				index++
			}
		} else {
			/* Character is not a percent sign. */
			if input[index] == '+' {
				newString[newIndex] = ' '
				newIndex++
			} else {
				newString[newIndex] = input[index]
				newIndex++
			}
			index++
		}
	}

	return string(newString[0:newIndex])
}

//HTMLEntitiesDecode decodes html entities to UTF-8 chars
func HTMLEntitiesDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		//If the start of a html encoded entity
		if input[index] == '&' && index+2 < inputLength {
			//If next char is # then we continue parsing
			if input[index+1] == '#' {
				if input[index+2] == 'X' || input[index+2] == 'x' {
					//Try numeric parsing
					left := index + 2
					right := left

					//While we have valid digits move the right pointer right
					//Unless we already have 4 digits
					for right+1 < inputLength && left-right < 4 && validHex(input[right+1]) {
						right++
					}

					//If we have at least one digit we decode
					if right > left {

						code, xv, fact := 0, 0, 1

						for i := 0; left < right-i; i++ {

							//map A-F to 10-16
							//map a-f to 10-16
							//map '1'-'9' to 1-9
							if input[right-i] >= 97 {
								xv = int(input[right-i]) - 97 + 10
							} else if input[right] >= 65 {
								xv = int(input[right-i]) - 65 + 10
							} else {
								xv = int(input[right-i]) - 48
							}

							code += xv * fact
							fact *= 16
						}

						unicodeString := string(code)
						for i := 0; i < len(unicodeString); i++ {
							newString[newIndex] = unicodeString[i]
							newIndex++
						}

						index += 3 + (right - left)

						if index < inputLength && input[index] == ';' {
							index++
						}
					} else {
						//if we have no valid digits we add the & and # to the new string
						newString[newIndex] = input[index]
						newIndex++
						index++

						newString[newIndex] = input[index]
						newIndex++
						index++
					}

				} else {
					//Try numeric parsing
					left := index + 1
					right := left

					//While we have valid digits move the right pointer right
					//Unless we already have 4 digits
					for right+1 < inputLength && left-right < 4 && isDigit(input[right+1]) {
						right++
					}

					//If we have at least one digit we decode
					if right > left {

						code, xv, fact := 0, 0, 1

						for i := 0; left < right-i; i++ {

							//map '1'-'9' to 1-9
							//map '1'-'9' to 1-9
							xv = int(input[right-i]) - 48

							code += xv * fact
							fact *= 10
						}

						unicodeString := string(code)
						for i := 0; i < len(unicodeString); i++ {
							newString[newIndex] = unicodeString[i]
							newIndex++
						}

						index += 2 + (right - left)

						if index < inputLength && input[index] == ';' {
							index++
						}
					} else {
						//if we have no valid digits we add the & and # to the new string
						newString[newIndex] = input[index]
						newIndex++
						index++

						newString[newIndex] = input[index]
						newIndex++
						index++
					}
				}
			} else {
				//Try to match to predefined xml entities
				//TODO implement full html predefined entities list
				match := false
				if index+3 < inputLength {
					if input[index+1:index+3] == "gt" {
						newString[newIndex] = '>'
						newIndex++
						index += 3
						match = true
					} else if input[index+1:index+3] == "lt" {
						newString[newIndex] = '<'
						newIndex++
						index += 3
						match = true
					} else if index+4 < inputLength {
						if input[index+1:index+4] == "amp" {
							newString[newIndex] = '&'
							newIndex++
							index += 4
							match = true
						} else if index+5 <= inputLength {
							if input[index+1:index+5] == "quot" {
								newString[newIndex] = '"'
								newIndex++
								index += 5
								match = true
							} else if input[index+1:index+5] == "apos" {
								newString[newIndex] = '\''
								newIndex++
								index += 5
								match = true
							} else if input[index+1:index+5] == "nbsp" {
								newString[newIndex] = '\xa0'
								newIndex++
								index += 5
								match = true
							}
						}

					}
				}

				if match {
					if index < inputLength && input[index] == ';' {
						index++
					}
				} else {
					//Not part of html encoded entity so copy char
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			}
		} else {
			//Otherwise we add the & and we go to the next char
			newString[newIndex] = input[index]
			newIndex++
			index++
		}
	}

	return string(newString[0:newIndex])
}

//JSDecode converts a Javascript encoded string into a UTF-8 encoded string
func JSDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		if input[index] == '\\' {

			/* \uHHHH unicode escape sequence*/
			if index+5 < inputLength && input[index+1] == 'u' &&
				validHex(input[index+2]) && validHex(input[index+3]) &&
				validHex(input[index+4]) && validHex(input[index+5]) {

				/* Use only the lower byte. */
				newString[newIndex] = x2c(input[index+4], input[index+5])

				/* Full width ASCII (ff01 - ff5e) needs 0x20 added */
				if newString[newIndex] > 0x00 && newString[newIndex] < 0x5f &&
					(input[index+2] == 'f' || input[index+2] == 'F') &&
					(input[index+3] == 'f' || input[index+3] == 'F') {

					newString[newIndex] += 0x20
				}

				newIndex++
				index += 6
				/* \xHH hex secapte sequence*/
			} else if index+3 < inputLength && input[index+1] == 'x' &&
				validHex(input[index+2]) && validHex(input[index+3]) {

				newString[newIndex] = x2c(input[index+2], input[index+3])
				newIndex++
				index += 4

				/* \OOO (only one byte, \000 - \377) */
			} else if index+1 < inputLength && isODidit(input[index+1]) {

				//TODO check this alloc, it is probably not necessary
				buf := make([]byte, 4)
				j := 0
				for index+1+j < inputLength && j < 3 {
					buf[j] = input[index+1+j]
					j++
					if index+1+j < inputLength && !isODidit(input[index+1+j]) {
						break
					}
				}

				if j > 0 {
					/* Do not use 3 characters if we will be > 1 byte */
					if (j == 3) && (buf[0] > '3') {
						j = 2
						buf[j] = '\x00'
					}

					code, xv, fact := 0, 0, 1

					for i := 1; i <= j; i++ {

						//map '1'-'9' to 1-9
						xv = int(buf[j-i]) - 48

						code += xv * fact
						fact *= 8
					}

					newString[newIndex] = byte(code)
					newIndex++
					index += 1 + j
				}

				/* \C */
			} else if index+1 < inputLength {
				c := input[index+1]
				switch input[index+1] {
				case 'a':
					c = '\a'
					break
				case 'b':
					c = '\b'
					break
				case 'f':
					c = '\f'
					break
				case 'n':
					c = '\n'
					break
				case 'r':
					c = '\r'
					break
				case 't':
					c = '\t'
					break
				case 'v':
					c = '\v'
					break
					/* The remaining (\?,\\,\',\") are just a removal
					 * of the escape char which is default.
					 */
				}

				newString[newIndex] = c
				newIndex++
				index += 2
			} else {
				/* Not enough bytes */
				for index < inputLength {
					newString[newIndex] = input[index]
					newIndex++
					index++
				}
			}
		} else {
			newString[newIndex] = input[index]
			newIndex++
			index++
		}
	}

	return string(newString[0:newIndex])
}

//CSSDecode converts css escaped chars to a utf-8 string
func CSSDecode(input string) string {
	if input == "" {
		return ""
	}

	inputLength := len(input)
	newString := make([]byte, inputLength)

	index, newIndex := 0, 0

	for index < inputLength {
		// Is the character a backslash?
		if input[index] == '\\' {

			// Is there at least one more byte?
			if index+1 < inputLength {
				// We are not going to need the backslash.
				index++

				// We are not going to need the backslash.
				j := 0
				for j < 6 && index+j < inputLength && validHex(input[index+j]) {
					j++
				}

				// We have at least one valid hexadecimal character.
				if j > 0 {
					fullcheck := false

					// For now just use the last two bytes
					switch j {
					case 1:
						// Number of hex characters
						newString[newIndex] = xsingle2c(input[index])
						newIndex++
						break

					case 2:
						//Use the last two from the end
						newString[newIndex] = x2c(input[index+j-2], input[index+j-1])
						newIndex++
						break
					case 3:
						//Use the last two from the end
						newString[newIndex] = x2c(input[index+j-2], input[index+j-1])
						newIndex++
						break
					case 4:
						// Use the last two from the end, but request a full width check.
						newString[newIndex] = x2c(input[index+j-2], input[index+j-1])
						fullcheck = true
						break
					case 5:
						/* Use the last two from the end, but request
						* a full width check if the number is greater
						* or equal to 0xFFFF.
						 */
						newString[newIndex] = x2c(input[index+j-2], input[index+j-1])

						if input[index] == '0' {
							fullcheck = true
						} else {
							newIndex++
						}

						break
					case 6:
						/* Use the last two from the end, but request
						 * a full width check if the number is greater
						 * or equal to 0xFFFF.
						 */
						newString[newIndex] = x2c(input[index+j-2], input[index+j-1])
						if input[index] == '0' && input[index+1] == '0' {
							fullcheck = true
						} else {
							newIndex++
						}
					}

					// Full width ASCII (0xff01 - 0xff5e) needs 0x20 added
					if fullcheck {
						if (newString[newIndex] > 0x00) && (newString[newIndex] < 0x5f) &&
							(input[index+j-3] == 'f' || input[index+j-3] == 'F') &&
							(input[index+j-4] == 'f' || input[index+j-4] == 'F') {

							newString[newIndex] += 0x20
						}

						newIndex++
					}

					// We must ignore a single whitespace after a hex escape
					if index+j < inputLength && isspace(input[index+j]) {
						j++
					}

					// Move over.
					index += j

				} else if input[index] == '\n' { // No hexadecimal digits after backslash

					// A newline character following backslash is ignored.
					index++

				} else { // The character after backslash is not a hexadecimal digit, nor a newline.

					// Use one character after backslash as is.
					newString[newIndex] = input[index]
					newIndex++
					index++
				}

			} else { // No characters after backslash.

				// Do not include backslash in output (continuation to nothing)
				index++
			}

		} else { // Character is not a backslash.
			//Copy one normal character to output.
			newString[newIndex] = input[index]
			newIndex++
			index++
		}
	}

	return string(newString[0:newIndex])
}
