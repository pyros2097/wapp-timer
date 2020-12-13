package main

import (
	"math"
	"strconv"
	"time"

	. "github.com/pyros2097/wapp"
)

func getTimeRemaining(secs int) (int, int, int) {
	hours := math.Floor(float64(secs / (60 * 60)))
	minutes := math.Floor(float64(secs % (60 * 60) / 60))
	seconds := math.Ceil(float64(secs % (60 * 60) % 60))
	return int(hours), int(minutes), int(seconds)
}

func pad(i int) string {
	s := strconv.Itoa(i)
	if len(s) == 1 {
		return "0" + s
	}
	return s
}

func Modal(c *RenderContext, close func()) UI {
	return Col(Css("z-40 overflow-auto left-0 top-0 bottom-0 right-0 w-full h-full fixed"),
		Col(Css("z-50 relative p-3 mx-auto my-0 w-6/12"),
			Div(Css("bg-white rounded shadow-lg border flex flex-col overflow-hidden"),
				Div(Css("flex flex-row justify-start border-b"),
					Div(Css("flex-1 px-6 py-3 text-xl font-bold"),
						Text("Title of the modal"),
					),
					Div(Css("px-6 py-1 text-2xl font-bold cursor-pointer select-none"), OnClick(close),
						Text("x"),
					),
				),
				Div(Css("p-6 flex-grow"),
					Text("You are watching this text in tailwind css model with alpine JS."),
				),
				Div(Css("px-6 py-3 border-t"),
					Div(Css("flex justify-end"),
						Div(Css("bg-gray-700 text-gray-100 rounded px-4 py-2 mr-1"), OnClick(close),
							Text("Cancel"),
						),
						Div(Css("bg-blue-600 text-gray-200 rounded px-4 py-2"),
							Text("Save"),
						),
					),
				),
			),
		),
	)
}

func Index(c *RenderContext) UI {
	endTime, setEndTime := c.UseState(10)
	hours, minutes, seconds := getTimeRemaining(endTime().(int))
	running, setRunning := c.UseState(false)
	showModal, setShowModal := c.UseState(false)
	toggleTimer := func() {
		if !running().(bool) {
			setRunning(true)
			go func() {
				for running().(bool) {
					nextTick := endTime().(int) - 1
					setEndTime(nextTick)
					if nextTick == 0 {
						setRunning(false)
					}
					time.Sleep(time.Second * 1)
				}
			}()
		} else {
			setRunning(false)
		}
	}
	open := func() {
		setShowModal(true)
	}
	close := func() {
		setShowModal(false)
	}
	return Col(Css("bg-black w-full h-full"),
		Row(Css("flex-1 text-10xl text-white leading"),
			Text(pad(hours)+":"+pad(minutes)+":"+pad(seconds)),
		),
		Row(
			Div(Css("bg-gray-300 text-gray-900 rounded hover:bg-gray-200 px-4 py-2 focus:outline-none cursor-pointer select-none"), OnClick(toggleTimer),
				IfElse(!running().(bool), Text("Start"), Text("Stop")),
			),
			Div(Css("bg-gray-300 text-gray-900 rounded hover:bg-gray-200 px-4 py-2 focus:outline-none cursor-pointer select-none ml-10"), OnClick(open),
				Text("Set Timer"),
			),
		),
		If(showModal().(bool), Modal(c, close)),
	)
}

func main() {
	Route("/", Index)
	Run()
}
