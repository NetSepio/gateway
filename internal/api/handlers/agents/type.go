package agents

type AgentUpdateDTO struct {
	Name         *string `json:"name"`
	Clients      *string `json:"clients"`
	AvatarImg    *string `json:"avatar_img"`
	CoverImg     *string `json:"cover_img"`
	VoiceModel   *string `json:"voice_model"`
	Organization *string `json:"organization"`
}