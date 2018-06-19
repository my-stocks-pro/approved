package httputils

func FullRequest(approved TestType, dataFromRedis []string) {
	newIDs := []string{}

	for _, image := range approved.Data {
		if Contains(newIDs, image.Media_id) == false {
			newIDs = append(newIDs, image.Media_id)
		}
	}

	Req(newIDs)

}