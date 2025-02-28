interface EnvConfig {
  API_BASE_URL: string;
}

const defaultConfig: EnvConfig = {
  API_BASE_URL: "http://localhost:8082/api/v1",
};

// Function to load runtime config
const loadRuntimeConfig = async (): Promise<EnvConfig> => {
  try {
    // Try to fetch runtime config
    const response = await fetch("/config.json");
    if (response.ok) {
      const runtimeConfig = await response.json();
      console.log("Loaded runtime config:", runtimeConfig);
      return { ...defaultConfig, ...runtimeConfig };
    }
  } catch (error) {
    console.warn("Failed to load runtime config, using default:", error);
  }

  // Fallback to environment variables
  return {
    API_BASE_URL:
      import.meta.env.VITE_API_BASE_URL || defaultConfig.API_BASE_URL,
  };
};

// Export a function to get the config
export const getConfig = async (): Promise<EnvConfig> => {
  const config = await loadRuntimeConfig();
  return config;
};

const transformApiUrl = (url: string): string => {
  // Remove trailing slash if present
  if (url.endsWith("/")) {
    url = url.slice(0, -1);
  }

  // Ensure URL has the correct protocol
  if (!url.startsWith("http://") && !url.startsWith("https://")) {
    url = `https://${url}`;
  }

  // Add /api/v1 path if not present
  if (!url.includes("/api/v1")) {
    url = `${url}/api/v1`;
  }

  return url;
};

// For backward compatibility
export const env: EnvConfig = defaultConfig;
export const apiBaseUrl = transformApiUrl(defaultConfig.API_BASE_URL);
