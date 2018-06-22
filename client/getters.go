package client

func (r *TypeRedis) GET() ([]string, error) {

	//TODO get data from Redis (api call to server)
	//redis.GET(currDateStr)
	dataFromRedis := []string{"111", "222", "333"}

	return dataFromRedis, nil
}
