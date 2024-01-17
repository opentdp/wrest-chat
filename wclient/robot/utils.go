package robot

func sliceContains(slice []string, val string) bool {

	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false

}

func sliceRemove(slice []string, val string) []string {

	for i, v := range slice {
		if v == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice

}
