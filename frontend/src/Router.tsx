import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { HomePage } from './pages/Home.page';
import { AdminLayout } from './components/admin/AdminLayout';

const router = createBrowserRouter([
  {
    path: '/',
    element: <AdminLayout />,
  },
]);

export function Router() {
  return <RouterProvider router={router} />;
}
