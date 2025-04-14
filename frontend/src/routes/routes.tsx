import { createBrowserRouter } from "react-router"
import UserDetail from "../component/UserDetail"
import { fetchUserById } from "../api"
import App from "../App"

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
            }
        ]
    }      
])

export default router