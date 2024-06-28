package service

import (
	"easy-nav/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndexList(c *gin.Context) {

	list, err := models.GetIndexList()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": list,
		"msg":    "success",
	})
	return
}

func AddUrl(c *gin.Context) {
	in := new(models.IndexReply)
	err := c.ShouldBind(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
		return
	} else {

		index := &models.IndexReply{
			Name: in.Name,
			Url:  in.Url,
		}
		cErr := models.CreateIndex(index)
		if cErr != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  cErr.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": index,
			})
			return
		}
	}
}

func DelIndexById(c *gin.Context) {
	id := c.Param("id")

	err := models.DelIndexById(&id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
	return
}

func UpdateIndex(c *gin.Context) {
	in := new(models.IndexReply)
	err := c.ShouldBind(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
		return
	} else {
		cErr := models.UpdateIndex(in)
		if cErr != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  cErr.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": in,
			})
			return
		}
	}

}
