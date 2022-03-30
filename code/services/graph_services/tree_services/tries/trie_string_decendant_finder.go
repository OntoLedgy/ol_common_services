package tries

import "sort"

func FindAllDescendants(
	parentString string,
	stringListToProcess []string) []string {

	var descendants []string

	parentStringLength := len(parentString)

	for _, stringInList := range stringListToProcess {

		if len(stringInList) > parentStringLength {

			if stringInList[:parentStringLength] == parentString &&
				len(stringInList) > parentStringLength {

				descendants = append(descendants, stringInList)

			}
		}

	}

	return descendants
}

func SortStringList(
	stringListToSort []string) []string {

	sort.Strings(stringListToSort)

	return stringListToSort
}
