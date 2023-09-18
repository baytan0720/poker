package container

type ContainerSlice []*Container

func (s ContainerSlice) Len() int {
	return len(s)
}

func (s ContainerSlice) Less(i, j int) bool {
	return s[i].CreatedAt.After(s[j].CreatedAt)
}

func (s ContainerSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
