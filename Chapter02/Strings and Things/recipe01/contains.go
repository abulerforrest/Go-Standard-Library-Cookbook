package main

import (
	"fmt"
	"strings"
)

func convertTagsToWP(tagsToBeConverted []string, slug string, s []string) []string {
	addLast := ""
	if s[1] != s[len(s)-1] {
		addLast = "-" + s[len(s)-1]
	}
	tagsToBeConverted = append(tagsToBeConverted, slug+"-"+s[1]+addLast)
	return tagsToBeConverted
}

/* having a little fun with mfn tags :)
-- Converting from API to 'WP taxonomy style' tags.
*/
func main() {
	mfnTagsArr := []string{
		":regulatory",
		":regulatory:mar",
		":regulatory:vpml",
		":regulatory:lhfi",
		":regulatory:listing",
		":correction",
		"sub:report",
		"sub:report:annual",
		"sub:report:interim",
		"sub:report:interim:q1",
		"sub:report:interim:q2",
		"sub:report:interim:q3",
		"sub:report:interim:q4",
		"sub:ca",
		"sub:ca:shares",
		"sub:ca:shares:issuance",
		"sub:ca:shares:repurchase",
		"sub:ca:shares:rights",
		"sub:ca:ma",
		"sub:ca:prospectus",
		"sub:ca:other",
		"sub:ci",
		"sub:ci:gm",
		"sub:ci:gm:notice",
		"sub:ci:gm:info",
		"sub:ci:insider",
		"sub:ci:shareholder-announcement",
		"sub:ci:calendar",
		"sub:ci:presentation",
		"sub:ci:nomination",
		"sub:ci:earnings",
		"sub:ci:sales",
		"sub:ci:sales:order",
		"sub:ci:staff",
		"sub:ci:other",
	}

	baseSlug := "mfn"
	var baseTags []string
	var reportTags []string
	var corporateActionTags []string
	var corporateInfoTags []string

	baseNeedle := ":"
	reportNeedle := "sub:report"
	corporateActionNeedle := "sub:ca"
	corporateInfoNeedle := "sub:ci"

	for _, tagString := range mfnTagsArr {
		isBase := strings.HasPrefix(tagString, baseNeedle)
		isReport := strings.Contains(tagString, reportNeedle)
		isCorporateAction := strings.Contains(tagString, corporateActionNeedle)
		isCorporateInfo := strings.Contains(tagString, corporateInfoNeedle)

		// split by baseNeedle
		s := strings.Split(tagString, baseNeedle)

		if isBase {
			baseTags = convertTagsToWP(baseTags, baseSlug, s)
			fmt.Printf("The tag \"%s\" starts with \"%s\": %t \n", tagString, baseNeedle, isBase)
		}

		if isReport {
			reportTags = convertTagsToWP(reportTags, baseSlug, s)
			fmt.Printf("The tag \"%s\" contains \"%s\": %t \n", tagString, reportNeedle, isReport)
		}

		if isCorporateAction {
			corporateActionTags = convertTagsToWP(corporateActionTags, baseSlug, s)
			fmt.Printf("The tag \"%s\" contains \"%s\": %t \n", tagString, corporateActionNeedle, isCorporateAction)
		}

		if isCorporateInfo {
			corporateInfoTags = convertTagsToWP(corporateInfoTags, baseSlug, s)
			fmt.Printf("The tag \"%s\" contains \"%s\": %t \n", tagString, corporateInfoNeedle, isCorporateInfo)
		}
	}

	fmt.Println("// BASE TAGS")
	for i := 0; i < len(baseTags); i++ {
		fmt.Printf("%v (is a base tag) \n", baseTags[i])
	}
	fmt.Println("// REPORT TAGS")
	for i := 0; i < len(reportTags); i++ {
		fmt.Printf("%v (is a report tag) \n", reportTags[i])
	}
	fmt.Println("// CORPORATE ACTION TAGS")
	for i := 0; i < len(corporateActionTags); i++ {
		fmt.Printf("%v (is a corporate action tag) \n", corporateActionTags[i])
	}
	fmt.Println("// CORPORATE INFORMATION TAGS")
	for i := 0; i < len(corporateInfoTags); i++ {
		fmt.Printf("%v (is a corporate information tag) \n", corporateInfoTags[i])
	}
}
