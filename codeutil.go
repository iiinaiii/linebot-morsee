package linebot

import (
	"strings"
)

type TranslateStatus int

const (
	StatusKana TranslateStatus = iota
	StatusAlphabet
	StatusNumber
	StatusMorse
	StatusInvalid
	StatusNone
)

type Language int

const (
	LanguageKana Language = iota
	LanguageAlphabet
)

func getMorsable(singleWord rune) Morsable {
	singleWordStr := string(singleWord)

	switch singleWordStr {
	case "イ", "い":
		return KANA_I
	case "ロ", "ろ":
		return KANA_RO
	case "ハ", "は":
		return KANA_HA
	case "ニ", "に":
		return KANA_NI
	case "ホ", "ほ":
		return KANA_HO
	case "ヘ", "へ":
		return KANA_HE
	case "ト", "と":
		return KANA_TO
	case "チ", "ち":
		return KANA_CHI
	case "リ", "り":
		return KANA_RI
	case "ヌ", "ぬ":
		return KANA_NU
	case "ル", "る":
		return KANA_RU
	case "ヲ", "を":
		return KANA_WO
	case "ワ", "わ":
		return KANA_WA
	case "カ", "か":
		return KANA_KA
	case "ヨ", "よ":
		return KANA_YO
	case "タ", "た":
		return KANA_TA
	case "レ", "れ":
		return KANA_RE
	case "ソ", "そ":
		return KANA_SO
	case "ツ", "つ":
		return KANA_TU
	case "ネ", "ね":
		return KANA_NE
	case "ナ", "な":
		return KANA_NA
	case "ラ", "ら":
		return KANA_RA
	case "ム", "む":
		return KANA_MU
	case "ウ", "う":
		return KANA_U
	case "ヰ", "ゐ":
		return KANA_WI
	case "ノ", "の":
		return KANA_NO
	case "オ", "お":
		return KANA_O
	case "ク", "く":
		return KANA_KU
	case "ヤ", "や":
		return KANA_YA
	case "マ", "ま":
		return KANA_MA
	case "ケ", "け":
		return KANA_KE
	case "フ", "ふ":
		return KANA_HU
	case "コ", "こ":
		return KANA_KO
	case "エ", "え":
		return KANA_E
	case "テ", "て":
		return KANA_TE
	case "ア", "あ":
		return KANA_A
	case "サ", "さ":
		return KANA_SA
	case "キ", "き":
		return KANA_KI
	case "ユ", "ゆ":
		return KANA_YU
	case "メ", "め":
		return KANA_ME
	case "ミ", "み":
		return KANA_MI
	case "シ", "し":
		return KANA_SHI
	case "ヱ", "ゑ":
		return KANA_WE
	case "ヒ", "ひ":
		return KANA_HI
	case "モ", "も":
		return KANA_MO
	case "セ", "せ":
		return KANA_SE
	case "ス", "す":
		return KANA_SU
	case "ン", "ん":
		return KANA_N
	case "ー":
		return TYOUON
	case "゛":
		return DAKUTEN
	case "゜":
		return HANDAKUTEN
	case "、":
		return KUGIRITEN
	case " ", "　":
		return SPACE
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

// TranslateToMorse function
func TranslateToMorse(words string) string {
	upperWords := strings.ToUpper(words)
	wordArray := strings.Split(upperWords, "\n")

	var result string
	for i, word := range wordArray {
		spacedWordArray := strings.Split(word, DelimiterSpace2)
		for j, spacedWord := range spacedWordArray {
			for k, character := range spacedWord {
				if k != 0 {
					result += DelimiterSlash2
				}
				morsable := getMorsable(character)
				if morsable == NO_DEF {
					return GetNoDefMessage()
				}

				result += morsable.GetMorseCode()
			}

			if j < len(spacedWordArray)-1 {
				result += DelimiterSpace2
			}
		}
		if i < len(wordArray)-1 {
			result += "\n"
		}
	}

	return result
}

// TranslateToWords function
func TranslateToWords(morse string, language Language) string {
	retWords := make([]string, 0)
	morseWords := strings.Split(morse, "\n")
	for _, morseWord := range morseWords {
		// format delimiter
		formatMorse := strings.Replace(morseWord, DelimiterSlash, DelimiterSpace2, -1)
		formatMorse = strings.Replace(formatMorse, DelimiterSlash2, DelimiterSpace2, -1)
		formatMorse = strings.Replace(formatMorse, DelimiterSpace, DelimiterSpace2, -1)

		// translate morse→words
		singleMorses := strings.Split(formatMorse, DelimiterSpace2)
		var tempWords string
		for _, singleMorse := range singleMorses {
			tempWords += translateToWord(singleMorse, language)
		}
		retWords = append(retWords, tempWords)
	}

	return strings.Join(retWords, "\n")
}

func translateToWord(singleMorse string, language Language) string {
	switch language {
	case LanguageAlphabet:
		switch singleMorse {
		case A.GetMorseCode():
			return A.Character
		case B.GetMorseCode():
			return B.Character
		case C.GetMorseCode():
			return C.Character
		case D.GetMorseCode():
			return D.Character
		case E.GetMorseCode():
			return E.Character
		case F.GetMorseCode():
			return F.Character
		case G.GetMorseCode():
			return G.Character
		case H.GetMorseCode():
			return H.Character
		case I.GetMorseCode():
			return I.Character
		case J.GetMorseCode():
			return J.Character
		case K.GetMorseCode():
			return K.Character
		case L.GetMorseCode():
			return L.Character
		case M.GetMorseCode():
			return M.Character
		case N.GetMorseCode():
			return N.Character
		case O.GetMorseCode():
			return O.Character
		case P.GetMorseCode():
			return P.Character
		case Q.GetMorseCode():
			return Q.Character
		case R.GetMorseCode():
			return R.Character
		case S.GetMorseCode():
			return S.Character
		case T.GetMorseCode():
			return T.Character
		case U.GetMorseCode():
			return U.Character
		case V.GetMorseCode():
			return V.Character
		case W.GetMorseCode():
			return W.Character
		case X.GetMorseCode():
			return X.Character
		case Y.GetMorseCode():
			return Y.Character
		case Z.GetMorseCode():
			return Z.Character
		}
	case LanguageKana:
		switch singleMorse {
		case KANA_I.GetMorseCode():
			return KANA_I.Katakana
		case KANA_RO.GetMorseCode():
			return KANA_RO.Katakana
		case KANA_HA.GetMorseCode():
			return KANA_HA.Katakana
		case KANA_NI.GetMorseCode():
			return KANA_NI.Katakana
		case KANA_HO.GetMorseCode():
			return KANA_HO.Katakana
		case KANA_HE.GetMorseCode():
			return KANA_HE.Katakana
		case KANA_TO.GetMorseCode():
			return KANA_TO.Katakana
		case KANA_CHI.GetMorseCode():
			return KANA_CHI.Katakana
		case KANA_RI.GetMorseCode():
			return KANA_RI.Katakana
		case KANA_NU.GetMorseCode():
			return KANA_NU.Katakana
		case KANA_RU.GetMorseCode():
			return KANA_RU.Katakana
		case KANA_WO.GetMorseCode():
			return KANA_WO.Katakana
		case KANA_WA.GetMorseCode():
			return KANA_WA.Katakana
		case KANA_KA.GetMorseCode():
			return KANA_KA.Katakana
		case KANA_YO.GetMorseCode():
			return KANA_YO.Katakana
		case KANA_TA.GetMorseCode():
			return KANA_TA.Katakana
		case KANA_RE.GetMorseCode():
			return KANA_RE.Katakana
		case KANA_SO.GetMorseCode():
			return KANA_SO.Katakana
		case KANA_TU.GetMorseCode():
			return KANA_TU.Katakana
		case KANA_NE.GetMorseCode():
			return KANA_NE.Katakana
		case KANA_NA.GetMorseCode():
			return KANA_NA.Katakana
		case KANA_RA.GetMorseCode():
			return KANA_RA.Katakana
		case KANA_MU.GetMorseCode():
			return KANA_MU.Katakana
		case KANA_U.GetMorseCode():
			return KANA_U.Katakana
		case KANA_WI.GetMorseCode():
			return KANA_WI.Katakana
		case KANA_NO.GetMorseCode():
			return KANA_NO.Katakana
		case KANA_O.GetMorseCode():
			return KANA_O.Katakana
		case KANA_KU.GetMorseCode():
			return KANA_KU.Katakana
		case KANA_YA.GetMorseCode():
			return KANA_YA.Katakana
		case KANA_MA.GetMorseCode():
			return KANA_MA.Katakana
		case KANA_KE.GetMorseCode():
			return KANA_KE.Katakana
		case KANA_HU.GetMorseCode():
			return KANA_HU.Katakana
		case KANA_KO.GetMorseCode():
			return KANA_KO.Katakana
		case KANA_E.GetMorseCode():
			return KANA_E.Katakana
		case KANA_TE.GetMorseCode():
			return KANA_TE.Katakana
		case KANA_A.GetMorseCode():
			return KANA_A.Katakana
		case KANA_SA.GetMorseCode():
			return KANA_SA.Katakana
		case KANA_KI.GetMorseCode():
			return KANA_KI.Katakana
		case KANA_YU.GetMorseCode():
			return KANA_YU.Katakana
		case KANA_ME.GetMorseCode():
			return KANA_ME.Katakana
		case KANA_MI.GetMorseCode():
			return KANA_MI.Katakana
		case KANA_SHI.GetMorseCode():
			return KANA_SHI.Katakana
		case KANA_WE.GetMorseCode():
			return KANA_WE.Katakana
		case KANA_HI.GetMorseCode():
			return KANA_HI.Katakana
		case KANA_MO.GetMorseCode():
			return KANA_MO.Katakana
		case KANA_SE.GetMorseCode():
			return KANA_SE.Katakana
		case KANA_SU.GetMorseCode():
			return KANA_SU.Katakana
		case KANA_N.GetMorseCode():
			return KANA_N.Katakana
		case TYOUON.GetMorseCode():
			return TYOUON.Katakana
		case DAKUTEN.GetMorseCode():
			return DAKUTEN.Katakana
		case HANDAKUTEN.GetMorseCode():
			return HANDAKUTEN.Katakana
		case KUGIRITEN.GetMorseCode():
			return KUGIRITEN.Katakana
		case KANA_OTHER_GA.GetMorseCode():
			return KANA_OTHER_GA.Katakana
		case KANA_OTHER_GI.GetMorseCode():
			return KANA_OTHER_GI.Katakana
		case KANA_OTHER_GU.GetMorseCode():
			return KANA_OTHER_GU.Katakana
		case KANA_OTHER_GE.GetMorseCode():
			return KANA_OTHER_GE.Katakana
		case KANA_OTHER_GO.GetMorseCode():
			return KANA_OTHER_GO.Katakana
		case KANA_OTHER_ZA.GetMorseCode():
			return KANA_OTHER_ZA.Katakana
		case KANA_OTHER_ZI.GetMorseCode():
			return KANA_OTHER_ZI.Katakana
		case KANA_OTHER_ZU.GetMorseCode():
			return KANA_OTHER_ZU.Katakana
		case KANA_OTHER_ZE.GetMorseCode():
			return KANA_OTHER_ZE.Katakana
		case KANA_OTHER_ZO.GetMorseCode():
			return KANA_OTHER_ZO.Katakana
		case KANA_OTHER_DA.GetMorseCode():
			return KANA_OTHER_DA.Katakana
		case KANA_OTHER_DI.GetMorseCode():
			return KANA_OTHER_DI.Katakana
		case KANA_OTHER_DU.GetMorseCode():
			return KANA_OTHER_DU.Katakana
		case KANA_OTHER_DE.GetMorseCode():
			return KANA_OTHER_DE.Katakana
		case KANA_OTHER_DO.GetMorseCode():
			return KANA_OTHER_DO.Katakana
		case KANA_OTHER_BA.GetMorseCode():
			return KANA_OTHER_BA.Katakana
		case KANA_OTHER_BI.GetMorseCode():
			return KANA_OTHER_BI.Katakana
		case KANA_OTHER_BU.GetMorseCode():
			return KANA_OTHER_BU.Katakana
		case KANA_OTHER_BE.GetMorseCode():
			return KANA_OTHER_BE.Katakana
		case KANA_OTHER_BO.GetMorseCode():
			return KANA_OTHER_BO.Katakana
		case KANA_OTHER_VU.GetMorseCode():
			return KANA_OTHER_VU.Katakana
		case KANA_OTHER_PA.GetMorseCode():
			return KANA_OTHER_PA.Katakana
		case KANA_OTHER_PI.GetMorseCode():
			return KANA_OTHER_PI.Katakana
		case KANA_OTHER_PU.GetMorseCode():
			return KANA_OTHER_PU.Katakana
		case KANA_OTHER_PE.GetMorseCode():
			return KANA_OTHER_PE.Katakana
		case KANA_OTHER_PO.GetMorseCode():
			return KANA_OTHER_PO.Katakana
		case KANA_OTHER_A.GetMorseCode():
			return KANA_OTHER_A.Katakana
		case KANA_OTHER_I.GetMorseCode():
			return KANA_OTHER_I.Katakana
		case KANA_OTHER_U.GetMorseCode():
			return KANA_OTHER_U.Katakana
		case KANA_OTHER_E.GetMorseCode():
			return KANA_OTHER_E.Katakana
		case KANA_OTHER_O.GetMorseCode():
			return KANA_OTHER_O.Katakana
		case KANA_OTHER_KA.GetMorseCode():
			return KANA_OTHER_KA.Katakana
		case KANA_OTHER_KE.GetMorseCode():
			return KANA_OTHER_KE.Katakana
		case KANA_OTHER_TU.GetMorseCode():
			return KANA_OTHER_TU.Katakana
		case KANA_OTHER_YA.GetMorseCode():
			return KANA_OTHER_YA.Katakana
		case KANA_OTHER_YU.GetMorseCode():
			return KANA_OTHER_YU.Katakana
		case KANA_OTHER_YO.GetMorseCode():
			return KANA_OTHER_YO.Katakana
		case KANA_OTHER_WA.GetMorseCode():
			return KANA_OTHER_WA.Katakana
		}
	}

	switch singleMorse {
	case ZERO.GetMorseCode():
		return ZERO.Character
	case ONE.GetMorseCode():
		return ONE.Character
	case TWO.GetMorseCode():
		return TWO.Character
	case THREE.GetMorseCode():
		return THREE.Character
	case FOUR.GetMorseCode():
		return FOUR.Character
	case FIVE.GetMorseCode():
		return FIVE.Character
	case SIX.GetMorseCode():
		return SIX.Character
	case SEVEN.GetMorseCode():
		return SEVEN.Character
	case EIGHT.GetMorseCode():
		return EIGHT.Character
	case NINE.GetMorseCode():
		return NINE.Character
	}

	return "?"
}

// GetTranslateStatus function
func GetTranslateStatus(words string) TranslateStatus {
	if words == "" {
		return StatusNone
	}

	status := StatusNone
	for i, word := range words {
		if i == 0 {
			status = getSingleTranslateStatus(word)
			continue
		}

		if string(word) == DelimiterSpace || string(word) == DelimiterSpace2 || string(word) == "\n" {
			continue
		}

		morsable := getMorsable(word)
		switch morsable {
		case TYOUON:
			continue
		}

		targetStatus := getSingleTranslateStatus(word)
		if targetStatus != status {
			if status == StatusNumber && (targetStatus == StatusAlphabet || targetStatus == StatusKana) {
				status = targetStatus
				continue
			}

			if (status == StatusAlphabet || status == StatusKana) && targetStatus == StatusNumber {
				continue
			}
			return StatusInvalid
		}
	}
	return status
}

func getSingleTranslateStatus(singleWord rune) TranslateStatus {
	switch string(singleWord) {
	case Dot, Dot2, Dash, Dash2, Dash3, Dash4, DelimiterSlash,
		DelimiterSlash2, DelimiterSpace, DelimiterSpace2:
		return StatusMorse
	}

	switch getMorsable(singleWord).(type) {
	case Alphabet:
		return StatusAlphabet
	case Kana, KanaOther:
		return StatusKana
	case Number:
		return StatusNumber
	case NoDef:
		return StatusInvalid
	}

	return StatusInvalid
}
