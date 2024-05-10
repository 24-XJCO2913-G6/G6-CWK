package handlers

import (
	. "main/backend/models"
)

func SearchBlogs(text string) ([]Blog_display, error) {
	var blogs []Blog
	var blogDisplays []Blog_display

	// 查询匹配的记录
	err := Db.Where("title = ? OR content = ?", text, text).Find(&blogs)
	if err != nil {
		return nil, err
	}

	// 转换查询结果为 Blog_display 结构
	if len(blogs) != 0 {
		for _, blog := range blogs {
			blogDisplay := Blog_display{
				// 这里根据实际需求设置 Author、Photo 等字段的值
				Pub_time:   blog.Pub_time,
				Visibility: blog.Visibility,
				Content:    blog.Content,
				Picture:    blog.Picture,
				Title:      blog.Title,
			}
			blogDisplays = append(blogDisplays, blogDisplay)
		}
	}

	return blogDisplays, nil
}
