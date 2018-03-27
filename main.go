package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Web App api
	router.GET("/api/config", getConfig)
	router.POST("/api/signup", saveSignup)
	router.GET("/api/message", sendMessage)
	router.GET("/api/waitlist", addToWaitList)

	// Admin api
	router.POST("/admin/api/signup", getSignups)
	router.PUT("/admin/api/signup", updateSignup)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

func writeSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK,
		gin.H{"status": "success", "data": data})
}

func writeFail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest,
		gin.H{"status": "fail", "data": data})
}

func writeError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError,
		gin.H{"status": "error", "message": err.Error()})
}

func getConfig(c *gin.Context) {
	config, err := readConfig("webapp.yaml")
	if err == nil {
		writeSuccess(c, config)
	} else {
		writeError(c, err)
	}
}

func saveSignup(c *gin.Context) {
	var sf SignUpForm
	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := c.ShouldBindJSON(&sf); err != nil {
		log.Println(bodyBytes)
		writeFail(c, errors.Wrap(err, "request json missing fields or invalid").Error())
		return
	}

	if err := sf.save(); err == nil {
		writeSuccess(c, sf)
	} else {
		writeError(c, err)
	}
}

func sendMessage(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		writeFail(c, errors.Wrap(err, "request json missing fields or invalid").Error())
		return
	}

	if err := msg.send(); err != nil {
		writeError(c, err)
		return
	}

	writeSuccess(c, nil)
}

func addToWaitList(c *gin.Context) {
	var wle WaitListEntry
	if err := c.ShouldBindJSON(&wle); err != nil {
		writeFail(c, errors.Wrap(err, "request json missing fields or invalid").Error())
		return
	}

	if err := wle.save(); err != nil {
		writeError(c, err)
		return
	}

	writeSuccess(c, nil)
}

func getSignups(c *gin.Context) {
	sfs, err := getAllSignups()
	if err != nil {
		writeError(c, err)
		return
	}

	writeSuccess(c, sfs)
}

func updateSignup(c *gin.Context) {
	var sf SignUpForm
	if err := c.ShouldBindJSON(&sf); err != nil {
		writeFail(c, errors.Wrap(err, "request json missing fields or invalid").Error())
		return
	}

	if err := sf.update(); err != nil {
		writeError(c, err)
		return
	}

	writeSuccess(c, nil)
}
