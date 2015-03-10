package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"code.google.com/p/google-api-go-client/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

// YT Comments (v2) JSON struct
type YoutubeFeed struct {
	Version  string `json:"version"`
	Encoding string `json:"encoding"`
	Feed     struct {
		Xmlns           string `json:"xmlns"`
		XmlnsOpenSearch string `json:"xmlns$openSearch"`
		XmlnsYt         string `json:"xmlns$yt"`
		ID              struct {
			T string `json:"$t"`
		} `json:"id"`
		Updated struct {
			T string `json:"$t"`
		} `json:"updated"`
		Category []struct {
			Scheme string `json:"scheme"`
			Term   string `json:"term"`
		} `json:"category"`
		Logo struct {
			T string `json:"$t"`
		} `json:"logo"`
		Link []struct {
			Rel  string `json:"rel"`
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"link"`
		Author []struct {
			Name struct {
				T string `json:"$t"`
			} `json:"name"`
			Uri struct {
				T string `json:"$t"`
			} `json:"uri"`
		} `json:"author"`
		Generator struct {
			T       string `json:"$t"`
			Version string `json:"version"`
			Uri     string `json:"uri"`
		} `json:"generator"`
		OpenSearchTotalResults struct {
			T int `json:"$t"`
		} `json:"openSearchTotalResults"`
		OpenSearchStartIndex struct {
			T int `json:"$t"`
		} `json:"openSearch$startIndex"`
		OpenSearchItemsPerPage struct {
			T int `json:"$t"`
		} `json:"openSearch$itemsPerPage"`
		Entry []struct {
			ID struct {
				T string `json:"$t"`
			} `json:"id"`
			Published struct {
				T string `json:"$t"`
			} `json:"published"`
			Updated struct {
				T string `json:"$t"`
			} `json:"updated"`
			Category []struct {
				Scheme string `json:"scheme"`
				Term   string `json:"term"`
			} `json:"category"`
			Title struct {
				T    string `json:"$t"`
				Type string `json:"type"`
			} `json:"title"`
			Content struct {
				T    string `json:"$t"`
				Type string `json:"type"`
			} `json:"content"`
			Link []struct {
				Rel  string `json:"rel"`
				Type string `json:"type"`
				Href string `json:"href"`
			} `json:"link"`
			Author []struct {
				Name struct {
					T string `json:"$t"`
				} `json:"name"`
				Uri struct {
					T string `json:"$t"`
				} `json:"uri"`
			} `json:"author"`
			YtChannelID struct {
				T string `json:"$t"`
			} `json:"yt$channelId"`
			YtGooglePlusUserID struct {
				T string `json:"$t"`
			} `json:"yt$googlePlusUserId"`
			YtReplyCount struct {
				T int `json:"$t"`
			} `json:"yt$replyCount"`
			YtVideoid struct {
				T string `json:"$t"`
			} `json:"yt$videoid"`
		} `json:"entry"`
	} `json:"feed"`
}

// Distilled Comment Type
type Comment struct {
	ID         string
	Published  string
	Title      string
	Content    string
	AuthorName string
}

// Pull the comments for a given YouTube video
func GetCommentsV2(videoID string) []Comment {
	var comments = []Comment{}
	var feed YoutubeFeed

	url := "https://gdata.youtube.com/feeds/api/videos/" + videoID + "/comments?v=2&alt=json"

	for url != "" {
		data, hasErr := fetchJSON(url)

		if hasErr == false {
			json.Unmarshal(data, &feed)

			for _, entry := range feed.Feed.Entry {
				thisComment := Comment{
					ID:         entry.ID.T,
					Published:  entry.Published.T,
					Title:      entry.Title.T,
					Content:    entry.Content.T,
					AuthorName: entry.Author[0].Name.T,
				}

				comments = append(comments, thisComment)
			}

			url = ""
			for _, link := range feed.Feed.Link {
				if link.Rel == "next" {
					url = link.Href
				}
			}
		}
	}

	return comments
}

// A distilled record of video metadata
type VideoMetadata struct {
	Title         string
	VideoViews    uint64
	ChannelID     string
	ChannelTitle  string
	TotalComments uint64
	PublishedAt   string
}

// Returns a subset of video information from the YouTube API
func GetVideoInfo(videoID string) VideoMetadata {
	client := &http.Client{
		Transport: &transport.APIKey{Key: *YouTubeKey},
	}

	youtubeService, err := youtube.New(client)
	if err != nil {
		panic(err)
	}

	call := youtubeService.Videos.List("id,snippet,statistics").Id(videoID)
	resp, err := call.Do()
	if err != nil {
		panic(err)
	}

	video := resp.Items[0]

	return VideoMetadata{
		Title:         video.Snippet.Title,
		ChannelID:     video.Snippet.ChannelId,
		ChannelTitle:  video.Snippet.ChannelTitle,
		TotalComments: video.Statistics.CommentCount,
		PublishedAt:   video.Snippet.PublishedAt,
		VideoViews:    video.Statistics.ViewCount,
	}
}

func fetchJSON(url string) ([]byte, bool) {
	r, _ := http.Get(url)
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err == nil {
		return body, false
	}

	return []byte{}, true
}
