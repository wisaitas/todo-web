import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import AddressForm, { AddressData } from "../components/AddressForm";
import AuthLayout from "../components/AuthLayout";
import Button from "../components/Button";
import Input from "../components/Input";


const Register = () => {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
        confirmPassword: ''
      });
    
      // State for addresses
      const [addresses, setAddresses] = useState<AddressData[]>([
        { province: '', district: '', subdistrict: '', addressDetails: '' }
      ]);
      
      // State for showing second address
      const [showSecondAddress, setShowSecondAddress] = useState(false);
      
      const [passwordMatch, setPasswordMatch] = useState(true);
    
      // Check if passwords match
      useEffect(() => {
        if (formData.confirmPassword) {
          setPasswordMatch(formData.password === formData.confirmPassword);
        } else {
          setPasswordMatch(true);
        }
      }, [formData.password, formData.confirmPassword]);
    
      const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { id, value } = e.target;
        setFormData(prev => ({ ...prev, [id]: value }));
      };
    
      const handleAddressChange = (index: number, field: keyof AddressData, value: string) => {
        const newAddresses = [...addresses];
        newAddresses[index] = { ...newAddresses[index], [field]: value };
        setAddresses(newAddresses);
      };
    
      const handleAddSecondAddress = () => {
        if (!showSecondAddress) {
          setShowSecondAddress(true);
          setAddresses([...addresses, { province: '', district: '', subdistrict: '', addressDetails: '' }]);
        }
      };
    
      const handleRemoveSecondAddress = () => {
        if (showSecondAddress) {
          setShowSecondAddress(false);
          setAddresses(addresses.slice(0, 1));
        }
      };
    
      const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        
        if (!passwordMatch) {
          alert("Passwords do not match!");
          return;
        }
        
        console.log("Form submitted:", {
          ...formData,
          addresses
        });
        // Add registration logic here
      };
    
      return (
        <AuthLayout title="Register">
          <form className="space-y-6" onSubmit={handleSubmit}>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <Input
                id="username"
                label="Username"
                placeholder="Enter your username"
                value={formData.username}
                onChange={handleChange}
              />
              
              <Input
                id="email"
                label="Email"
                type="email"
                placeholder="Enter your email"
                value={formData.email}
                onChange={handleChange}
              />
            </div>
            
            <div className="space-y-4">
              <AddressForm 
                address={addresses[0]}
                onChange={(field, value) => handleAddressChange(0, field, value)}
                index={0}
              />
              
              {showSecondAddress && (
                <AddressForm 
                  address={addresses[1]}
                  onChange={(field, value) => handleAddressChange(1, field, value)}
                  index={1}
                  isSecondary
                  onRemove={handleRemoveSecondAddress}
                />
              )}
              
              {!showSecondAddress && (
                <div className="flex justify-center">
                  <button
                    type="button"
                    onClick={handleAddSecondAddress}
                    className="inline-flex items-center px-4 py-2 text-sm font-medium text-blue-600 bg-white border border-blue-600 rounded-md hover:bg-blue-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  >
                    <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                    </svg>
                    Add Secondary Address
                  </button>
                </div>
              )}
            </div>
            
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <Input
                id="password"
                label="Password"
                type="password"
                placeholder="Enter your password"
                value={formData.password}
                onChange={handleChange}
              />
              
              <div>
                <Input
                  id="confirmPassword"
                  label="Confirm Password"
                  type="password"
                  placeholder="Confirm your password"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                />
                {!passwordMatch && (
                  <p className="mt-1 text-sm text-red-600">Passwords do not match</p>
                )}
              </div>
            </div>
            
            <div>
              <Button type="submit" fullWidth>
                Register
              </Button>
            </div>
            
            <div className="text-center">
              <p className="text-sm text-gray-600">
                Already have an account?{" "}
                <Link to="/login" className="font-medium text-blue-600 hover:text-blue-800 hover:underline">
                  Login here
                </Link>
              </p>
            </div>
          </form>
        </AuthLayout>
      );
};

export default Register; 