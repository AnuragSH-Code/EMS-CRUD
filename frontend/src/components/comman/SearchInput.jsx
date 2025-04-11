import React from "react";

function SearchInput({ value, onChange }) {
  return (
    <input
      className="bg-white-off text-lg px-4 py-2 rounded-md max-w-72 border border-gray-300"
      placeholder="Search by name"
      type="text"
      name="search"
      id="search"
      value={value}
      onChange={onChange}
    />
  );
}

export default SearchInput;
