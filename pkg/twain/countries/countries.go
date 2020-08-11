// Package countries wraps the countries constants (prefixed with TWCY_) from
// the TWAIN header file.
package countries

import "github.com/mikerourke/go-twain/pkg/twain"

type Country twain.UInt16

const (
	// Afghanistan is an alias for  TWCY_AFGHANISTAN.
	Afghanistan Country = 1001

	// Algeria is an alias for  TWCY_ALGERIA.
	Algeria Country = 213

	// AmericanSamoa is an alias for  TWCY_AMERICANSAMOA.
	AmericanSamoa Country = 684

	// Andorra is an alias for  TWCY_ANDORRA.
	Andorra Country = 33

	// Angola is an alias for  TWCY_ANGOLA.
	Angola Country = 1002

	// Anguilla is an alias for  TWCY_ANGUILLA.
	Anguilla Country = 8090

	// Antigua is an alias for  TWCY_ANTIGUA.
	Antigua Country = 8091

	// Argentina is an alias for  TWCY_ARGENTINA.
	Argentina Country = 54

	// Aruba is an alias for  TWCY_ARUBA.
	Aruba Country = 297

	// AscensionIs is an alias for  TWCY_ASCENSIONI.
	AscensionIs Country = 247

	// Australia is an alias for  TWCY_AUSTRALIA.
	Australia Country = 61

	// Austria is an alias for  TWCY_AUSTRIA.
	Austria Country = 43

	// Bahamas is an alias for  TWCY_BAHAMAS.
	Bahamas Country = 8092

	// Bahrain is an alias for  TWCY_BAHRAIN.
	Bahrain Country = 973

	// Bangladesh is an alias for  TWCY_BANGLADESH.
	Bangladesh Country = 880

	// Barbados is an alias for  TWCY_BARBADOS.
	Barbados Country = 8093

	// Belgium is an alias for  TWCY_BELGIUM.
	Belgium Country = 32

	// Belize is an alias for  TWCY_BELIZE.
	Belize Country = 501

	// Benin is an alias for  TWCY_BENIN.
	Benin Country = 229

	// Bermuda is an alias for  TWCY_BERMUDA.
	Bermuda Country = 8094

	// Bhutan is an alias for  TWCY_BHUTAN.
	Bhutan Country = 1003

	// Bolivia is an alias for  TWCY_BOLIVIA.
	Bolivia Country = 591

	// Botswana is an alias for  TWCY_BOTSWANA.
	Botswana Country = 267

	// Britain is an alias for  TWCY_BRITAIN.
	Britain Country = 6

	// BritVirginIs is an alias for  TWCY_BRITVIRGINIS.
	BritVirginIs Country = 8095

	// Brazil is an alias for  TWCY_BRAZIL.
	Brazil Country = 55

	// Brunei is an alias for  TWCY_BRUNEI.
	Brunei Country = 673

	// Bulgaria is an alias for  TWCY_BULGARIA.
	Bulgaria Country = 359

	// BurkinaFaso is an alias for  TWCY_BURKINAFASO.
	BurkinaFaso Country = 1004

	// Burma is an alias for  TWCY_BURMA.
	Burma Country = 1005

	// Burundi is an alias for  TWCY_BURUNDI.
	Burundi Country = 1006

	// Camaroon is an alias for  TWCY_CAMAROON.
	Camaroon Country = 237

	// Canada is an alias for  TWCY_CANADA.
	Canada Country = 2

	// CapeVerdeIs is an alias for  TWCY_CAPEVERDEIS.
	CapeVerdeIs Country = 238

	// CaymanIs is an alias for  TWCY_CAYMANIS.
	CaymanIs Country = 8096

	// Centralafrep is an alias for  TWCY_CENTRALAFREP.
	Centralafrep Country = 1007

	// Chad is an alias for  TWCY_CHAD.
	Chad Country = 1008

	// Chile is an alias for  TWCY_CHILE.
	Chile Country = 56

	// China is an alias for  TWCY_CHINA.
	China Country = 86

	// ChristmasIs is an alias for  TWCY_CHRISTMASIS.
	ChristmasIs Country = 1009

	// CocosIs is an alias for  TWCY_COCOSIS.
	CocosIs Country = 1009

	// Colombia is an alias for  TWCY_COLOMBIA.
	Colombia Country = 57

	// Comoros is an alias for  TWCY_COMOROS.
	Comoros Country = 1010

	// Congo is an alias for  TWCY_CONGO.
	Congo Country = 1011

	// CookIs is an alias for  TWCY_COOKIS.
	CookIs Country = 1012

	// Costarica is an alias for  TWCY_COSTARICA.
	Costarica Country = 506

	// Cuba is an alias for  TWCY_CUBA.
	Cuba Country = 5

	// Cyprus is an alias for  TWCY_CYPRUS.
	Cyprus Country = 357

	// Czechoslovakia is an alias for  TWCY_CZECHOSLOVAKIA.
	Czechoslovakia Country = 42

	// Denmark is an alias for  TWCY_DENMARK.
	Denmark Country = 45

	// Djibouti is an alias for  TWCY_DJIBOUTI.
	Djibouti Country = 1013

	// Dominica is an alias for  TWCY_DOMINICA.
	Dominica Country = 8097

	// Domincanrep is an alias for  TWCY_DOMINCANREP.
	Domincanrep Country = 8098

	// EasterIs is an alias for  TWCY_EASTERIS.
	EasterIs Country = 1014

	// Ecuador is an alias for  TWCY_ECUADOR.
	Ecuador Country = 593

	// Egypt is an alias for  TWCY_EGYPT.
	Egypt Country = 20

	// Elsalvador is an alias for  TWCY_ELSALVADOR.
	Elsalvador Country = 503

	// Eqguinea is an alias for  TWCY_EQGUINEA.
	Eqguinea Country = 1015

	// Ethiopia is an alias for  TWCY_ETHIOPIA.
	Ethiopia Country = 251

	// FalklandIs is an alias for  TWCY_FALKLANDIS.
	FalklandIs Country = 1016

	// FaeroeIs is an alias for  TWCY_FAEROEIS.
	FaeroeIs Country = 298

	// Fijiislands is an alias for  TWCY_FIJIISLANDS.
	Fijiislands Country = 679

	// Finland is an alias for  TWCY_FINLAND.
	Finland Country = 358

	// France is an alias for  TWCY_FRANCE.
	France Country = 33

	// Frantilles is an alias for  TWCY_FRANTILLES.
	Frantilles Country = 596

	// Frguiana is an alias for  TWCY_FRGUIANA.
	Frguiana Country = 594

	// Frpolyneisa is an alias for  TWCY_FRPOLYNEISA.
	Frpolyneisa Country = 689

	// FutanaIs is an alias for  TWCY_FUTANAIS.
	FutanaIs Country = 1043

	// Gabon is an alias for  TWCY_GABON.
	Gabon Country = 241

	// Gambia is an alias for  TWCY_GAMBIA.
	Gambia Country = 220

	// Germany is an alias for  TWCY_GERMANY.
	Germany Country = 49

	// Ghana is an alias for  TWCY_GHANA.
	Ghana Country = 233

	// Gibralter is an alias for  TWCY_GIBRALTER.
	Gibralter Country = 350

	// Greece is an alias for  TWCY_GREECE.
	Greece Country = 30

	// Greenland is an alias for  TWCY_GREENLAND.
	Greenland Country = 299

	// Grenada is an alias for  TWCY_GRENADA.
	Grenada Country = 8099

	// Grenedines is an alias for  TWCY_GRENEDINES.
	Grenedines Country = 8015

	// Guadeloupe is an alias for  TWCY_GUADELOUPE.
	Guadeloupe Country = 590

	// Guam is an alias for  TWCY_GUAM.
	Guam Country = 671

	// Guantanamobay is an alias for  TWCY_GUANTANAMOBAY.
	Guantanamobay Country = 5399

	// Guatemala is an alias for  TWCY_GUATEMALA.
	Guatemala Country = 502

	// Guinea is an alias for  TWCY_GUINEA.
	Guinea Country = 224

	// Guineabissau is an alias for  TWCY_GUINEABISSAU.
	Guineabissau Country = 1017

	// Guyana is an alias for  TWCY_GUYANA.
	Guyana Country = 592

	// Haiti is an alias for  TWCY_HAITI.
	Haiti Country = 509

	// Honduras is an alias for  TWCY_HONDURAS.
	Honduras Country = 504

	// Hongkong is an alias for  TWCY_HONGKONG.
	Hongkong Country = 852

	// Hungary is an alias for  TWCY_HUNGARY.
	Hungary Country = 36

	// Iceland is an alias for  TWCY_ICELAND.
	Iceland Country = 354

	// India is an alias for  TWCY_INDIA.
	India Country = 91

	// Indonesia is an alias for  TWCY_INDONESIA.
	Indonesia Country = 62

	// Iran is an alias for  TWCY_IRAN.
	Iran Country = 98

	// Iraq is an alias for  TWCY_IRAQ.
	Iraq Country = 964

	// Ireland is an alias for  TWCY_IRELAND.
	Ireland Country = 353

	// Israel is an alias for  TWCY_ISRAEL.
	Israel Country = 972

	// Italy is an alias for  TWCY_ITALY.
	Italy Country = 39

	// Ivorycoast is an alias for  TWCY_IVORYCOAST.
	Ivorycoast Country = 225

	// Jamaica is an alias for  TWCY_JAMAICA.
	Jamaica Country = 8010

	// Japan is an alias for  TWCY_JAPAN.
	Japan Country = 81

	// Jordan is an alias for  TWCY_JORDAN.
	Jordan Country = 962

	// Kenya is an alias for  TWCY_KENYA.
	Kenya Country = 254

	// Kiribati is an alias for  TWCY_KIRIBATI.
	Kiribati Country = 1018

	// Korea is an alias for  TWCY_KOREA.
	Korea Country = 82

	// Kuwait is an alias for  TWCY_KUWAIT.
	Kuwait Country = 965

	// Laos is an alias for  TWCY_LAOS.
	Laos Country = 1019

	// Lebanon is an alias for  TWCY_LEBANON.
	Lebanon Country = 1020

	// Liberia is an alias for  TWCY_LIBERIA.
	Liberia Country = 231

	// Libya is an alias for  TWCY_LIBYA.
	Libya Country = 218

	// Liechtenstein is an alias for  TWCY_LIECHTENSTEIN.
	Liechtenstein Country = 41

	// Luxenbourg is an alias for  TWCY_LUXENBOURG.
	Luxenbourg Country = 352

	// Macao is an alias for  TWCY_MACAO.
	Macao Country = 853

	// Madagascar is an alias for  TWCY_MADAGASCAR.
	Madagascar Country = 1021

	// Malawi is an alias for  TWCY_MALAWI.
	Malawi Country = 265

	// Malaysia is an alias for  TWCY_MALAYSIA.
	Malaysia Country = 60

	// Maldives is an alias for  TWCY_MALDIVES.
	Maldives Country = 960

	// Mali is an alias for  TWCY_MALI.
	Mali Country = 1022

	// Malta is an alias for  TWCY_MALTA.
	Malta Country = 356

	// MarshallIs is an alias for  TWCY_MARSHALLIS.
	MarshallIs Country = 692

	// Mauritania is an alias for  TWCY_MAURITANIA.
	Mauritania Country = 1023

	// Mauritius is an alias for  TWCY_MAURITIUS.
	Mauritius Country = 230

	// Mexico is an alias for  TWCY_MEXICO.
	Mexico Country = 3

	// Micronesia is an alias for  TWCY_MICRONESIA.
	Micronesia Country = 691

	// Miquelon is an alias for  TWCY_MIQUELON.
	Miquelon Country = 508

	// Monaco is an alias for  TWCY_MONACO.
	Monaco Country = 33

	// Mongolia is an alias for  TWCY_MONGOLIA.
	Mongolia Country = 1024

	// Montserrat is an alias for  TWCY_MONTSERRAT.
	Montserrat Country = 8011

	// Morocco is an alias for  TWCY_MOROCCO.
	Morocco Country = 212

	// Mozambique is an alias for  TWCY_MOZAMBIQUE.
	Mozambique Country = 1025

	// Namibia is an alias for  TWCY_NAMIBIA.
	Namibia Country = 264

	// Nauru is an alias for  TWCY_NAURU.
	Nauru Country = 1026

	// Nepal is an alias for  TWCY_NEPAL.
	Nepal Country = 977

	// Netherlands is an alias for  TWCY_NETHERLANDS.
	Netherlands Country = 31

	// NethAntilles is an alias for  TWCY_NETHANTILLES.
	NethAntilles Country = 599

	// Nevis is an alias for  TWCY_NEVIS.
	Nevis Country = 8012

	// NewCaledonia is an alias for  TWCY_NEWCALEDONIA.
	NewCaledonia Country = 687

	// NewZealand is an alias for  TWCY_NEWZEALAND.
	NewZealand Country = 64

	// Nicaragua is an alias for  TWCY_NICARAGUA.
	Nicaragua Country = 505

	// Niger is an alias for  TWCY_NIGER.
	Niger Country = 227

	// Nigeria is an alias for  TWCY_NIGERIA.
	Nigeria Country = 234

	// Niue is an alias for  TWCY_NIUE.
	Niue Country = 1027

	// NorfolkIs is an alias for  TWCY_NORFOLKI.
	NorfolkIs Country = 1028

	// Norway is an alias for  TWCY_NORWAY.
	Norway Country = 47

	// Oman is an alias for  TWCY_OMAN.
	Oman Country = 968

	// Pakistan is an alias for  TWCY_PAKISTAN.
	Pakistan Country = 92

	// Palau is an alias for  TWCY_PALAU.
	Palau Country = 1029

	// Panama is an alias for  TWCY_PANAMA.
	Panama Country = 507

	// Paraguay is an alias for  TWCY_PARAGUAY.
	Paraguay Country = 595

	// Peru is an alias for  TWCY_PERU.
	Peru Country = 51

	// Phillippines is an alias for  TWCY_PHILLIPPINES.
	Phillippines Country = 63

	// PitcairnIs is an alias for  TWCY_PITCAIRNIS.
	PitcairnIs Country = 1030

	// PNewGuinea is an alias for  TWCY_PNEWGUINEA.
	PNewGuinea Country = 675

	// Poland is an alias for  TWCY_POLAND.
	Poland Country = 48

	// Portugal is an alias for  TWCY_PORTUGAL.
	Portugal Country = 351

	// Qatar is an alias for  TWCY_QATAR.
	Qatar Country = 974

	// ReunionIs is an alias for  TWCY_REUNIONI.
	ReunionIs Country = 1031

	// Romania is an alias for  TWCY_ROMANIA.
	Romania Country = 40

	// Rwanda is an alias for  TWCY_RWANDA.
	Rwanda Country = 250

	// Saipan is an alias for  TWCY_SAIPAN.
	Saipan Country = 670

	// SanMarino is an alias for  TWCY_SANMARINO.
	SanMarino Country = 39

	// SaoTome is an alias for  TWCY_SAOTOME.
	SaoTome Country = 1033

	// SaudiArabia is an alias for  TWCY_SAUDIARABIA.
	SaudiArabia Country = 966

	// Senegal is an alias for  TWCY_SENEGAL.
	Senegal Country = 221

	// SeychellesIs is an alias for  TWCY_SEYCHELLESIS.
	SeychellesIs Country = 1034

	// SierraLeone is an alias for  TWCY_SIERRALEONE.
	SierraLeone Country = 1035

	// Singapore is an alias for  TWCY_SINGAPORE.
	Singapore Country = 65

	// SolomonIs is an alias for  TWCY_SOLOMONIS.
	SolomonIs Country = 1036

	// Somali is an alias for  TWCY_SOMALI.
	Somali Country = 1037

	// SouthAfrica is an alias for  TWCY_SOUTHAFRICA.
	SouthAfrica Country = 27

	// Spain is an alias for  TWCY_SPAIN.
	Spain Country = 34

	// SriLanka is an alias for  TWCY_SRILANKA.
	SriLanka Country = 94

	// StHelena is an alias for  TWCY_STHELENA.
	StHelena Country = 1032

	// StKitts is an alias for  TWCY_STKITTS.
	StKitts Country = 8013

	// StLucia is an alias for  TWCY_STLUCIA.
	StLucia Country = 8014

	// StPierre is an alias for  TWCY_STPIERRE.
	StPierre Country = 508

	// StVincent is an alias for  TWCY_STVINCENT.
	StVincent Country = 8015

	// Sudan is an alias for  TWCY_SUDAN.
	Sudan Country = 1038

	// Suriname is an alias for  TWCY_SURINAME.
	Suriname Country = 597

	// Swaziland is an alias for  TWCY_SWAZILAND.
	Swaziland Country = 268

	// Sweden is an alias for  TWCY_SWEDEN.
	Sweden Country = 46

	// Switzerland is an alias for  TWCY_SWITZERLAND.
	Switzerland Country = 41

	// Syria is an alias for  TWCY_SYRIA.
	Syria Country = 1039

	// Taiwan is an alias for  TWCY_TAIWAN.
	Taiwan Country = 886

	// Tanzania is an alias for  TWCY_TANZANIA.
	Tanzania Country = 255

	// Thailand is an alias for  TWCY_THAILAND.
	Thailand Country = 66

	// Tobago is an alias for  TWCY_TOBAGO.
	Tobago Country = 8016

	// Togo is an alias for  TWCY_TOGO.
	Togo Country = 228

	// TongaIs is an alias for  TWCY_TONGAIS.
	TongaIs Country = 676

	// Trinidad is an alias for  TWCY_TRINIDAD.
	Trinidad Country = 8016

	// Tunisia is an alias for  TWCY_TUNISIA.
	Tunisia Country = 216

	// Turkey is an alias for  TWCY_TURKEY.
	Turkey Country = 90

	// TurksCaicos is an alias for  TWCY_TURKSCAICOS.
	TurksCaicos Country = 8017

	// Tuvalu is an alias for  TWCY_TUVALU.
	Tuvalu Country = 1040

	// Uganda is an alias for  TWCY_UGANDA.
	Uganda Country = 256

	// USSR is an alias for  TWCY_USSR.
	USSR Country = 7

	// UAEmirates is an alias for  TWCY_UAEMIRATES.
	UAEmirates Country = 971

	// UnitedKingdom is an alias for  TWCY_UNITEDKINGDOM.
	UnitedKingdom Country = 44

	// USA is an alias for  TWCY_USA.
	USA Country = 1

	// Uruguay is an alias for  TWCY_URUGUAY.
	Uruguay Country = 598

	// Vanuatu is an alias for  TWCY_VANUATU.
	Vanuatu Country = 1041

	// VaticanCity is an alias for  TWCY_VATICANCITY.
	VaticanCity Country = 39

	// Venezuela is an alias for  TWCY_VENEZUELA.
	Venezuela Country = 58

	// Wake is an alias for  TWCY_WAKE.
	Wake Country = 1042

	// WallisIs is an alias for  TWCY_WALLISIS.
	WallisIs Country = 1043

	// WesternSahara is an alias for  TWCY_WESTERNSAHARA.
	WesternSahara Country = 1044

	// WesternSamoa is an alias for  TWCY_WESTERNSAMOA.
	WesternSamoa Country = 1045

	// Yemen is an alias for  TWCY_YEMEN.
	Yemen Country = 1046

	// Yugoslavia is an alias for  TWCY_YUGOSLAVIA.
	Yugoslavia Country = 38

	// Zaire is an alias for  TWCY_ZAIRE.
	Zaire Country = 243

	// Zambia is an alias for  TWCY_ZAMBIA.
	Zambia Country = 260

	// Zimbabwe is an alias for  TWCY_ZIMBABWE.
	Zimbabwe Country = 263

	// Albania is an alias for  TWCY_ALBANIA.
	Albania Country = 355

	// Armenia is an alias for  TWCY_ARMENIA.
	Armenia Country = 374

	// Azerbaijan is an alias for  TWCY_AZERBAIJAN.
	Azerbaijan Country = 994

	// Belarus is an alias for  TWCY_BELARUS.
	Belarus Country = 375

	// BosniaHerzgo is an alias for  TWCY_BOSNIAHERZGO.
	BosniaHerzgo Country = 387

	// Cambodia is an alias for  TWCY_CAMBODIA.
	Cambodia Country = 855

	// Croatia is an alias for  TWCY_CROATIA.
	Croatia Country = 385

	// CzechRepublic is an alias for  TWCY_CZECHREPUBLIC.
	CzechRepublic Country = 420

	// Diegogarcia is an alias for  TWCY_DIEGOGARCIA.
	Diegogarcia Country = 246

	// Eritrea is an alias for  TWCY_ERITREA.
	Eritrea Country = 291

	// Estonia is an alias for  TWCY_ESTONIA.
	Estonia Country = 372

	// Georgia is an alias for  TWCY_GEORGIA.
	Georgia Country = 995

	// Latvia is an alias for  TWCY_LATVIA.
	Latvia Country = 371

	// Lesotho is an alias for  TWCY_LESOTHO.
	Lesotho Country = 266

	// Lithuania is an alias for  TWCY_LITHUANIA.
	Lithuania Country = 370

	// Macedonia is an alias for  TWCY_MACEDONIA.
	Macedonia Country = 389

	// MayotteIs is an alias for  TWCY_MAYOTTEIS.
	MayotteIs Country = 269

	// Moldova is an alias for  TWCY_MOLDOVA.
	Moldova Country = 373

	// Myanmar is an alias for  TWCY_MYANMAR.
	Myanmar Country = 95

	// NorthKorea is an alias for  TWCY_NORTHKOREA.
	NorthKorea Country = 850

	// PuertoRico is an alias for  TWCY_PUERTORICO.
	PuertoRico Country = 787

	// Russia is an alias for  TWCY_RUSSIA.
	Russia Country = 7

	// Serbia is an alias for  TWCY_SERBIA.
	Serbia Country = 381

	// Slovakia is an alias for  TWCY_SLOVAKIA.
	Slovakia Country = 421

	// Slovenia is an alias for  TWCY_SLOVENIA.
	Slovenia Country = 386

	// SouthKorea is an alias for  TWCY_SOUTHKOREA.
	SouthKorea Country = 82

	// Ukraine is an alias for  TWCY_UKRAINE.
	Ukraine Country = 380

	// USVirginIs is an alias for  TWCY_USVIRGINIS.
	USVirginIs Country = 340

	// Vietnam is an alias for  TWCY_VIETNAM.
	Vietnam Country = 84
)
