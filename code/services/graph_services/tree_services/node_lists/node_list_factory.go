package node_lists

func Create(list_of_strings []string) *NodeLists {

	node_list := new(NodeLists)

	node_list.PopulateList(list_of_strings)

	return node_list

}
