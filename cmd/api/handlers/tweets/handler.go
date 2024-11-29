package tweets

import tweets "uala/pkg/tweets/ports"

type Handler struct {
	TweetsService tweets.TweetsService
}
