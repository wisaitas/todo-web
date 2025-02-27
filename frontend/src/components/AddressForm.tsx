import { useEffect, useState } from 'react';
import { locationApi } from '../api/locationApi';
import { District, Province, SubDistrict } from '../api/types';
import Select from './Select';
import TextArea from './TextArea';

export interface AddressData {
  province_id: number;
  district_id: number;
  sub_district_id: number;
  address: string;
}

interface AddressFormProps {
  address: AddressData;
  onChange: (field: keyof AddressData, value: number | string) => void;
  index: number;
  isSecondary?: boolean;
  onRemove?: () => void;
}

const AddressForm = ({ 
  address, 
  onChange, 
  index, 
  isSecondary = false, 
  onRemove 
}: AddressFormProps) => {
  const [provinces, setProvinces] = useState<Province[]>([]);
  const [districts, setDistricts] = useState<District[]>([]);
  const [subDistricts, setSubDistricts] = useState<SubDistrict[]>([]);
  const [loading, setLoading] = useState({
    provinces: false,
    districts: false,
    subDistricts: false
  });
  const [error, setError] = useState<string | null>(null);

  // Fetch provinces on component mount
  useEffect(() => {
    const fetchProvinces = async () => {
      try {
        setLoading(prev => ({ ...prev, provinces: true }));
        const data = await locationApi.getProvinces();
        setProvinces(data);
        setError(null);
      } catch (err) {
        setError('Failed to load provinces');
        console.error(err);
      } finally {
        setLoading(prev => ({ ...prev, provinces: false }));
      }
    };

    fetchProvinces();
  }, []);

  // Fetch districts when province changes
  useEffect(() => {
    if (address.province_id) {
      const fetchDistricts = async () => {
        try {
          setLoading(prev => ({ ...prev, districts: true }));
          const data = await locationApi.getDistrictsByProvince(address.province_id);
          setDistricts(data);
          setError(null);
          
          // Reset district and subdistrict if province changes
          if (address.district_id && !data.some(d => d.id === address.district_id)) {
            onChange('district_id', 0);
            onChange('sub_district_id', 0);
          }
        } catch (err) {
          setError('Failed to load districts');
          console.error(err);
        } finally {
          setLoading(prev => ({ ...prev, districts: false }));
        }
      };

      fetchDistricts();
    } else {
      setDistricts([]);
    }
  }, [address.province_id]);

  // Fetch subdistricts when district changes
  useEffect(() => {
    if (address.district_id) {
      const fetchSubDistricts = async () => {
        try {
          setLoading(prev => ({ ...prev, subDistricts: true }));
          const data = await locationApi.getSubDistrictsByDistrict(address.district_id);
          setSubDistricts(data);
          setError(null);
          
          // Reset subdistrict if district changes
          if (address.sub_district_id && !data.some(s => s.id === address.sub_district_id)) {
            onChange('sub_district_id', 0);
          }
        } catch (err) {
          setError('Failed to load sub-districts');
          console.error(err);
        } finally {
          setLoading(prev => ({ ...prev, subDistricts: false }));
        }
      };

      fetchSubDistricts();
    } else {
      setSubDistricts([]);
    }
  }, [address.district_id]);

  return (
    <div className="space-y-4 p-4 border border-gray-200 rounded-lg bg-gray-50">
      <div className="flex justify-between items-center">
        <h2 className="text-xl font-semibold text-gray-700">
          {isSecondary ? "Secondary Address" : "Primary Address"}
        </h2>
        {isSecondary && onRemove && (
          <button
            type="button"
            onClick={onRemove}
            className="text-red-500 hover:text-red-700 text-sm font-medium"
          >
            Remove
          </button>
        )}
      </div>
      
      {error && (
        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
          {error}
        </div>
      )}
      
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Select
          id={`province-${index}`}
          label="Province"
          options={provinces.map(p => ({ value: p.id.toString(), label: p.name_en }))}
          value={address.province_id ? address.province_id.toString() : ''}
          onChange={(e) => onChange('province_id', parseInt(e.target.value))}
          placeholder={loading.provinces ? "Loading provinces..." : "Select Province"}
          disabled={loading.provinces}
        />
        
        <Select
          id={`district-${index}`}
          label="District"
          options={districts.map(d => ({ value: d.id.toString(), label: d.name_en }))}
          value={address.district_id ? address.district_id.toString() : ''}
          onChange={(e) => onChange('district_id', parseInt(e.target.value))}
          placeholder={loading.districts ? "Loading districts..." : "Select District"}
          disabled={!address.province_id || loading.districts}
        />
        
        <Select
          id={`subdistrict-${index}`}
          label="Subdistrict"
          options={subDistricts.map(s => ({ value: s.id.toString(), label: s.name_en }))}
          value={address.sub_district_id ? address.sub_district_id.toString() : ''}
          onChange={(e) => onChange('sub_district_id', parseInt(e.target.value))}
          placeholder={loading.subDistricts ? "Loading sub-districts..." : "Select Subdistrict"}
          disabled={!address.district_id || loading.subDistricts}
        />
      </div>
      
      <div>
        <TextArea
          id={`address-${index}`}
          label="Address Details"
          rows={3}
          placeholder="Enter your address details"
          value={address.address || ''}
          onChange={(e) => onChange('address', e.target.value)}
          maxLength={200}
        />
      </div>
    </div>
  );
};

export default AddressForm; 