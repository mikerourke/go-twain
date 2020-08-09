package languages

import "github.com/mikerourke/go-twain/pkg/twain"

type Language twain.UInt16

const (
	// UserLocale is an alias for  TWLG_USERLOCALE.
	UserLocale Language = -1

	// DAN is an alias for  TWLG_DAN.
	DAN Language = 0

	// DUT is an alias for  TWLG_DUT.
	DUT Language = 1

	// ENG is an alias for  TWLG_ENG.
	ENG Language = 2

	// FCF is an alias for  TWLG_FCF.
	FCF Language = 3

	// FIN is an alias for  TWLG_FIN.
	FIN Language = 4

	// FRN is an alias for  TWLG_FRN.
	FRN Language = 5

	// GER is an alias for  TWLG_GER.
	GER Language = 6

	// ICE is an alias for  TWLG_ICE.
	ICE Language = 7

	// ITN is an alias for  TWLG_ITN.
	ITN Language = 8

	// NOR is an alias for  TWLG_NOR.
	NOR Language = 9

	// POR is an alias for  TWLG_POR.
	POR Language = 10

	// SPA is an alias for  TWLG_SPA.
	SPA Language = 11

	// SWE is an alias for  TWLG_SWE.
	SWE Language = 12

	// USA is an alias for  TWLG_USA.
	USA Language = 13

	// Afrikaans is an alias for  TWLG_AFRIKAANS.
	Afrikaans Language = 14

	// Albania is an alias for  TWLG_ALBANIA.
	Albania Language = 15

	// Arabic is an alias for  TWLG_ARABIC.
	Arabic Language = 16

	// ArabicAlgeria is an alias for  TWLG_ARABIC_ALGERIA.
	ArabicAlgeria Language = 17

	// ArabicBahrain is an alias for  TWLG_ARABIC_BAHRAIN.
	ArabicBahrain Language = 18

	// ArabicEgypt is an alias for  TWLG_ARABIC_EGYPT.
	ArabicEgypt Language = 19

	// ArabicIraq is an alias for  TWLG_ARABIC_IRAQ.
	ArabicIraq Language = 20

	// ArabicJordan is an alias for  TWLG_ARABIC_JORDAN.
	ArabicJordan Language = 21

	// ArabicKuwait is an alias for  TWLG_ARABIC_KUWAIT.
	ArabicKuwait Language = 22

	// ArabicLebanon is an alias for  TWLG_ARABIC_LEBANON.
	ArabicLebanon Language = 23

	// ArabicLibya is an alias for  TWLG_ARABIC_LIBYA.
	ArabicLibya Language = 24

	// ArabicMorocco is an alias for  TWLG_ARABIC_MOROCCO.
	ArabicMorocco Language = 25

	// ArabicOman is an alias for  TWLG_ARABIC_OMAN.
	ArabicOman Language = 26

	// ArabicQatar is an alias for  TWLG_ARABIC_QATAR.
	ArabicQatar Language = 27

	// ArabicSaudiArabia is an alias for  TWLG_ARABIC_SAUDIARABIA.
	ArabicSaudiArabia Language = 28

	// ArabicSyria is an alias for  TWLG_ARABIC_SYRIA.
	ArabicSyria Language = 29

	// ArabicTunisia is an alias for  TWLG_ARABIC_TUNISIA.
	ArabicTunisia Language = 30

	// ArabicUae is an alias for  TWLG_ARABIC_UAE.
	ArabicUae Language = 31

	// ArabicYemen is an alias for  TWLG_ARABIC_YEMEN.
	ArabicYemen Language = 32

	// Basque is an alias for  TWLG_BASQUE.
	Basque Language = 33

	// Byelorussian is an alias for  TWLG_BYELORUSSIAN.
	Byelorussian Language = 34

	// Bulgarian is an alias for  TWLG_BULGARIAN.
	Bulgarian Language = 35

	// Catalan is an alias for  TWLG_CATALAN.
	Catalan Language = 36

	// Chinese is an alias for  TWLG_CHINESE.
	Chinese Language = 37

	// ChineseHongKong is an alias for  TWLG_CHINESE_HONGKONG.
	ChineseHongKong Language = 38

	// ChinesePRC is an alias for  TWLG_CHINESE_PRC.
	ChinesePRC Language = 39

	// ChineseSingapore is an alias for  TWLG_CHINESE_SINGAPORE.
	ChineseSingapore Language = 40

	// ChineseSimplified is an alias for  TWLG_CHINESE_SIMPLIFIED.
	ChineseSimplified Language = 41

	// ChineseTaiwan is an alias for  TWLG_CHINESE_TAIWAN.
	ChineseTaiwan Language = 42

	// ChineseTraditional is an alias for  TWLG_CHINESE_TRADITIONAL.
	ChineseTraditional Language = 43

	// Croatia is an alias for  TWLG_CROATIA.
	Croatia Language = 44

	// Czech is an alias for  TWLG_CZECH.
	Czech Language = 45

	// Danish is an alias for  TWLG_DANISH.
	Danish Language = DAN

	// Dutch is an alias for  TWLG_DUTCH.
	Dutch Language = DUT

	// DutchBelgian is an alias for  TWLG_DUTCH_BELGIAN.
	DutchBelgian Language = 46

	// English is an alias for  TWLG_ENGLISH.
	English Language = ENG

	// EnglishAustralian is an alias for  TWLG_ENGLISH_AUSTRALIAN.
	EnglishAustralian Language = 47

	// EnglishCanadian is an alias for  TWLG_ENGLISH_CANADIAN.
	EnglishCanadian Language = 48

	// EnglishIreland is an alias for  TWLG_ENGLISH_IRELAND.
	EnglishIreland Language = 49

	// EnglishNewZealand is an alias for  TWLG_ENGLISH_NEWZEALAND.
	EnglishNewZealand Language = 50

	// EnglishSouthAfrica is an alias for  TWLG_ENGLISH_SOUTHAFRICA.
	EnglishSouthAfrica Language = 51

	// EnglishUK is an alias for  TWLG_ENGLISH_UK.
	EnglishUK Language = 52

	// EnglishUSA is an alias for  TWLG_ENGLISH_USA.
	EnglishUSA Language = USA

	// Estonian is an alias for  TWLG_ESTONIAN.
	Estonian Language = 53

	// Faeroese is an alias for  TWLG_FAEROESE.
	Faeroese Language = 54

	// Farsi is an alias for  TWLG_FARSI.
	Farsi Language = 55

	// Finnish is an alias for  TWLG_FINNISH.
	Finnish Language = FIN

	// French is an alias for  TWLG_FRENCH.
	French Language = FRN

	// FrenchBelgian is an alias for  TWLG_FRENCH_BELGIAN.
	FrenchBelgian Language = 56

	// FrenchCanadian is an alias for  TWLG_FRENCH_CANADIAN.
	FrenchCanadian Language = FCF

	// FrenchLuxembourg is an alias for  TWLG_FRENCH_LUXEMBOURG.
	FrenchLuxembourg Language = 57

	// FrenchSwiss is an alias for  TWLG_FRENCH_SWISS.
	FrenchSwiss Language = 58

	// German is an alias for  TWLG_GERMAN.
	German Language = GER

	// GermanAustrian is an alias for  TWLG_GERMAN_AUSTRIAN.
	GermanAustrian Language = 59

	// GermanLuxembourg is an alias for  TWLG_GERMAN_LUXEMBOURG.
	GermanLuxembourg Language = 60

	// GermanLiechtenstein is an alias for  TWLG_GERMAN_LIECHTENSTEIN.
	GermanLiechtenstein Language = 61

	// GermanSwiss is an alias for  TWLG_GERMAN_SWISS.
	GermanSwiss Language = 62

	// Greek is an alias for  TWLG_GREEK.
	Greek Language = 63

	// Hebrew is an alias for  TWLG_HEBREW.
	Hebrew Language = 64

	// Hungarian is an alias for  TWLG_HUNGARIAN.
	Hungarian Language = 65

	// Icelandic is an alias for  TWLG_ICELANDIC.
	Icelandic Language = ICE

	// Indonesian is an alias for  TWLG_INDONESIAN.
	Indonesian Language = 66

	// Italian is an alias for  TWLG_ITALIAN.
	Italian Language = ITN

	// ItalianSwiss is an alias for  TWLG_ITALIAN_SWISS.
	ItalianSwiss Language = 67

	// Japanese is an alias for  TWLG_JAPANESE.
	Japanese Language = 68

	// Korean is an alias for  TWLG_KOREAN.
	Korean Language = 69

	// KoreanJohab is an alias for  TWLG_KOREAN_JOHAB.
	KoreanJohab Language = 70

	// Latvian is an alias for  TWLG_LATVIAN.
	Latvian Language = 71

	// Lithuanian is an alias for  TWLG_LITHUANIAN.
	Lithuanian Language = 72

	// Norwegian is an alias for  TWLG_NORWEGIAN.
	Norwegian Language = NOR

	// NorwegianBokmal is an alias for  TWLG_NORWEGIAN_BOKMAL.
	NorwegianBokmal Language = 73

	// NorwegianNynorsk is an alias for  TWLG_NORWEGIAN_NYNORSK.
	NorwegianNynorsk Language = 74

	// Polish is an alias for  TWLG_POLISH.
	Polish Language = 75

	// Portuguese is an alias for  TWLG_PORTUGUESE.
	Portuguese Language = POR

	// PortugueseBrazil is an alias for  TWLG_PORTUGUESE_BRAZIL.
	PortugueseBrazil Language = 76

	// Romanian is an alias for  TWLG_ROMANIAN.
	Romanian Language = 77

	// Russian is an alias for  TWLG_RUSSIAN.
	Russian Language = 78

	// SerbianLatin is an alias for  TWLG_SERBIAN_LATIN.
	SerbianLatin Language = 79

	// Slovak is an alias for  TWLG_SLOVAK.
	Slovak Language = 80

	// Slovenian is an alias for  TWLG_SLOVENIAN.
	Slovenian Language = 81

	// Spanish is an alias for  TWLG_SPANISH.
	Spanish Language = SPA

	// SpanishMexican is an alias for  TWLG_SPANISH_MEXICAN.
	SpanishMexican Language = 82

	// SpanishModern is an alias for  TWLG_SPANISH_MODERN.
	SpanishModern Language = 83

	// Swedish is an alias for  TWLG_SWEDISH.
	Swedish Language = SWE

	// Thai is an alias for  TWLG_THAI.
	Thai Language = 84

	// Turkish is an alias for  TWLG_TURKISH.
	Turkish Language = 85

	// Ukranian is an alias for  TWLG_UKRANIAN.
	Ukranian Language = 86

	// Assamese is an alias for  TWLG_ASSAMESE.
	Assamese Language = 87

	// Bengali is an alias for  TWLG_BENGALI.
	Bengali Language = 88

	// Bihari is an alias for  TWLG_BIHARI.
	Bihari Language = 89

	// Bodo is an alias for  TWLG_BODO.
	Bodo Language = 90

	// Dogri is an alias for  TWLG_DOGRI.
	Dogri Language = 91

	// Gujarati is an alias for  TWLG_GUJARATI.
	Gujarati Language = 92

	// Haryanvi is an alias for  TWLG_HARYANVI.
	Haryanvi Language = 93

	// Hindi is an alias for  TWLG_HINDI.
	Hindi Language = 94

	// Kannada is an alias for  TWLG_KANNADA.
	Kannada Language = 95

	// Kashmiri is an alias for  TWLG_KASHMIRI.
	Kashmiri Language = 96

	// Malayalam is an alias for  TWLG_MALAYALAM.
	Malayalam Language = 97

	// Marathi is an alias for  TWLG_MARATHI.
	Marathi Language = 98

	// Marwari is an alias for  TWLG_MARWARI.
	Marwari Language = 99

	// Meghalayan is an alias for  TWLG_MEGHALAYAN.
	Meghalayan Language = 100

	// Mizo is an alias for  TWLG_MIZO.
	Mizo Language = 101

	// Naga is an alias for  TWLG_NAGA.
	Naga Language = 102

	// Orissi is an alias for  TWLG_ORISSI.
	Orissi Language = 103

	// Punjabi is an alias for  TWLG_PUNJABI.
	Punjabi Language = 104

	// Pushtu is an alias for  TWLG_PUSHTU.
	Pushtu Language = 105

	// SerbianCyrillic is an alias for  TWLG_SERBIAN_CYRILLIC.
	SerbianCyrillic Language = 106

	// Sikkimi is an alias for  TWLG_SIKKIMI.
	Sikkimi Language = 107

	// SwedishFinland is an alias for  TWLG_SWEDISH_FINLAND.
	SwedishFinland Language = 108

	// Tamil is an alias for  TWLG_TAMIL.
	Tamil Language = 109

	// Telugu is an alias for  TWLG_TELUGU.
	Telugu Language = 110

	// Tripuri is an alias for  TWLG_TRIPURI.
	Tripuri Language = 111

	// Urdu is an alias for  TWLG_URDU.
	Urdu Language = 112

	// Vietnamese is an alias for  TWLG_VIETNAMESE.
	Vietnamese Language = 113
)
