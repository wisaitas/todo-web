// Base response types
export interface BaseResponse {
  id: string;
  created_at: string;
  updated_at: string;
}

// Province types
export interface Province {
  id: number;
  name_th: string;
  name_en: string;
}

// District types
export interface District {
  id: number;
  name_th: string;
  name_en: string;
  province_id: number;
}

// SubDistrict types
export interface SubDistrict {
  id: number;
  name_th: string;
  name_en: string;
  district_id: number;
  zip_code: number;
}

// API response wrappers
export interface ApiSuccessResponse<T> {
  message: string;
  data: T;
}

export interface ApiErrorResponse {
  message: string;
}

// Auth types
export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  access_token: string;
  refresh_token: string;
}

export interface RegisterAddressRequest {
  province_id: number;
  district_id: number;
  sub_district_id: number;
  address?: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
  confirm_password: string;
  addresses: RegisterAddressRequest[];
}

export interface RegisterResponse extends BaseResponse {
  username: string;
  email: string;
}

// User types
export interface AddressResponse extends BaseResponse {
  province_id: number;
  district_id: number;
  sub_district_id: number;
  address: string;
}

export interface UserResponse extends BaseResponse {
  username: string;
  email: string;
  addresses: AddressResponse[];
}
