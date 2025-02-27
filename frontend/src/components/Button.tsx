interface ButtonProps {
  children: React.ReactNode;
  type?: "button" | "submit" | "reset";
  fullWidth?: boolean;
  variant?: "primary" | "secondary" | "outline";
  onClick?: () => void;
}

const Button = ({ 
  children, 
  type = "button", 
  fullWidth = false, 
  variant = "primary", 
  onClick 
}: ButtonProps) => {
  const baseClasses = "px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors";
  
  const variantClasses = {
    primary: "text-white bg-blue-600 hover:bg-blue-700",
    secondary: "text-white bg-gray-600 hover:bg-gray-700",
    outline: "text-blue-600 bg-transparent border border-blue-600 hover:bg-blue-50"
  };
  
  const widthClass = fullWidth ? "w-full" : "";
  
  return (
    <button
      type={type}
      className={`${baseClasses} ${variantClasses[variant]} ${widthClass}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default Button; 