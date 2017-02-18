package linebot

const (
	Dot             = "・"
	Dot2            = "."
	Dash            = "－"
	Dash2           = "-"
	Dash3           = "_"
	Dash4           = "ー"
	DelimiterSlash  = "／"
	DelimiterSlash2 = "/"
	DelimiterSpace  = "　"
	DelimiterSpace2 = " "
)

type Morsable interface {
  GetMorseCode() string
}

// Alhabet enum
type Alphabet struct {
	Character string
	Code      string
}

func (a Alphabet) GetMorseCode() string {
  return a.Code
}

var (
	A = Alphabet{"A", Dot + Dash}
	B = Alphabet{"B", Dash + Dot + Dot + Dot}
	C = Alphabet{"C", Dash + Dot + Dash + Dot}
	D = Alphabet{"D", Dash + Dot + Dot}
	E = Alphabet{"E", Dot}
	F = Alphabet{"F", Dot + Dot + Dash + Dot}
	G = Alphabet{"G", Dash + Dash + Dot}
	H = Alphabet{"H", Dot + Dot + Dot + Dot}
	I = Alphabet{"I", Dot + Dot}
	J = Alphabet{"J", Dot + Dash + Dash + Dash}
	K = Alphabet{"K", Dash + Dot + Dash}
	L = Alphabet{"L", Dot + Dash + Dot + Dot}
	M = Alphabet{"M", Dash + Dash}
	N = Alphabet{"N", Dash + Dot}
	O = Alphabet{"O", Dash + Dash + Dash}
	P = Alphabet{"P", Dot + Dash + Dash + Dot}
	Q = Alphabet{"Q", Dash + Dash + Dot + Dash}
	R = Alphabet{"R", Dot + Dash + Dot}
	S = Alphabet{"S", Dot + Dot + Dot}
	T = Alphabet{"T", Dash}
	U = Alphabet{"U", Dot + Dot + Dash}
	V = Alphabet{"V", Dot + Dot + Dot + Dash}
	W = Alphabet{"W", Dot + Dash + Dash}
	X = Alphabet{"X", Dash + Dot + Dot + Dash}
	Y = Alphabet{"Y", Dash + Dot + Dash + Dash}
	Z = Alphabet{"Z", Dash + Dash + Dot + Dot}
)

type Kana struct {
  Katakana string
  Hiragana string
  Code string
}

func (k Kana) GetMorseCode() string {
  return k.Code
}

