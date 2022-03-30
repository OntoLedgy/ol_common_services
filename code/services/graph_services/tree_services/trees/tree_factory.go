package trees

func Create(
	rootNode any) *Trees {

	tree :=
		new(
			Trees)

	tree.
		RootNodeValue =
		rootNode

	return tree
}
