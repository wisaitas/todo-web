interface SelectOption {
  value: string;
  label: string;
}

interface SelectProps {
  id: string;
  label: string;
  options: SelectOption[];
  value: string;
  onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  placeholder?: string;
  disabled?: boolean;
}

const Select = ({ 
  id, 
  label, 
  options, 
  value, 
  onChange, 
  placeholder = "Select an option", 
  disabled = false 
}: SelectProps) => {
  return (
    <div className={`space-y-2 ${disabled ? 'opacity-60' : ''}`}>
      <label htmlFor={id} className="block text-sm font-medium text-gray-700">
        {label}
      </label>
      <select
        id={id}
        value={value}
        onChange={onChange}
        disabled={disabled}
        className={`
          w-full px-4 py-2 border rounded-md 
          focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 
          bg-white
          ${disabled 
            ? 'border-gray-200 text-gray-400 bg-gray-50 cursor-not-allowed' 
            : 'border-gray-300 text-gray-700 cursor-pointer hover:border-gray-400'
          }
          transition-colors duration-200
        `}
      >
        <option value="" disabled>
          {placeholder}
        </option>
        {options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </div>
  );
};

export default Select; 