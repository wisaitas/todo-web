import { createBrowserRouter, Navigate } from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute';
import Login from '../pages/Login';
import Register from '../pages/Register';
import Home from '../pages/Home';
// Import your dashboard or other protected pages
// import Dashboard from '../pages/Dashboard';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Navigate to="/login" replace />
  },
  {
    path: '/login',
    element: <Login />
  },
  {
    path: '/register',
    element: <Register />
  },
  // Protected routes
  {
    element: <ProtectedRoute />,
    children: [
      {
        path: '/home',
        element: <Home /> // Replace with your Dashboard component
      },
      // Add more protected routes here
    ]
  }
]);

export default router; 