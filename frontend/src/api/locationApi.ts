import { apiClient } from "./client";
import { District, Province, SubDistrict } from "./types";

export const locationApi = {
  /**
   * Get all provinces
   */
  getProvinces: async (): Promise<Province[]> => {
    return apiClient.get<Province[]>("/provinces");
  },

  /**
   * Get districts by province ID
   */
  getDistrictsByProvince: async (provinceId: number): Promise<District[]> => {
    return apiClient.get<District[]>("/districts", { province_id: provinceId });
  },

  /**
   * Get sub-districts by district ID
   */
  getSubDistrictsByDistrict: async (
    districtId: number
  ): Promise<SubDistrict[]> => {
    return apiClient.get<SubDistrict[]>("/sub-districts", {
      district_id: districtId,
    });
  },
};
