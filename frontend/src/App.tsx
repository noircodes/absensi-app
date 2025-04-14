import { Outlet } from 'react-router'
import './App.css'

function App() {
  console.log("App component rendered")

  return (
    <>
      <div>
        <h1>Welcome to the Teams App</h1>
      <Outlet />
    </div>
    </>
  )
}

export default App