package api

type clientConfig struct{
	serverUrl string
	apiKey string
}

func(cfg *clientConfig)NewClient()(*Client, error){
	c :=&Client{
		
	}
}
