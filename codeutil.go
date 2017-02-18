package linebot

import(
  "strings"
)

func getMorsable(singleWord rune) Morsable {
  singleWordStr := string(singleWord)

  switch singleWordStr {
  case "イ","い":
    return KANA_I
  case "ロ", "ろ":
    return KANA_RO
  case "ハ","は":
    return KANA_HA
  case "二","に":
    return KANA_NI
  case "ホ","ほ":
    return KANA_HO
  case "ヘ","へ":
    return KANA_HE
  case "ト","と":
    return KANA_TO
  case "チ","ち":
    return KANA_CHI
  case "リ","り":
    return KANA_RI
  case "ヌ","ぬ":
    return KANA_NU
  case "ル","る":
    return KANA_RU
  case "ヲ","を":
    return KANA_WO
  case "ワ","わ":
    return KANA_WA
  case "カ","か":
    return KANA_KA
  case "ヨ","よ":
    return KANA_YO
  case "タ","た":
    return KANA_TA
  case "レ","れ":
    return KANA_RE
  case "ソ","そ":
    return KANA_SO
  case "ツ","つ":
    return KANA_TU
  case "ネ","ね":
    return KANA_NE
  case "ナ","な":
    return KANA_NA
  case "ラ","ら":
    return KANA_RA
  case "ム","む":
    return KANA_MU
  case "ウ","う":
    return KANA_U
  case "ヰ","ゐ":
    return KANA_WI
  case "ノ","の":
    return KANA_NO
  case "オ","お":
    return KANA_O
  case "ク","く":
    return KANA_KU
  case "ヤ","や":
    return KANA_YA
  case "マ","ま":
    return KANA_MA
  case "ケ","け":
    return KANA_KE
  case "フ","ふ":
    return KANA_HU
  case "コ","こ":
    return KANA_KO
  case "エ","え":
    return KANA_E
  case "テ","て":
    return KANA_TE
  case "ア","あ":
    return KANA_A
  case "サ","さ":
    return KANA_SA
  case "キ","き":
    return KANA_KI
  case "ユ","ゆ":
    return KANA_YU
  case "メ","め":
    return KANA_ME
  case "ミ","み":
    return KANA_MI
  case "シ","し":
    return KANA_SHI
  case "ヱ","ゑ":
    return KANA_WE
  case "ヒ","ひ":
    return KANA_HI
  case "モ","も":
    return KANA_MO
  case "セ","せ":
    return KANA_SE
  case "ス","す":
    return KANA_SU
  case "ン","ん":
    return KANA_N
  case "ー":
    return TYOUON
  case "゛":
    return DAKUTEN
  case "゜":
    return HANDAKUTEN
  case "、":
    return KUGIRITEN
  case "ガ", "が":
    return KANA_OTHER_GA
  case "ギ", "ぎ":
    return KANA_OTHER_GI
  case "グ", "ぐ":
    return KANA_OTHER_GU
  case "ゲ", "げ":
    return KANA_OTHER_GE
  case "ゴ", "ご":
    return KANA_OTHER_GO
  case "ザ", "ざ":
    return KANA_OTHER_ZA
  case "ジ", "じ":
    return KANA_OTHER_ZI
  case "ズ", "ず":
    return KANA_OTHER_ZU
  case "ゼ", "ぜ":
    return KANA_OTHER_ZE
  case "ゾ", "ぞ":
    return KANA_OTHER_ZO
  case "ダ", "だ":
    return KANA_OTHER_DA
  case "ヂ", "ぢ":
    return KANA_OTHER_DI
  case "ヅ", "づ":
    return KANA_OTHER_DU
  case "デ", "で":
    return KANA_OTHER_DE
  case "ド", "ど":
    return KANA_OTHER_DO
  case "バ", "ば":
    return KANA_OTHER_BA
  case "ビ", "び":
    return KANA_OTHER_BI
  case "ブ", "ぶ":
    return KANA_OTHER_BU
  case "ベ", "べ":
    return KANA_OTHER_BE
  case "ボ", "ぼ":
    return KANA_OTHER_BO
  case "ヴ", "ゔ":
    return KANA_OTHER_VU
  case "パ", "ぱ":
    return KANA_OTHER_PA
  case "ピ", "ぴ":
    return KANA_OTHER_PI
  case "プ", "ぷ":
    return KANA_OTHER_PU
  case "ペ", "ぺ":
    return KANA_OTHER_PE
  case "ポ", "ぽ":
    return KANA_OTHER_PO
  case "ァ", "ぁ":
    return KANA_OTHER_A
  case "ィ", "ぃ":
    return KANA_OTHER_I
  case "ゥ", "ぅ":
    return KANA_OTHER_U
  case "ェ", "ぇ":
    return KANA_OTHER_E
  case "ォ", "ぉ":
    return KANA_OTHER_O
  case "ヵ":
    return KANA_OTHER_KA
  case "ヶ":
    return KANA_OTHER_KE
  case "ッ", "っ":
    return KANA_OTHER_TU
  case "ャ", "ゃ":
    return KANA_OTHER_YA
  case "ュ", "ゅ":
    return KANA_OTHER_YU
  case "ョ", "ょ":
    return KANA_OTHER_YO
  case "ヮ", "ゎ":
    return KANA_OTHER_WA
  }

	upper := strings.ToUpper(singleWordStr)
	switch upper {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "D":
		return D
	case "E":
		return E
	case "F":
		return F
	case "G":
		return G
	case "H":
		return H
	case "I":
		return I
	case "J":
		return J
	case "K":
		return K
	case "L":
		return L
	case "M":
		return M
	case "N":
		return N
	case "O":
		return O
	case "P":
		return P
	case "Q":
		return Q
	case "R":
		return R
	case "S":
		return S
	case "T":
		return T
	case "U":
		return U
	case "V":
		return V
	case "W":
		return W
	case "X":
		return X
	case "Y":
		return Y
	case "Z":
		return Z
	}

  switch singleWordStr {
  case "0":
    return ZERO
  case "1":
    return ONE
  case "2":
    return TWO
  case "3":
    return THREE
  case "4":
    return FOUR
  case "5":
    return FIVE
  case "6":
    return SIX
  case "7":
    return SEVEN
  case "8":
    return EIGHT
  case "9":
    return NINE
  }
	return NO_DEF
}

func TranslateToCode(words string) string {
	upperWords := strings.ToUpper(words)
	wordArray := strings.Split(upperWords, "\n")

	var result string
	for i, word := range wordArray {
		for j, character := range word {
			if j != 0 {
				result += DelimiterSlash2
			}
      morsable := getMorsable(character)
      if(morsable == NO_DEF){
        return GetNoDefMessage()
      }

			result += morsable.GetMorseCode()
		}
    if(i < len(wordArray)-1){
		    result += "\n"
    }
	}

	return result

}
