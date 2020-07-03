package config

import eureka "github.com/xuanbo/eureka-client"

func ConnectToEureka() *eureka.Client {
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://eurekaUser:eurekaPass@localhost:3007/eureka/",
		App:                   "go-example",
		Port:                  8080,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})

	return client
}
