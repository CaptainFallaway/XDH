import { useState } from 'preact/hooks'
import { }

export function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <button className="btn" onClick={() => setCount(count + 1)}>{count}</button> 
    </>
  )
}
