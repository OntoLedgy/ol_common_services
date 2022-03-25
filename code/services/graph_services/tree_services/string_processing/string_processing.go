package string_processing

import "sort"

func Find_all_descendants(
	parent_string string,
	string_list_to_process []string) []string {

	var descendants []string

	parent_string_length := len(parent_string)

	for _, string_in_list := range string_list_to_process {

		if len(string_in_list) > parent_string_length {

			if string_in_list[:parent_string_length] == parent_string &&
				len(string_in_list) > parent_string_length {

				descendants = append(descendants, string_in_list)

			}
		}

	}

	return descendants
}

func Sort_string_list(
	string_list_to_sort []string) []string {

	sort.Strings(string_list_to_sort)

	return string_list_to_sort
}
