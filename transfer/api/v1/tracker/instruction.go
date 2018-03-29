package tracker

// RedirectRequest represents `GET /api/v1/tracker/instruction` request.
type InstructionRequest struct {
	EncryptedMarker string
}

// RedirectRequest represents `GET /api/v1/tracker/instruction` response.
type InstructionResponse struct {
	EncryptedMarker string
	Error           error
}
