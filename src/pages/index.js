import React, { useRef, useState } from "react"
import './index.css';

const secondsToTime = secs => {
  let hours = Math.floor(secs / (60 * 60));
  let divisor_for_minutes = secs % (60 * 60);
  let minutes = Math.floor(divisor_for_minutes / 60);
  let divisor_for_seconds = divisor_for_minutes % 60;
  let seconds = Math.ceil(divisor_for_seconds);
  hours += "";
  minutes += "";
  seconds += "";
  hours = hours.padStart(2, "0");
  minutes = minutes.padStart(2, "0");
  seconds = seconds.padStart(2, "0");
  return {
    h: hours,
    m: minutes,
    s: seconds
  };
};

const IndexPage = () => {
  const timerRef = useRef(null)
  const [seconds, setSeconds] = useState(5400);
  const { h, m, s } = secondsToTime(seconds);
  const toggleTimer = () => {
    if (timerRef.current) {
      clearInterval(timerRef.current)
      timerRef.current = null
      setSeconds(5400)
    } else {
      timerRef.current = setInterval(() => {
        setSeconds((seconds) => {
          return seconds > 0 ? seconds - 1 : 0
        })
      }, 1000)
    }
  }
  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ flex: 1, fontSize: "325px", fontStyle: "sans-serif", color: 'white', lineHeight: 2 }}>{`${h}:${m}:${s} `}</div>
      <button onClick={toggleTimer}>{timerRef.current ? 'Stop' : 'Start'}</button>
    </div>
  )
}

export default IndexPage

