import { Outlet } from 'react-router'
import './App.css'
import { ThemeProvider } from './components/theme-provider'


export default function App() {
  console.log("App component rendered")

  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <Outlet />
    </ThemeProvider>
  )
}