package main

import (
	"regexp"
	"strings"
)

type collaboratorFragment struct {
	Name  string
	Email string
}

type taskBodyScraper struct {
	rotatesPhraseMatcher *regexp.Regexp
	nameWithEmailMatcher *regexp.Regexp
}

func NewScraper() taskBodyScraper {
	rotatesPhraseMatcher := regexp.MustCompile(`(?i)rotates +among +(.+)$`)
	nameWithEmailMatcher := regexp.MustCompile(`^(?i)(.+\S)\s+\((\S+@\S+)\)$`)
	return taskBodyScraper{
		rotatesPhraseMatcher: rotatesPhraseMatcher,
		nameWithEmailMatcher: nameWithEmailMatcher,
	}
}

func (scraper *taskBodyScraper) ScrapeTaskBody(body string) []collaboratorFragment {
	var collaborators []collaboratorFragment
	groups := scraper.rotatesPhraseMatcher.FindStringSubmatch(body)
	if groups == nil {
		return nil
	}
	nameList := strings.Split(groups[1], ",")
	for _, name := range nameList {
		name = strings.TrimSpace(name)
		var collaborator collaboratorFragment
		groups := scraper.nameWithEmailMatcher.FindStringSubmatch(name)
		if groups != nil {
			collaborator.Name = groups[1]
			collaborator.Email = groups[2]
		} else {
			if strings.ContainsRune(name, '@') {
				collaborator.Email = name
			} else {
				collaborator.Name = name
			}
		}
		if collaborator != (collaboratorFragment{}) {
			collaborators = append(collaborators, collaborator)
		}
	}

	return collaborators
}
