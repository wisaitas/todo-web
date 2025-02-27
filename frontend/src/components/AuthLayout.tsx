import React from 'react';

interface AuthLayoutProps {
  children: React.ReactNode;
  title: string;
}

const AuthLayout = ({ children, title }: AuthLayoutProps) => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
      <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-lg">
        <h1 className="text-3xl font-bold text-center text-gray-800">{title}</h1>
        {children}
      </div>
    </div>
  );
};

export default AuthLayout; 