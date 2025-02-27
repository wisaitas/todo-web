import { useState } from "react";
import { Link } from "react-router-dom";
import AuthLayout from "../components/AuthLayout";
import Button from "../components/Button";
import Input from "../components/Input";

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
  
    const handleSubmit = (e: React.FormEvent) => {
      e.preventDefault();
      console.log({ username, password });
      // Add login logic here
    };
  
    return (
      <AuthLayout title="Login">
        <form className="space-y-6" onSubmit={handleSubmit}>
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
            <Button type="submit" fullWidth>
              Sign in
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