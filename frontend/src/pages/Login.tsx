import { useState } from "react";
import { Link, useLocation, useNavigate } from "react-router-dom";
import { authApi } from "../api/authApi";
import AuthLayout from "../components/AuthLayout";
import Button from "../components/Button";
import Input from "../components/Input";
import { setCookie } from '../utils/cookies';

const Login = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [successMessage] = useState<string | null>(
        location.state?.message || null
    );
  
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        if (!username || !password) {
            setError('Please enter both username and password');
            return;
        }
        
        try {
            setIsSubmitting(true);
            setError(null);
            
            const response = await authApi.login({ username, password });
            
            // Store tokens in cookies instead of localStorage
            // Set access token to expire in 1 day
            setCookie('accessToken', response.access_token, 1);
            // Set refresh token to expire in 7 days
            setCookie('refreshToken', response.refresh_token, 7);
            
            // Redirect to home or dashboard
            navigate('/dashboard');
            
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message);
            } else {
                setError('Invalid username or password');
            }
            console.error(err);
        } finally {
            setIsSubmitting(false);
        }
    };
  
    return (
        <AuthLayout title="Login">
            <form className="space-y-6" onSubmit={handleSubmit}>
                {successMessage && (
                    <div className="bg-green-50 border border-green-200 text-green-700 px-4 py-3 rounded">
                        {successMessage}
                    </div>
                )}
                
                {error && (
                    <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
                        {error}
                    </div>
                )}
                
                <Input
                    id="username"
                    label="Username"
                    placeholder="Enter your username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                
                <Input
                    id="password"
                    label="Password"
                    type="password"
                    placeholder="Enter your password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                
                <div className="flex items-center justify-between">
                    <div className="flex items-center">
                        <input
                            id="remember-me"
                            name="remember-me"
                            type="checkbox"
                            className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                        />
                        <label htmlFor="remember-me" className="ml-2 block text-sm text-gray-900">
                            Remember me
                        </label>
                    </div>

                    <div className="text-sm">
                        <a href="#" className="font-medium text-blue-600 hover:text-blue-800 hover:underline">
                            Forgot password?
                        </a>
                    </div>
                </div>
                
                <div>
                    <Button type="submit" fullWidth disabled={isSubmitting}>
                        {isSubmitting ? 'Signing in...' : 'Sign in'}
                    </Button>
                </div>
                
                <div className="text-center">
                    <p className="text-sm text-gray-600">
                        Don't have an account?{" "}
                        <Link to="/register" className="font-medium text-blue-600 hover:text-blue-800 hover:underline">
                            Register here
                        </Link>
                    </p>
                </div>
            </form>
        </AuthLayout>
    );
};

export default Login; 