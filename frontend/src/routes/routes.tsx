import { createBrowserRouter, RouterProvider } from "react-router"
import UserDetail from "../components/UserDetail"
import { fetchUserById, fetchUsers } from "../api"
import UserPaging from "../components/UserPaging"
import { HomePage } from "@/pages/Home.page"

const router = createBrowserRouter([
    {
        path: '/',
        element: <HomePage />,
        children: [
            {
                index: true,
                Component: UserPaging
            },
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
            // {
            //     path: 'dashboard',
            //     Component: Dashboard
            // }
        ]
    }      
])

export default function Router() {
    return (
        <RouterProvider router={router} />
    )
}