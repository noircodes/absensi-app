import { createBrowserRouter } from "react-router"
import UserDetail from "../components/UserDetail"
import { fetchUserById } from "../api"
import App from "../App"
import Dashboard from "@/app/dashboard/Dashboard"

const router = createBrowserRouter([
    {
        path: '/',
        Component: App,
        children: [
            {
                path: '/user/:id',
                loader: async ({ params }) => {
                    if (!params.id) {
                    throw new Error('User ID is required')
                    }
                    return fetchUserById(params.id)
                },
                Component: UserDetail,
            },
            {
                path: '/dashboard',
                Component: Dashboard
            }
        ]
    }      
])

export default router