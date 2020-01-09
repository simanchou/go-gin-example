package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/simanchou/go-gin-example/models"
	"github.com/simanchou/go-gin-example/pkg/e"
	"github.com/simanchou/go-gin-example/pkg/setting"
	"github.com/simanchou/go-gin-example/pkg/util"
	"github.com/unknwon/com"
)

// GetArticle get one article
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id must be greater than 0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticles get one more articles
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("state must be 0 or 1")
	}

	var tagID int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagID = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagID

		valid.Min(tagID, 1, "tag_id").Message("tag id must be greater than 0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.Key: %s, err.Message: %s\n", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle add article
func AddArticle(c *gin.Context) {
	tagID := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagID, 1, "tag_id").Message("tag id must be greater than 0")
	valid.Required(title, "title").Message("title is need")
	valid.Required(desc, "desc").Message("desc is need")
	valid.Required(content, "content").Message("content is need")
	valid.Required(createdBy, "created_by").Message("created_by is need")
	valid.Range(state, 0, 1, "state").Message("state must be 0 or 1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagByID(tagID) {
			data := make(map[string]interface{})
			data["tag_id"] = tagID
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.Key: %s, err.Message: %s\n", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// EditArticle modify article
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagID := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("state must be 0 or 1")
	}

	valid.Min(id, 1, "id").Message("id must be greater than 0")
	valid.MaxSize(title, 100, "title").Message("title can't greater than 100 characters")
	valid.MaxSize(desc, 255, "desc").Message("desc can't greater than 100 characters")
	valid.MaxSize(content, 65535, "content").Message("content can't greater than 65535 characters")
	valid.Required(modifiedBy, "modified_by").Message("modified_by is need")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("modified_by can't greater than 100 characters")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagID) {
				data := make(map[string]interface{})
				if tagID > 0 {
					data["tag_id"] = tagID
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy
				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.Key: %s, err.Message: %s\n", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteArticle delete article
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id must be greater than 0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.Key: %s, err.Message: %s\n", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
