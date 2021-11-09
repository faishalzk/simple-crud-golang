package language

var languages []Language

// GetLanguage get language by id
func GetLanguage(id int) Language {
	if id < len(languages) && id >= 0{
		return languages[id]
	}
	// If no match found, return 'empty' Language
	return Language{}
}

// GetLanguages get list of languages
func GetLanguages() []Language {
	return languages
}

// CreateLanguage creates language
func CreateLanguage(language Language) Language{
	languages = append(languages, language)
	return language
}

// UpdateLanguage updates language by id
func UpdateLanguage(id int, language Language) Language{
	languageTarget := languages[id]

	languageTarget.Language = language.Language
	languageTarget.Appeared = language.Appeared
	languageTarget.Functional = language.Functional
	languageTarget.ObjectOriented = language.ObjectOriented
	languageTarget.Created = language.Created
	languageTarget.Relation = language.Relation

	languages[id] = languageTarget

	return languages[id]
}

// DeleteLanguage deletes language by id
func DeleteLanguage(id int) []Language {
	languages = append(languages[:id], languages[id+1:]...)
	return languages
}
