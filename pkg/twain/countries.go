package twain

// Country is an alias for TWCY_ values.
type Country uint16

const (
	// CountryAfghanistan is an alias for TWCY_AFGHANISTAN.
	CountryAfghanistan Country = 1001

	// CountryAlgeria is an alias for TWCY_ALGERIA.
	CountryAlgeria Country = 213

	// CountryAmericanSamoa is an alias for TWCY_AMERICANSAMOA.
	CountryAmericanSamoa Country = 684

	// CountryAndorra is an alias for TWCY_ANDORRA.
	CountryAndorra Country = 33

	// CountryAngola is an alias for TWCY_ANGOLA.
	CountryAngola Country = 1002

	// CountryAnguilla is an alias for TWCY_ANGUILLA.
	CountryAnguilla Country = 8090

	// CountryAntigua is an alias for TWCY_ANTIGUA.
	CountryAntigua Country = 8091

	// CountryArgentina is an alias for TWCY_ARGENTINA.
	CountryArgentina Country = 54

	// CountryAruba is an alias for TWCY_ARUBA.
	CountryAruba Country = 297

	// CountryAscensionIs is an alias for TWCY_ASCENSIONI.
	CountryAscensionIs Country = 247

	// CountryAustralia is an alias for TWCY_AUSTRALIA.
	CountryAustralia Country = 61

	// CountryAustria is an alias for TWCY_AUSTRIA.
	CountryAustria Country = 43

	// CountryBahamas is an alias for TWCY_BAHAMAS.
	CountryBahamas Country = 8092

	// CountryBahrain is an alias for TWCY_BAHRAIN.
	CountryBahrain Country = 973

	// CountryBangladesh is an alias for TWCY_BANGLADESH.
	CountryBangladesh Country = 880

	// CountryBarbados is an alias for TWCY_BARBADOS.
	CountryBarbados Country = 8093

	// CountryBelgium is an alias for TWCY_BELGIUM.
	CountryBelgium Country = 32

	// CountryBelize is an alias for TWCY_BELIZE.
	CountryBelize Country = 501

	// CountryBenin is an alias for TWCY_BENIN.
	CountryBenin Country = 229

	// CountryBermuda is an alias for TWCY_BERMUDA.
	CountryBermuda Country = 8094

	// CountryBhutan is an alias for TWCY_BHUTAN.
	CountryBhutan Country = 1003

	// CountryBolivia is an alias for TWCY_BOLIVIA.
	CountryBolivia Country = 591

	// CountryBotswana is an alias for TWCY_BOTSWANA.
	CountryBotswana Country = 267

	// CountryBritain is an alias for TWCY_BRITAIN.
	CountryBritain Country = 6

	// CountryBritVirginIs is an alias for TWCY_BRITVIRGINIS.
	CountryBritVirginIs Country = 8095

	// CountryBrazil is an alias for TWCY_BRAZIL.
	CountryBrazil Country = 55

	// CountryBrunei is an alias for TWCY_BRUNEI.
	CountryBrunei Country = 673

	// CountryBulgaria is an alias for TWCY_BULGARIA.
	CountryBulgaria Country = 359

	// CountryBurkinaFaso is an alias for TWCY_BURKINAFASO.
	CountryBurkinaFaso Country = 1004

	// CountryBurma is an alias for TWCY_BURMA.
	CountryBurma Country = 1005

	// CountryBurundi is an alias for TWCY_BURUNDI.
	CountryBurundi Country = 1006

	// CountryCamaroon is an alias for TWCY_CAMAROON.
	CountryCamaroon Country = 237

	// CountryCanada is an alias for TWCY_CANADA.
	CountryCanada Country = 2

	// CountryCapeVerdeIs is an alias for TWCY_CAPEVERDEIS.
	CountryCapeVerdeIs Country = 238

	// CountryCaymanIs is an alias for TWCY_CAYMANIS.
	CountryCaymanIs Country = 8096

	// CountryCentralafrep is an alias for TWCY_CENTRALAFREP.
	CountryCentralafrep Country = 1007

	// CountryChad is an alias for TWCY_CHAD.
	CountryChad Country = 1008

	// CountryChile is an alias for TWCY_CHILE.
	CountryChile Country = 56

	// CountryChina is an alias for TWCY_CHINA.
	CountryChina Country = 86

	// CountryChristmasIs is an alias for TWCY_CHRISTMASIS.
	CountryChristmasIs Country = 1009

	// CountryCocosIs is an alias for TWCY_COCOSIS.
	CountryCocosIs Country = 1009

	// CountryColombia is an alias for TWCY_COLOMBIA.
	CountryColombia Country = 57

	// CountryComoros is an alias for TWCY_COMOROS.
	CountryComoros Country = 1010

	// CountryCongo is an alias for TWCY_CONGO.
	CountryCongo Country = 1011

	// CountryCookIs is an alias for TWCY_COOKIS.
	CountryCookIs Country = 1012

	// CountryCostarica is an alias for TWCY_COSTARICA.
	CountryCostarica Country = 506

	// CountryCuba is an alias for TWCY_CUBA.
	CountryCuba Country = 5

	// CountryCyprus is an alias for TWCY_CYPRUS.
	CountryCyprus Country = 357

	// CountryCzechoslovakia is an alias for TWCY_CZECHOSLOVAKIA.
	CountryCzechoslovakia Country = 42

	// CountryDenmark is an alias for TWCY_DENMARK.
	CountryDenmark Country = 45

	// CountryDjibouti is an alias for TWCY_DJIBOUTI.
	CountryDjibouti Country = 1013

	// CountryDominica is an alias for TWCY_DOMINICA.
	CountryDominica Country = 8097

	// CountryDomincanrep is an alias for TWCY_DOMINCANREP.
	CountryDomincanrep Country = 8098

	// CountryEasterIs is an alias for TWCY_EASTERIS.
	CountryEasterIs Country = 1014

	// CountryEcuador is an alias for TWCY_ECUADOR.
	CountryEcuador Country = 593

	// CountryEgypt is an alias for TWCY_EGYPT.
	CountryEgypt Country = 20

	// CountryElsalvador is an alias for TWCY_ELSALVADOR.
	CountryElsalvador Country = 503

	// CountryEqguinea is an alias for TWCY_EQGUINEA.
	CountryEqguinea Country = 1015

	// CountryEthiopia is an alias for TWCY_ETHIOPIA.
	CountryEthiopia Country = 251

	// CountryFalklandIs is an alias for TWCY_FALKLANDIS.
	CountryFalklandIs Country = 1016

	// CountryFaeroeIs is an alias for TWCY_FAEROEIS.
	CountryFaeroeIs Country = 298

	// CountryFijiislands is an alias for TWCY_FIJIISLANDS.
	CountryFijiislands Country = 679

	// CountryFinland is an alias for TWCY_FINLAND.
	CountryFinland Country = 358

	// CountryFrance is an alias for TWCY_FRANCE.
	CountryFrance Country = 33

	// CountryFrantilles is an alias for TWCY_FRANTILLES.
	CountryFrantilles Country = 596

	// CountryFrguiana is an alias for TWCY_FRGUIANA.
	CountryFrguiana Country = 594

	// CountryFrpolyneisa is an alias for TWCY_FRPOLYNEISA.
	CountryFrpolyneisa Country = 689

	// CountryFutanaIs is an alias for TWCY_FUTANAIS.
	CountryFutanaIs Country = 1043

	// CountryGabon is an alias for TWCY_GABON.
	CountryGabon Country = 241

	// CountryGambia is an alias for TWCY_GAMBIA.
	CountryGambia Country = 220

	// CountryGermany is an alias for TWCY_GERMANY.
	CountryGermany Country = 49

	// CountryGhana is an alias for TWCY_GHANA.
	CountryGhana Country = 233

	// CountryGibralter is an alias for TWCY_GIBRALTER.
	CountryGibralter Country = 350

	// CountryGreece is an alias for TWCY_GREECE.
	CountryGreece Country = 30

	// CountryGreenland is an alias for TWCY_GREENLAND.
	CountryGreenland Country = 299

	// CountryGrenada is an alias for TWCY_GRENADA.
	CountryGrenada Country = 8099

	// CountryGrenedines is an alias for TWCY_GRENEDINES.
	CountryGrenedines Country = 8015

	// CountryGuadeloupe is an alias for TWCY_GUADELOUPE.
	CountryGuadeloupe Country = 590

	// CountryGuam is an alias for TWCY_GUAM.
	CountryGuam Country = 671

	// CountryGuantanamobay is an alias for TWCY_GUANTANAMOBAY.
	CountryGuantanamobay Country = 5399

	// CountryGuatemala is an alias for TWCY_GUATEMALA.
	CountryGuatemala Country = 502

	// CountryGuinea is an alias for TWCY_GUINEA.
	CountryGuinea Country = 224

	// CountryGuineabissau is an alias for TWCY_GUINEABISSAU.
	CountryGuineabissau Country = 1017

	// CountryGuyana is an alias for TWCY_GUYANA.
	CountryGuyana Country = 592

	// CountryHaiti is an alias for TWCY_HAITI.
	CountryHaiti Country = 509

	// CountryHonduras is an alias for TWCY_HONDURAS.
	CountryHonduras Country = 504

	// CountryHongkong is an alias for TWCY_HONGKONG.
	CountryHongkong Country = 852

	// CountryHungary is an alias for TWCY_HUNGARY.
	CountryHungary Country = 36

	// CountryIceland is an alias for TWCY_ICELAND.
	CountryIceland Country = 354

	// CountryIndia is an alias for TWCY_INDIA.
	CountryIndia Country = 91

	// CountryIndonesia is an alias for TWCY_INDONESIA.
	CountryIndonesia Country = 62

	// CountryIran is an alias for TWCY_IRAN.
	CountryIran Country = 98

	// CountryIraq is an alias for TWCY_IRAQ.
	CountryIraq Country = 964

	// CountryIreland is an alias for TWCY_IRELAND.
	CountryIreland Country = 353

	// CountryIsrael is an alias for TWCY_ISRAEL.
	CountryIsrael Country = 972

	// CountryItaly is an alias for TWCY_ITALY.
	CountryItaly Country = 39

	// CountryIvorycoast is an alias for TWCY_IVORYCOAST.
	CountryIvorycoast Country = 225

	// CountryJamaica is an alias for TWCY_JAMAICA.
	CountryJamaica Country = 8010

	// CountryJapan is an alias for TWCY_JAPAN.
	CountryJapan Country = 81

	// CountryJordan is an alias for TWCY_JORDAN.
	CountryJordan Country = 962

	// CountryKenya is an alias for TWCY_KENYA.
	CountryKenya Country = 254

	// CountryKiribati is an alias for TWCY_KIRIBATI.
	CountryKiribati Country = 1018

	// CountryKorea is an alias for TWCY_KOREA.
	CountryKorea Country = 82

	// CountryKuwait is an alias for TWCY_KUWAIT.
	CountryKuwait Country = 965

	// CountryLaos is an alias for TWCY_LAOS.
	CountryLaos Country = 1019

	// CountryLebanon is an alias for TWCY_LEBANON.
	CountryLebanon Country = 1020

	// CountryLiberia is an alias for TWCY_LIBERIA.
	CountryLiberia Country = 231

	// CountryLibya is an alias for TWCY_LIBYA.
	CountryLibya Country = 218

	// CountryLiechtenstein is an alias for TWCY_LIECHTENSTEIN.
	CountryLiechtenstein Country = 41

	// CountryLuxenbourg is an alias for TWCY_LUXENBOURG.
	CountryLuxenbourg Country = 352

	// CountryMacao is an alias for TWCY_MACAO.
	CountryMacao Country = 853

	// CountryMadagascar is an alias for TWCY_MADAGASCAR.
	CountryMadagascar Country = 1021

	// CountryMalawi is an alias for TWCY_MALAWI.
	CountryMalawi Country = 265

	// CountryMalaysia is an alias for TWCY_MALAYSIA.
	CountryMalaysia Country = 60

	// CountryMaldives is an alias for TWCY_MALDIVES.
	CountryMaldives Country = 960

	// CountryMali is an alias for TWCY_MALI.
	CountryMali Country = 1022

	// CountryMalta is an alias for TWCY_MALTA.
	CountryMalta Country = 356

	// CountryMarshallIs is an alias for TWCY_MARSHALLIS.
	CountryMarshallIs Country = 692

	// CountryMauritania is an alias for TWCY_MAURITANIA.
	CountryMauritania Country = 1023

	// CountryMauritius is an alias for TWCY_MAURITIUS.
	CountryMauritius Country = 230

	// CountryMexico is an alias for TWCY_MEXICO.
	CountryMexico Country = 3

	// CountryMicronesia is an alias for TWCY_MICRONESIA.
	CountryMicronesia Country = 691

	// CountryMiquelon is an alias for TWCY_MIQUELON.
	CountryMiquelon Country = 508

	// CountryMonaco is an alias for TWCY_MONACO.
	CountryMonaco Country = 33

	// CountryMongolia is an alias for TWCY_MONGOLIA.
	CountryMongolia Country = 1024

	// CountryMontserrat is an alias for TWCY_MONTSERRAT.
	CountryMontserrat Country = 8011

	// CountryMorocco is an alias for TWCY_MOROCCO.
	CountryMorocco Country = 212

	// CountryMozambique is an alias for TWCY_MOZAMBIQUE.
	CountryMozambique Country = 1025

	// CountryNamibia is an alias for TWCY_NAMIBIA.
	CountryNamibia Country = 264

	// CountryNauru is an alias for TWCY_NAURU.
	CountryNauru Country = 1026

	// CountryNepal is an alias for TWCY_NEPAL.
	CountryNepal Country = 977

	// CountryNetherlands is an alias for TWCY_NETHERLANDS.
	CountryNetherlands Country = 31

	// CountryNethAntilles is an alias for TWCY_NETHANTILLES.
	CountryNethAntilles Country = 599

	// CountryNevis is an alias for TWCY_NEVIS.
	CountryNevis Country = 8012

	// CountryNewCaledonia is an alias for TWCY_NEWCALEDONIA.
	CountryNewCaledonia Country = 687

	// CountryNewZealand is an alias for TWCY_NEWZEALAND.
	CountryNewZealand Country = 64

	// CountryNicaragua is an alias for TWCY_NICARAGUA.
	CountryNicaragua Country = 505

	// CountryNiger is an alias for TWCY_NIGER.
	CountryNiger Country = 227

	// CountryNigeria is an alias for TWCY_NIGERIA.
	CountryNigeria Country = 234

	// CountryNiue is an alias for TWCY_NIUE.
	CountryNiue Country = 1027

	// CountryNorfolkIs is an alias for TWCY_NORFOLKI.
	CountryNorfolkIs Country = 1028

	// CountryNorway is an alias for TWCY_NORWAY.
	CountryNorway Country = 47

	// CountryOman is an alias for TWCY_OMAN.
	CountryOman Country = 968

	// CountryPakistan is an alias for TWCY_PAKISTAN.
	CountryPakistan Country = 92

	// CountryPalau is an alias for TWCY_PALAU.
	CountryPalau Country = 1029

	// CountryPanama is an alias for TWCY_PANAMA.
	CountryPanama Country = 507

	// CountryParaguay is an alias for TWCY_PARAGUAY.
	CountryParaguay Country = 595

	// CountryPeru is an alias for TWCY_PERU.
	CountryPeru Country = 51

	// CountryPhillippines is an alias for TWCY_PHILLIPPINES.
	CountryPhillippines Country = 63

	// CountryPitcairnIs is an alias for TWCY_PITCAIRNIS.
	CountryPitcairnIs Country = 1030

	// CountryPNewGuinea is an alias for TWCY_PNEWGUINEA.
	CountryPNewGuinea Country = 675

	// CountryPoland is an alias for TWCY_POLAND.
	CountryPoland Country = 48

	// CountryPortugal is an alias for TWCY_PORTUGAL.
	CountryPortugal Country = 351

	// CountryQatar is an alias for TWCY_QATAR.
	CountryQatar Country = 974

	// CountryReunionIs is an alias for TWCY_REUNIONI.
	CountryReunionIs Country = 1031

	// CountryRomania is an alias for TWCY_ROMANIA.
	CountryRomania Country = 40

	// CountryRwanda is an alias for TWCY_RWANDA.
	CountryRwanda Country = 250

	// CountrySaipan is an alias for TWCY_SAIPAN.
	CountrySaipan Country = 670

	// CountrySanMarino is an alias for TWCY_SANMARINO.
	CountrySanMarino Country = 39

	// CountrySaoTome is an alias for TWCY_SAOTOME.
	CountrySaoTome Country = 1033

	// CountrySaudiArabia is an alias for TWCY_SAUDIARABIA.
	CountrySaudiArabia Country = 966

	// CountrySenegal is an alias for TWCY_SENEGAL.
	CountrySenegal Country = 221

	// CountrySeychellesIs is an alias for TWCY_SEYCHELLESIS.
	CountrySeychellesIs Country = 1034

	// CountrySierraLeone is an alias for TWCY_SIERRALEONE.
	CountrySierraLeone Country = 1035

	// CountrySingapore is an alias for TWCY_SINGAPORE.
	CountrySingapore Country = 65

	// CountrySolomonIs is an alias for TWCY_SOLOMONIS.
	CountrySolomonIs Country = 1036

	// CountrySomali is an alias for TWCY_SOMALI.
	CountrySomali Country = 1037

	// CountrySouthAfrica is an alias for TWCY_SOUTHAFRICA.
	CountrySouthAfrica Country = 27

	// CountrySpain is an alias for TWCY_SPAIN.
	CountrySpain Country = 34

	// CountrySriLanka is an alias for TWCY_SRILANKA.
	CountrySriLanka Country = 94

	// CountryStHelena is an alias for TWCY_STHELENA.
	CountryStHelena Country = 1032

	// CountryStKitts is an alias for TWCY_STKITTS.
	CountryStKitts Country = 8013

	// CountryStLucia is an alias for TWCY_STLUCIA.
	CountryStLucia Country = 8014

	// CountryStPierre is an alias for TWCY_STPIERRE.
	CountryStPierre Country = 508

	// CountryStVincent is an alias for TWCY_STVINCENT.
	CountryStVincent Country = 8015

	// CountrySudan is an alias for TWCY_SUDAN.
	CountrySudan Country = 1038

	// CountrySuriname is an alias for TWCY_SURINAME.
	CountrySuriname Country = 597

	// CountrySwaziland is an alias for TWCY_SWAZILAND.
	CountrySwaziland Country = 268

	// CountrySweden is an alias for TWCY_SWEDEN.
	CountrySweden Country = 46

	// CountrySwitzerland is an alias for TWCY_SWITZERLAND.
	CountrySwitzerland Country = 41

	// CountrySyria is an alias for TWCY_SYRIA.
	CountrySyria Country = 1039

	// CountryTaiwan is an alias for TWCY_TAIWAN.
	CountryTaiwan Country = 886

	// CountryTanzania is an alias for TWCY_TANZANIA.
	CountryTanzania Country = 255

	// CountryThailand is an alias for TWCY_THAILAND.
	CountryThailand Country = 66

	// CountryTobago is an alias for TWCY_TOBAGO.
	CountryTobago Country = 8016

	// CountryTogo is an alias for TWCY_TOGO.
	CountryTogo Country = 228

	// CountryTongaIs is an alias for TWCY_TONGAIS.
	CountryTongaIs Country = 676

	// CountryTrinidad is an alias for TWCY_TRINIDAD.
	CountryTrinidad Country = 8016

	// CountryTunisia is an alias for TWCY_TUNISIA.
	CountryTunisia Country = 216

	// CountryTurkey is an alias for TWCY_TURKEY.
	CountryTurkey Country = 90

	// CountryTurksCaicos is an alias for TWCY_TURKSCAICOS.
	CountryTurksCaicos Country = 8017

	// CountryTuvalu is an alias for TWCY_TUVALU.
	CountryTuvalu Country = 1040

	// CountryUganda is an alias for TWCY_UGANDA.
	CountryUganda Country = 256

	// CountryUSSR is an alias for TWCY_USSR.
	CountryUSSR Country = 7

	// CountryUAEmirates is an alias for TWCY_UAEMIRATES.
	CountryUAEmirates Country = 971

	// CountryUnitedKingdom is an alias for TWCY_UNITEDKINGDOM.
	CountryUnitedKingdom Country = 44

	// CountryUSA is an alias for TWCY_USA.
	CountryUSA Country = 1

	// CountryUruguay is an alias for TWCY_URUGUAY.
	CountryUruguay Country = 598

	// CountryVanuatu is an alias for TWCY_VANUATU.
	CountryVanuatu Country = 1041

	// CountryVaticanCity is an alias for TWCY_VATICANCITY.
	CountryVaticanCity Country = 39

	// CountryVenezuela is an alias for TWCY_VENEZUELA.
	CountryVenezuela Country = 58

	// CountryWake is an alias for TWCY_WAKE.
	CountryWake Country = 1042

	// CountryWallisIs is an alias for TWCY_WALLISIS.
	CountryWallisIs Country = 1043

	// CountryWesternSahara is an alias for TWCY_WESTERNSAHARA.
	CountryWesternSahara Country = 1044

	// CountryWesternSamoa is an alias for TWCY_WESTERNSAMOA.
	CountryWesternSamoa Country = 1045

	// CountryYemen is an alias for TWCY_YEMEN.
	CountryYemen Country = 1046

	// CountryYugoslavia is an alias for TWCY_YUGOSLAVIA.
	CountryYugoslavia Country = 38

	// CountryZaire is an alias for TWCY_ZAIRE.
	CountryZaire Country = 243

	// CountryZambia is an alias for TWCY_ZAMBIA.
	CountryZambia Country = 260

	// CountryZimbabwe is an alias for TWCY_ZIMBABWE.
	CountryZimbabwe Country = 263

	// CountryAlbania is an alias for TWCY_ALBANIA.
	CountryAlbania Country = 355

	// CountryArmenia is an alias for TWCY_ARMENIA.
	CountryArmenia Country = 374

	// CountryAzerbaijan is an alias for TWCY_AZERBAIJAN.
	CountryAzerbaijan Country = 994

	// CountryBelarus is an alias for TWCY_BELARUS.
	CountryBelarus Country = 375

	// CountryBosniaHerzgo is an alias for TWCY_BOSNIAHERZGO.
	CountryBosniaHerzgo Country = 387

	// CountryCambodia is an alias for TWCY_CAMBODIA.
	CountryCambodia Country = 855

	// CountryCroatia is an alias for TWCY_CROATIA.
	CountryCroatia Country = 385

	// CountryCzechRepublic is an alias for TWCY_CZECHREPUBLIC.
	CountryCzechRepublic Country = 420

	// CountryDiegogarcia is an alias for TWCY_DIEGOGARCIA.
	CountryDiegogarcia Country = 246

	// CountryEritrea is an alias for TWCY_ERITREA.
	CountryEritrea Country = 291

	// CountryEstonia is an alias for TWCY_ESTONIA.
	CountryEstonia Country = 372

	// CountryGeorgia is an alias for TWCY_GEORGIA.
	CountryGeorgia Country = 995

	// CountryLatvia is an alias for TWCY_LATVIA.
	CountryLatvia Country = 371

	// CountryLesotho is an alias for TWCY_LESOTHO.
	CountryLesotho Country = 266

	// CountryLithuania is an alias for TWCY_LITHUANIA.
	CountryLithuania Country = 370

	// CountryMacedonia is an alias for TWCY_MACEDONIA.
	CountryMacedonia Country = 389

	// CountryMayotteIs is an alias for TWCY_MAYOTTEIS.
	CountryMayotteIs Country = 269

	// CountryMoldova is an alias for TWCY_MOLDOVA.
	CountryMoldova Country = 373

	// CountryMyanmar is an alias for TWCY_MYANMAR.
	CountryMyanmar Country = 95

	// CountryNorthKorea is an alias for TWCY_NORTHKOREA.
	CountryNorthKorea Country = 850

	// CountryPuertoRico is an alias for TWCY_PUERTORICO.
	CountryPuertoRico Country = 787

	// CountryRussia is an alias for TWCY_RUSSIA.
	CountryRussia Country = 7

	// CountrySerbia is an alias for TWCY_SERBIA.
	CountrySerbia Country = 381

	// CountrySlovakia is an alias for TWCY_SLOVAKIA.
	CountrySlovakia Country = 421

	// CountrySlovenia is an alias for TWCY_SLOVENIA.
	CountrySlovenia Country = 386

	// CountrySouthKorea is an alias for TWCY_SOUTHKOREA.
	CountrySouthKorea Country = 82

	// CountryUkraine is an alias for TWCY_UKRAINE.
	CountryUkraine Country = 380

	// CountryUSVirginIs is an alias for TWCY_USVIRGINIS.
	CountryUSVirginIs Country = 340

	// CountryVietnam is an alias for TWCY_VIETNAM.
	CountryVietnam Country = 84
)
