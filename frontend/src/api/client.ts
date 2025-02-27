import { apiBaseUrl } from "../env";
import { getCookie } from "../utils/cookies";
import { ApiErrorResponse, ApiSuccessResponse } from "./types";

class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
    console.log("API Client initialized with base URL:", baseUrl);
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;
    console.log("Making API request to:", url);

    // Get the access token from cookies
    const accessToken = getCookie("accessToken");

    // Create a Headers object instead of a plain object
    const headers = new Headers(options.headers);

    // Set content type if not already set
    if (!headers.has("Content-Type")) {
      headers.set("Content-Type", "application/json");
    }

    // Add Authorization header if token exists
    if (accessToken) {
      headers.set("Authorization", `Bearer ${accessToken}`);
    }

    const config = {
      ...options,
      headers,
    };

    try {
      const response = await fetch(url, config);

      if (!response.ok) {
        // Try to parse error as JSON, but handle non-JSON responses too
        let errorMessage = "An error occurred";
        try {
          const errorData: ApiErrorResponse = await response.json();
          errorMessage = errorData.message || errorMessage;
        } catch (parseError) {
          // If response is not JSON, use text content or status text
          const textContent = await response.text();
          errorMessage =
            textContent ||
            response.statusText ||
            `HTTP error ${response.status}`;
        }

        console.error("API error:", errorMessage);
        throw new Error(errorMessage);
      }

      const contentType = response.headers.get("content-type");
      if (contentType && contentType.includes("application/json")) {
        const data: ApiSuccessResponse<T> = await response.json();
        return data.data;
      } else {
        throw new Error("Unexpected response format: not JSON");
      }
    } catch (error) {
      console.error("API request failed:", error);
      if (error instanceof Error) {
        throw new Error(error.message);
      }
      throw new Error("An unknown error occurred");
    }
  }

  public async get<T>(
    endpoint: string,
    params?: Record<string, string | number>
  ): Promise<T> {
    const url = params
      ? `${endpoint}?${new URLSearchParams(this.convertParamsToString(params))}`
      : endpoint;

    return this.request<T>(url, { method: "GET" });
  }

  public async post<T>(endpoint: string, data: any): Promise<T> {
    return this.request<T>(endpoint, {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  public async put<T>(endpoint: string, data: any): Promise<T> {
    return this.request<T>(endpoint, {
      method: "PUT",
      body: JSON.stringify(data),
    });
  }

  public async delete<T>(endpoint: string): Promise<T> {
    return this.request<T>(endpoint, { method: "DELETE" });
  }

  private convertParamsToString(
    params: Record<string, string | number>
  ): Record<string, string> {
    return Object.entries(params).reduce((acc, [key, value]) => {
      acc[key] = String(value);
      return acc;
    }, {} as Record<string, string>);
  }
}

export const apiClient = new ApiClient(apiBaseUrl);