var (
  KANA_I = Kana{"イ", "い", Dot + Dash}
  KANA_RO = Kana{"ロ", "ろ", Dot + Dash + Dot + Dash}
  KANA_HA = Kana{"ハ", "は", Dash + Dot + Dot + Dot}
  KANA_NI = Kana{"二", "に", Dash + Dot + Dash + Dot}
  KANA_HO = Kana{"ホ", "ほ", Dash + Dot + Dot + Dash}
  KANA_HE = Kana{"ヘ", "へ", Dot + Dash}
  KANA_TO = Kana{"ト", "と", Dot + Dot + Dash + Dot + Dot}
  KANA_CHI = Kana{"チ", "ち", Dot + Dot + Dash + Dot}
  KANA_RI = Kana{"リ", "り", Dash + Dash + Dot}
  KANA_NU = Kana{"ヌ", "ぬ", Dot + Dot + Dot + Dot}
  KANA_RU = Kana{"ル", "る", Dash + Dot + Dash + Dash + Dot}
  KANA_WO = Kana{"ヲ", "を", Dot + Dash + Dash + Dash}
  KANA_WA = Kana{"ワ", "わ", Dash + Dot + Dash}
  KANA_KA = Kana{"カ", "か", Dot + Dash + Dot + Dot}
  KANA_YO = Kana{"ヨ", "よ", Dash + Dash}
  KANA_TA = Kana{"タ", "た", Dash + Dot}
  KANA_RE = Kana{"レ", "れ", Dash + Dash + Dash}
  KANA_SO = Kana{"ソ", "そ", Dash + Dash + Dash + Dot}
  KANA_TU = Kana{"ツ", "つ", Dot + Dash + Dash + Dot}
  KANA_NE = Kana{"ネ", "ね", Dash + Dash + Dot + Dash}
  KANA_NA = Kana{"ナ", "な", Dot + Dash + Dot}
  KANA_RA = Kana{"ラ", "ら", Dot + Dot + Dot}
  KANA_MU = Kana{"ム", "む", Dash}
  KANA_U = Kana{"ウ", "う", Dot + Dot + Dash}
  KANA_WI = Kana{"ヰ", "ゐ", Dot + Dash + Dot + Dot + Dash}
  KANA_NO = Kana{"ノ", "の", Dot + Dot + Dash + Dash}
  KANA_O = Kana{"オ", "お", Dot + Dash + Dot + Dot + Dot}
  KANA_KU = Kana{"ク", "く", Dot + Dot + Dot + Dash}
  KANA_YA = Kana{"ヤ", "や", Dot + Dash + Dash}
  KANA_MA = Kana{"マ", "ま", Dash + Dot + Dot + Dash}
  KANA_KE = Kana{"ケ", "け", Dash + Dot + Dash + Dash}
  KANA_HU = Kana{"フ", "ふ", Dash + Dash + Dot + Dot}
  KANA_KO = Kana{"コ", "こ", Dash + Dash + Dash + Dash}
  KANA_E = Kana{"エ", "え", Dash + Dot + Dash + Dash + Dash}
  KANA_TE = Kana{"テ", "て", Dot + Dash + Dot + Dash + Dash}
  KANA_A = Kana{"ア", "あ", Dash + Dash + Dot + Dash + Dash}
  KANA_SA = Kana{"サ", "さ", Dash + Dot + Dash + Dot + Dash}
  KANA_KI = Kana{"キ", "き", Dash + Dot + Dash + Dot + Dot}
  KANA_YU = Kana{"ユ", "ゆ", Dash + Dot + Dot + Dash + Dash}
  KANA_ME = Kana{"メ", "め", Dash + Dot + Dot + Dot + Dash}
  KANA_MI = Kana{"ミ", "み", Dot + Dot + Dash + Dot + Dash}
  KANA_SHI = Kana{"シ", "し", Dash + Dash + Dot + Dash + Dot}
  KANA_WE = Kana{"ヱ", "ゑ", Dot + Dash + Dash + Dot + Dot}
  KANA_HI = Kana{"ヒ", "ひ", Dash + Dash + Dot + Dot + Dash}
  KANA_MO = Kana{"モ", "も", Dash + Dot + Dot + Dash + Dot}
  KANA_SE = Kana{"セ", "せ", Dot + Dash + Dash + Dash + Dot}
  KANA_SU = Kana{"ス", "す", Dash + Dash + Dash + Dot + Dash}
  KANA_N = Kana{"ン", "ん", Dot + Dash + Dot + Dash + Dot}
  TYOUON = Kana{"ー", "ー", Dot + Dash + Dash + Dot + Dash}
  DAKUTEN = Kana{"゛", "゛", Dot + Dot}
  HANDAKUTEN = Kana{"゜", "゜", Dot + Dot + Dash + Dash + Dot}
  KUGIRITEN = Kana{"、", "、", Dot + Dash + Dot + Dash + Dot + Dash}
)

type KanaOther struct {
  Katakana string
  Hiragana string
  Code string
  Type int
}

func (ko KanaOther) GetMorseCode() string {
  switch ko.Type {
  case 0:
    // 拗音
    return ko.Code
  case 1:
    // 濁音
    return ko.Code + DelimiterSlash2 + DAKUTEN.Code
  case 2:
    // 半濁音
    return ko.Code + DelimiterSlash2 + HANDAKUTEN.Code
  }

  return ""
}

