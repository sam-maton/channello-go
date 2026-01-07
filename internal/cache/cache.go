package cache

import (
	"encoding/json"
	"os"
	"time"
)

type Video struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	ID   struct {
		Kind    string `json:"kind"`
		VideoID string `json:"videoId"`
	} `json:"id"`
	Snippet struct {
		PublishedAt time.Time `json:"publishedAt"`
		ChannelID   string    `json:"channelId"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Thumbnails  struct {
			Default struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"default"`
			Medium struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"high"`
		} `json:"thumbnails"`
		ChannelTitle         string    `json:"channelTitle"`
		LiveBroadcastContent string    `json:"liveBroadcastContent"`
		PublishTime          time.Time `json:"publishTime"`
	} `json:"snippet"`
}

type ChannelCache struct {
	Date int64   `json:"date"`
	Data []Video `json:"data"`
}

type ChannelFeed struct {
	ChannelID    string
	ChannelTitle string
	CachedAt     time.Time
	Videos       []Video
}

type Cache map[string]ChannelCache

const CacheExpiryDuration = 12 * time.Hour

func LoadCache(filepath string) (Cache, error) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var cache Cache

	err = json.Unmarshal(data, &cache)

	if err != nil {
		return nil, err
	}

	return cache, nil
}

func (c Cache) GetChannelFeeds(channelIDs []string) ([]ChannelFeed, []string) {
	feeds := make([]ChannelFeed, 0, len(channelIDs))
	nonCached := []string{}
	for _, id := range channelIDs {
		if channelCache, exists := c[id]; exists {
			now := time.Now().UnixMilli()
			if now-channelCache.Date > CacheExpiryDuration.Milliseconds() {
				nonCached = append(nonCached, id)
				continue
			}
			feed := ChannelFeed{
				ChannelID: id,
				CachedAt:  time.UnixMilli(channelCache.Date),
				Videos:    channelCache.Data,
			}

			if len(channelCache.Data) > 0 {
				feed.ChannelTitle = channelCache.Data[0].Snippet.ChannelTitle
			}

			feeds = append(feeds, feed)
		} else {
			nonCached = append(nonCached, id)
		}
	}

	return feeds, nonCached
}
