// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports unicode'. DO NOT EDIT.

import (
	"reflect"
	"unicode"
)

func init() {
	Symbols["unicode"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ASCII_Hex_Digit":                    reflect.ValueOf(&unicode.ASCII_Hex_Digit).Elem(),
		"Adlam":                              reflect.ValueOf(&unicode.Adlam).Elem(),
		"Ahom":                               reflect.ValueOf(&unicode.Ahom).Elem(),
		"Anatolian_Hieroglyphs":              reflect.ValueOf(&unicode.Anatolian_Hieroglyphs).Elem(),
		"Arabic":                             reflect.ValueOf(&unicode.Arabic).Elem(),
		"Armenian":                           reflect.ValueOf(&unicode.Armenian).Elem(),
		"Avestan":                            reflect.ValueOf(&unicode.Avestan).Elem(),
		"AzeriCase":                          reflect.ValueOf(&unicode.AzeriCase).Elem(),
		"Balinese":                           reflect.ValueOf(&unicode.Balinese).Elem(),
		"Bamum":                              reflect.ValueOf(&unicode.Bamum).Elem(),
		"Bassa_Vah":                          reflect.ValueOf(&unicode.Bassa_Vah).Elem(),
		"Batak":                              reflect.ValueOf(&unicode.Batak).Elem(),
		"Bengali":                            reflect.ValueOf(&unicode.Bengali).Elem(),
		"Bhaiksuki":                          reflect.ValueOf(&unicode.Bhaiksuki).Elem(),
		"Bidi_Control":                       reflect.ValueOf(&unicode.Bidi_Control).Elem(),
		"Bopomofo":                           reflect.ValueOf(&unicode.Bopomofo).Elem(),
		"Brahmi":                             reflect.ValueOf(&unicode.Brahmi).Elem(),
		"Braille":                            reflect.ValueOf(&unicode.Braille).Elem(),
		"Buginese":                           reflect.ValueOf(&unicode.Buginese).Elem(),
		"Buhid":                              reflect.ValueOf(&unicode.Buhid).Elem(),
		"C":                                  reflect.ValueOf(&unicode.C).Elem(),
		"Canadian_Aboriginal":                reflect.ValueOf(&unicode.Canadian_Aboriginal).Elem(),
		"Carian":                             reflect.ValueOf(&unicode.Carian).Elem(),
		"CaseRanges":                         reflect.ValueOf(&unicode.CaseRanges).Elem(),
		"Categories":                         reflect.ValueOf(&unicode.Categories).Elem(),
		"Caucasian_Albanian":                 reflect.ValueOf(&unicode.Caucasian_Albanian).Elem(),
		"Cc":                                 reflect.ValueOf(&unicode.Cc).Elem(),
		"Cf":                                 reflect.ValueOf(&unicode.Cf).Elem(),
		"Chakma":                             reflect.ValueOf(&unicode.Chakma).Elem(),
		"Cham":                               reflect.ValueOf(&unicode.Cham).Elem(),
		"Cherokee":                           reflect.ValueOf(&unicode.Cherokee).Elem(),
		"Co":                                 reflect.ValueOf(&unicode.Co).Elem(),
		"Common":                             reflect.ValueOf(&unicode.Common).Elem(),
		"Coptic":                             reflect.ValueOf(&unicode.Coptic).Elem(),
		"Cs":                                 reflect.ValueOf(&unicode.Cs).Elem(),
		"Cuneiform":                          reflect.ValueOf(&unicode.Cuneiform).Elem(),
		"Cypriot":                            reflect.ValueOf(&unicode.Cypriot).Elem(),
		"Cyrillic":                           reflect.ValueOf(&unicode.Cyrillic).Elem(),
		"Dash":                               reflect.ValueOf(&unicode.Dash).Elem(),
		"Deprecated":                         reflect.ValueOf(&unicode.Deprecated).Elem(),
		"Deseret":                            reflect.ValueOf(&unicode.Deseret).Elem(),
		"Devanagari":                         reflect.ValueOf(&unicode.Devanagari).Elem(),
		"Diacritic":                          reflect.ValueOf(&unicode.Diacritic).Elem(),
		"Digit":                              reflect.ValueOf(&unicode.Digit).Elem(),
		"Duployan":                           reflect.ValueOf(&unicode.Duployan).Elem(),
		"Egyptian_Hieroglyphs":               reflect.ValueOf(&unicode.Egyptian_Hieroglyphs).Elem(),
		"Elbasan":                            reflect.ValueOf(&unicode.Elbasan).Elem(),
		"Ethiopic":                           reflect.ValueOf(&unicode.Ethiopic).Elem(),
		"Extender":                           reflect.ValueOf(&unicode.Extender).Elem(),
		"FoldCategory":                       reflect.ValueOf(&unicode.FoldCategory).Elem(),
		"FoldScript":                         reflect.ValueOf(&unicode.FoldScript).Elem(),
		"Georgian":                           reflect.ValueOf(&unicode.Georgian).Elem(),
		"Glagolitic":                         reflect.ValueOf(&unicode.Glagolitic).Elem(),
		"Gothic":                             reflect.ValueOf(&unicode.Gothic).Elem(),
		"Grantha":                            reflect.ValueOf(&unicode.Grantha).Elem(),
		"GraphicRanges":                      reflect.ValueOf(&unicode.GraphicRanges).Elem(),
		"Greek":                              reflect.ValueOf(&unicode.Greek).Elem(),
		"Gujarati":                           reflect.ValueOf(&unicode.Gujarati).Elem(),
		"Gurmukhi":                           reflect.ValueOf(&unicode.Gurmukhi).Elem(),
		"Han":                                reflect.ValueOf(&unicode.Han).Elem(),
		"Hangul":                             reflect.ValueOf(&unicode.Hangul).Elem(),
		"Hanunoo":                            reflect.ValueOf(&unicode.Hanunoo).Elem(),
		"Hatran":                             reflect.ValueOf(&unicode.Hatran).Elem(),
		"Hebrew":                             reflect.ValueOf(&unicode.Hebrew).Elem(),
		"Hex_Digit":                          reflect.ValueOf(&unicode.Hex_Digit).Elem(),
		"Hiragana":                           reflect.ValueOf(&unicode.Hiragana).Elem(),
		"Hyphen":                             reflect.ValueOf(&unicode.Hyphen).Elem(),
		"IDS_Binary_Operator":                reflect.ValueOf(&unicode.IDS_Binary_Operator).Elem(),
		"IDS_Trinary_Operator":               reflect.ValueOf(&unicode.IDS_Trinary_Operator).Elem(),
		"Ideographic":                        reflect.ValueOf(&unicode.Ideographic).Elem(),
		"Imperial_Aramaic":                   reflect.ValueOf(&unicode.Imperial_Aramaic).Elem(),
		"In":                                 reflect.ValueOf(unicode.In),
		"Inherited":                          reflect.ValueOf(&unicode.Inherited).Elem(),
		"Inscriptional_Pahlavi":              reflect.ValueOf(&unicode.Inscriptional_Pahlavi).Elem(),
		"Inscriptional_Parthian":             reflect.ValueOf(&unicode.Inscriptional_Parthian).Elem(),
		"Is":                                 reflect.ValueOf(unicode.Is),
		"IsControl":                          reflect.ValueOf(unicode.IsControl),
		"IsDigit":                            reflect.ValueOf(unicode.IsDigit),
		"IsGraphic":                          reflect.ValueOf(unicode.IsGraphic),
		"IsLetter":                           reflect.ValueOf(unicode.IsLetter),
		"IsLower":                            reflect.ValueOf(unicode.IsLower),
		"IsMark":                             reflect.ValueOf(unicode.IsMark),
		"IsNumber":                           reflect.ValueOf(unicode.IsNumber),
		"IsOneOf":                            reflect.ValueOf(unicode.IsOneOf),
		"IsPrint":                            reflect.ValueOf(unicode.IsPrint),
		"IsPunct":                            reflect.ValueOf(unicode.IsPunct),
		"IsSpace":                            reflect.ValueOf(unicode.IsSpace),
		"IsSymbol":                           reflect.ValueOf(unicode.IsSymbol),
		"IsTitle":                            reflect.ValueOf(unicode.IsTitle),
		"IsUpper":                            reflect.ValueOf(unicode.IsUpper),
		"Javanese":                           reflect.ValueOf(&unicode.Javanese).Elem(),
		"Join_Control":                       reflect.ValueOf(&unicode.Join_Control).Elem(),
		"Kaithi":                             reflect.ValueOf(&unicode.Kaithi).Elem(),
		"Kannada":                            reflect.ValueOf(&unicode.Kannada).Elem(),
		"Katakana":                           reflect.ValueOf(&unicode.Katakana).Elem(),
		"Kayah_Li":                           reflect.ValueOf(&unicode.Kayah_Li).Elem(),
		"Kharoshthi":                         reflect.ValueOf(&unicode.Kharoshthi).Elem(),
		"Khmer":                              reflect.ValueOf(&unicode.Khmer).Elem(),
		"Khojki":                             reflect.ValueOf(&unicode.Khojki).Elem(),
		"Khudawadi":                          reflect.ValueOf(&unicode.Khudawadi).Elem(),
		"L":                                  reflect.ValueOf(&unicode.L).Elem(),
		"Lao":                                reflect.ValueOf(&unicode.Lao).Elem(),
		"Latin":                              reflect.ValueOf(&unicode.Latin).Elem(),
		"Lepcha":                             reflect.ValueOf(&unicode.Lepcha).Elem(),
		"Letter":                             reflect.ValueOf(&unicode.Letter).Elem(),
		"Limbu":                              reflect.ValueOf(&unicode.Limbu).Elem(),
		"Linear_A":                           reflect.ValueOf(&unicode.Linear_A).Elem(),
		"Linear_B":                           reflect.ValueOf(&unicode.Linear_B).Elem(),
		"Lisu":                               reflect.ValueOf(&unicode.Lisu).Elem(),
		"Ll":                                 reflect.ValueOf(&unicode.Ll).Elem(),
		"Lm":                                 reflect.ValueOf(&unicode.Lm).Elem(),
		"Lo":                                 reflect.ValueOf(&unicode.Lo).Elem(),
		"Logical_Order_Exception":            reflect.ValueOf(&unicode.Logical_Order_Exception).Elem(),
		"Lower":                              reflect.ValueOf(&unicode.Lower).Elem(),
		"LowerCase":                          reflect.ValueOf(unicode.LowerCase),
		"Lt":                                 reflect.ValueOf(&unicode.Lt).Elem(),
		"Lu":                                 reflect.ValueOf(&unicode.Lu).Elem(),
		"Lycian":                             reflect.ValueOf(&unicode.Lycian).Elem(),
		"Lydian":                             reflect.ValueOf(&unicode.Lydian).Elem(),
		"M":                                  reflect.ValueOf(&unicode.M).Elem(),
		"Mahajani":                           reflect.ValueOf(&unicode.Mahajani).Elem(),
		"Malayalam":                          reflect.ValueOf(&unicode.Malayalam).Elem(),
		"Mandaic":                            reflect.ValueOf(&unicode.Mandaic).Elem(),
		"Manichaean":                         reflect.ValueOf(&unicode.Manichaean).Elem(),
		"Marchen":                            reflect.ValueOf(&unicode.Marchen).Elem(),
		"Mark":                               reflect.ValueOf(&unicode.Mark).Elem(),
		"Masaram_Gondi":                      reflect.ValueOf(&unicode.Masaram_Gondi).Elem(),
		"MaxASCII":                           reflect.ValueOf(unicode.MaxASCII),
		"MaxCase":                            reflect.ValueOf(unicode.MaxCase),
		"MaxLatin1":                          reflect.ValueOf(unicode.MaxLatin1),
		"MaxRune":                            reflect.ValueOf(unicode.MaxRune),
		"Mc":                                 reflect.ValueOf(&unicode.Mc).Elem(),
		"Me":                                 reflect.ValueOf(&unicode.Me).Elem(),
		"Meetei_Mayek":                       reflect.ValueOf(&unicode.Meetei_Mayek).Elem(),
		"Mende_Kikakui":                      reflect.ValueOf(&unicode.Mende_Kikakui).Elem(),
		"Meroitic_Cursive":                   reflect.ValueOf(&unicode.Meroitic_Cursive).Elem(),
		"Meroitic_Hieroglyphs":               reflect.ValueOf(&unicode.Meroitic_Hieroglyphs).Elem(),
		"Miao":                               reflect.ValueOf(&unicode.Miao).Elem(),
		"Mn":                                 reflect.ValueOf(&unicode.Mn).Elem(),
		"Modi":                               reflect.ValueOf(&unicode.Modi).Elem(),
		"Mongolian":                          reflect.ValueOf(&unicode.Mongolian).Elem(),
		"Mro":                                reflect.ValueOf(&unicode.Mro).Elem(),
		"Multani":                            reflect.ValueOf(&unicode.Multani).Elem(),
		"Myanmar":                            reflect.ValueOf(&unicode.Myanmar).Elem(),
		"N":                                  reflect.ValueOf(&unicode.N).Elem(),
		"Nabataean":                          reflect.ValueOf(&unicode.Nabataean).Elem(),
		"Nd":                                 reflect.ValueOf(&unicode.Nd).Elem(),
		"New_Tai_Lue":                        reflect.ValueOf(&unicode.New_Tai_Lue).Elem(),
		"Newa":                               reflect.ValueOf(&unicode.Newa).Elem(),
		"Nko":                                reflect.ValueOf(&unicode.Nko).Elem(),
		"Nl":                                 reflect.ValueOf(&unicode.Nl).Elem(),
		"No":                                 reflect.ValueOf(&unicode.No).Elem(),
		"Noncharacter_Code_Point":            reflect.ValueOf(&unicode.Noncharacter_Code_Point).Elem(),
		"Number":                             reflect.ValueOf(&unicode.Number).Elem(),
		"Nushu":                              reflect.ValueOf(&unicode.Nushu).Elem(),
		"Ogham":                              reflect.ValueOf(&unicode.Ogham).Elem(),
		"Ol_Chiki":                           reflect.ValueOf(&unicode.Ol_Chiki).Elem(),
		"Old_Hungarian":                      reflect.ValueOf(&unicode.Old_Hungarian).Elem(),
		"Old_Italic":                         reflect.ValueOf(&unicode.Old_Italic).Elem(),
		"Old_North_Arabian":                  reflect.ValueOf(&unicode.Old_North_Arabian).Elem(),
		"Old_Permic":                         reflect.ValueOf(&unicode.Old_Permic).Elem(),
		"Old_Persian":                        reflect.ValueOf(&unicode.Old_Persian).Elem(),
		"Old_South_Arabian":                  reflect.ValueOf(&unicode.Old_South_Arabian).Elem(),
		"Old_Turkic":                         reflect.ValueOf(&unicode.Old_Turkic).Elem(),
		"Oriya":                              reflect.ValueOf(&unicode.Oriya).Elem(),
		"Osage":                              reflect.ValueOf(&unicode.Osage).Elem(),
		"Osmanya":                            reflect.ValueOf(&unicode.Osmanya).Elem(),
		"Other":                              reflect.ValueOf(&unicode.Other).Elem(),
		"Other_Alphabetic":                   reflect.ValueOf(&unicode.Other_Alphabetic).Elem(),
		"Other_Default_Ignorable_Code_Point": reflect.ValueOf(&unicode.Other_Default_Ignorable_Code_Point).Elem(),
		"Other_Grapheme_Extend":              reflect.ValueOf(&unicode.Other_Grapheme_Extend).Elem(),
		"Other_ID_Continue":                  reflect.ValueOf(&unicode.Other_ID_Continue).Elem(),
		"Other_ID_Start":                     reflect.ValueOf(&unicode.Other_ID_Start).Elem(),
		"Other_Lowercase":                    reflect.ValueOf(&unicode.Other_Lowercase).Elem(),
		"Other_Math":                         reflect.ValueOf(&unicode.Other_Math).Elem(),
		"Other_Uppercase":                    reflect.ValueOf(&unicode.Other_Uppercase).Elem(),
		"P":                                  reflect.ValueOf(&unicode.P).Elem(),
		"Pahawh_Hmong":                       reflect.ValueOf(&unicode.Pahawh_Hmong).Elem(),
		"Palmyrene":                          reflect.ValueOf(&unicode.Palmyrene).Elem(),
		"Pattern_Syntax":                     reflect.ValueOf(&unicode.Pattern_Syntax).Elem(),
		"Pattern_White_Space":                reflect.ValueOf(&unicode.Pattern_White_Space).Elem(),
		"Pau_Cin_Hau":                        reflect.ValueOf(&unicode.Pau_Cin_Hau).Elem(),
		"Pc":                                 reflect.ValueOf(&unicode.Pc).Elem(),
		"Pd":                                 reflect.ValueOf(&unicode.Pd).Elem(),
		"Pe":                                 reflect.ValueOf(&unicode.Pe).Elem(),
		"Pf":                                 reflect.ValueOf(&unicode.Pf).Elem(),
		"Phags_Pa":                           reflect.ValueOf(&unicode.Phags_Pa).Elem(),
		"Phoenician":                         reflect.ValueOf(&unicode.Phoenician).Elem(),
		"Pi":                                 reflect.ValueOf(&unicode.Pi).Elem(),
		"Po":                                 reflect.ValueOf(&unicode.Po).Elem(),
		"Prepended_Concatenation_Mark":       reflect.ValueOf(&unicode.Prepended_Concatenation_Mark).Elem(),
		"PrintRanges":                        reflect.ValueOf(&unicode.PrintRanges).Elem(),
		"Properties":                         reflect.ValueOf(&unicode.Properties).Elem(),
		"Ps":                                 reflect.ValueOf(&unicode.Ps).Elem(),
		"Psalter_Pahlavi":                    reflect.ValueOf(&unicode.Psalter_Pahlavi).Elem(),
		"Punct":                              reflect.ValueOf(&unicode.Punct).Elem(),
		"Quotation_Mark":                     reflect.ValueOf(&unicode.Quotation_Mark).Elem(),
		"Radical":                            reflect.ValueOf(&unicode.Radical).Elem(),
		"Regional_Indicator":                 reflect.ValueOf(&unicode.Regional_Indicator).Elem(),
		"Rejang":                             reflect.ValueOf(&unicode.Rejang).Elem(),
		"ReplacementChar":                    reflect.ValueOf(unicode.ReplacementChar),
		"Runic":                              reflect.ValueOf(&unicode.Runic).Elem(),
		"S":                                  reflect.ValueOf(&unicode.S).Elem(),
		"STerm":                              reflect.ValueOf(&unicode.STerm).Elem(),
		"Samaritan":                          reflect.ValueOf(&unicode.Samaritan).Elem(),
		"Saurashtra":                         reflect.ValueOf(&unicode.Saurashtra).Elem(),
		"Sc":                                 reflect.ValueOf(&unicode.Sc).Elem(),
		"Scripts":                            reflect.ValueOf(&unicode.Scripts).Elem(),
		"Sentence_Terminal":                  reflect.ValueOf(&unicode.Sentence_Terminal).Elem(),
		"Sharada":                            reflect.ValueOf(&unicode.Sharada).Elem(),
		"Shavian":                            reflect.ValueOf(&unicode.Shavian).Elem(),
		"Siddham":                            reflect.ValueOf(&unicode.Siddham).Elem(),
		"SignWriting":                        reflect.ValueOf(&unicode.SignWriting).Elem(),
		"SimpleFold":                         reflect.ValueOf(unicode.SimpleFold),
		"Sinhala":                            reflect.ValueOf(&unicode.Sinhala).Elem(),
		"Sk":                                 reflect.ValueOf(&unicode.Sk).Elem(),
		"Sm":                                 reflect.ValueOf(&unicode.Sm).Elem(),
		"So":                                 reflect.ValueOf(&unicode.So).Elem(),
		"Soft_Dotted":                        reflect.ValueOf(&unicode.Soft_Dotted).Elem(),
		"Sora_Sompeng":                       reflect.ValueOf(&unicode.Sora_Sompeng).Elem(),
		"Soyombo":                            reflect.ValueOf(&unicode.Soyombo).Elem(),
		"Space":                              reflect.ValueOf(&unicode.Space).Elem(),
		"Sundanese":                          reflect.ValueOf(&unicode.Sundanese).Elem(),
		"Syloti_Nagri":                       reflect.ValueOf(&unicode.Syloti_Nagri).Elem(),
		"Symbol":                             reflect.ValueOf(&unicode.Symbol).Elem(),
		"Syriac":                             reflect.ValueOf(&unicode.Syriac).Elem(),
		"Tagalog":                            reflect.ValueOf(&unicode.Tagalog).Elem(),
		"Tagbanwa":                           reflect.ValueOf(&unicode.Tagbanwa).Elem(),
		"Tai_Le":                             reflect.ValueOf(&unicode.Tai_Le).Elem(),
		"Tai_Tham":                           reflect.ValueOf(&unicode.Tai_Tham).Elem(),
		"Tai_Viet":                           reflect.ValueOf(&unicode.Tai_Viet).Elem(),
		"Takri":                              reflect.ValueOf(&unicode.Takri).Elem(),
		"Tamil":                              reflect.ValueOf(&unicode.Tamil).Elem(),
		"Tangut":                             reflect.ValueOf(&unicode.Tangut).Elem(),
		"Telugu":                             reflect.ValueOf(&unicode.Telugu).Elem(),
		"Terminal_Punctuation":               reflect.ValueOf(&unicode.Terminal_Punctuation).Elem(),
		"Thaana":                             reflect.ValueOf(&unicode.Thaana).Elem(),
		"Thai":                               reflect.ValueOf(&unicode.Thai).Elem(),
		"Tibetan":                            reflect.ValueOf(&unicode.Tibetan).Elem(),
		"Tifinagh":                           reflect.ValueOf(&unicode.Tifinagh).Elem(),
		"Tirhuta":                            reflect.ValueOf(&unicode.Tirhuta).Elem(),
		"Title":                              reflect.ValueOf(&unicode.Title).Elem(),
		"TitleCase":                          reflect.ValueOf(unicode.TitleCase),
		"To":                                 reflect.ValueOf(unicode.To),
		"ToLower":                            reflect.ValueOf(unicode.ToLower),
		"ToTitle":                            reflect.ValueOf(unicode.ToTitle),
		"ToUpper":                            reflect.ValueOf(unicode.ToUpper),
		"TurkishCase":                        reflect.ValueOf(&unicode.TurkishCase).Elem(),
		"Ugaritic":                           reflect.ValueOf(&unicode.Ugaritic).Elem(),
		"Unified_Ideograph":                  reflect.ValueOf(&unicode.Unified_Ideograph).Elem(),
		"Upper":                              reflect.ValueOf(&unicode.Upper).Elem(),
		"UpperCase":                          reflect.ValueOf(unicode.UpperCase),
		"UpperLower":                         reflect.ValueOf(unicode.UpperLower),
		"Vai":                                reflect.ValueOf(&unicode.Vai).Elem(),
		"Variation_Selector":                 reflect.ValueOf(&unicode.Variation_Selector).Elem(),
		"Version":                            reflect.ValueOf(unicode.Version),
		"Warang_Citi":                        reflect.ValueOf(&unicode.Warang_Citi).Elem(),
		"White_Space":                        reflect.ValueOf(&unicode.White_Space).Elem(),
		"Yi":                                 reflect.ValueOf(&unicode.Yi).Elem(),
		"Z":                                  reflect.ValueOf(&unicode.Z).Elem(),
		"Zanabazar_Square":                   reflect.ValueOf(&unicode.Zanabazar_Square).Elem(),
		"Zl":                                 reflect.ValueOf(&unicode.Zl).Elem(),
		"Zp":                                 reflect.ValueOf(&unicode.Zp).Elem(),
		"Zs":                                 reflect.ValueOf(&unicode.Zs).Elem(),

		// type definitions
		"CaseRange":   reflect.ValueOf((*unicode.CaseRange)(nil)),
		"Range16":     reflect.ValueOf((*unicode.Range16)(nil)),
		"Range32":     reflect.ValueOf((*unicode.Range32)(nil)),
		"RangeTable":  reflect.ValueOf((*unicode.RangeTable)(nil)),
		"SpecialCase": reflect.ValueOf((*unicode.SpecialCase)(nil)),

		// interface wrapper definitions

	}
}
