package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	txt := readInput()
	commandPattern := regexp.MustCompile(`\([^\)]+?[\)]|[\""].+?[\""]|[^ ]+`)
	foundCommands := commandPattern.FindAllString(txt, -1)
	foundCommands = findReplaceCommand(foundCommands)
	/*for _, i := range foundCommands {
		fmt.Println(i)
	}*/
	foundCommands = findReplaceCommandMulti(foundCommands)

	//checkArticle := regexp.MustCompile(`(^|\s)([aA])(\s[aeiouhAEIOUH])`)
	//foundArticle := checkArticle.FindAllString(txt) - 1
	//foundArticle = replaceArticle(foundArticle)

	//checkSingleQuotes := regexp.MustCompile(`(')(\s+)(.+?)(\s+)(')`)
	//foundSingleQuetes := checkSingleQuotes.FindAllString(txt, -1)
	//foundSingleQuetes = replaceSingleQuotes(foundSingleQuetes)
	//fmt.Printf("Typeof %T", foundSingleQuetes)
}

func readInput() string {
	//if len(os.Args) > 1 {
	//	fmt.Println("Error. No filename")
	//}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %q", err)
	}
	//Readfile returns bytes, so convert into string for further use
	return string(file)
}



func findReplaceCommand(txt []string) []string {
	for i, foundCommand := range txt {
		if foundCommand == "(hex)" && i >= 1 {
			txt[i-1] = hexToDec(txt[i-1])
			if i < len(txt)-1 {
				txt = RemoveIndex(txt, i)
				i--
			} else {
				txt = txt[:len(txt)-1]
			}
		} else if foundCommand == "(bin)" && i >= 1 {
			txt[i-1] = binToDec(txt[i-1])
			if i < len(txt)-1 {
				txt = RemoveIndex(txt, i-1)
				i--
			} else {
				txt = txt[:len(txt)-1]
			}
		} else if foundCommand == "(low)" && i >= 1 {
			txt[i-1] = strings.ToLower(txt[i-1])
			if i < len(txt)-1 {
				txt = RemoveIndex(txt, i)
				i--
			} else {
				txt = txt[:len(txt)-1]
			}
		} else if foundCommand == "(up)" && i >= 1 {
			txt[i-1] = strings.ToUpper(txt[i-1])
			if i < len(txt)-1 {
				txt = RemoveIndex(txt, i)
				i--
			} else {
				txt = txt[:len(txt)-1]
			}
		} else if foundCommand == "(cap)" && i >= 1 {
			txt[i-1] = strings.Title(txt[i-1])
			if i < len(txt)-1 {
				txt = RemoveIndex(txt, i)
				i--
			} else {
				txt = txt[:len(txt)-1]
			}
		}
		fmt.Println(txt)
	}
	return txt
}

func findReplaceCommandMulti(txt []string) []string {
	findNumber := regexp.MustCompile(`\d+`)
	for i, foundCommand := range txt {
		num, _ := strconv.Atoi(findNumber.FindString(foundCommand))
		if strings.Contains(foundCommand, "up") && i >= 1 {
			if num > i {
				num = i
			}
			for j := 0; j < num; j++ {
				txt[i-j-1] = strings.ToUpper(txt[i-j-1])
			}
			txt = RemoveIndex(txt, i)
			i--
		} else if strings.Contains(foundCommand, "low") && i >= 1 {
			if num > i {
				num = i
			}
			for j := 0; j < num; j++ {
				txt[i-j-1] = strings.ToLower(txt[i-j-1])
			}
			txt = RemoveIndex(txt, i)
			i--
		} else if strings.Contains(foundCommand, "cap") && i >= 1 {
			if num > i {
				num = i
			}
			for j := 0; j < num; j++ {
				txt[i-j-1] = strings.Title(txt[i-j-1])
			}
			txt = RemoveIndex(txt, i)
			i--
		}
	}
	return txt
}
*/

func hexToDec(s string) string {
	result_dec, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Printf("ERROR. The string was not converted from hex to dec: %q\n", err)
		return s
	}
	return strconv.Itoa(int(result_dec))
}

func binToDec(s string) string {
	result_dec, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Printf("ERROR. The string was not converted from bin to dec: %q\n", err)
		return s
	}
	return strconv.Itoa(int(result_dec))
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

/*
func replaceArticle(txt []string) []string {
	for _, article := range txt {
		txt = article.ReplaceString(txt, article, `$1${2}n$3`, 1)
	}
	// r, _ := regexp.Compile(`(^|\s)([aA])(\s[aeiouhAEIOUH])`)
	//result := r.ReplaceAllString(text, `$1${2}n$3`)
	return txt
}

// Regex for ”
func replaceSingleQuotes(txt [] string) string {
	//r, _ := regexp.Compile(`(')(\s+)(.+?)(\s+)(')`)
	result := r.ReplaceAllString(txt, `$1$3$5`)
	return result
}

// Regex for . , ! ? : ;
func fixPunctuations(txt []string) []string {
	for _, punc := range txt{
		punc.ReplaceAllString(txt, `$2 `, 1)
	}
	return txt

	//r, _ := regexp.Compile(`(\s*)([\.,!\?:;]+)(\s*)`)
	//result := r.ReplaceAllString(text, `$2 `)
	//return result
}

*/
