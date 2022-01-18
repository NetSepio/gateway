package delegateartifactcreation

type DelegateArtifactCreationRequest struct {
	CreatorAddress string `json:"creatorAddress"`
	MetaDataHash   string `json:"metaDataHash"`
}
