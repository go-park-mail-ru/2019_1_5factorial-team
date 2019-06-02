package validator

import (
	"fmt"
	"regexp"
	"strings"
)

const expBad = `(\s+|^)([пПnрРp]?[3ЗзВBвПnпрРpPАaAаОoO0о]?[сСcCиИuUОoO0оАaAаыЫуУyтТT]?|\w*[оаАaAО0oO])[Ппn][иИuUeEеЕ][зЗ3][ДдDd]\w*[\?\,\.\!\;\-]*|(\s+|^)\w{0,4}[оОoO0иИuUаАaAcCсСзЗ3тТTуУy]?[XxХх][уУy][йЙеЕeёЁEeяЯ9юЮиИuU]\w*[\?\,\.\;\-\!]*|(\s+|^)[бпПnБ6][лЛ][яЯ9]+([дтДТDT]\w*)?[\?\,\.\;\!\-]*|(\s+|^)\w*[бпПnБ6][лЛ][яЯ9][дтДТDT]\w+[\?\,\.\;\-\!]*|(\s+|^)(\w*[оОoO0ъЪьыЫЬаАaAзЗ3уУy])?[еЕeEиИuUёЁ][бБ6пП]([оОoO0ыЫаАaAнНHиИuUуУyлЛеЕeкКkKE]\w*)?[\?\,\!\.\;\-]*|\s*^[ШшЩщ][лЛ][юЮ][хХxX]?[шШщЩ]?[кКkK]?\w*[\?\,\!\.\;\-]*|\s*^[сСcC][цЦ]?[уyУ]+[чЧ]?[КkKк]*\w*[\?\,\!\.\;\-]*|\s*^[пПn][uUИи][Дд][aAАаоОoO0][Рpр]\w*[\?\,\!\.\;\-]*|\s*^[гГ][ОoOоаАaA][НHн][Дд][oOО0о][нНH]\w*[\?\,\!\.\;\-]*|\s*^\w*[3Зз][аАaAоОoO0][лK][уyУ][пПn]\w*[\?\,\!\.\;\-]*`

var (
	//expBadCompile = regexp.MustCompile(expBad)
	expBadCompile *regexp.Regexp
	symbols       = map[string][]string{
		"а": {"а", "a", "@"},
		"б": {"б", "6", "b"},
		"в": {"в", "b", "v"},
		"г": {"г", "r", "g"},
		"д": {"д", "d", "g"},
		"е": {"е", "e"},
		"ё": {"ё", "е", "e"},
		"ж": {"ж", "zh", "*"},
		"з": {"з", "3", "z"},
		"и": {"и", "u", "i"},
		"й": {"й", "u", "y", "i"},
		"к": {"к", "k", "i{", "|{"},
		"л": {"л", "l", "ji"},
		"м": {"м", "m"},
		"н": {"н", "h", "n"},
		"о": {"о", "o", "0"},
		"п": {"п", "n", "p"},
		"р": {"р", "r", "p"},
		"с": {"с", "c", "s"},
		"т": {"т", "m", "t"},
		"у": {"у", "y", "u"},
		"ф": {"ф", "f"},
		"х": {"х", "x", "h", "к", "k", "}{"},
		"ц": {"ц", "c", "u,"},
		"ч": {"ч", "ch"},
		"ш": {"ш", "sh"},
		"щ": {"щ", "sch"},
		"ь": {"ь", "b"},
		"ы": {"ы", "bi"},
		"ъ": {"ъ"},
		"э": {"э", "е", "e"},
		"ю": {"ю", "io"},
		"я": {"я", "ya"},
	}

	wordToReplace = []string{
		"бля",
		"хуй",
		"заебись",
		"мудак",
		"пизда",
		"ебал",
		"ахуе",
		"пидор",
		"яндекс",
		"пидор",
		"залупа",
		"хер",
	}
)

// ((?:х|x|h|к|k|}{)\W*[уyu]\W*[йuyi])
// |
// ((?:б|6|b)\W*[лlji]\W*[яya])

// (
// 	(?:`символ`|) \W* [символы] \W* []
// )

func init() {
	//for first, k := range symbols {
	//	fmt.Println(first, k)
	//}

	//fmt.Println("=============")
	regExp := ""

	for _, word := range wordToReplace {
		wordSlice := strings.Split(word, "")

		val, _ := symbols[wordSlice[0]]
		rule := strings.Join(val, "|")

		regExp = fmt.Sprintf(`%s((?i)(?:%s)`, regExp, rule)

		for i, wordSymbol := range wordSlice {
			if i == 0 {
				continue
			}
			//fmt.Println(wordSymbol)

			val, _ := symbols[wordSymbol]
			rule = strings.Join(val, "")

			regExp = fmt.Sprintf(`%s\W*[%s]`, regExp, rule)

		}
		regExp = fmt.Sprintf(`%s)|`, regExp)
		//fmt.Println(regExp)
	}
	fmt.Println(regExp)

	//expBadCompile = regexp.MustCompile(regExp)
}

func ReplaceBadWords(input string) string {

	//return string(expBadCompile.ReplaceAll([]byte(input), []byte("***")))

	return string(expBadCompile.ReplaceAllFunc([]byte(input), func(bytes []byte) []byte {
		return []byte(strings.Repeat("*", len(string(bytes))/2))
	}))
}
