import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { authApi } from "../api/authApi";
import { RegisterRequest } from "../api/types";
import AddressForm, { AddressData } from "../components/AddressForm";
import AuthLayout from "../components/AuthLayout";
import Button from "../components/Button";
import Input from "../components/Input";

const Register = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
        confirmPassword: ''
    });
    
    // State for addresses
    const [addresses, setAddresses] = useState<AddressData[]>([
        { province_id: 0, district_id: 0, sub_district_id: 0, address: '' }
    ]);
    
    // State for showing second address
    const [showSecondAddress, setShowSecondAddress] = useState(false);
    
    const [passwordMatch, setPasswordMatch] = useState(true);
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [error, setError] = useState<string | null>(null);
    
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
    
    const handleAddressChange = (index: number, field: keyof AddressData, value: number | string) => {
        const newAddresses = [...addresses];
        newAddresses[index] = { ...newAddresses[index], [field]: value };
        setAddresses(newAddresses);
    };
    
    const handleAddSecondAddress = () => {
        if (!showSecondAddress) {
            setShowSecondAddress(true);
            setAddresses([...addresses, { province_id: 0, district_id: 0, sub_district_id: 0, address: '' }]);
        }
    };
    
    const handleRemoveSecondAddress = () => {
        if (showSecondAddress) {
            setShowSecondAddress(false);
            setAddresses(addresses.slice(0, 1));
        }
    };
    
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        if (!passwordMatch) {
            setError("Passwords do not match!");
            return;
        }
        
        // Validate required fields
        if (!formData.username || !formData.email || !formData.password) {
            setError("Please fill in all required fields");
            return;
        }
        
        // Validate addresses
        const validAddresses = addresses.filter(addr => 
            addr.province_id && addr.district_id && addr.sub_district_id
        );
        
        if (validAddresses.length === 0) {
            setError("Please complete at least one address");
            return;
        }
        
        try {
            setIsSubmitting(true);
            setError(null);
            
            const registerData: RegisterRequest = {
                username: formData.username,
                email: formData.email,
                password: formData.password,
                confirm_password: formData.confirmPassword,
                addresses: validAddresses.map(addr => ({
                    province_id: addr.province_id,
                    district_id: addr.district_id,
                    sub_district_id: addr.sub_district_id,
                    address: addr.address || undefined
                }))
            };
            
            await authApi.register(registerData);
            
            // Redirect to login page on successful registration
            navigate('/login', { state: { message: 'Registration successful! Please log in.' } });
            
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message);
            } else {
                setError('An error occurred during registration');
            }
            console.error(err);
        } finally {
            setIsSubmitting(false);
        }
    };
    
    return (
        <AuthLayout title="Register">
            <form className="space-y-6" onSubmit={handleSubmit}>
                {error && (
                    <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
                        {error}
                    </div>
                )}
                
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
                    <Button type="submit" fullWidth disabled={isSubmitting}>
                        {isSubmitting ? 'Registering...' : 'Register'}
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