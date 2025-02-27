interface TextAreaProps {
  id: string;
  label: string;
  placeholder?: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
  rows?: number;
  maxLength?: number;
  showCount?: boolean;
}

const TextArea = ({
  id,
  label,
  placeholder = "",
  value,
  onChange,
  rows = 3,
  maxLength,
  showCount = true
}: TextAreaProps) => {
  const isOverLimit = maxLength ? value.length > maxLength : false;
  
  return (
    <div className="space-y-2 relative">
      <label htmlFor={id} className="block text-sm font-medium text-gray-700">
        {label}
      </label>
      
      {showCount && maxLength && (
        <div className={`absolute top-0 right-0 text-sm ${isOverLimit ? 'text-red-500 font-medium' : 'text-gray-500'}`}>
          {value.length}/{maxLength}
        </div>
      )}
      
      <textarea
        id={id}
        rows={rows}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
        className={`
          w-full px-4 py-2 border rounded-md 
          focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500
          ${isOverLimit ? 'border-red-500 focus:ring-red-500 focus:border-red-500' : 'border-gray-300'}
        `}
      />
      
      {isOverLimit && (
        <p className="text-red-500 text-sm mt-1">
          Text exceeds maximum length ({maxLength} characters)
        </p>
      )}
    </div>
  );
};

export default TextArea; 