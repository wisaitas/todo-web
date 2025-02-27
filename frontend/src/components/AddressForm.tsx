import { useEffect, useState } from 'react';
import Select from './Select';
import TextArea from './TextArea';

// Mock data for demonstration - in a real app, you would fetch this from an API
const provinces = [
  { value: "bangkok", label: "Bangkok" },
  { value: "chiang_mai", label: "Chiang Mai" },
  { value: "phuket", label: "Phuket" }
];

const districtsByProvince: Record<string, Array<{ value: string, label: string }>> = {
  bangkok: [
    { value: "bangkok_pathumwan", label: "Pathumwan" },
    { value: "bangkok_bangrak", label: "Bang Rak" }
  ],
  chiang_mai: [
    { value: "chiang_mai_muang", label: "Muang Chiang Mai" },
    { value: "chiang_mai_mae_rim", label: "Mae Rim" }
  ],
  phuket: [
    { value: "phuket_muang", label: "Muang Phuket" },
    { value: "phuket_kathu", label: "Kathu" }
  ]
};

const subdistrictsByDistrict: Record<string, Array<{ value: string, label: string }>> = {
  bangkok_pathumwan: [
    { value: "pathumwan_1", label: "Pathumwan Subdistrict 1" },
    { value: "pathumwan_2", label: "Pathumwan Subdistrict 2" }
  ],
  bangkok_bangrak: [
    { value: "bangrak_1", label: "Bang Rak Subdistrict 1" },
    { value: "bangrak_2", label: "Bang Rak Subdistrict 2" }
  ],
  chiang_mai_muang: [
    { value: "muang_cm_1", label: "Muang CM Subdistrict 1" },
    { value: "muang_cm_2", label: "Muang CM Subdistrict 2" }
  ],
  chiang_mai_mae_rim: [
    { value: "mae_rim_1", label: "Mae Rim Subdistrict 1" },
    { value: "mae_rim_2", label: "Mae Rim Subdistrict 2" }
  ],
  phuket_muang: [
    { value: "muang_pk_1", label: "Muang Phuket Subdistrict 1" },
    { value: "muang_pk_2", label: "Muang Phuket Subdistrict 2" }
  ],
  phuket_kathu: [
    { value: "kathu_1", label: "Kathu Subdistrict 1" },
    { value: "kathu_2", label: "Kathu Subdistrict 2" }
  ]
};

export interface AddressData {
  province: string;
  district: string;
  subdistrict: string;
  addressDetails: string;
}

interface AddressFormProps {
  address: AddressData;
  onChange: (field: keyof AddressData, value: string) => void;
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
  const [availableDistricts, setAvailableDistricts] = useState<Array<{ value: string, label: string }>>([]);
  const [availableSubdistricts, setAvailableSubdistricts] = useState<Array<{ value: string, label: string }>>([]);

  // Update available districts when province changes
  useEffect(() => {
    if (address.province) {
      setAvailableDistricts(districtsByProvince[address.province] || []);
      
      // Reset district and subdistrict if province changes and current district is not valid
      if (address.district && !districtsByProvince[address.province]?.some(d => d.value === address.district)) {
        onChange('district', '');
        onChange('subdistrict', '');
      }
    } else {
      setAvailableDistricts([]);
    }
  }, [address.province]);

  // Update available subdistricts when district changes
  useEffect(() => {
    if (address.district) {
      setAvailableSubdistricts(subdistrictsByDistrict[address.district] || []);
      
      // Reset subdistrict if district changes and current subdistrict is not valid
      if (address.subdistrict && !subdistrictsByDistrict[address.district]?.some(s => s.value === address.subdistrict)) {
        onChange('subdistrict', '');
      }
    } else {
      setAvailableSubdistricts([]);
    }
  }, [address.district]);

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
      
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Select
          id={`province-${index}`}
          label="Province"
          options={provinces}
          value={address.province}
          onChange={(e) => onChange('province', e.target.value)}
          placeholder="Select Province"
        />
        
        <Select
          id={`district-${index}`}
          label="District"
          options={availableDistricts}
          value={address.district}
          onChange={(e) => onChange('district', e.target.value)}
          placeholder="Select District"
          disabled={!address.province}
        />
        
        <Select
          id={`subdistrict-${index}`}
          label="Subdistrict"
          options={availableSubdistricts}
          value={address.subdistrict}
          onChange={(e) => onChange('subdistrict', e.target.value)}
          placeholder="Select Subdistrict"
          disabled={!address.district}
        />
      </div>
      
      <div>
        <TextArea
          id={`addressDetails-${index}`}
          label="Address Details"
          rows={3}
          placeholder="Enter your address details"
          value={address.addressDetails}
          onChange={(e) => onChange('addressDetails', e.target.value)}
          maxLength={200}
        />
      </div>
    </div>
  );
};

export default AddressForm; 