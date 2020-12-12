package main

import (
	"math"
	"strconv"
	"time"

	. "github.com/pyros2097/wapp"
)

// const SetTimerModal = ({ setSeconds, onRequestClose }) => {
// 	const [value, setValue] = useState(0)
// 	return (
// 	  <Modal
// 		styles={{
// 		  overlay: {
// 			position: 'fixed',
// 			top: 0,
// 			left: 0,
// 			right: 0,
// 			bottom: 0,
// 			zIndex: 99999999,
// 			overflow: 'hidden',
// 			perspective: 1300,
// 			backgroundColor: 'rgba(0, 0, 0, 0.3)'
// 		  },
// 		  content: {
// 			position: 'relative',
// 			margin: '100px',
// 			width: '400px',
// 			border: '1px solid rgba(0, 0, 0, .2)',
// 			background: '#fff',
// 			overflow: 'auto',
// 			borderRadius: '4px',
// 			outline: 'none',
// 			boxShadow: '0 5px 10px rgba(0, 0, 0, .3)',
// 		  }
// 		}}
// 		onRequestClose={onRequestClose}
// 		effect={Effect.Newspaper}>
// 		<div style={{ display: 'flex', flexDirection: 'column' }}>
// 		  <div style={{ display: 'flex', flexDirection: 'column', padding: '20px 0px 0px 20px' }}>
// 			<p style={{ fontFamily: "Helvetica", fontSize: "24px" }}>Please enter the timer duration in seconds?</p>
// 			<input style={{ width: "200px" }} type="number" onChange={(e) => setValue(parseInt(e.target.value))} />
// 		  </div>
// 		  <div style={{ display: 'flex', flexDirection: 'row', justifyContent: 'flex-end', marginTop: '100px' }}>
// 			<button style={{ width: '50%', height: '50px ' }} onClick={() => {
// 			  setSeconds(value);
// 			  ModalManager.close();
// 			}}>Save</button>
// 			<button style={{ width: '50%', height: '50px ' }} onClick={ModalManager.close}>Cancel</button>
// 		  </div>
// 		</div>
// 	  </Modal>
// 	)
//   }

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

func Index(c *RenderContext) UI {
	endTime, setEndTime := c.UseState(10)
	hours, minutes, seconds := getTimeRemaining(endTime().(int))
	running, setRunning := c.UseState(false)
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
	return Col(Css("bg-black w-full h-full"),
		Row(Css("flex-1 text-6xl text-white leading"),
			Text(pad(hours)+":"+pad(minutes)+":"+pad(seconds)),
		),
		Row(
			Div(Css("bg-gray-300 text-gray-900 rounded hover:bg-gray-200 px-4 py-2 focus:outline-none cursor-pointer select-none"), OnClick(toggleTimer),
				IfElse(!running().(bool), Text("Start"), Text("Stop")),
			),
			Div(Css("bg-gray-300 text-gray-900 rounded hover:bg-gray-200 px-4 py-2 focus:outline-none cursor-pointer select-none ml-10"),
				Text("Set Timer"),
			),
		),
	)
}
