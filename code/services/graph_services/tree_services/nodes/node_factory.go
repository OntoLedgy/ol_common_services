package nodes

func Create_node(
	string_to_add string) *TreeNodes {

	node :=
		new(
			TreeNodes)

	node.
		Value =
		string_to_add

	return node
}