var (
  KANA_OTHER_GA = KanaOther{"ガ", "が", Dot + Dash + Dot + Dot, 1}
  KANA_OTHER_GI = KanaOther{"ギ", "ぎ", Dash + Dot + Dash + Dot + Dot, 1}
  KANA_OTHER_GU = KanaOther{"グ", "ぐ", Dot + Dot + Dot + Dash, 1}
  KANA_OTHER_GE = KanaOther{"ゲ", "げ", Dash + Dot + Dash + Dash, 1}
  KANA_OTHER_GO = KanaOther{"ゴ", "ご", Dash + Dash + Dash + Dash, 1}
  KANA_OTHER_ZA = KanaOther{"ザ", "ざ", Dash + Dot + Dash + Dot + Dash, 1}
  KANA_OTHER_ZI = KanaOther{"ジ", "じ", Dash + Dash + Dot + Dash + Dot, 1}
  KANA_OTHER_ZU = KanaOther{"ズ", "ず", Dash + Dash + Dash + Dot + Dash, 1}
  KANA_OTHER_ZE = KanaOther{"ゼ", "ぜ", Dot + Dash + Dash + Dash + Dot, 1}
  KANA_OTHER_ZO = KanaOther{"ゾ", "ぞ", Dash + Dash + Dash + Dot, 1}
  KANA_OTHER_DA = KanaOther{"ダ", "だ", Dash + Dot, 1}
  KANA_OTHER_DI = KanaOther{"ヂ", "ぢ", Dot + Dot + Dash + Dot, 1}
  KANA_OTHER_DU = KanaOther{"ヅ", "づ", Dot + Dash + Dash + Dot, 1}
  KANA_OTHER_DE = KanaOther{"デ", "で", Dot + Dash + Dot + Dash + Dash, 1}
  KANA_OTHER_DO = KanaOther{"ド", "ど", Dot + Dot + Dash + Dot + Dot, 1}
  KANA_OTHER_BA = KanaOther{"バ", "ば", Dash + Dot + Dot + Dot, 1}
  KANA_OTHER_BI = KanaOther{"ビ", "び", Dash + Dash + Dot + Dot + Dash, 1}
  KANA_OTHER_BU = KanaOther{"ブ", "ぶ", Dash + Dash + Dot + Dot, 1}
  KANA_OTHER_BE = KanaOther{"ベ", "べ", Dot, 1}
  KANA_OTHER_BO = KanaOther{"ボ", "ぼ", Dash + Dot + Dot, 1}
  KANA_OTHER_VU = KanaOther{"ヴ", "ゔ", Dot + Dot + Dash, 1}
  KANA_OTHER_PA = KanaOther{"パ", "ぱ", Dash + Dot + Dot + Dot, 2}
  KANA_OTHER_PI = KanaOther{"ピ", "ぴ", Dash + Dash + Dot + Dot + Dash, 2}
  KANA_OTHER_PU = KanaOther{"プ", "ぷ", Dash + Dash + Dot + Dot, 2}
  KANA_OTHER_PE = KanaOther{"ペ", "ぺ", Dot, 2}
  KANA_OTHER_PO = KanaOther{"ポ", "ぽ", Dash + Dot + Dot, 2}
  KANA_OTHER_A = KanaOther{"ァ", "ぁ", Dash + Dash + Dot + Dash + Dash, 0}
  KANA_OTHER_I = KanaOther{"ィ", "ぃ", Dot + Dash, 0}
  KANA_OTHER_U = KanaOther{"ゥ", "ぅ", Dot + Dot + Dash, 0}
  KANA_OTHER_E = KanaOther{"ェ", "ぇ", Dash + Dot + Dash + Dash + Dash, 0}
  KANA_OTHER_O = KanaOther{"ォ", "ぉ", Dot + Dash + Dot + Dot + Dot, 0}
  KANA_OTHER_KA = KanaOther{"ヵ", "ヵ", Dot + Dash + Dot + Dot, 0}
  KANA_OTHER_KE = KanaOther{"ヶ", "ヶ", Dash + Dot + Dash + Dash, 0}
  KANA_OTHER_TU = KanaOther{"ッ", "っ", Dot + Dash + Dash + Dot, 0}
  KANA_OTHER_YA = KanaOther{"ャ", "ゃ", Dot + Dash + Dash, 0}
  KANA_OTHER_YU = KanaOther{"ュ", "ゅ", Dash + Dot + Dot + Dash + Dash, 0}
  KANA_OTHER_YO = KanaOther{"ョ", "ょ", Dash + Dash, 0}
  KANA_OTHER_WA = KanaOther{"ヮ", "ゎ", Dash + Dash, 0}
)

type Number struct {
  Character string
  Code string
}

func (n Number) GetMorseCode() string {
  return n.Code
}

var (
  ZERO = Number{"0", Dash + Dash + Dash + Dash + Dash}
  ONE = Number{"1", Dot + Dash + Dash + Dash + Dash}
  TWO = Number{"2", Dot + Dot + Dash + Dash + Dash}
  THREE = Number{"3", Dot + Dot + Dot + Dash + Dash}
  FOUR = Number{"4", Dot + Dot + Dot + Dot + Dash}
  FIVE = Number{"5", Dot + Dot + Dot + Dot + Dot}
  SIX = Number{"6", Dash + Dot + Dot + Dot + Dot}
  SEVEN = Number{"7", Dash + Dash + Dot + Dot + Dot}
  EIGHT = Number{"8", Dash + Dash + Dash + Dot + Dot}
  NINE = Number{"9", Dash + Dash + Dash + Dash + Dot}
)

type NoDef struct {
  Character string
  Code string
}

func (nodef NoDef) GetMorseCode() string {
  return ""
}

func GetNoDefMessage() string {
  return "ヘンカンデキマセン"
}

var (
  NO_DEF = NoDef{"", ""}
)
