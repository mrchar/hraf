package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mrchar/hraf"
)

func main() {
	server := hraf.NewServer()
	server.Register("todo", &Controller{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Run("localhost:8080", true)
		if err != nil {
			log.Panic(err)
		}
		os.Exit(0)
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGKILL, syscall.SIGTERM)
	<-exit

	err := server.Close()
	if err != nil {
		log.Panic(err)
	}
	wg.Wait()
}

// TODO 描述一个计划项目
type TODO struct {
	Title    string
	Message  string
	Deadline *time.Time
}

// Controller 用于接收http请求、处理并返回结果
type Controller struct {
	hraf.Basic
	TODOs []TODO
}

// GetAllTODO 获取所有的TODO
func (c *Controller) GetAllTODO(ctx hraf.Context) {
	err := ctx.Respond(map[string]interface{}{"data": c.TODOs})
	if err != nil {
		log.Panic(err)
	}
}

// GetOne 获取某一个TODO
func (c *Controller) GetOne(ctx hraf.Context) {
	params := struct {
		Title string `hraf:"title,url"`
	}{}

	err := ctx.Params(&params)
	if err != nil {
		log.Println(err)
		ctx.Error(http.StatusBadRequest, fmt.Errorf("获取参数出错：%w", err))
		return
	}

	for i := range c.TODOs {
		if params.Title == c.TODOs[i].Title {
			err := ctx.Respond(c.TODOs[i])
			if err != nil {
				log.Panic(err)
			}
			return
		}
	}

	err = fmt.Errorf("没有title为%s的TODO", params.Title)
	ctx.Error(http.StatusNotFound, err)
}

// AppendOne 添加一个TODO
func (c *Controller) AppendOne(ctx hraf.Context) {
	params := TODO{}
	err := ctx.Params(&params)
	if err != nil {
		log.Println(err)
		ctx.Error(http.StatusBadRequest, fmt.Errorf("获取参数出错：%w", err))
		return
	}

	for i := range c.TODOs {
		if params.Title == c.TODOs[i].Title {
			err := errors.New("There is a TODO with the same name")
			log.Println(err)
			ctx.Error(http.StatusBadRequest, err)
			return
		}
	}

	c.TODOs = append(c.TODOs, params)
	err = ctx.Respond(map[string]string{"message": "ok"})
	if err != nil {
		log.Panic(err)
	}
}

// UpdateOne 修改一个TODO
func (c *Controller) UpdateOne(ctx hraf.Context) {
	params := TODO{}
	err := ctx.Params(&params)
	if err != nil {
		log.Println(err)
		ctx.Error(http.StatusBadRequest, fmt.Errorf("获取参数出错：%w", err))
		return
	}

	for i := range c.TODOs {
		if params.Title == c.TODOs[i].Title {
			c.TODOs[i] = params
			err = ctx.Respond(map[string]string{"message": "ok"})
			if err != nil {
				log.Panic(err)
			}
			return
		}
	}

	err = fmt.Errorf("没有title为%s的TODO", params.Title)
	ctx.Error(http.StatusNotFound, err)
}

// DeleteOne 删除某一个TODO
func (c *Controller) DeleteOne(ctx hraf.Context) {
	params := struct {
		Title string `hraf:"title,url"`
	}{}

	err := ctx.Params(&params)
	if err != nil {
		log.Println(err)
		ctx.Error(http.StatusBadRequest, fmt.Errorf("获取参数出错：%w", err))
		return
	}

	for i := range c.TODOs {
		if params.Title == c.TODOs[i].Title {
			c.TODOs = append(c.TODOs[:i], c.TODOs[i+1:]...)
			err := ctx.Respond(map[string]string{"message": "ok"})
			if err != nil {
				log.Panic(err)
			}
			return
		}
	}

	err = fmt.Errorf("没有title为%s的TODO", params.Title)
	ctx.Error(http.StatusNotFound, err)
}
