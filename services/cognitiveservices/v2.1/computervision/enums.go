package computervision

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DescriptionExclude enumerates the values for description exclude.
type DescriptionExclude string

const (
	// Celebrities ...
	Celebrities DescriptionExclude = "Celebrities"
	// Landmarks ...
	Landmarks DescriptionExclude = "Landmarks"
)

// PossibleDescriptionExcludeValues returns an array of possible values for the DescriptionExclude const type.
func PossibleDescriptionExcludeValues() []DescriptionExclude {
	return []DescriptionExclude{Celebrities, Landmarks}
}

// Details enumerates the values for details.
type Details string

const (
	// DetailsCelebrities ...
	DetailsCelebrities Details = "Celebrities"
	// DetailsLandmarks ...
	DetailsLandmarks Details = "Landmarks"
)

// PossibleDetailsValues returns an array of possible values for the Details const type.
func PossibleDetailsValues() []Details {
	return []Details{DetailsCelebrities, DetailsLandmarks}
}

// Gender enumerates the values for gender.
type Gender string

const (
	// Female ...
	Female Gender = "Female"
	// Male ...
	Male Gender = "Male"
)

// PossibleGenderValues returns an array of possible values for the Gender const type.
func PossibleGenderValues() []Gender {
	return []Gender{Female, Male}
}

// OcrLanguages enumerates the values for ocr languages.
type OcrLanguages string

const (
	// Ar ...
	Ar OcrLanguages = "ar"
	// Cs ...
	Cs OcrLanguages = "cs"
	// Da ...
	Da OcrLanguages = "da"
	// De ...
	De OcrLanguages = "de"
	// El ...
	El OcrLanguages = "el"
	// En ...
	En OcrLanguages = "en"
	// Es ...
	Es OcrLanguages = "es"
	// Fi ...
	Fi OcrLanguages = "fi"
	// Fr ...
	Fr OcrLanguages = "fr"
	// Hu ...
	Hu OcrLanguages = "hu"
	// It ...
	It OcrLanguages = "it"
	// Ja ...
	Ja OcrLanguages = "ja"
	// Ko ...
	Ko OcrLanguages = "ko"
	// Nb ...
	Nb OcrLanguages = "nb"
	// Nl ...
	Nl OcrLanguages = "nl"
	// Pl ...
	Pl OcrLanguages = "pl"
	// Pt ...
	Pt OcrLanguages = "pt"
	// Ro ...
	Ro OcrLanguages = "ro"
	// Ru ...
	Ru OcrLanguages = "ru"
	// Sk ...
	Sk OcrLanguages = "sk"
	// SrCyrl ...
	SrCyrl OcrLanguages = "sr-Cyrl"
	// SrLatn ...
	SrLatn OcrLanguages = "sr-Latn"
	// Sv ...
	Sv OcrLanguages = "sv"
	// Tr ...
	Tr OcrLanguages = "tr"
	// Unk ...
	Unk OcrLanguages = "unk"
	// ZhHans ...
	ZhHans OcrLanguages = "zh-Hans"
	// ZhHant ...
	ZhHant OcrLanguages = "zh-Hant"
)

// PossibleOcrLanguagesValues returns an array of possible values for the OcrLanguages const type.
func PossibleOcrLanguagesValues() []OcrLanguages {
	return []OcrLanguages{Ar, Cs, Da, De, El, En, Es, Fi, Fr, Hu, It, Ja, Ko, Nb, Nl, Pl, Pt, Ro, Ru, Sk, SrCyrl, SrLatn, Sv, Tr, Unk, ZhHans, ZhHant}
}

// TextOperationStatusCodes enumerates the values for text operation status codes.
type TextOperationStatusCodes string

const (
	// Failed ...
	Failed TextOperationStatusCodes = "Failed"
	// NotStarted ...
	NotStarted TextOperationStatusCodes = "NotStarted"
	// Running ...
	Running TextOperationStatusCodes = "Running"
	// Succeeded ...
	Succeeded TextOperationStatusCodes = "Succeeded"
)

