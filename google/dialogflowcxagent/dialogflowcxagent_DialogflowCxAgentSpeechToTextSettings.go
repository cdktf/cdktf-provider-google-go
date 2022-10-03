package dialogflowcxagent


type DialogflowCxAgentSpeechToTextSettings struct {
	// Whether to use speech adaptation for speech recognition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_agent#enable_speech_adaptation DialogflowCxAgent#enable_speech_adaptation}
	EnableSpeechAdaptation interface{} `field:"optional" json:"enableSpeechAdaptation" yaml:"enableSpeechAdaptation"`
}

