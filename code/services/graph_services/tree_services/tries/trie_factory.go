package tries

func Create(listOfStrings []string) *Tries {

	sortedStringList :=
		SortStringList(
			listOfStrings)

	trie :=
		new(
			Tries)

	trie.
		PopulateList(
			sortedStringList)

	trie.
		ProcessStringList(
			sortedStringList)

	return trie

}
