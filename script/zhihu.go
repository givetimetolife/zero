package script

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const ZHIHU_LINK = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50"

type (
	zhihu struct{}

	// zhihuData struct {
	// 	Data []struct {
	// 		Title       string `json:"target.title,omitempty"`
	// 		ExcerptArea string `json:"target.excerpt,omitempty"`
	// 		Link        string `json:"target.url"`
	// 		Metadata    string `json:"metadata,omitempty"`
	// 	} `json:"data"`
	// }

	zhihuData struct {
		Data []struct {
			Type      string `json:"type"`
			StyleType string `json:"style_type"`
			ID        string `json:"id"`
			CardID    string `json:"card_id"`
			Target    struct {
				ID            int    `json:"id"`
				Title         string `json:"title"`
				URL           string `json:"url"`
				Type          string `json:"type"`
				Created       int    `json:"created"`
				AnswerCount   int    `json:"answer_count"`
				FollowerCount int    `json:"follower_count"`
				Author        struct {
					Type      string `json:"type"`
					UserType  string `json:"user_type"`
					ID        string `json:"id"`
					URLToken  string `json:"url_token"`
					URL       string `json:"url"`
					Name      string `json:"name"`
					Headline  string `json:"headline"`
					AvatarURL string `json:"avatar_url"`
				} `json:"author"`
				BoundTopicIds []int  `json:"bound_topic_ids"`
				CommentCount  int    `json:"comment_count"`
				IsFollowing   bool   `json:"is_following"`
				Excerpt       string `json:"excerpt"`
			} `json:"target"`
			AttachedInfo string `json:"attached_info"`
			DetailText   string `json:"detail_text"`
			Trend        int    `json:"trend"`
			Debut        bool   `json:"debut"`
			Children     []struct {
				Type      string `json:"type"`
				Thumbnail string `json:"thumbnail"`
			} `json:"children"`
		} `json:"data"`
		Paging struct {
			IsEnd    bool   `json:"is_end"`
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"paging"`
		FreshText string `json:"fresh_text"`
	}
)

func (z zhihu) Data(ctx context.Context) ([]*ZeroData, error) {

	responseBody := func() (string, error) {
		res, err := http.Get(ZHIHU_LINK)
		if err != nil {
			return "", err
		}

		defer res.Body.Close()
		dataB, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}

		newDataB, err := strconv.Unquote(strings.Replace(strconv.Quote(string(dataB)), `\\u`, `\u`, -1))
		if err != nil {
			return "", err
		}
		return newDataB, nil
	}

	var (
		zhihu *zhihuData
		body  string
		err   error
	)

	if body, err = responseBody(); err != nil {
		return nil, err
	}
	fmt.Println(body)

	zhihu = &zhihuData{}
	if err = json.Unmarshal([]byte(body), zhihu); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", zhihu)
	return zhihu.convertToZd(), nil
}
