package language

import (
	"testing"
)

var lang1 = Language{
	Language: "Fortran",
	Appeared: 1957,
	Functional: false,
	ObjectOriented: true,
	Created: []string{"John Backus"},
	Relation: &Relation{
		InfluencedBy: []string{"Speedcoding"},
		Influences: []string{"ALGOL 58", "BASIC C", "C", "Chapel", "CMS-2", "DOPE", "Fortress",
			"PL/I", "PACT I", "MUMPS", "IDL", "Ratfor"},
	},
}

func TestCreateLanguage(t *testing.T) {
	lang2 := Language{
		Language: "Pascal",
		Appeared: 1970,
		Functional: false,
		ObjectOriented: false,
		Created: []string{"Niklaus Wirth"},
		Relation: &Relation{
			InfluencedBy: []string{"ALGOL W", "Simula 67"},
			Influences: []string{"Ada", "C/AL", "Component Pascal", "Go", "Java"},
		},
	}
	res := CreateLanguage(lang2)
	if res.Language != "Pascal" {
		t.Errorf("Expected language name = %s, got: %s", "Pascal", res.Language)
	}
}

func TestGetLanguage(t *testing.T) {
	res := GetLanguage(0)
	if res.Language != "Pascal" {
		t.Errorf("Expected GetLanguages name = %s, got: %s", "Pascal", res.Language)
	}
}

func TestGetLanguages(t *testing.T) {
	CreateLanguage(lang1)
	allLang := GetLanguages()
	if len(allLang) != 2 {
		t.Errorf("Expected len(GetLanguages()) = %d, got: %d", 2, len(allLang))
	}
}

func TestUpdateLanguage(t *testing.T) {
	lang1Target := lang1
	lang1Target.Appeared = 1958

	UpdateLanguage(0, lang1Target)
	lang1Updated := GetLanguage(0)
	if lang1Updated.Appeared != 1958 {
		t.Errorf("Expected GetLanguages appeared = %d, got: %d", 1958, lang1Updated.Appeared)
	}

	allLang := GetLanguages()
	if len(allLang) != 2 {
		t.Errorf("Expected len(GetLanguages()) = %d, got: %d", 2, len(allLang))
	}
}

func TestDeleteLanguage(t *testing.T) {
	DeleteLanguage(0)
	allLang := GetLanguages()
	if len(allLang) != 1 {
		t.Errorf("Expected GetLanguages = %d, got: %d", 1, len(allLang))
	}
}
