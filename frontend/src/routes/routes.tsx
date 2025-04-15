import { createBrowserRouter } from "react-router"
import UserDetail from "../components/UserDetail"
import { fetchUserById, fetchUsers } from "../api"
import App from "../App"
import Dashboard from "@/app/dashboard/Dashboard"
import UserPaging from "@/components/UserPaging"

const router = createBrowserRouter([
    {
        path: '/',
        Component: App,
        children: [
            {
                path: 'user',
                children: [
                    {
                        index: true,
                        Component: UserPaging,
                        loader: async () => {
                            return fetchUsers()
                        },
                    },
                    {
                        path: ':id',
                        loader: async ({ params }) => {
                            if (!params.id) {
                            throw new Error('User ID is required')
                            }
                            return fetchUserById(params.id)
                        },
                        Component: UserDetail,
                    },
                ]
            },
            {
                path: 'dashboard',
                Component: Dashboard
            }
        ]
    }      
])

export default router