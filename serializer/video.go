/*
 * @Author: xuyong
 * @Date: 2024-08-13 20:15:35
 * @LastEditors: xuyong
 */
package serializer

import "singo/model"

type Video struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Info string `json:"info"`
	CreatedAt int64 `json:"created_at"`
}

func BindVideo (video model.Video) Video {
	return Video{
		ID: video.ID,
		Title: video.Title,
		Info: video.Info,
		CreatedAt: video.CreatedAt.Unix(),
	}
}

func BindVideoResponse(video model.Video) Response {
	return Response {
		Data: BindVideo(video),
	}
}