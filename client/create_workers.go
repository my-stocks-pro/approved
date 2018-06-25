package client

func (a *ApprovedType) BaseWorker() {
	var newIDs []string

	for resp := range a.BaseChan {
		for _, image := range resp.Data {
			if Contains(a.Redis.ListIDS, image.Media_id) == false {
				newIDs = append(newIDs, image.Media_id)
			}
		}
		a.Redis.ListIDS = append(a.Redis.ListIDS, newIDs...)

		//data, errMarshal := json.Marshal(a.Redis.ListIDS)
		//if errMarshal != nil {
		//	panic(errMarshal)
		//}
		//
		//_, err := redis.POST(data)
		//if err != nil {
		//	fmt.Println(err)
		//}
	}
}

func (a *ApprovedType) workerSlack() {
	//for resp := range a.ChanSlack {
	//	d, e := json.Marshal(resp)
	//	if e != nil {
	//		panic(e)
	//	}
	//	slack.Post(d)
	//}
}