// PossibleTextOperationStatusCodesValues returns an array of possible values for the TextOperationStatusCodes const type.
func PossibleTextOperationStatusCodesValues() []TextOperationStatusCodes {
	return []TextOperationStatusCodes{Failed, NotStarted, Running, Succeeded}
}

// TextRecognitionMode enumerates the values for text recognition mode.
type TextRecognitionMode string

const (
	// Handwritten ...
	Handwritten TextRecognitionMode = "Handwritten"
	// Printed ...
	Printed TextRecognitionMode = "Printed"
)

// PossibleTextRecognitionModeValues returns an array of possible values for the TextRecognitionMode const type.
func PossibleTextRecognitionModeValues() []TextRecognitionMode {
	return []TextRecognitionMode{Handwritten, Printed}
}

// TextRecognitionResultConfidenceClass enumerates the values for text recognition result confidence class.
type TextRecognitionResultConfidenceClass string

const (
	// High ...
	High TextRecognitionResultConfidenceClass = "High"
	// Low ...
	Low TextRecognitionResultConfidenceClass = "Low"
)

// PossibleTextRecognitionResultConfidenceClassValues returns an array of possible values for the TextRecognitionResultConfidenceClass const type.
func PossibleTextRecognitionResultConfidenceClassValues() []TextRecognitionResultConfidenceClass {
	return []TextRecognitionResultConfidenceClass{High, Low}
}

// TextRecognitionResultDimensionUnit enumerates the values for text recognition result dimension unit.
type TextRecognitionResultDimensionUnit string

const (
	// Inch ...
	Inch TextRecognitionResultDimensionUnit = "inch"
	// Pixel ...
	Pixel TextRecognitionResultDimensionUnit = "pixel"
)

// PossibleTextRecognitionResultDimensionUnitValues returns an array of possible values for the TextRecognitionResultDimensionUnit const type.
func PossibleTextRecognitionResultDimensionUnitValues() []TextRecognitionResultDimensionUnit {
	return []TextRecognitionResultDimensionUnit{Inch, Pixel}
}

// VisualFeatureTypes enumerates the values for visual feature types.
type VisualFeatureTypes string

const (
	// VisualFeatureTypesAdult ...
	VisualFeatureTypesAdult VisualFeatureTypes = "Adult"
	// VisualFeatureTypesBrands ...
	VisualFeatureTypesBrands VisualFeatureTypes = "Brands"
	// VisualFeatureTypesCategories ...
	VisualFeatureTypesCategories VisualFeatureTypes = "Categories"
	// VisualFeatureTypesColor ...
	VisualFeatureTypesColor VisualFeatureTypes = "Color"
	// VisualFeatureTypesDescription ...
	VisualFeatureTypesDescription VisualFeatureTypes = "Description"
	// VisualFeatureTypesFaces ...
	VisualFeatureTypesFaces VisualFeatureTypes = "Faces"
	// VisualFeatureTypesImageType ...
	VisualFeatureTypesImageType VisualFeatureTypes = "ImageType"
	// VisualFeatureTypesObjects ...
	VisualFeatureTypesObjects VisualFeatureTypes = "Objects"
	// VisualFeatureTypesTags ...
	VisualFeatureTypesTags VisualFeatureTypes = "Tags"
)

// PossibleVisualFeatureTypesValues returns an array of possible values for the VisualFeatureTypes const type.
func PossibleVisualFeatureTypesValues() []VisualFeatureTypes {
	return []VisualFeatureTypes{VisualFeatureTypesAdult, VisualFeatureTypesBrands, VisualFeatureTypesCategories, VisualFeatureTypesColor, VisualFeatureTypesDescription, VisualFeatureTypesFaces, VisualFeatureTypesImageType, VisualFeatureTypesObjects, VisualFeatureTypesTags}
}
