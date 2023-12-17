package aiprovider

import (

	// "pingyingqi/service/CodeAi/provider"
	_ "pingyingqi/utils/logging"
)

var AiHelper aiHelper

type aiHelper struct {
	Provider []AiProvider
}

func init() {
	AiHelper = aiHelper{}
}
