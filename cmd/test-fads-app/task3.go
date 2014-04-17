package main

import (
	"reflect"

	"github.com/go-hep/fads"
	"github.com/go-hep/fwk"
)

type task3 struct {
	fwk.TaskBase

	parts string
}

func (tsk *task3) Configure(ctx fwk.Context) fwk.Error {
	var err fwk.Error
	msg := ctx.Msg()
	msg.Infof("configure...\n")

	tsk.parts = "/fads/StableParticles"
	err = tsk.DeclProp("Output", &tsk.parts)
	if err != nil {
		return err
	}

	err = tsk.DeclOutPort(tsk.parts, reflect.TypeOf([]fads.Candidate{}))
	if err != nil {
		return err
	}

	msg.Infof("configure... [done]\n")
	return err
}

func (tsk *task3) StartTask(ctx fwk.Context) fwk.Error {
	msg := ctx.Msg()
	msg.Infof("start...\n")
	return nil
}

func (tsk *task3) StopTask(ctx fwk.Context) fwk.Error {
	msg := ctx.Msg()
	msg.Infof("stop...\n")
	return nil
}

func (tsk *task3) Process(ctx fwk.Context) fwk.Error {
	msg := ctx.Msg()
	msg.Infof("proc...\n")
	store := ctx.Store()

	parts := make([]fads.Candidate, 0)
	err := store.Put(tsk.parts, parts)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	fwk.Register(reflect.TypeOf(task3{}))
}
