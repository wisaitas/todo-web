import { deleteCookie, getCookie } from "./cookies";

/**
 * Check if the user is authenticated
 */
export const isAuthenticated = (): boolean => {
  return !!getCookie("accessToken");
};

/**
 * Log out the user by removing auth cookies
 */
export const logout = (): void => {
  deleteCookie("accessToken");
  deleteCookie("refreshToken");
  // Optionally redirect to login page
  window.location.href = "/login";
};

/**
 * Get the current access token
 */
export const getAccessToken = (): string | null => {
  return getCookie("accessToken");
};

/**
 * Get the current refresh token
 */
export const getRefreshToken = (): string | null => {
  return getCookie("refreshToken");
};
