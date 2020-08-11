package twain

// Language is an alias for the TWLG_ values.
type Language int16

const (
	// LanguageUserLocale is an alias for  TWLG_USERLOCALE.
	LanguageUserLocale Language = -1

	// LanguageDAN is an alias for  TWLG_DAN.
	LanguageDAN Language = 0

	// LanguageDUT is an alias for  TWLG_DUT.
	LanguageDUT Language = 1

	// LanguageENG is an alias for  TWLG_ENG.
	LanguageENG Language = 2

	// LanguageFCF is an alias for  TWLG_FCF.
	LanguageFCF Language = 3

	// LanguageFIN is an alias for  TWLG_FIN.
	LanguageFIN Language = 4

	// LanguageFRN is an alias for  TWLG_FRN.
	LanguageFRN Language = 5

	// LanguageGER is an alias for  TWLG_GER.
	LanguageGER Language = 6

	// LanguageICE is an alias for  TWLG_ICE.
	LanguageICE Language = 7

	// LanguageITN is an alias for  TWLG_ITN.
	LanguageITN Language = 8

	// LanguageNOR is an alias for  TWLG_NOR.
	LanguageNOR Language = 9

	// LanguagePOR is an alias for  TWLG_POR.
	LanguagePOR Language = 10

	// LanguageSPA is an alias for  TWLG_SPA.
	LanguageSPA Language = 11

	// LanguageSWE is an alias for  TWLG_SWE.
	LanguageSWE Language = 12

	// LanguageUSA is an alias for  TWLG_USA.
	LanguageUSA Language = 13

	// LanguageAfrikaans is an alias for  TWLG_AFRIKAANS.
	LanguageAfrikaans Language = 14

	// LanguageAlbania is an alias for  TWLG_ALBANIA.
	LanguageAlbania Language = 15

	// LanguageArabic is an alias for  TWLG_ARABIC.
	LanguageArabic Language = 16

	// LanguageArabicAlgeria is an alias for  TWLG_ARABIC_ALGERIA.
	LanguageArabicAlgeria Language = 17

	// LanguageArabicBahrain is an alias for  TWLG_ARABIC_BAHRAIN.
	LanguageArabicBahrain Language = 18

	// LanguageArabicEgypt is an alias for  TWLG_ARABIC_EGYPT.
	LanguageArabicEgypt Language = 19

	// LanguageArabicIraq is an alias for  TWLG_ARABIC_IRAQ.
	LanguageArabicIraq Language = 20

	// LanguageArabicJordan is an alias for  TWLG_ARABIC_JORDAN.
	LanguageArabicJordan Language = 21

	// LanguageArabicKuwait is an alias for  TWLG_ARABIC_KUWAIT.
	LanguageArabicKuwait Language = 22

	// LanguageArabicLebanon is an alias for  TWLG_ARABIC_LEBANON.
	LanguageArabicLebanon Language = 23

	// LanguageArabicLibya is an alias for  TWLG_ARABIC_LIBYA.
	LanguageArabicLibya Language = 24

	// LanguageArabicMorocco is an alias for  TWLG_ARABIC_MOROCCO.
	LanguageArabicMorocco Language = 25

	// LanguageArabicOman is an alias for  TWLG_ARABIC_OMAN.
	LanguageArabicOman Language = 26

	// LanguageArabicQatar is an alias for  TWLG_ARABIC_QATAR.
	LanguageArabicQatar Language = 27

	// LanguageArabicSaudiArabia is an alias for  TWLG_ARABIC_SAUDIARABIA.
	LanguageArabicSaudiArabia Language = 28

	// LanguageArabicSyria is an alias for  TWLG_ARABIC_SYRIA.
	LanguageArabicSyria Language = 29

	// LanguageArabicTunisia is an alias for  TWLG_ARABIC_TUNISIA.
	LanguageArabicTunisia Language = 30

	// LanguageArabicUae is an alias for  TWLG_ARABIC_UAE.
	LanguageArabicUae Language = 31

	// LanguageArabicYemen is an alias for  TWLG_ARABIC_YEMEN.
	LanguageArabicYemen Language = 32

	// LanguageBasque is an alias for  TWLG_BASQUE.
	LanguageBasque Language = 33

	// LanguageByelorussian is an alias for  TWLG_BYELORUSSIAN.
	LanguageByelorussian Language = 34

	// LanguageBulgarian is an alias for  TWLG_BULGARIAN.
	LanguageBulgarian Language = 35

	// LanguageCatalan is an alias for  TWLG_CATALAN.
	LanguageCatalan Language = 36

	// LanguageChinese is an alias for  TWLG_CHINESE.
	LanguageChinese Language = 37

	// LanguageChineseHongKong is an alias for  TWLG_CHINESE_HONGKONG.
	LanguageChineseHongKong Language = 38

	// LanguageChinesePRC is an alias for  TWLG_CHINESE_PRC.
	LanguageChinesePRC Language = 39

	// LanguageChineseSingapore is an alias for  TWLG_CHINESE_SINGAPORE.
	LanguageChineseSingapore Language = 40

	// LanguageChineseSimplified is an alias for  TWLG_CHINESE_SIMPLIFIED.
	LanguageChineseSimplified Language = 41

	// LanguageChineseTaiwan is an alias for  TWLG_CHINESE_TAIWAN.
	LanguageChineseTaiwan Language = 42

	// LanguageChineseTraditional is an alias for  TWLG_CHINESE_TRADITIONAL.
	LanguageChineseTraditional Language = 43

	// LanguageCroatia is an alias for  TWLG_CROATIA.
	LanguageCroatia Language = 44

	// LanguageCzech is an alias for  TWLG_CZECH.
	LanguageCzech Language = 45

	// LanguageDanish is an alias for  TWLG_DANISH.
	LanguageDanish Language = LanguageDAN

	// LanguageDutch is an alias for  TWLG_DUTCH.
	LanguageDutch Language = LanguageDUT

	// LanguageDutchBelgian is an alias for  TWLG_DUTCH_BELGIAN.
	LanguageDutchBelgian Language = 46

	// LanguageEnglish is an alias for  TWLG_ENGLISH.
	LanguageEnglish Language = LanguageENG

	// LanguageEnglishAustralian is an alias for  TWLG_ENGLISH_AUSTRALIAN.
	LanguageEnglishAustralian Language = 47

	// LanguageEnglishCanadian is an alias for  TWLG_ENGLISH_CANADIAN.
	LanguageEnglishCanadian Language = 48

	// LanguageEnglishIreland is an alias for  TWLG_ENGLISH_IRELAND.
	LanguageEnglishIreland Language = 49

	// LanguageEnglishNewZealand is an alias for  TWLG_ENGLISH_NEWZEALAND.
	LanguageEnglishNewZealand Language = 50

	// LanguageEnglishSouthAfrica is an alias for  TWLG_ENGLISH_SOUTHAFRICA.
	LanguageEnglishSouthAfrica Language = 51

	// LanguageEnglishUK is an alias for  TWLG_ENGLISH_UK.
	LanguageEnglishUK Language = 52

	// LanguageEnglishUSA is an alias for  TWLG_ENGLISH_USA.
	LanguageEnglishUSA Language = LanguageUSA

	// LanguageEstonian is an alias for  TWLG_ESTONIAN.
	LanguageEstonian Language = 53

	// LanguageFaeroese is an alias for  TWLG_FAEROESE.
	LanguageFaeroese Language = 54

	// LanguageFarsi is an alias for  TWLG_FARSI.
	LanguageFarsi Language = 55

	// LanguageFinnish is an alias for  TWLG_FINNISH.
	LanguageFinnish Language = LanguageFIN

	// LanguageFrench is an alias for  TWLG_FRENCH.
	LanguageFrench Language = LanguageFRN

	// LanguageFrenchBelgian is an alias for  TWLG_FRENCH_BELGIAN.
	LanguageFrenchBelgian Language = 56

	// LanguageFrenchCanadian is an alias for  TWLG_FRENCH_CANADIAN.
	LanguageFrenchCanadian Language = LanguageFCF

	// LanguageFrenchLuxembourg is an alias for  TWLG_FRENCH_LUXEMBOURG.
	LanguageFrenchLuxembourg Language = 57

	// LanguageFrenchSwiss is an alias for  TWLG_FRENCH_SWISS.
	LanguageFrenchSwiss Language = 58

	// LanguageGerman is an alias for  TWLG_GERMAN.
	LanguageGerman Language = LanguageGER

	// LanguageGermanAustrian is an alias for  TWLG_GERMAN_AUSTRIAN.
	LanguageGermanAustrian Language = 59

	// LanguageGermanLuxembourg is an alias for  TWLG_GERMAN_LUXEMBOURG.
	LanguageGermanLuxembourg Language = 60

	// LanguageGermanLiechtenstein is an alias for  TWLG_GERMAN_LIECHTENSTEIN.
	LanguageGermanLiechtenstein Language = 61

	// LanguageGermanSwiss is an alias for  TWLG_GERMAN_SWISS.
	LanguageGermanSwiss Language = 62

	// LanguageGreek is an alias for  TWLG_GREEK.
	LanguageGreek Language = 63

	// LanguageHebrew is an alias for  TWLG_HEBREW.
	LanguageHebrew Language = 64

	// LanguageHungarian is an alias for  TWLG_HUNGARIAN.
	LanguageHungarian Language = 65

	// LanguageIcelandic is an alias for  TWLG_ICELANDIC.
	LanguageIcelandic Language = LanguageICE

	// LanguageIndonesian is an alias for  TWLG_INDONESIAN.
	LanguageIndonesian Language = 66

	// LanguageItalian is an alias for  TWLG_ITALIAN.
	LanguageItalian Language = LanguageITN

	// LanguageItalianSwiss is an alias for  TWLG_ITALIAN_SWISS.
	LanguageItalianSwiss Language = 67

	// LanguageJapanese is an alias for  TWLG_JAPANESE.
	LanguageJapanese Language = 68

	// LanguageKorean is an alias for  TWLG_KOREAN.
	LanguageKorean Language = 69

	// LanguageKoreanJohab is an alias for  TWLG_KOREAN_JOHAB.
	LanguageKoreanJohab Language = 70

	// LanguageLatvian is an alias for  TWLG_LATVIAN.
	LanguageLatvian Language = 71

	// LanguageLithuanian is an alias for  TWLG_LITHUANIAN.
	LanguageLithuanian Language = 72

	// LanguageNorwegian is an alias for  TWLG_NORWEGIAN.
	LanguageNorwegian Language = LanguageNOR

	// LanguageNorwegianBokmal is an alias for  TWLG_NORWEGIAN_BOKMAL.
	LanguageNorwegianBokmal Language = 73

	// LanguageNorwegianNynorsk is an alias for  TWLG_NORWEGIAN_NYNORSK.
	LanguageNorwegianNynorsk Language = 74

	// LanguagePolish is an alias for  TWLG_POLISH.
	LanguagePolish Language = 75

	// LanguagePortuguese is an alias for  TWLG_PORTUGUESE.
	LanguagePortuguese Language = LanguagePOR

	// LanguagePortugueseBrazil is an alias for  TWLG_PORTUGUESE_BRAZIL.
	LanguagePortugueseBrazil Language = 76

	// LanguageRomanian is an alias for  TWLG_ROMANIAN.
	LanguageRomanian Language = 77

	// LanguageRussian is an alias for  TWLG_RUSSIAN.
	LanguageRussian Language = 78

	// LanguageSerbianLatin is an alias for  TWLG_SERBIAN_LATIN.
	LanguageSerbianLatin Language = 79

	// LanguageSlovak is an alias for  TWLG_SLOVAK.
	LanguageSlovak Language = 80

	// LanguageSlovenian is an alias for  TWLG_SLOVENIAN.
	LanguageSlovenian Language = 81

	// LanguageSpanish is an alias for  TWLG_SPANISH.
	LanguageSpanish Language = LanguageSPA

	// LanguageSpanishMexican is an alias for  TWLG_SPANISH_MEXICAN.
	LanguageSpanishMexican Language = 82

	// LanguageSpanishModern is an alias for  TWLG_SPANISH_MODERN.
	LanguageSpanishModern Language = 83

	// LanguageSwedish is an alias for  TWLG_SWEDISH.
	LanguageSwedish Language = LanguageSWE

	// LanguageThai is an alias for  TWLG_THAI.
	LanguageThai Language = 84

	// LanguageTurkish is an alias for  TWLG_TURKISH.
	LanguageTurkish Language = 85

	// LanguageUkranian is an alias for  TWLG_UKRANIAN.
	LanguageUkranian Language = 86

	// LanguageAssamese is an alias for  TWLG_ASSAMESE.
	LanguageAssamese Language = 87

	// LanguageBengali is an alias for  TWLG_BENGALI.
	LanguageBengali Language = 88

	// LanguageBihari is an alias for  TWLG_BIHARI.
	LanguageBihari Language = 89

	// LanguageBodo is an alias for  TWLG_BODO.
	LanguageBodo Language = 90

	// LanguageDogri is an alias for  TWLG_DOGRI.
	LanguageDogri Language = 91

	// LanguageGujarati is an alias for  TWLG_GUJARATI.
	LanguageGujarati Language = 92

	// LanguageHaryanvi is an alias for  TWLG_HARYANVI.
	LanguageHaryanvi Language = 93

	// LanguageHindi is an alias for  TWLG_HINDI.
	LanguageHindi Language = 94

	// LanguageKannada is an alias for  TWLG_KANNADA.
	LanguageKannada Language = 95

	// LanguageKashmiri is an alias for  TWLG_KASHMIRI.
	LanguageKashmiri Language = 96

	// LanguageMalayalam is an alias for  TWLG_MALAYALAM.
	LanguageMalayalam Language = 97

	// LanguageMarathi is an alias for  TWLG_MARATHI.
	LanguageMarathi Language = 98

	// LanguageMarwari is an alias for  TWLG_MARWARI.
	LanguageMarwari Language = 99

	// LanguageMeghalayan is an alias for  TWLG_MEGHALAYAN.
	LanguageMeghalayan Language = 100

	// LanguageMizo is an alias for  TWLG_MIZO.
	LanguageMizo Language = 101

	// LanguageNaga is an alias for  TWLG_NAGA.
	LanguageNaga Language = 102

	// LanguageOrissi is an alias for  TWLG_ORISSI.
	LanguageOrissi Language = 103

	// LanguagePunjabi is an alias for  TWLG_PUNJABI.
	LanguagePunjabi Language = 104

	// LanguagePushtu is an alias for  TWLG_PUSHTU.
	LanguagePushtu Language = 105

	// LanguageSerbianCyrillic is an alias for  TWLG_SERBIAN_CYRILLIC.
	LanguageSerbianCyrillic Language = 106

	// LanguageSikkimi is an alias for  TWLG_SIKKIMI.
	LanguageSikkimi Language = 107

	// LanguageSwedishFinland is an alias for  TWLG_SWEDISH_FINLAND.
	LanguageSwedishFinland Language = 108

	// LanguageTamil is an alias for  TWLG_TAMIL.
	LanguageTamil Language = 109

	// LanguageTelugu is an alias for  TWLG_TELUGU.
	LanguageTelugu Language = 110

	// LanguageTripuri is an alias for  TWLG_TRIPURI.
	LanguageTripuri Language = 111

	// LanguageUrdu is an alias for  TWLG_URDU.
	LanguageUrdu Language = 112

	// LanguageVietnamese is an alias for  TWLG_VIETNAMESE.
	LanguageVietnamese Language = 113
)
